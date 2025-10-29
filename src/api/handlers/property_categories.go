package handlers

import (
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/helper"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/services"
	"github.com/gin-gonic/gin"
)

type PropertyCategoryHandler struct {
	services *services.PropertyCategoryService
}

func NewPropertyCategoryHandler(cfg *config.Config) *PropertyCategoryHandler {
	return &PropertyCategoryHandler{
		services: services.NewPropertyCategoryService(cfg),
	}
}

//

// CreatePropertyCategory godoc
// @summary create a PropertyCategory
// @Description create a PropertyCategory
// @Tags PropertyCategories
// @Accept json
// @Produces json
// @Param Request body dto.CreatePropertyCategoryRequest true "create a PropertyCategory"
// @Success 201 {object} helper.BaseHttpResponse{Result=dto.PropertyCategoryResponse} "PropertyCategory Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/property-categories/ [post]
// @Security BearerAuth
func (h *PropertyCategoryHandler) Create(c *gin.Context) {
	Create(c, h.services.CreatePropertyCategory)

}

// UpdatePropertyCategory godoc
// @summary Update a PropertyCategory
// @Description Update a PropertyCategory
// @Tags PropertyCategories
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Param Request body dto.UpdatePropertyCategoryRequest true "Update a PropertyCategory"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.PropertyCategoryResponse} "PropertyCategory Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/property-categories/{id} [put]
// @Security BearerAuth
func (h *PropertyCategoryHandler) Update(c *gin.Context) {
	Update(c, h.services.UpdatePropertyCategory)
}

// DeletePropertyCategory godoc
// @summary Delete a PropertyCategory
// @Description Delete a PropertyCategory
// @Tags PropertyCategories
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/property-categories/{id} [delete]
// @Security BearerAuth
func (h *PropertyCategoryHandler) Delete(c *gin.Context) {
	Delete(c, h.services.DeletePropertyCategory)
}

// GET

// GetPropertyCategory godoc
// @summary GET a PropertyCategory
// @Description get a PropertyCategory
// @Tags PropertyCategories
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.PropertyCategoryResponse} "PropertyCategory Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/property-categories/{id} [get]
// @Security BearerAuth
func (h *PropertyCategoryHandler) GetById(c *gin.Context) {
	GetById(c, h.services.GetPropertyCategoryById)
}

// Get property-categories godoc
// @summary Get property-categories
// @Description Get property-categories
// @Tags PropertyCategories
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.PropertyCategoryResponse]} "PropertyCategory Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/property-categories/get-by-filter [post]
// @Security BearerAuth
func (h *PropertyCategoryHandler) GetByFilter(c *gin.Context) {

	GetByFilter(c, h.services.GetByFilter)
}
