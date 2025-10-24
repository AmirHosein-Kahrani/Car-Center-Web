package main

import (
	"github.com/AmirHosein-Kahrani/Car-Center-Web/api"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/cache"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/db"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/db/migrations"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/logging"
)

// @Car Center API
// @version 1.0
// @description Car Store api
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:5005
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

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
	migrations.Up_1()

	api.InitServer(cfg)

}
