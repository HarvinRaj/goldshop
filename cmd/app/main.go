package main

import (
	"github.com/HarvinRaj/goldshop/configs"
	"github.com/HarvinRaj/goldshop/internal/db"
	"github.com/HarvinRaj/goldshop/internal/routes"
	_ "github.com/joho/godotenv/autoload"
	"log"
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

	if err := routes.NewRoutes(config); err != nil {
		log.Fatal(err)
	}

	if err := db.NewDBConnection(config); err != nil {
		log.Fatal(err)
	}

}
