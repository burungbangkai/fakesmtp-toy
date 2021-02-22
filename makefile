# user service make file

clean:
	rm pkg/adapter/grpc/*.pb.go

init:
	protoc -I=proto --go_out=plugins=grpc:. proto/*.proto