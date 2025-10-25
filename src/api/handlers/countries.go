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

type CountryHandler struct {
	services *services.CountryService
}

func NewCountryHandler(cfg *config.Config) *CountryHandler {
	return &CountryHandler{services: services.NewCountryService(cfg)}
}

// CreateCountry godoc
// @summary create a country
// @Description create a country
// @Tags Countries
// @Accept json
// @Produces json
// @Param Request body dto.CreateUpdateCountryRequest true "create a country"
// @Success 201 {object} helper.BaseHttpResponse{Result=dto.CountryResponse} "Country Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/countries/ [post]
// @Security BearerAuth
func (h *CountryHandler) Create(c *gin.Context) {
	req := dto.CreateUpdateCountryRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}

	res, err := h.services.CreateCountry(c, &req)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, 0))
}

// UpdateCountry godoc
// @summary Update a country
// @Description Update a country
// @Tags Countries
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Param name body string true "name"
// @Success 201 {object} helper.BaseHttpResponse{Result=dto.CountryResponse} "Country Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/countries/{id} [put]
// @Security BearerAuth
func (h *CountryHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	req := dto.CreateUpdateCountryRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}

	res, err := h.services.UpdateCountry(c, id, &req)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))
}

// DeleteCountry godoc
// @summary Delete a country
// @Description Delete a country
// @Tags Countries
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 201 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/countries/{id} [delete]
// @Security BearerAuth
func (h *CountryHandler) Delete(c *gin.Context) {

	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, 121))
		return
	}

	err := h.services.DeleteCountry(c, id)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, 222))
}

// GET

// GetCountry godoc
// @summary GET a country
// @Description get a country
// @Tags Countries
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.CountryResponse} "Country Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/countries/{id} [get]
// @Security BearerAuth
func (h *CountryHandler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, 121))
		return
	}

	res, err := h.services.GetCountryById(c, id)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, 0))
}

// TODO: DON'T FORGET TO WRITE CODE
// GET BY FILTER
func (h *CountryHandler) GetByFilter(c *gin.Context) {}



