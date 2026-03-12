package main

import (
	"context"
	"time"

	"blog-server/api/handler"
	"blog-server/config"
	"blog-server/datastore"
	"blog-server/logger"
	"blog-server/repository"
	"blog-server/usecase"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func providerFiberApp(cfg *config.Config, log logger.Logger) (*fiber.App, error) {
	fiberCfg := fiber.Config{
		EnableTrustedProxyCheck: true,
		ProxyHeader:             fiber.HeaderXForwardedFor,
		ErrorHandler:            handler.ErrorHandler(log, cfg),
		BodyLimit:               10 * 1024 * 1024,
	}

	app := fiber.New(fiberCfg)
	return app, nil
}

func runServerLifecycle(lc fx.Lifecycle, app *fiber.App, cfg *config.Config, log logger.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Info("Server is starting")
				if err := app.Listen(cfg.Server.Addr()); err != nil {
					log.Error("Server startup failed", logger.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			timeout := cfg.Server.GracefulShutdown
			if timeout <= 0 {
				timeout = 5 * time.Second
			}
			shutdownCtx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()

			if err := app.ShutdownWithContext(shutdownCtx); err != nil {
				log.Error("Server shutdown failed", logger.Error(err))
			} else {
				log.Info("Server has been shut down successfully.")
			}
			return nil
		},
	})
}

func main() {
	app := fx.New(
		fx.Options(
			config.Module(),
			datastore.Module(),
			logger.Module(),
			repository.Module(),
			usecase.Module(),
			handler.Module(),
		),
		fx.Provide(
			providerFiberApp,
		),
		fx.Invoke(
			runServerLifecycle,
		),
	)
	app.Run()
}
