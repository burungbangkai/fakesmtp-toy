package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/burungbangkai/fakesmtp/internal/port"
	"github.com/burungbangkai/fakesmtp/internal/usecase"

	pgrpc "github.com/burungbangkai/fakesmtp/pkg/adapter/grpc"
	empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type server struct {
	gui usecase.GetUserInboxs
	aui usecase.AddUserInbox
	dui usecase.DeleteUserInbox
	uui usecase.UpdateUserInbox
}

func (s server) GetUserInboxs(ctx context.Context, req *pgrpc.GetUserInboxsReq) (resp *pgrpc.GetUserInboxsResp, err error) {
	ibxs, err := s.gui(ctx, req.GetUserId())
	if err != nil {
		return
	}
	resp = &pgrpc.GetUserInboxsResp{}
	if len(ibxs) > 0 {
		resp.UserId = ibxs[0].UserID
	}
	for _, ibx := range ibxs {
		resp.Inboxs = append(resp.GetInboxs(), &pgrpc.UserInbox{
			Name:     ibx.InboxName,
			Password: ibx.InboxPassword,
			UserName: ibx.InboxUserName,
		})
	}
	return
}

func (s server) AddUserInbox(ctx context.Context, req *pgrpc.AddUserInboxReq) (resp *pgrpc.AddUserInboxResp, err error) {
	ibx, err := s.aui(ctx, req.GetUserId())
	if err != nil {
		return
	}
	resp = &pgrpc.AddUserInboxResp{
		UserId: ibx.UserID,
		Inbox: &pgrpc.UserInbox{
			UserName: ibx.InboxUserName,
			Name:     ibx.InboxName,
			Password: ibx.InboxPassword,
		},
	}
	return
}

func (s server) DeleteUserInbox(ctx context.Context, req *pgrpc.DeleteUserInboxReq) (resp *empty.Empty, err error) {
	err = s.dui(ctx, req.GetUserId(), req.GetInboxName())
	if err != nil {
		return
	}
	resp = &empty.Empty{}
	return
}

func (s server) UpdateUserInbox(ctx context.Context, req *pgrpc.UpdateUserInboxReq) (resp *pgrpc.UpdateUserInboxResp, err error) {
	ibx, err := s.uui(ctx, req.GetUserId(), req.GetInboxName(), req.GetNewInboxName())
	if err != nil {
		return
	}
	resp = &pgrpc.UpdateUserInboxResp{
		UserId: ibx.UserID,
		Inbox: &pgrpc.UserInbox{
			UserName: ibx.InboxUserName,
			Name:     ibx.InboxName,
			Password: ibx.InboxPassword,
		},
	}
	return
}

func NewServer(port int,
	gui usecase.GetUserInboxs,
	aui usecase.AddUserInbox,
	dui usecase.DeleteUserInbox,
	uui usecase.UpdateUserInbox,
) (port.Serve, port.StopGracefully) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	svc := &server{
		gui: gui,
		aui: aui,
		dui: dui,
		uui: uui,
	}
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
