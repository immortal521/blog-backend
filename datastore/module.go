// Package datastore
package datastore

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("datastore", fx.Provide(
		NewClient,
	))
}
