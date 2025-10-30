package api

import (
	"fmt"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/middlewares"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/routers"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/validations"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/docs"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	cors "github.com/rs/cors/wrapper/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var Logger = logging.NewLogger(config.GetConfig())

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

	err := r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		Logger.Fatal(logging.General, logging.StartUp, err.Error(), nil)
	}
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {

	api := r.Group("/api")

	v1 := api.Group("/v1/")
	{

		//-----------------------------------
		// Test
		health := v1.Group("/health")
		test_router := v1.Group("/test" /*, middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"})*/)
		// User
		users := v1.Group("/users")
		// Base
		countries := v1.Group("/countries", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		cities := v1.Group("/cities", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		files := v1.Group("/files", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		companies := v1.Group("/companies", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		colors := v1.Group("/colors", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))

		// Property
		properties := v1.Group("/properties", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		propertyCategories := v1.Group("/property-categories", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		// Car
		carTypes := v1.Group("/car-types", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		gearboxes := v1.Group("/gearboxes", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModels := v1.Group("/car-models", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelColors := v1.Group("/car-model-colors", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))

		//-----------------------------------
		// Test
		routers.TestRouter(test_router)
		routers.Health(health)
		// User
		routers.User(users, cfg)
		// Base
		routers.Country(countries, cfg)
		routers.City(cities, cfg)
		routers.File(files, cfg)
		routers.Company(companies, cfg)
		routers.Color(colors, cfg)

		// Property
		routers.Property(properties, cfg)
		routers.Property(propertyCategories, cfg)
		// Car
		routers.CarType(carTypes, cfg)
		routers.Gearbox(gearboxes, cfg)
		routers.CarModel(carModels, cfg)
		routers.CarModelColor(carModelColors, cfg)

		//-----------------------------------
	}
}

func RegisterValidator() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validations.IranianPhone_validator, true)
		if err != nil {
			Logger.Error(logging.Validation, logging.StartUp, err.Error(), nil)
		}
		err = val.RegisterValidation("valid_pass", validations.PasswordValidator, true)
		if err != nil {
			Logger.Error(logging.Validation, logging.StartUp, err.Error(), nil)
		}
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
