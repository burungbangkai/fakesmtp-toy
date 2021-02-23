package usecase

import (
	"context"

	"github.com/burungbangkai/fakesmtp/internal/model"
	"github.com/burungbangkai/fakesmtp/internal/port"
	"github.com/segmentio/ksuid"
)

type AddUserInbox func(ctx context.Context, userID string) (resp *model.UserInboxConfig, err error)

func NewUserInboxAdder(
	randomSeed int64,
	inserter port.InsertUserInboxConfig,
	passwordLength int,
	userNameLength int,
) AddUserInbox {
	var randStrGen generateRandomString = newRandomStringGenerator(randomSeed)
	return func(ctx context.Context, userID string) (resp *model.UserInboxConfig, err error) {
		// generate unique random inbox name,
		cfg := &model.UserInboxConfig{
			UserID:    userID,
			InboxName: ksuid.New().String(),
		}
		// random user and password
		cfg.InboxPassword = randStrGen(passwordLength)
		cfg.InboxUserName = randStrGen(userNameLength)
		// then store to db
		err = inserter(ctx, cfg)
		if err != nil {
			return
		}
		// ... ;)
		return
	}
}
