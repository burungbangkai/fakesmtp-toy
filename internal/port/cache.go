package port

import (
	"context"

	"github.com/burungbangkai/fakesmtp/internal/model"
)

type GetUserInboxConfigByKey func(ctx context.Context, key string) (*model.UserInboxConfig, error)
