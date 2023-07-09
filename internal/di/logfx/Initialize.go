package logfx

import (
	"fxdemo/config"
	"fxdemo/internal/pkg/log"
	"go.uber.org/fx"
)

var Module = fx.Provide(provideLog)

func provideLog(conf *config.Application) log.Logger {
	return log.NewLogger(conf.Logger)
}
