package handlers

import (
	"net/http"
	"strconv"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/helper"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/services"
	"github.com/gin-gonic/gin"
)

type CityHandler struct {
	services *services.CityService
}

func NewCityHandler(cfg *config.Config) *CityHandler {
	return &CityHandler{
		services: services.NewCityService(cfg),
	}
}

//

// CreateCity godoc
// @summary create a City
// @Description create a City
// @Tags Cities
// @Accept json
// @Produces json
// @Param Request body dto.CreateUpdateCityRequest true "create a City"
// @Success 201 {object} helper.BaseHttpResponse{Result=dto.CityResponse} "City Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/cities/ [post]
// @Security BearerAuth
func (h *CityHandler) Create(c *gin.Context) {
	req := dto.CreateUpdateCityRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}

	res, err := h.services.CreateCity(c, &req)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, 0))
}

// UpdateCity godoc
// @summary Update a City
// @Description Update a City
// @Tags Cities
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Param Request body dto.CreateUpdateCityRequest true "Update a City"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.CityResponse} "City Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/cities/{id} [put]
// @Security BearerAuth
func (h *CityHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	req := dto.CreateUpdateCityRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}

	res, err := h.services.UpdateCity(c, id, &req)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))
}

// DeleteCity godoc
// @summary Delete a City
// @Description Delete a City
// @Tags Cities
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/cities/{id} [delete]
// @Security BearerAuth
func (h *CityHandler) Delete(c *gin.Context) {

	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, 121))
		return
	}

	err := h.services.DeleteCity(c, id)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, 222))
}

// GET

// GetCity godoc
// @summary GET a City
// @Description get a City
// @Tags Cities
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.CityResponse} "City Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/cities/{id} [get]
// @Security BearerAuth
func (h *CityHandler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, 121))
		return
	}

	res, err := h.services.GetCityById(c, id)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, 0))
}

// Get cities godoc
// @summary Get cities
// @Description Get cities
// @Tags cities
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CityResponse]} "City Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/cities/get-by-filter [post]
// @Security BearerAuth
func (h *CityHandler) GetByFilter(c *gin.Context) {

	req := dto.PaginationInputWithFilter{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}

	res, err := h.services.GetByFilter(c, &req)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))

}
