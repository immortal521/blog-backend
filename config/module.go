package config

import (
	"log"

	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("config", fx.Provide(func(lc fx.Lifecycle, shutdowner fx.Shutdowner) (*Config, error) {
		return Load(Options{
			WatchFile: true,
			OnChange: func(cfg *Config) {
				log.Printf("1")
				log.Println("[config] changed, triggering restart")
				if err := shutdowner.Shutdown(); err != nil {
					log.Printf("[config] shutdown error: %v", err)
				}
			},
		})
	}))
}
