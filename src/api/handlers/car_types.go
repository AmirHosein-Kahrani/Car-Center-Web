package handlers

import (
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/helper"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/services"
	"github.com/gin-gonic/gin"
)

type CarTypeHandler struct {
	services *services.CarTypeService
}

func NewCarTypeHandler(cfg *config.Config) *CarTypeHandler {
	return &CarTypeHandler{
		services: services.NewCarTypeService(cfg),
	}
}

//

// CreateCarType godoc
// @summary create a CarType
// @Description create a CarType
// @Tags CarTypes
// @Accept json
// @Produces json
// @Param Request body dto.CreateCarTypeRequest true "create a CarType"
// @Success 201 {object} helper.BaseHttpResponse{Result=dto.CarTypeResponse} "CarType Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-types/ [post]
// @Security BearerAuth
func (h *CarTypeHandler) Create(c *gin.Context) {
	Create(c, h.services.CreateCarType)

}

// UpdateCarType godoc
// @summary Update a CarType
// @Description Update a CarType
// @Tags CarTypes
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Param Request body dto.UpdateCarTypeRequest true "Update a CarType"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.CarTypeResponse} "CarType Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-types/{id} [put]
// @Security BearerAuth
func (h *CarTypeHandler) Update(c *gin.Context) {
	Update(c, h.services.UpdateCarType)
}

// DeleteCarType godoc
// @summary Delete a CarType
// @Description Delete a CarType
// @Tags CarTypes
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-types/{id} [delete]
// @Security BearerAuth
func (h *CarTypeHandler) Delete(c *gin.Context) {
	Delete(c, h.services.DeleteCarType)
}

// GET

// GetCarType godoc
// @summary GET a CarType
// @Description get a CarType
// @Tags CarTypes
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.CarTypeResponse} "CarType Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-types/{id} [get]
// @Security BearerAuth
func (h *CarTypeHandler) GetById(c *gin.Context) {
	GetById(c, h.services.GetCarTypeById)
}

// Get car-types godoc
// @summary Get car-types
// @Description Get car-types
// @Tags CarTypes
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarTypeResponse]} "CarType Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-types/get-by-filter [post]
// @Security BearerAuth
func (h *CarTypeHandler) GetByFilter(c *gin.Context) {

	GetByFilter(c, h.services.GetByFilter)
}
