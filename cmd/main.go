package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/sullyh7/aetoons/internal/api"
	"github.com/sullyh7/aetoons/internal/api/handlers"
	"github.com/sullyh7/aetoons/internal/api/models"
	"github.com/sullyh7/aetoons/internal/config"
	"github.com/sullyh7/aetoons/internal/db"
	"github.com/sullyh7/aetoons/internal/db/repositories"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config.NewConfig()
	DB, err := db.Connect(cfg)
	DB.AutoMigrate(&models.Show{}, &models.Episode{})
	if err != nil {
		log.Fatal(err)
	}

	showRepository := repositories.NewShowRepository(DB)
	showHandler := handlers.NewShowHandler(showRepository)

	api := api.NewAPI(cfg, showHandler)
	if err := api.Start(":8080"); err != nil {
		log.Fatal(err)
	}

}
