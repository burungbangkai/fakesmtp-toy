package model

type Config struct {
	SMTPPort        int
	MaxEmailSize    int
	NatsURI         string
	NatsClusterID   string
	NatsMaxInflight int
}
