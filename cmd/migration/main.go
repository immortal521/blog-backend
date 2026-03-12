package main

import (
	"context"
	"log"

	"blog-server/config"
	"blog-server/datastore"
	"blog-server/ent"
	"blog-server/ent/migrate"
)

func main() {
	cfg := config.MustLoad()
	client, err := datastore.NewClient(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer func() {
		if err := client.Close(); err != nil {
			log.Printf("failed to close database client: %v", err)
		}
	}()
	createSchema(client)
}

func createSchema(client *ent.Client) {
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(false),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
