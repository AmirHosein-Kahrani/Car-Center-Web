package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/helper"
	"github.com/gin-gonic/gin"
)

// var Logger  = logging.NewLogger(config.GetConfig())

func Create[Ti any, To any](c *gin.Context, caller func(ctx context.Context, req *Ti) (*To, error)) {

	req := new(Ti)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithError(nil, false, helper.ValidationError, err))
		return
	}

	res, err := caller(c, req)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternatlError, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, 0))
}

func Update[Ti any, To any](c *gin.Context, caller func(ctx context.Context, id int, req *Ti) (*To, error)) {

	id, _ := strconv.Atoi(c.Params.ByName("id"))

	req := new(Ti)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithError(nil, false, helper.ValidationError, err))
		return
	}

	res, err := caller(c, id, req)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternatlError, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))

}

func Delete(c *gin.Context, caller func(ctx context.Context, id int) error) {

	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, helper.ValidationError))
		return
	}

	err := caller(c, id)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternatlError, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, 222))

}

func GetById[To any](c *gin.Context, caller func(c context.Context, id int) (*To, error)) {

	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, helper.ValidationError))
		return
	}

	res, err := caller(c, id)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternatlError, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, 0))

}

func GetByFilter[Ti any, To any](c *gin.Context, caller func(c context.Context, req *Ti) (*To, error)) {

	req := new(Ti)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithError(nil, false, helper.ValidationError, err))
		return
	}

	res, err := caller(c, req)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternatlError, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))

}
