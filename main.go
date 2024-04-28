package main

import (
	"context"
	"log"

	"github.com/ddr4869/ganache-explorer/config"
	"github.com/ddr4869/ganache-explorer/ent"
	"github.com/ddr4869/ganache-explorer/internal"
)

func main() {
	cfg := config.Init()

	client, err := ent.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres password=1821 sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	router, err := internal.NewRestController(cfg)
	if err != nil {
		log.Fatalf("failed creating server: %v", err)
	}
	router.Start()
}
