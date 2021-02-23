package port

import (
	"context"

	"github.com/burungbangkai/fakesmtp/internal/model"
)

type (
	InsertUserInboxConfig func(ctx context.Context, cfg *model.UserInboxConfig) error
)
