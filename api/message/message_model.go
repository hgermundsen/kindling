package message

import "time"

// Message mirrors how messages are stored in the db.
//
// Not storing time zone info here because we are only interested in how long
// ago a message was posted -- not at what time in a specific time zone. The
// front-end displays posts with something like: "Posted 8 hours ago," so the
// fact that Postgres normalizes the times that it stores to UTC is fine.
type Message struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Upvotes      int       `json:"upvotes"`
	Downvotes    int       `json:"downvotes"`
	CreationTime time.Time `json:"creationTime"`
}
