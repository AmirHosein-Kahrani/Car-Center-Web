package handlers

import (
	"net/http"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/helper"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUsersHandler(cfg *config.Config) *UserHandler {
	service := services.NewUserService(cfg)
	return &UserHandler{service: service}
}

// LoginByUsernameRequest godoc
// @Summary LoginByUsernameRequest
// @Description LoginByUsernameRequest
// @Tags Users
// @Accept json
// @Produce json
// @Param Request body dto.LoginByUsernameRequest true "LoginByUsernameRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/login-by-username [post]
func (h *UserHandler) LoginByUserName(c *gin.Context) {

	req := new(dto.LoginByUsernameRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(
				nil, false, helper.ValidationError, err))
		return
	}
	token, err := h.service.LoginByUserName(req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(
				nil, false, helper.InternatlError, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(token, true, helper.Success))
}

// RegisterByUserName godoc
// @Summary RegisterByUserName
// @Description RegisterByUserName
// @Tags Users
// @Accept json
// @Produce json
// @Param Request body dto.RegisterUserByUsernameRequest true "RegisterUserByUsernameRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/register-by-username [post]
func (h *UserHandler) RegisterByUserName(c *gin.Context) {

	req := new(dto.RegisterUserByUsernameRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(
				nil, false, helper.ValidationError, err))
		return
	}
	err = h.service.RegisterByUserName(req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(
				nil, false, helper.InternatlError, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, helper.Success))
}

// RegisterLoginByMobileRequest godoc
// @Summary RegisterLoginByMobileRequest
// @Description RegisterLoginByMobileRequest
// @Tags Users
// @Accept json
// @Produce json
// @Param Request body dto.RegisterLoginByMobileRequest true "RegisterLoginByMobileRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/login-by-mobile [post]
func (h *UserHandler) RegisterLoginByMobileRequest(c *gin.Context) {
	req := new(dto.RegisterLoginByMobileRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(
				nil, false, helper.ValidationError, err))
		return
	}
	token, err := h.service.RegisterLoginByMobileNumber(req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(
				nil, false, helper.InternatlError, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(token, true, helper.Success))
}

// SendOtp godoc
// @Summary Send Otp To User
// @Description Send Otp To User
// @Tags Users
// @Accept json
// @Produce json
// @Param Request body dto.GetOtpRequest true "GetOtpRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/send-otp [post]
func (h *UserHandler) SendOtp(c *gin.Context) {
	req := new(dto.GetOtpRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(
				nil, false, helper.ValidationError, err))
		return
	}
	err = h.service.SetOtp(req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(
				nil, false, helper.InternatlError, err))
		return
	}
	// *Call internal SMS   Service
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, helper.Success))
}
