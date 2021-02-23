package usecase

import (
	"context"

	"github.com/burungbangkai/fakesmtp/internal/model"
)

type GetUserInboxs func(ctx context.Context, userID string) (resp []*model.UserInboxConfig, err error)
