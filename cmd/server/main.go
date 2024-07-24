package main

import (
	"log"
	"net/http"

	"github.com/mfritschdotgo/techchallengefase2/configs"
	"github.com/mfritschdotgo/techchallengefase2/database"
	"github.com/mfritschdotgo/techchallengefase2/routes"
)

func main() {
	config := configs.GetConfig()

	client, err := database.ConnectDatabase(config.MONGO_USER, config.MONGO_PASSWORD, config.MONGO_HOST, config.MONGO_PORT, config.MONGO_DATABASE)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	db := client.Database(config.MONGO_DATABASE)

	r := routes.SetupRoutes(db)
	log.Fatal(http.ListenAndServe(":9090", r))
}
