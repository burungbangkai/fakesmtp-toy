package usecase

import (
	"context"
	"fmt"

	"github.com/burungbangkai/fakesmtp/internal/model"
	"github.com/burungbangkai/fakesmtp/internal/port"
)

type (
	ProcessEmailSession func(ctx context.Context, userName, password string) (*model.UserInboxConfig, error)
)

func NewEmailSessionProcessor(
	getFromCache port.GetUserInboxConfigByKey,
) ProcessEmailSession {
	return func(ctx context.Context, userName, password string) (cfg *model.UserInboxConfig, err error) {
		// get inbox config from cache/repo by userName and password
		cfg, err = getFromCache(ctx, createUserInboxConfigKey(userName, password))
		// if error reject
		if err != nil {
			return
		}
		// if no matched cfg return invalid credential error
		if cfg == nil {
			err = InvalidEmailSessionCredential
			return
		}
		// else return the config
		return
	}
}

func createUserInboxConfigKey(un, pass string) string {
	return fmt.Sprintf("%s:%s", un, pass)
}
