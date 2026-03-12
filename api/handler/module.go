package handler

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func RegisterRoutes(app *fiber.App,
	postHandler *PostHandler,
) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	RegisterPostRoute(v1, postHandler)
}

func Module() fx.Option {
	return fx.Module(
		"handler",
		fx.Provide(
			NewPostHandler,
		),
		fx.Invoke(
			RegisterRoutes,
		))
}

