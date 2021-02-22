package config

import (
	"github.com/burungbangkai/fakesmtp/internal/model"
	"github.com/burungbangkai/fakesmtp/internal/port"
	"github.com/kelseyhightower/envconfig"
)

var EnvConfigLoader port.LoadConfig = func(cfg *model.Config) error {
	return envconfig.Process("", cfg)
}
