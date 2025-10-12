package main

import (
	"log"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/api"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/cache"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/db"
)

func main() {
	cfg := config.GetConfig()

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil{
		log.Fatal(err)
	}


	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil{

	println("postgresql error")
		log.Fatal(err)
	}


	api.InitServer(cfg)

}