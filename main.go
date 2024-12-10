package main

import (
	"fmt"
	"log"

	"github.com/sullyh7/aetoons/api"
	"github.com/sullyh7/aetoons/api/service"
	"github.com/sullyh7/aetoons/config"
	"github.com/sullyh7/aetoons/model"
)

const ADDR = ":8080"
const SQLITE_DB = "./app.db"

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	store := model.NewMySQLStore(SQLITE_DB)

	vimeoService, err := service.NewVimeoService(config.VimeoAccessToken)
	if err != nil {
		log.Fatal(err)
	}
	server := api.NewServer(store, vimeoService, config)

	server.SetupRoutes()

	fmt.Printf("Server running on %v\n", ADDR)
	if err := server.Start(ADDR); err != nil {
		log.Fatalln(err)
	}
}
