package main

import (
	"go-streams/internal/applogger"
	"go-streams/internal/web"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			applogger.CreateLogger,
			web.CreateWebApp,
		),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
		fx.Invoke(web.RegisterStaticFiles),
		fx.Invoke(web.StartServer),
	).Run()
}
