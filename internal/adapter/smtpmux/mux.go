package smtpmux

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/burungbangkai/fakesmtp/internal/model"
	"github.com/burungbangkai/fakesmtp/internal/port"
	"github.com/burungbangkai/fakesmtp/internal/usecase"
	"github.com/emersion/go-smtp"
	esmtp "github.com/emersion/go-smtp"
)

type (
	be struct {
		receiveEmail usecase.ReceiveEmail
		loginInbox   usecase.ProcessEmailSession
	}
	session struct {
		*model.UserInboxConfig
		fn usecase.ReceiveEmail
	}
)

func (b be) Login(state *esmtp.ConnectionState, username, password string) (esmtp.Session, error) {
	ctx := context.Background()
	cfg, err := b.loginInbox(ctx, username, password)
	if err != nil {
		return nil, err
	}
	return &session{cfg, b.receiveEmail}, nil
}

func (b be) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	return nil, smtp.ErrAuthRequired
}

func (s *session) Mail(from string, opts smtp.MailOptions) error {
	return nil
}

func (s *session) Rcpt(to string) error {
	return nil
}

func (s *session) Data(r io.Reader) error {
	ctx := context.Background()
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return s.fn(ctx, s.UserInboxConfig, b)
}

func (s *session) Reset() {}

func (s *session) Logout() error {
	return nil
}

func New(
	port int,
	hostname string,
	re usecase.ReceiveEmail,
	pes usecase.ProcessEmailSession,
) (port.Serve, port.StopGracefully) {
	mux := &be{receiveEmail: re, loginInbox: pes}
	srv := esmtp.NewServer(mux)
	srv.Addr = fmt.Sprintf(":%d", port)
	srv.Domain = hostname
	srv.AllowInsecureAuth = true
	return func() {
			fmt.Println("running at", srv.Addr)
			_ = srv.ListenAndServe()
		}, func() {
			_ = srv.Close()
		}
}
