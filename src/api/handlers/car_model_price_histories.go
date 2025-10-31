package handlers

import (
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/helper"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/services"
	"github.com/gin-gonic/gin"
)

type CarModelPriceHistoryHandler struct {
	services *services.CarModelPriceHistoryService
}

func NewCarModelPriceHistoryHandler(cfg *config.Config) *CarModelPriceHistoryHandler {
	return &CarModelPriceHistoryHandler{
		services: services.NewCarModelPriceHistoryService(cfg),
	}
}

// CreateCarModelPriceHistory godoc
// @summary create a CarModelPriceHistory
// @Description create a CarModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept json
// @Produces json
// @Param Request body dto.CreateCarModelPriceHistoryRequest true "create a CarModelPriceHistory"
// @Success 201 {object} helper.BaseHttpResponse{Result=dto.CarModelPriceHistoryResponse} "CarModelPriceHistory Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-price-histories/ [post]
// @Security BearerAuth
func (h *CarModelPriceHistoryHandler) Create(c *gin.Context) {
	Create(c, h.services.CreateCarModelPriceHistory)

}

// UpdateCarModelPriceHistory godoc
// @summary Update a CarModelPriceHistory
// @Description Update a CarModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Param Request body dto.UpdateCarModelPriceHistoryRequest true "Update a CarModelPriceHistory"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.CarModelPriceHistoryResponse} "CarModelPriceHistory Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-price-histories/{id} [put]
// @Security BearerAuth
func (h *CarModelPriceHistoryHandler) Update(c *gin.Context) {
	Update(c, h.services.UpdateCarModelPriceHistory)
}

// DeleteCarModelPriceHistory godoc
// @summary Delete a CarModelPriceHistory
// @Description Delete a CarModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-price-histories/{id} [delete]
// @Security BearerAuth
func (h *CarModelPriceHistoryHandler) Delete(c *gin.Context) {
	Delete(c, h.services.DeleteCarModelPriceHistory)
}

// GET

// GetCarModelPriceHistory godoc
// @summary GET a CarModelPriceHistory
// @Description get a CarModelPriceHistory
// @Tags CarModelPriceHistories
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.CarModelPriceHistoryResponse} "CarModelPriceHistory Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-price-histories/{id} [get]
// @Security BearerAuth
func (h *CarModelPriceHistoryHandler) GetById(c *gin.Context) {
	GetById(c, h.services.GetCarModelPriceHistoryById)
}

// Get car-model-price-histories godoc
// @summary Get car-model-price-histories
// @Description Get car-model-price-histories
// @Tags CarModelPriceHistories
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelPriceHistoryResponse]} "CarModelPriceHistory Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-price-histories/get-by-filter [post]
// @Security BearerAuth
func (h *CarModelPriceHistoryHandler) GetByFilter(c *gin.Context) {

	GetByFilter(c, h.services.GetByFilter)
}
