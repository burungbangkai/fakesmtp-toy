package usecase

import (
	"context"

	"github.com/burungbangkai/fakesmtp/internal/model"
)

type AddUserInbox func(ctx context.Context, userID string) (resp *model.UserInboxConfig, err error)

func NewUserInboxAdder() AddUserInbox {
	return func(ctx context.Context, userID string) (resp *model.UserInboxConfig, err error) {
		// generate random inbox name, user/password
		// then store to db
		return
	}
}
