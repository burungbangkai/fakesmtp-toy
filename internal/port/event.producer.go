package port

import (
	"context"

	"github.com/burungbangkai/fakesmtp/internal/model"
)

type SendRawEmail func(ctx context.Context, msg *model.RawEmail) error
