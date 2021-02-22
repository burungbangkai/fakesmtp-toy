package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/burungbangkai/fakesmtp/internal/port"
	pgrpc "github.com/burungbangkai/fakesmtp/pkg/adapter/grpc"
	empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) GetUserInboxs(ctx context.Context, req *pgrpc.GetUserInboxsReq) (resp *pgrpc.GetUserInboxsResp, err error) {
	return nil, nil
}

func (s *server) AddUserInbox(ctx context.Context, req *pgrpc.AddUserInboxReq) (resp *pgrpc.AddUserInboxResp, err error) {
	return nil, nil
}

func (s *server) DeleteUserInbox(ctx context.Context, req *pgrpc.DeleteUserInboxReq) (resp *empty.Empty, err error) {
	return nil, nil
}

func (s *server) UpdateUserInbox(ctx context.Context, req *pgrpc.UpdateUserInboxReq) (resp *pgrpc.UpdateUserInboxResp, err error) {
	return nil, nil
}

func NewServer(port int) (port.Serve, port.StopGracefully) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	svc := &server{}
	s := grpc.NewServer()
	pgrpc.RegisterSmtpServiceServer(s, svc)
	return func() {
			fmt.Println("running at", lis.Addr().String())
			err = s.Serve(lis)
			if err != nil {
				fmt.Println(err.Error())
			}
		}, func() {
			s.GracefulStop()
		}
}
