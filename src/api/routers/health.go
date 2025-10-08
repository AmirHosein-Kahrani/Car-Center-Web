package routers

import (
	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/handlers"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup){
	handler := handlers.NewHealthHandler()
	
	r.GET("/", handler.Health)
}