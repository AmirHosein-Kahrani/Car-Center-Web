package handlers

import (
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/helper"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/services"
	"github.com/gin-gonic/gin"
)

type CarModelPropertyHandler struct {
	services *services.CarModelPropertyService
}

func NewCarModelPropertyHandler(cfg *config.Config) *CarModelPropertyHandler {
	return &CarModelPropertyHandler{
		services: services.NewCarModelPropertyService(cfg),
	}
}

//

// CreateCarModelProperty godoc
// @summary create a CarModelProperty
// @Description create a CarModelProperty
// @Tags CarModelProperties
// @Accept json
// @Produces json
// @Param Request body dto.CreateCarModelPropertyRequest true "create a CarModelProperty"
// @Success 201 {object} helper.BaseHttpResponse{Result=dto.CarModelPropertyResponse} "CarModelProperty Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-properties/ [post]
// @Security BearerAuth
func (h *CarModelPropertyHandler) Create(c *gin.Context) {
	Create(c, h.services.CreateCarModelProperty)

}

// UpdateCarModelProperty godoc
// @summary Update a CarModelProperty
// @Description Update a CarModelProperty
// @Tags CarModelProperties
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Param Request body dto.UpdateCarModelPropertyRequest true "Update a CarModelProperty"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.CarModelPropertyResponse} "CarModelProperty Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-properties/{id} [put]
// @Security BearerAuth
func (h *CarModelPropertyHandler) Update(c *gin.Context) {
	Update(c, h.services.UpdateCarModelProperty)
}

// DeleteCarModelProperty godoc
// @summary Delete a CarModelProperty
// @Description Delete a CarModelProperty
// @Tags CarModelProperties
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-properties/{id} [delete]
// @Security BearerAuth
func (h *CarModelPropertyHandler) Delete(c *gin.Context) {
	Delete(c, h.services.DeleteCarModelProperty)
}

// GET

// GetCarModelProperty godoc
// @summary GET a CarModelProperty
// @Description get a CarModelProperty
// @Tags CarModelProperties
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.CarModelPropertyResponse} "CarModelProperty Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-properties/{id} [get]
// @Security BearerAuth
func (h *CarModelPropertyHandler) GetById(c *gin.Context) {
	GetById(c, h.services.GetCarModelPropertyById)
}

// Get CarModelProperty godoc
// @summary Get CarModelProperty
// @Description Get CarModelProperty
// @Tags CarModelProperties
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelPropertyResponse]} "CarModelProperty Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-properties/get-by-filter [post]
// @Security BearerAuth
func (h *CarModelPropertyHandler) GetByFilter(c *gin.Context) {

	GetByFilter(c, h.services.GetByFilter)
}
