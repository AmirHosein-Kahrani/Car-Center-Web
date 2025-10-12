package db

import (
	"fmt"
	"log"
	"time"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitDb(cfg *config.Config) error {

	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.PassWord, cfg.Postgres.DbName,
		cfg.Postgres.SSLMode,
	)

	dbClient, err := gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDb, _ := dbClient.DB()
	err = sqlDb.Ping()
	if err != nil {
		return err
	}

	sqlDb.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDb.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)

	log.Println("connection with DB Established")
	return nil
}

func GetDb() *gorm.DB {
	return dbClient
}


func CloseDb(){
	con, _ := dbClient.DB()
	con.Close()
}