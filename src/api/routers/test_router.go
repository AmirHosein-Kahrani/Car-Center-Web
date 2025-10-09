package routers

import (
	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/handlers"
	"github.com/gin-gonic/gin"
)

func TestRouter(r *gin.RouterGroup) {
	h := handlers.NewTestHandler()

	r.GET("/", h.Test)
	r.GET("/users", h.Users)
	r.GET("/user/:id", h.UserById)
	r.GET("/user/get-user-by-name/:username", h.UserByUsername)
	r.GET("/user/:id/accounts", h.Accounts)
	r.POST("/user/add-user", h.AddUser)

	// 03
	r.POST("/binder/header1", h.HeaderBinder1)
	r.POST("/binder/header2", h.HeaderBinder2)

	r.POST("/binder/query1", h.QueryBinder1)
	r.POST("/binder/query2", h.QueryBinder2)
	r.POST("/binder/uri1/:id/:name", h.UriBinder)

	r.POST("/binder/body", h.BodyBinder)

	r.POST("/binder/form", h.FormBinder)


	r.POST("/binder/File", h.FileBinder)


	r.GET("/pass", h.PassHandler)

}
