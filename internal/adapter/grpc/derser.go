package grpc

import (
	"github.com/burungbangkai/fakesmtp/internal/port"
	"github.com/burungbangkai/fakesmtp/internal/usecase"
	"google.golang.org/protobuf/proto"
)

var Serializer port.Serialize = func(msg interface{}) ([]byte, error) {
	pmsg, ok := msg.(proto.Message)
	if !ok {
		return nil, usecase.InvalidTypeErr
	}
	return proto.Marshal(pmsg)
}

var Deserializer port.Deserialize = func(b []byte, msg interface{}) (err error) {
	pmsg, ok := msg.(proto.Message)
	if !ok {
		return usecase.InvalidTypeErr
	}
	return proto.Unmarshal(b, pmsg)
}
