package api

import (
	"fmt"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/routers"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/gin-gonic/gin"
) 


func InitServer(){
	cfg := config.GetConfig()
	r := gin.New()
	// r1 := gin.Default()

	r.Use(gin.Logger(), gin.Recovery())
	v1 := r.Group("api/v1/")
	{
		health := v1.Group("/health")


		routers.Health(health)
	}
	

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}


