package migrations

import (
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/db"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/data/models"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/logging"
)

var cfg = config.GetConfig()
var logger = logging.NewLogger(cfg)

func Up_1() {
	database := db.GetDb()
	tables := []interface{}{}
	country := models.Country{}
	city := models.City{}

	if !database.Migrator().HasTable(country) {
		logger.Info(logging.Postgres, logging.Migration, "Table Country Created", nil)

		tables = append(tables, country)
	}

	if !database.Migrator().HasTable(city) {
		logger.Info(logging.Postgres, logging.Migration, "Table City Created", nil)

		tables = append(tables, city)
	}

	database.Migrator().CreateTable(tables...)

	logger.Info(logging.Postgres, logging.Migration, "Tables Created", nil)

}

func Down_1() {

}
