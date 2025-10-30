package handlers

import (
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/helper"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/services"
	"github.com/gin-gonic/gin"
)

type CarModelColorHandler struct {
	services *services.CarModelColorService
}

func NewCarModelColorHandler(cfg *config.Config) *CarModelColorHandler {
	return &CarModelColorHandler{
		services: services.NewCarModelColorService(cfg),
	}
}

//

// CreateCarModelColor godoc
// @summary create a CarModelColor
// @Description create a CarModelColor
// @Tags CarModelColors
// @Accept json
// @Produces json
// @Param Request body dto.CreateCarModelColorRequest true "create a CarModelColor"
// @Success 201 {object} helper.BaseHttpResponse{Result=dto.CarModelColorResponse} "CarModelColor Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-colors/ [post]
// @Security BearerAuth
func (h *CarModelColorHandler) Create(c *gin.Context) {
	Create(c, h.services.CreateCarModelColor)

}

// UpdateCarModelColor godoc
// @summary Update a CarModelColor
// @Description Update a CarModelColor
// @Tags CarModelColors
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Param Request body dto.UpdateCarModelColorRequest true "Update a CarModelColor"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.CarModelColorResponse} "CarModelColor Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-colors/{id} [put]
// @Security BearerAuth
func (h *CarModelColorHandler) Update(c *gin.Context) {
	Update(c, h.services.UpdateCarModelColor)
}

// DeleteCarModelColor godoc
// @summary Delete a CarModelColor
// @Description Delete a CarModelColor
// @Tags CarModelColors
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-colors/{id} [delete]
// @Security BearerAuth
func (h *CarModelColorHandler) Delete(c *gin.Context) {
	Delete(c, h.services.DeleteCarModelColor)
}

// GET

// GetCarModelColor godoc
// @summary GET a CarModelColor
// @Description get a CarModelColor
// @Tags CarModelColors
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.CarModelColorResponse} "CarModelColor Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-colors/{id} [get]
// @Security BearerAuth
func (h *CarModelColorHandler) GetById(c *gin.Context) {
	GetById(c, h.services.GetCarModelColorById)
}

// Get car-model-colors godoc
// @summary Get car-model-colors
// @Description Get car-model-colors
// @Tags CarModelColors
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelColorResponse]} "CarModelColor Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-colors/get-by-filter [post]
// @Security BearerAuth
func (h *CarModelColorHandler) GetByFilter(c *gin.Context) {

	GetByFilter(c, h.services.GetByFilter)
}
