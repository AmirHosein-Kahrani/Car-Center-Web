package handlers

import (
	"net/http"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/helper"
	"github.com/gin-gonic/gin"
)

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
}

func (h TestHandler) Users(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "Users",
	})
}

// UserById godocs
// @Summary UserById
// @Description UserById
// @Tags Test
// @Param id path int true "user id"
// @Accept json
// @Produce json
// @Success 200 {object} helper.BaseHttpResponse "Succuss"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/test/user/{id} [get]
func (h TestHandler) UserById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"result": "UserById",
		"id":     id,
	})
}

func (h TestHandler) UserByUsername(c *gin.Context) {
	username := c.Param("username")
	c.JSON(http.StatusOK, gin.H{
		"result":   "UserByUsername",
		"username": username,
	})
}

func (h TestHandler) Accounts(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"result": "Accounts",
		"id":     id,
	})
}

func (h TestHandler) AddUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "AddUser",
	})
}

// S23 CH3
type header struct {
	ApplicationType string
	Browser         string
}

func (h *TestHandler) HeaderBinder1(c *gin.Context) {
	userId := c.GetHeader("UserId")

	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"userId": userId,
	})
}

func (h *TestHandler) HeaderBinder2(c *gin.Context) {
	obj := header{}
	err := c.BindHeader(&obj)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "HeaderBinder1",
			"err":    err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"Header": obj,
	})
}

func (h *TestHandler) QueryBinder1(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{
		"result": "QueryBinder1",
		"name":   name,
		"id":     id,
	})
}

func (h *TestHandler) QueryBinder2(c *gin.Context) {
	ids := c.QueryArray("id")
	names := c.QueryArray("name")

	c.JSON(http.StatusOK, gin.H{
		"result": "QueryBinder2",
		"names":  names,
		"ids":    ids,
	})
}

func (h *TestHandler) UriBinder(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")

	c.JSON(http.StatusOK, gin.H{
		"result": "UriBinder",
		"name":   name,
		"id":     id,
	})
}

type PersonData struct {
	FirstName    string `json:"first_name" binding:"required,alpha,min=4,max=10"`
	LastName     string `json:"last_name" binding:"required,alpha,min=6,max=20"`
	MobileNumber string `json:"mobile_number" binding:"required,mobile"`
}

// BodyBinder godocs
// @summery BodyBinder
// @Description BodyBinder
// @Tags Test
// @Param person body PersonData true "person data"
// @Accept json
// @Produce json
// @Succuss 200 {object} helper.BaseHttpResponse "Succuss"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/test/binder/body [post]
// @Security AuthBearer
func (h *TestHandler) BodyBinder(c *gin.Context) {
	p := PersonData{}
	// c.BindJSON(&p)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"result": "BodyBinder",
		"Person": p,
	}, true, 0))
}

func (h *TestHandler) FormBinder(c *gin.Context) {
	p := PersonData{}
	// c.BindJSON(&p)
	c.Bind(&p)

	c.JSON(http.StatusOK, gin.H{
		"result": "FormBinder",
		"Person": p,
	})
}

func (h *TestHandler) FileBinder(c *gin.Context) {
	file, _ := c.FormFile("myfile")
	err := c.SaveUploadedFile(file, "file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "FileBinder",
		"file":   file.Filename,
	})
}

type Password struct {
	Pass string `json:"password" binding:"valid_pass"`
}

func (h *TestHandler) PassHandler(c *gin.Context) {

	pass := Password{}
	// c.BindJSON(&p)
	err := c.ShouldBindJSON(&pass)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error validate": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Password",
		"Pass":   pass,
	})
}
