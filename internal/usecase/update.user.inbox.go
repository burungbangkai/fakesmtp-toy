package usecase

import (
	"context"

	"github.com/burungbangkai/fakesmtp/internal/model"
)

type UpdateUserInbox func(ctx context.Context, userID, inboxName, newInboxName string) (resp *model.UserInboxConfig, err error)
