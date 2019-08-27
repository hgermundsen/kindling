package message

import (
	"github.com/nchaloult/kindling/db"
)

// FetchAllMessages returns all messages stored in the db
func FetchAllMessages() ([]Message, error) {
	// Execute sql statement
	rows, err := db.GetDB().Query("select * from message;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Will eventually be populated with results from above query
	output := make([]Message, 0)

	// Loop through each of the rows that satisfied the above query
	for rows.Next() {
		curMessage := Message{}

		// Populate curMessage's fields
		err = rows.Scan(
			&curMessage.ID,
			&curMessage.Title,
			&curMessage.Content,
			&curMessage.Upvotes,
			&curMessage.Downvotes,
			&curMessage.CreationTime,
		)

		output = append(output, curMessage)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return output, nil
}

// FetchMessageByID returns the message in the db with the provided id
func FetchMessageByID(id string) (Message, error) {
	// Execute sql statement
	rows, err := db.GetDB().Query("select * from message where id=$1;", id)
	if err != nil {
		// TODO: you can't return nil structs, so this is a stand-in solution
		// refactor this method to return a pointer to Message, maybe?
		return Message{}, err
	}
	defer rows.Close()

	// Will eventually be populated with results from above query
	output := Message{}

	// Loop through each of the rows that satisfied the above query
	// (should only be 1 row)
	for rows.Next() {
		// Populate output's fields
		err = rows.Scan(
			&output.ID,
			&output.Title,
			&output.Content,
			&output.Upvotes,
			&output.Downvotes,
			&output.CreationTime,
		)
	}
	err = rows.Err()
	if err != nil {
		// TODO: you can't return nil structs, so this is a stand-in solution
		// refactor this method to return a pointer to Message, maybe?
		return Message{}, err
	}

	return output, nil
}
