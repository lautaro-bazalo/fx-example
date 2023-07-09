package main

import (
	"fxdemo/internal/di/configfx"
	"fxdemo/internal/di/dbfx"
	"fxdemo/internal/di/ginfx"
	"fxdemo/internal/di/logfx"
	"fxdemo/internal/di/userfx"
	"fxdemo/internal/pkg/server"
	"go.uber.org/fx"
)

func main() {

	fx.New(
		configfx.Module,
		ginfx.Moudule,
		userfx.Module,
		logfx.Module,
		dbfx.Module,
		fx.Invoke(server.RegisterService),
		fx.Invoke(server.ServerRun),
	).Run()

}
