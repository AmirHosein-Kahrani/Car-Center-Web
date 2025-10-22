package routers

import (
	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/handlers"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup, cfg *config.Config) {

	var h = handlers.NewUsersHandler(cfg)

	router.POST("/send-otp", h.SendOtp)
	router.POST("/login-by-username", h.LoginByUserName)
	router.POST("/register-by-username", h.RegisterByUserName)
	router.POST("/login-by-mobile", h.RegisterLoginByMobileRequest)
}
