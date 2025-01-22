package main

import (
	"log"

	"github.com/HarvinRaj/goldshop/configs"
	"github.com/HarvinRaj/goldshop/internal/ctxutils"
	"github.com/HarvinRaj/goldshop/internal/db"
	"github.com/HarvinRaj/goldshop/internal/routes"
)

func main() {

	config, err := configs.LoadConfig("env.json")
	if err != nil {
		log.Fatal(err)
	}

	ctx := ctxutils.NewCTX()

	database, err := db.NewDBConnection(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	if err = routes.NewRoutes(config, database); err != nil {
		log.Fatal(err)
	}
}
