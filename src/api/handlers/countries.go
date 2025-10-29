package handlers

import (
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/helper"
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
	Create(c, h.services.CreateCountry)
}

// UpdateCountry godoc
// @summary Update a country
// @Description Update a country
// @Tags Countries
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Param Request body dto.CreateUpdateCountryRequest true "Update a Country"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.CountryResponse} "Country Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/countries/{id} [put]
// @Security BearerAuth
func (h *CountryHandler) Update(c *gin.Context) {
	Update(c, h.services.UpdateCountry)
}

// DeleteCountry godoc
// @summary Delete a country
// @Description Delete a country
// @Tags Countries
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/countries/{id} [delete]
// @Security BearerAuth
func (h *CountryHandler) Delete(c *gin.Context) {
	Delete(c, h.services.DeleteCountry)
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
	GetById(c, h.services.GetCountryById)
}

// Get Countries godoc
// @summary Get Countries
// @Description Get Countries
// @Tags Countries
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CountryResponse]} "Country Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/countries/get-by-filter [post]
// @Security BearerAuth
func (h *CountryHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.services.GetByFilter)
}
