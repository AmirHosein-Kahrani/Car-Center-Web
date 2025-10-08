package api

import (
	"fmt"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/routers"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/validations"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	// r1 := gin.Default()
	val, ok := binding.Validator.Engine().(*validator.Validate)

	if ok {
		val.RegisterValidation("mobile", validations.IranianPhone_validator, true)
	}
	r.Use(gin.Logger(), gin.Recovery())

	api := r.Group("/api")

	v1 := api.Group("/v1/")
	{
		health := v1.Group("/health")
		routers.Health(health)

		test_router := v1.Group("/test")
		routers.TestRouter(test_router)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
