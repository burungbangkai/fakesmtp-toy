package cache

import (
	"context"
	"sync"

	"github.com/burungbangkai/fakesmtp/internal/model"
	"github.com/burungbangkai/fakesmtp/internal/port"
	"github.com/burungbangkai/fakesmtp/internal/usecase"
)

func NewUserInboxConfigSychedMapCache() port.GetUserInboxConfigByKey {
	m := sync.Map{}
	return func(ctx context.Context, key string) (*model.UserInboxConfig, error) {
		inf, ok := m.Load(key)
		if !ok {
			return nil, nil
		}
		cfg, ok := inf.(*model.UserInboxConfig)
		if !ok {
			return nil, usecase.InvalidTypeErr
		}
		return cfg, nil
	}
}
