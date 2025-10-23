package middlewares

import (
	"net/http"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/helper"
	"github.com/gin-gonic/gin"
)

func CustomRecoveryHandler(c *gin.Context, err any) {

	if err, ok := err.(error); ok {
		httpResponse := helper.GenerateBaseResponseWithError(nil, false, -500, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
		return
	}
	httpResponse := helper.GenerateBaseResponseWithAnyError(nil, false, -500, err)
	c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
}
