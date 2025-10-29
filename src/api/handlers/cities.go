package handlers

import (
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/helper"
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
	Create(c, h.services.CreateCity)

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
	Update(c, h.services.UpdateCity)
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
	Delete(c, h.services.DeleteCity)
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
	GetById(c, h.services.GetCityById)
}

// Get cities godoc
// @summary Get cities
// @Description Get cities
// @Tags Cities
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CityResponse]} "City Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/cities/get-by-filter [post]
// @Security BearerAuth
func (h *CityHandler) GetByFilter(c *gin.Context) {

	GetByFilter(c, h.services.GetByFilter)
}
