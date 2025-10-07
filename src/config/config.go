package config

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/viper"
)



type Config struct{
	Server ServerConfig
	Postgres PostgresConfig
	Redis	RedisConfig	
}

type ServerConfig struct{

	Port string	
	RunMode string	
}


type PostgresConfig struct{
	Host string	
	Port string	
	User string	
	PassWord string	
	DbName string	
	SSLMode bool	
}




type RedisConfig struct{

	Host string	
	Port string	
	PassWord string	
	Db string	
	MinIdleconnection int
	Pollsize int
	PoolTimeout int
}

func GetConfig() *Config{
	cfgPath :=	getConfigPath(os.Getenv("APP_ENV"))
	v , err:= LoadConfig(cfgPath, "yml")
	if err != nil{
		log.Fatalf("Error In Load Config %v",err)
	}
	cfg, err := ParseConfig(v)
	if err != nil{
		log.Fatalf("Error In Parse Config %v",err)
	}
	return cfg
}

func ParseConfig(v *viper.Viper) (*Config, error){
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable To Parse Config: %v", err)
		return nil, err
	}
	return &cfg, nil
}





func LoadConfig(fileName , fileType string)(*viper.Viper, error){

	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(fileName)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil{
		log.Printf("Unable To Read Config: %v", err)
		if _ , ok := err.(viper.ConfigFileNotFoundError);ok {
			return nil, errors.New("config File Not Found")
		}
		return nil, err
	}
	return v, nil
}




func getConfigPath(env string) string{

	switch env {
	case "docker":
		return "config/config-docker"
	case "production":
		return "config/config-production"
	default:
		return "../config/config-development"
	}




}
