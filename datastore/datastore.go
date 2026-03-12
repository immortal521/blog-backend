package datastore

import (
	"fmt"

	"blog-server/config"
	"blog-server/ent"
	_ "blog-server/ent/runtime"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
)

func New(cfg *config.Config) string {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
	)

	return dsn
}

// NewClient returns an orm client
func NewClient(cfg *config.Config) (*ent.Client, error) {
	var entOptions []ent.Option
	entOptions = append(entOptions, ent.Debug())

	dsn := New(cfg)

	return ent.Open(dialect.Postgres, dsn, entOptions...)
}
