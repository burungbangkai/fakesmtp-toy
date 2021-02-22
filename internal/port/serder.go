package port

type (
	Serialize   func(msg interface{}) ([]byte, error)
	Deserialize func(b []byte, msg interface{}) (err error)
)
