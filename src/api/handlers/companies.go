package handlers

import (
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	_ "github.com/AmirHosein-Kahrani/Car-Center-Web/api/helper"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/services"
	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	services *services.CompanyService
}

func NewCompanyHandler(cfg *config.Config) *CompanyHandler {
	return &CompanyHandler{
		services: services.NewCompanyService(cfg),
	}
}

//

// CreateCompany godoc
// @summary create a Company
// @Description create a Company
// @Tags Companies
// @Accept json
// @Produces json
// @Param Request body dto.CreateCompanyRequest true "create a Company"
// @Success 201 {object} helper.BaseHttpResponse{Result=dto.CompanyResponse} "Company Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/companies/ [post]
// @Security BearerAuth
func (h *CompanyHandler) Create(c *gin.Context) {
	Create(c, h.services.CreateCompany)

}

// UpdateCompany godoc
// @summary Update a Company
// @Description Update a Company
// @Tags Companies
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Param Request body dto.UpdateCompanyRequest true "Update a Company"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.CompanyResponse} "Company Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/companies/{id} [put]
// @Security BearerAuth
func (h *CompanyHandler) Update(c *gin.Context) {
	Update(c, h.services.UpdateCompany)
}

// DeleteCompany godoc
// @summary Delete a Company
// @Description Delete a Company
// @Tags Companies
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/companies/{id} [delete]
// @Security BearerAuth
func (h *CompanyHandler) Delete(c *gin.Context) {
	Delete(c, h.services.DeleteCompany)
}

// GET

// GetCompany godoc
// @summary GET a Company
// @Description get a Company
// @Tags Companies
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.CompanyResponse} "Company Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/companies/{id} [get]
// @Security BearerAuth
func (h *CompanyHandler) GetById(c *gin.Context) {
	GetById(c, h.services.GetCompanyById)
}

// Get companies godoc
// @summary Get companies
// @Description Get companies
// @Tags Companies
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CompanyResponse]} "Company Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/companies/get-by-filter [post]
// @Security BearerAuth
func (h *CompanyHandler) GetByFilter(c *gin.Context) {

	GetByFilter(c, h.services.GetByFilter)
}
