package api

import (
	"fmt"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/middlewares"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/routers"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/validations"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/docs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	cors "github.com/rs/cors/wrapper/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {

	r := gin.New()
	// r1 := gin.Default()

	RegisterSwagger(r, cfg)

	RegisterValidator()

	// r.Use(gin.Logger(), gin.Recovery(), middlewares.TestMiddleware())
	// r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequest())
	r.Use(gin.Logger(), gin.CustomRecovery(middlewares.CustomRecoveryHandler), middlewares.LimitByRequest())

	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(cors.Default())

	RegisterRoutes(r, cfg)

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {

	api := r.Group("/api")

	v1 := api.Group("/v1/")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test" /*, middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"})*/)
		users := v1.Group("/users")
		countries := v1.Group("/countries", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		cities := v1.Group("/cities", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))

		routers.TestRouter(test_router)
		routers.Health(health)
		routers.User(users, cfg)
		routers.Country(countries, cfg)
		routers.City(cities, cfg)

	}
}

func RegisterValidator() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validations.IranianPhone_validator, true)
		val.RegisterValidation("valid_pass", validations.PasswordValidator, true)
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "Golang Web Api"
	docs.SwaggerInfo.Description = "Golang Web Api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.Port)

	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
