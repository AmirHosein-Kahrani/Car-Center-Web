package handlers

import (
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/helper"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/services"
	"github.com/gin-gonic/gin"
)

type CarModelImageHandler struct {
	services *services.CarModelImageService
}

func NewCarModelImageHandler(cfg *config.Config) *CarModelImageHandler {
	return &CarModelImageHandler{
		services: services.NewCarModelImageService(cfg),
	}
}

//

// CreateCarModelImage godoc
// @summary create a CarModelImage
// @Description create a CarModelImage
// @Tags CarModelImages
// @Accept json
// @Produces json
// @Param Request body dto.CreateCarModelImageRequest true "create a CarModelImage"
// @Success 201 {object} helper.BaseHttpResponse{Result=dto.CarModelImageResponse} "CarModelImage Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-images/ [post]
// @Security BearerAuth
func (h *CarModelImageHandler) Create(c *gin.Context) {
	Create(c, h.services.CreateCarModelImage)

}

// UpdateCarModelImage godoc
// @summary Update a CarModelImage
// @Description Update a CarModelImage
// @Tags CarModelImages
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Param Request body dto.UpdateCarModelImageRequest true "Update a CarModelImage"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.CarModelImageResponse} "CarModelImage Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-images/{id} [put]
// @Security BearerAuth
func (h *CarModelImageHandler) Update(c *gin.Context) {
	Update(c, h.services.UpdateCarModelImage)
}

// DeleteCarModelImage godoc
// @summary Delete a CarModelImage
// @Description Delete a CarModelImage
// @Tags CarModelImages
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-images/{id} [delete]
// @Security BearerAuth
func (h *CarModelImageHandler) Delete(c *gin.Context) {
	Delete(c, h.services.DeleteCarModelImage)
}

// GET

// GetCarModelImage godoc
// @summary GET a CarModelImage
// @Description get a CarModelImage
// @Tags CarModelImages
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.CarModelImageResponse} "CarModelImage Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-images/{id} [get]
// @Security BearerAuth
func (h *CarModelImageHandler) GetById(c *gin.Context) {
	GetById(c, h.services.GetCarModelImageById)
}

// Get car-model-images godoc
// @summary Get car-model-images
// @Description Get car-model-images
// @Tags CarModelImages
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelImageResponse]} "CarModelImage Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-images/get-by-filter [post]
// @Security BearerAuth
func (h *CarModelImageHandler) GetByFilter(c *gin.Context) {

	GetByFilter(c, h.services.GetByFilter)
}
