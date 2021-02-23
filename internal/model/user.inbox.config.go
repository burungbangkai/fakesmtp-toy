package model

type UserInboxConfig struct {
	UserID        string
	InboxName     string
	InboxUserName string
	InboxPassword string
}

type UserInboxEncrypted struct {
	UserID        string
	InboxName     string
	InboxUserName []byte
	InboxPassword []byte
}
