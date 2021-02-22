package model

type RawEmail struct {
	UserID    string
	InboxName string
	RawMsg    []byte
}
