package main

import (

	"github.com/AmirHosein-Kahrani/Car-Center-Web/api"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/cache"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/db"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/logging"
)

// @securityDefinition.apikey AuthBearer
// @in header
// @name authorization
func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.StartUp, err.Error(), nil)
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {

		logger.Fatal(logging.Postgres, logging.StartUp, err.Error(), nil)

	}

	api.InitServer(cfg)

}
