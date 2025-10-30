package handlers

import (
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/helper"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/services"
	"github.com/gin-gonic/gin"
)

type CarModelYearHandler struct {
	services *services.CarModelYearService
}

func NewCarModelYearHandler(cfg *config.Config) *CarModelYearHandler {
	return &CarModelYearHandler{
		services: services.NewCarModelYearService(cfg),
	}
}

//

// CreateCarModelYear godoc
// @summary create a CarModelYear
// @Description create a CarModelYear
// @Tags CarModelYears
// @Accept json
// @Produces json
// @Param Request body dto.CreateCarModelYearRequest true "create a CarModelYear"
// @Success 201 {object} helper.BaseHttpResponse{Result=dto.CarModelYearResponse} "CarModelYear Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-years/ [post]
// @Security BearerAuth
func (h *CarModelYearHandler) Create(c *gin.Context) {
	Create(c, h.services.CreateCarModelYear)

}

// UpdateCarModelYear godoc
// @summary Update a CarModelYear
// @Description Update a CarModelYear
// @Tags CarModelYears
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Param Request body dto.UpdateCarModelYearRequest true "Update a CarModelYear"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.CarModelYearResponse} "CarModelYear Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-years/{id} [put]
// @Security BearerAuth
func (h *CarModelYearHandler) Update(c *gin.Context) {
	Update(c, h.services.UpdateCarModelYear)
}

// DeleteCarModelYear godoc
// @summary Delete a CarModelYear
// @Description Delete a CarModelYear
// @Tags CarModelYears
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-years/{id} [delete]
// @Security BearerAuth
func (h *CarModelYearHandler) Delete(c *gin.Context) {
	Delete(c, h.services.DeleteCarModelYear)
}

// GET

// GetCarModelYear godoc
// @summary GET a CarModelYear
// @Description get a CarModelYear
// @Tags CarModelYears
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.CarModelYearResponse} "CarModelYear Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-years/{id} [get]
// @Security BearerAuth
func (h *CarModelYearHandler) GetById(c *gin.Context) {
	GetById(c, h.services.GetCarModelYearById)
}

// Get CarModelYear godoc
// @summary Get CarModelYear
// @Description Get CarModelYear
// @Tags CarModelYears
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelYearResponse]} "CarModelYear Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-years/get-by-filter [post]
// @Security BearerAuth
func (h *CarModelYearHandler) GetByFilter(c *gin.Context) {

	GetByFilter(c, h.services.GetByFilter)
}
