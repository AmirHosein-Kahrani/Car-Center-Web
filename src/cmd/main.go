package main

import (
	"github.com/AmirHosein-Kahrani/Car-Center-Web/api"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/cache")

func main() {
	cfg := config.GetConfig()
	api.InitServer(cfg)
	defer cache.CloseRedis()
	cache.InitRedis(cfg)

}