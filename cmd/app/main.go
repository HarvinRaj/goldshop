package main

import (
	"github.com/HarvinRaj/goldshop/internal/routes"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func init() {
	err := godotenv.Load("../../configs/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	port := os.Getenv("PORT")

	router := routes.NewRoutes()

	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to run server, %v", err)
	}

}
