package api

import (
	"fmt"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/middlewares"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/routers"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/validations"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	cors "github.com/rs/cors/wrapper/gin"
)

func InitServer(cfg *config.Config) {

	r := gin.New()
	// r1 := gin.Default()

	RegisterValidator()

	// r.Use(gin.Logger(), gin.Recovery(), middlewares.TestMiddleware())
	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequest())
	r.Use(cors.Default())

	RegisterRoutes(r)

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func RegisterRoutes(r *gin.Engine) {

	api := r.Group("/api")

	v1 := api.Group("/v1/")
	{
		health := v1.Group("/health")
		routers.Health(health)

		test_router := v1.Group("/test")
		routers.TestRouter(test_router)
	}
}

func RegisterValidator() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validations.IranianPhone_validator, true)
		val.RegisterValidation("valid_pass", validations.PasswordValidator, true)
	}
}
