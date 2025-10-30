package handlers

import (
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/helper"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/services"
	"github.com/gin-gonic/gin"
)

type ColorHandler struct {
	services *services.ColorService
}

func NewColorHandler(cfg *config.Config) *ColorHandler {
	return &ColorHandler{
		services: services.NewColorService(cfg),
	}
}

//

// CreateColor godoc
// @summary create a Color
// @Description create a Color
// @Tags Colors
// @Accept json
// @Produces json
// @Param Request body dto.CreateColorRequest true "create a Color"
// @Success 201 {object} helper.BaseHttpResponse{Result=dto.ColorResponse} "Color Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/colors/ [post]
// @Security BearerAuth
func (h *ColorHandler) Create(c *gin.Context) {
	Create(c, h.services.CreateColor)

}

// UpdateColor godoc
// @summary Update a Color
// @Description Update a Color
// @Tags Colors
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Param Request body dto.UpdateColorRequest true "Update a Color"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.ColorResponse} "Color Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/colors/{id} [put]
// @Security BearerAuth
func (h *ColorHandler) Update(c *gin.Context) {
	Update(c, h.services.UpdateColor)
}

// DeleteColor godoc
// @summary Delete a Color
// @Description Delete a Color
// @Tags Colors
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/colors/{id} [delete]
// @Security BearerAuth
func (h *ColorHandler) Delete(c *gin.Context) {
	Delete(c, h.services.DeleteColor)
}

// GET

// GetColor godoc
// @summary GET a Color
// @Description get a Color
// @Tags Colors
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.ColorResponse} "Color Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/colors/{id} [get]
// @Security BearerAuth
func (h *ColorHandler) GetById(c *gin.Context) {
	GetById(c, h.services.GetColorById)
}

// Get colors godoc
// @summary Get colors
// @Description Get colors
// @Tags Colors
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.ColorResponse]} "Color Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/colors/get-by-filter [post]
// @Security BearerAuth
func (h *ColorHandler) GetByFilter(c *gin.Context) {

	GetByFilter(c, h.services.GetByFilter)
}
