package main

import (
	"context"
	"github.com/HarvinRaj/goldshop/configs"
	"github.com/HarvinRaj/goldshop/internal/db"
	"github.com/HarvinRaj/goldshop/internal/routes"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"time"
)

var config *configs.Config

func init() {

	var err error

	config, err = configs.LoadConfig("env.json")
	if err != nil {
		log.Fatalf("Failed to load config, %v", err)
	}
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbConnection, err := db.NewDBConnection(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	if err = routes.NewRoutes(config, dbConnection); err != nil {
		log.Fatal(err)
	}
}
