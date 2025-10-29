package handlers

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/dto"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/api/helper"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/logging"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var Logger = logging.NewLogger(config.GetConfig())

type FileHandler struct {
	services *services.FileService
}

func NewFileHandler(cfg *config.Config) *FileHandler {

	return &FileHandler{
		services: services.NewFileService(cfg),
	}
}

// CreateFile godoc
// @summary create a file
// @Description create a file
// @Tags Files
// @Accept x-www-form-urlencoded
// @Produces json
// @Param file formData dto.UpdateFileRequest true "create a file"
// @Param file formData file true "create a file"
// @Success 201 {object} helper.BaseHttpResponse{Result=dto.FileResponse} "File Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/files/ [post]
// @Security BearerAuth
func (h *FileHandler) Create(c *gin.Context) {
	upload := dto.UploadFileRequest{}
	err := c.ShouldBind(&upload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	req := dto.CreateFileRequest{}
	req.Description = upload.Description
	req.MediaType = upload.File.Header.Get("Content-Type")
	req.Directory = "uploads"
	req.Name, err = saveUploadFile(upload.File, req.Directory)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternatlError, err))
		return
	}

	res, err := h.services.CreateFile(c, &req)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternatlError, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, helper.Success))
}

func saveUploadFile(file *multipart.FileHeader, directory string) (string, error) {

	randFileName := uuid.New()
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return "", err
	}
	fileName := file.Filename
	fileNameArr := strings.Split(fileName, ".")
	fileExt := fileNameArr[len(fileNameArr)-1]
	fileName = fmt.Sprintf("%s.%s", randFileName, fileExt)
	dst := fmt.Sprintf("%s/%s", directory, fileName)
	src, err := file.Open()
	if err != nil {
		return "", err
	}

	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return "", err
	}
	return fileName, nil

}

// --------------------------

// UpdateFile godoc
// @summary Update a file
// @Description Update a file
// @Tags Files
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Param Request body dto.UpdateFileRequest true "Update a File"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.FileResponse} "File Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/files/{id} [put]
// @Security BearerAuth
func (h *FileHandler) Update(c *gin.Context) {
	Update(c, h.services.UpdateFile)
}

// DeleteFile godoc
// @summary Delete a file
// @Description Delete a file
// @Tags Files
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/files/{id} [delete]
// @Security BearerAuth
func (h *FileHandler) Delete(c *gin.Context) {

	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, helper.ValidationError))
		return
	}
	file, err := h.services.GetFileById(c, id)
	if err != nil {
		Logger.Error(logging.IO, logging.RemoveFile, err.Error(), nil)
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, helper.NotFoundError))
		return
	}
	err = os.Remove(fmt.Sprintf("%s/%s", file.Directory, file.Name))

	if err != nil {
		Logger.Error(logging.IO, logging.RemoveFile, err.Error(), nil)
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, helper.InternatlError))
		return
	}
	err = h.services.DeleteFile(c, id)

	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternatlError, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, helper.Success))
}

// GET

// GetFile godoc
// @summary GET a file
// @Description get a file
// @Tags Files
// @Accept json
// @Produces json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{Result=dto.FileResponse} "File Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/files/{id} [get]
// @Security BearerAuth
func (h *FileHandler) GetById(c *gin.Context) {
	GetById(c, h.services.GetFileById)
}

// Get Files godoc
// @summary Get Files
// @Description Get Files
// @Tags Files
// @Accept json
// @Produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.FileResponse]} "File Response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Router /v1/files/get-by-filter [post]
// @Security BearerAuth
func (h *FileHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.services.GetByFilter)
}
