package message

import (
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

func TestGetAllMessages(t *testing.T) {
	verb := "GET"
	uri := "/api/message"

	// Get mock data that will be compared against response body
	mockDataHappyPath, _ := mockFetchAllMessages()
	// Test "happy path"
	testEndpointBehavior(
		t,
		"GetAllMessages (happy path) - ", verb, uri,
		mockConHappyPath.GetAllMessages,
		mockDataHappyPath,
		200,
	)

	// Test when no messages exist
	testEndpointBehavior(
		t,
		"GetAllMessages (not found) - ", verb, uri,
		mockConNotFound.GetAllMessages,
		"Not Found",
		404,
	)
}

// Helper function that hits an endpoint and tests for the expected status code
// and response body content
func testEndpointBehavior(
	t *testing.T,
	prefix, verb, uri string,
	mockHandler func(w http.ResponseWriter, r *http.Request, ps httprouter.Params),
	mockData interface{},
	wantStatusCode int,
) {
	request, err := http.NewRequest(verb, uri, nil)
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
	default:
		t.Fatalf("%sdidn't recognize provided HTTP verb: %s", prefix, verb)
	}

	// Hit the endpoint
	router.ServeHTTP(recorder, request)

	// Test status code
	if recorder.Code != wantStatusCode {
		t.Errorf("%sunexpected status code: %d; want: %d", prefix, recorder.Code, wantStatusCode)
	}

	// If the status code was a 404, then there's no need to look at the
	// response body, too
	if wantStatusCode != 404 {
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
