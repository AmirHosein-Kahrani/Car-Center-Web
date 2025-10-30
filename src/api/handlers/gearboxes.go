package handlers

import (
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/helper"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/services"
	"github.com/gin-gonic/gin"
)

type GearboxHandler struct {
	services *services.GearboxService
}

func NewGearboxHandler(cfg *config.Config) *GearboxHandler {
	return &GearboxHandler{
		services: services.NewGearboxService(cfg),
	}
}

//

// CreateGearbox godoc
// @summary create a Gearbox
// @Description create a Gearbox
// @Tags Gearboxes
// @Accept json
// @Produces json
// @Param Request body dto.CreateGearboxRequest true "create a Gearbox"
// @Success 201 {object} helper.BaseHttpResponse{Result=dto.GearboxResponse} "Gearbox Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/gearboxes/ [post]
// @Security BearerAuth
func (h *GearboxHandler) Create(c *gin.Context) {
	Create(c, h.services.CreateGearbox)

}

// UpdateGearbox godoc
// @summary Update a Gearbox
// @Description Update a Gearbox
// @Tags Gearboxes
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Param Request body dto.UpdateGearboxRequest true "Update a Gearbox"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.GearboxResponse} "Gearbox Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/gearboxes/{id} [put]
// @Security BearerAuth
func (h *GearboxHandler) Update(c *gin.Context) {
	Update(c, h.services.UpdateGearbox)
}

// DeleteGearbox godoc
// @summary Delete a Gearbox
// @Description Delete a Gearbox
// @Tags Gearboxes
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/gearboxes/{id} [delete]
// @Security BearerAuth
func (h *GearboxHandler) Delete(c *gin.Context) {
	Delete(c, h.services.DeleteGearbox)
}

// GET

// GetGearbox godoc
// @summary GET a Gearbox
// @Description get a Gearbox
// @Tags Gearboxes
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.GearboxResponse} "Gearbox Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/gearboxes/{id} [get]
// @Security BearerAuth
func (h *GearboxHandler) GetById(c *gin.Context) {
	GetById(c, h.services.GetGearboxById)
}

// Get gearboxes godoc
// @summary Get gearboxes
// @Description Get gearboxes
// @Tags Gearboxes
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.GearboxResponse]} "Gearbox Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/gearboxes/get-by-filter [post]
// @Security BearerAuth
func (h *GearboxHandler) GetByFilter(c *gin.Context) {

	GetByFilter(c, h.services.GetByFilter)
}
