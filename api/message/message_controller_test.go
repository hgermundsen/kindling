package message

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Mock functions

func mockFetchAllMessages() ([]Message, error) {
	return []Message{
		{
			ID:           1,
			Title:        "message 1",
			Content:      "contents of message 1",
			Upvotes:      2,
			Downvotes:    1,
			Flags:        0,
			CreationTime: time.Date(2019, time.August, 30, 14, 50, 0, 0, time.UTC),
		},
		{
			ID:           2,
			Title:        "message 2",
			Content:      "contents of message 2",
			Upvotes:      4,
			Downvotes:    0,
			Flags:        0,
			CreationTime: time.Date(2019, time.August, 30, 15, 20, 0, 0, time.UTC),
		},
	}, nil
}

func mockFetchMessageByID(id string) (Message, error) {
	return Message{
		ID:           1,
		Title:        "message 1",
		Content:      "contents of message 1",
		Upvotes:      2,
		Downvotes:    1,
		Flags:        0,
		CreationTime: time.Date(2019, time.August, 30, 14, 50, 0, 0, time.UTC),
	}, nil
}

func mockInsertMessage(message Message) error {
	return nil
}

// Mock structs from mock functions

var mockRepoHappyPath *Repo = NewRepo(
	mockFetchAllMessages,
	mockFetchMessageByID,
	mockInsertMessage,
)

var mockRepoNotFound *Repo = NewRepo(
	func() ([]Message, error) { return nil, nil },
	func(string) (Message, error) { return Message{}, nil },
	func(Message) error { return nil },
)

var mockConHappyPath *Controller = NewController(mockRepoHappyPath)

var mockConNotFound *Controller = NewController(mockRepoNotFound)

// Tests
//
// https://blog.questionable.services/article/testing-http-handlers-go/

// Test GET /api/message
func TestGetAllMessages(t *testing.T) {
	verb := "GET"
	uri := "/api/message"

	// Get mock data that will be compared against response body
	mockDataHappyPath, _ := mockFetchAllMessages()
	// Test "happy path"
	testEndpointBehavior(
		t,
		"GetAllMessages (happy path) - ", verb, uri,
		nil,
		mockConHappyPath.GetAllMessages,
		mockDataHappyPath,
		200,
	)

	// Test when no messages exist
	testEndpointBehavior(
		t,
		"GetAllMessages (not found) - ", verb, uri,
		nil,
		mockConNotFound.GetAllMessages,
		"Not Found",
		404,
	)
}

// Test GET /api/message/:id
func TestGetMessageByID(t *testing.T) {
	verb := "GET"
	uri := "/api/message/:id"

	// Get mock data that will be compared against response body
	mockDataHappyPath, _ := mockFetchMessageByID("dummyId")
	// Test happy path
	testEndpointBehavior(
		t,
		"GetMessageByID (happy path) - ", verb, uri,
		nil,
		mockConHappyPath.GetMessageByID,
		mockDataHappyPath,
		200,
	)

	// Test when no messages exist
	testEndpointBehavior(
		t,
		"GetMessageByID (not found) - ", verb, uri,
		nil,
		mockConNotFound.GetMessageByID,
		nil,
		404,
	)
}

// Test POST /api/message
func TestCreateMessage(t *testing.T) {
	verb := "POST"
	uri := "/api/message"

	mockReqBodyBare := []byte(`{
"title": "testing title",
"content": "testing content"
	}`)
	// Test happy path
	testEndpointBehavior(
		t,
		"CreateMessage (happy path) - ", verb, uri,
		mockReqBodyBare,
		mockConHappyPath.CreateMessage,
		nil,
		204,
	)

	mockReqBodyEmpty := []byte(`{}`)
	// Test empty request body
	testEndpointBehavior(
		t,
		"CreateMessage (empty request body) - ", verb, uri,
		mockReqBodyEmpty,
		mockConHappyPath.CreateMessage,
		nil,
		400,
	)

	mockReqBodyTitleNoContent := []byte(`{
"title": "testing title"
	}`)
	// Test request body w/ title, but no content
	testEndpointBehavior(
		t,
		"CreateMessage (request body w/ title, no content) - ", verb, uri,
		mockReqBodyTitleNoContent,
		mockConHappyPath.CreateMessage,
		nil,
		400,
	)
}

// Helper function that hits an endpoint and tests for the expected status code
// and response body content
//
// This might find its way into common.go
func testEndpointBehavior(
	t *testing.T,
	prefix, verb, uri string,
	reqBody []byte,
	mockHandler func(w http.ResponseWriter, r *http.Request, ps httprouter.Params),
	mockData interface{},
	wantStatusCode int,
) {
	// "Set" implementation in golang
	statusCodesNoBodyWhiteList := map[int]bool{204: true, 400: true, 404: true}

	request, err := http.NewRequest(verb, uri, bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatalf("%s%s", prefix, err)
	}

	// Create a recorder instead of a ResponseWriter so that we can test what
	// would normally be written to a ResponseWriter
	recorder := httptest.NewRecorder()

	router := httprouter.New()
	switch verb {
	case "GET":
		router.GET(uri, mockHandler)
	case "POST":
		router.POST(uri, mockHandler)
	default:
		t.Fatalf("%sdidn't recognize provided HTTP verb: %s", prefix, verb)
	}

	// Hit the endpoint
	router.ServeHTTP(recorder, request)

	// Test status code
	if recorder.Code != wantStatusCode {
		t.Errorf("%sunexpected status code: %d; want: %d", prefix, recorder.Code, wantStatusCode)
	}

	// If the status code wasn't in our list of status codes that mean response
	// bodies don't need to be checked (like 204, 404), then test the response
	// body
	if _, present := statusCodesNoBodyWhiteList[wantStatusCode]; !present {
		// Test response body
		mockJSON, err := json.Marshal(mockData)
		if err != nil {
			t.Errorf("%scouldn't marshal mock data into json:\n%s", prefix, err)
		}

		// Stringify and clean up results and expected results
		got := strings.TrimSpace(recorder.Body.String())
		want := strings.TrimSpace(string(mockJSON))

		if !strings.Contains(got, want) {
			t.Errorf("%sunexpected response body:\ngot: %s\n\nwant: %s", prefix, got, want)
		}
	}
}
