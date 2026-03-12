package usecase

import (
	"blog-server/usecase/post"

	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("usecase", fx.Provide(
		post.NewService,
	))
}
