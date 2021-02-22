package port

import "github.com/burungbangkai/fakesmtp/internal/model"

type LoadConfig func(*model.Config) error
