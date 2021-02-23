package usecase

import (
	"context"

	"github.com/burungbangkai/fakesmtp/internal/model"
)

type AddUserInbox func(ctx context.Context, userID string) (resp *model.UserInboxConfig, err error)
