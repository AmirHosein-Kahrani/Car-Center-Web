package handlers

import (
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/helper"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/services"
	"github.com/gin-gonic/gin"
)

type PersianYearHandler struct {
	services *services.PersianYearService
}

func NewPersianYearHandler(cfg *config.Config) *PersianYearHandler {
	return &PersianYearHandler{
		services: services.NewPersianYearService(cfg),
	}
}

//

// CreatePersianYear godoc
// @summary create a PersianYear
// @Description create a PersianYear
// @Tags PersianYears
// @Accept json
// @Produces json
// @Param Request body dto.CreatePersianYearRequest true "create a PersianYear"
// @Success 201 {object} helper.BaseHttpResponse{Result=dto.PersianYearResponse} "PersianYear Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/years/ [post]
// @Security BearerAuth
func (h *PersianYearHandler) Create(c *gin.Context) {
	Create(c, h.services.CreatePersianYear)

}

// UpdatePersianYear godoc
// @summary Update a PersianYear
// @Description Update a PersianYear
// @Tags PersianYears
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Param Request body dto.UpdatePersianYearRequest true "Update a PersianYear"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.PersianYearResponse} "PersianYear Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/years/{id} [put]
// @Security BearerAuth
func (h *PersianYearHandler) Update(c *gin.Context) {
	Update(c, h.services.UpdatePersianYear)
}

// DeletePersianYear godoc
// @summary Delete a PersianYear
// @Description Delete a PersianYear
// @Tags PersianYears
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/years/{id} [delete]
// @Security BearerAuth
func (h *PersianYearHandler) Delete(c *gin.Context) {
	Delete(c, h.services.DeletePersianYear)
}

// GET

// GetPersianYear godoc
// @summary GET a PersianYear
// @Description get a PersianYear
// @Tags PersianYears
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.PersianYearResponse} "PersianYear Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/years/{id} [get]
// @Security BearerAuth
func (h *PersianYearHandler) GetById(c *gin.Context) {
	GetById(c, h.services.GetPersianYearById)
}

// Get PersianYear godoc
// @summary Get PersianYear
// @Description Get PersianYear
// @Tags PersianYears
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.PersianYearResponse]} "PersianYear Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/years/get-by-filter [post]
// @Security BearerAuth
func (h *PersianYearHandler) GetByFilter(c *gin.Context) {

	GetByFilter(c, h.services.GetByFilter)
}
