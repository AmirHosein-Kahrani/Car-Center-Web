package routers

import (
	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/handlers"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/gin-gonic/gin"
)

func PropertyCategory(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewPropertyCategoryHandler(cfg)
	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}

func Property(r *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewPropertyHandler(cfg)
	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST("/get-by-filter", h.GetByFilter)
}
