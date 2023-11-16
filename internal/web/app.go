package web

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func CreateWebApp(lc fx.Lifecycle, logger *zap.Logger) *fiber.App {
	app := fiber.New()

	lc.Append(fx.Hook{OnStop: func(ctx context.Context) error {
		logger.Info("Shutting down")
		return app.Shutdown()
	}})
	app.Use(cors.New(cors.Config{
		//AllowCredentials: true,
		AllowOrigins: "*",
		//AllowHeaders:     "Origin, Content-Type, Accept, Accept-Language, Content-Length",
	}))
	return app
}

func RegisterStaticFiles(app *fiber.App, logger *zap.Logger) error {
	logger.Info("Setting up static routes")
	app.Static("/data", "./static")
	app.Static("/", "./web/static")

	return nil
}

func StartServer(app *fiber.App, logger *zap.Logger) {
	logger.Info("Starting up")
	if err := app.Listen(":3500"); err != nil {
		panic(err)
	}
}
