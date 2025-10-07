package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type HealthHandler struct{

}

func NewHealthHandler() *HealthHandler{
	return &HealthHandler{}
}


func (h *HealthHandler) Health(c *gin.Context){
	c.JSON(http.StatusOK, "Working")
}



func (h *HealthHandler) HealthPost(c *gin.Context){
	c.JSON(http.StatusOK, "Working Post")
}




func (h *HealthHandler) HealthPostByID(c *gin.Context){

	id := c.Params.ByName("id")
	c.JSON(http.StatusOK,fmt.Sprintf("Working Post BY ID: %s, %T", id, id))
}
