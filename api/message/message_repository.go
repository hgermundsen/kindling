package message

// FetchAllMessages returns all messages stored in the db
func FetchAllMessages() []string {
	dummyMessages := [2]string{"message one", "message two"}
	return dummyMessages[:]
}
