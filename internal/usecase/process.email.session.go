package usecase

import (
	"context"

	"github.com/burungbangkai/fakesmtp/internal/model"
)

type (
	ProcessEmailSession func(ctx context.Context, userName, password string) (*model.UserInboxConfig, error)
)

func NewEmailSessionProcessor() ProcessEmailSession {
	return func(ctx context.Context, userName, password string) (cfg *model.UserInboxConfig, err error) {
		// get inbox config from cache/repo by userName and password
		// if error reject
		// if no matched cfg return invalid credential error
		// else return the config
		return
	}
}
