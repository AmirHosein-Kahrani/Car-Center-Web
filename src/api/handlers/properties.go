package handlers

import (
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/helper"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/services"
	"github.com/gin-gonic/gin"
)

type PropertyHandler struct {
	services *services.PropertyService
}

func NewPropertyHandler(cfg *config.Config) *PropertyHandler {
	return &PropertyHandler{
		services: services.NewPropertyService(cfg),
	}
}

//

// CreateProperty godoc
// @summary create a Property
// @Description create a Property
// @Tags Properties
// @Accept json
// @Produces json
// @Param Request body dto.CreatePropertyRequest true "create a Property"
// @Success 201 {object} helper.BaseHttpResponse{Result=dto.PropertyResponse} "Property Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/properties/ [post]
// @Security BearerAuth
func (h *PropertyHandler) Create(c *gin.Context) {
	Create(c, h.services.CreateProperty)

}

// UpdateProperty godoc
// @summary Update a Property
// @Description Update a Property
// @Tags Properties
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Param Request body dto.UpdatePropertyRequest true "Update a Property"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.PropertyResponse} "Property Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/properties/{id} [put]
// @Security BearerAuth
func (h *PropertyHandler) Update(c *gin.Context) {
	Update(c, h.services.UpdateProperty)
}

// DeleteProperty godoc
// @summary Delete a Property
// @Description Delete a Property
// @Tags Properties
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/properties/{id} [delete]
// @Security BearerAuth
func (h *PropertyHandler) Delete(c *gin.Context) {
	Delete(c, h.services.DeleteProperty)
}

// GET

// GetProperty godoc
// @summary GET a Property
// @Description get a Property
// @Tags Properties
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.PropertyResponse} "Property Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/properties/{id} [get]
// @Security BearerAuth
func (h *PropertyHandler) GetById(c *gin.Context) {
	GetById(c, h.services.GetPropertyById)
}

// Get properties godoc
// @summary Get properties
// @Description Get properties
// @Tags Properties
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.PropertyResponse]} "Property Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/properties/get-by-filter [post]
// @Security BearerAuth
func (h *PropertyHandler) GetByFilter(c *gin.Context) {

	GetByFilter(c, h.services.GetByFilter)
}
