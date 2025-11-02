package handlers

import (
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/helper"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/services"
	"github.com/gin-gonic/gin"
)

type CarModelCommentHandler struct {
	services *services.CarModelCommentService
}

func NewCarModelCommentHandler(cfg *config.Config) *CarModelCommentHandler {
	return &CarModelCommentHandler{
		services: services.NewCarModelCommentService(cfg),
	}
}

//

// CreateCarModelComment godoc
// @summary create a CarModelComment
// @Description create a CarModelComment
// @Tags CarModelComments
// @Accept json
// @Produces json
// @Param Request body dto.CreateCarModelCommentRequest true "create a CarModelComment"
// @Success 201 {object} helper.BaseHttpResponse{Result=dto.CarModelCommentResponse} "CarModelComment Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-comments/ [post]
// @Security BearerAuth
func (h *CarModelCommentHandler) Create(c *gin.Context) {
	Create(c, h.services.CreateCarModelComment)

}

// UpdateCarModelComment godoc
// @summary Update a CarModelComment
// @Description Update a CarModelComment
// @Tags CarModelComments
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Param Request body dto.UpdateCarModelCommentRequest true "Update a CarModelComment"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.CarModelCommentResponse} "CarModelComment Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-comments/{id} [put]
// @Security BearerAuth
func (h *CarModelCommentHandler) Update(c *gin.Context) {
	Update(c, h.services.UpdateCarModelComment)
}

// DeleteCarModelComment godoc
// @summary Delete a CarModelComment
// @Description Delete a CarModelComment
// @Tags CarModelComments
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-comments/{id} [delete]
// @Security BearerAuth
func (h *CarModelCommentHandler) Delete(c *gin.Context) {
	Delete(c, h.services.DeleteCarModelComment)
}

// GET

// GetCarModelComment godoc
// @summary GET a CarModelComment
// @Description get a CarModelComment
// @Tags CarModelComments
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.CarModelCommentResponse} "CarModelComment Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-comments/{id} [get]
// @Security BearerAuth
func (h *CarModelCommentHandler) GetById(c *gin.Context) {
	GetById(c, h.services.GetCarModelCommentById)
}

// Get CarModelComment godoc
// @summary Get CarModelComment
// @Description Get CarModelComment
// @Tags CarModelComments
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelCommentResponse]} "CarModelComment Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/car-model-comments/get-by-filter [post]
// @Security BearerAuth
func (h *CarModelCommentHandler) GetByFilter(c *gin.Context) {

	GetByFilter(c, h.services.GetByFilter)
}
