package handlers

import (
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/helper"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/services"
	"github.com/gin-gonic/gin"
)

type CarModelHandler struct {
	services *services.CarModelService
}

func NewCarModelHandler(cfg *config.Config) *CarModelHandler {
	return &CarModelHandler{
		services: services.NewCarModelService(cfg),
	}
}

//

// CreateCarModel godoc
// @summary create a CarModel
// @Description create a CarModel
// @Tags CarModels
// @Accept json
// @Produces json
// @Param Request body dto.CreateCarModelRequest true "create a CarModel"
// @Success 201 {object} helper.BaseHttpResponse{Result=dto.CarModelResponse} "CarModel Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-models/ [post]
// @Security BearerAuth
func (h *CarModelHandler) Create(c *gin.Context) {
	Create(c, h.services.CreateCarModel)

}

// UpdateCarModel godoc
// @summary Update a CarModel
// @Description Update a CarModel
// @Tags CarModels
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Param Request body dto.UpdateCarModelRequest true "Update a CarModel"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.CarModelResponse} "CarModel Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-models/{id} [put]
// @Security BearerAuth
func (h *CarModelHandler) Update(c *gin.Context) {
	Update(c, h.services.UpdateCarModel)
}

// DeleteCarModel godoc
// @summary Delete a CarModel
// @Description Delete a CarModel
// @Tags CarModels
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-models/{id} [delete]
// @Security BearerAuth
func (h *CarModelHandler) Delete(c *gin.Context) {
	Delete(c, h.services.DeleteCarModel)
}

// GET

// GetCarModel godoc
// @summary GET a CarModel
// @Description get a CarModel
// @Tags CarModels
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.CarModelResponse} "CarModel Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-models/{id} [get]
// @Security BearerAuth
func (h *CarModelHandler) GetById(c *gin.Context) {
	GetById(c, h.services.GetCarModelById)
}

// Get car-models godoc
// @summary Get car-models
// @Description Get car-models
// @Tags CarModels
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelResponse]} "CarModel Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-models/get-by-filter [post]
// @Security BearerAuth
func (h *CarModelHandler) GetByFilter(c *gin.Context) {

	GetByFilter(c, h.services.GetByFilter)
}
