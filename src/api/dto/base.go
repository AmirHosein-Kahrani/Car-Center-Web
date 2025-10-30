package dto

import (
	"mime/multipart"
	"time"
)

type CreateUpdateCountryRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=20"`
}

type CountryResponse struct {
	Id        int               `json:"id"`
	Name      string            `json:"name"`
	Cities    []CityResponse    `json:"cities,omitempty"`
	Companies []CompanyResponse `json:"companies,omitempty"`
}

type CreateUpdateCityRequest struct {
	Name      string `json:"name" binding:"required,alpha,min=3,max=20"`
	CountryId int    `json:"countryId,omitempty"`
}

type CityResponse struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Country CountryResponse `json:"country,omitempty"`
}

// File
type FileFromRequest struct {
	File *multipart.FileHeader `json:"file" form:"file" binding:"required" swaggerignore:"true"`
}

type UploadFileRequest struct {
	FileFromRequest
	Description string `json:"description" form:"description" binding:"required"`
}

type CreateFileRequest struct {
	Directory   string `json:"name"`
	Name        string `json:"directory"`
	Description string `json:"description"`
	MediaType   string `json:"mediaType"`
}

type UpdateFileRequest struct {
	Description string `json:"description"`
}

type FileResponse struct {
	Id          int    `json:"id"`
	Directory   string `json:"name"`
	Name        string `json:"directory"`
	Description string `json:"description"`
	MediaType   string `json:"mediaType"`
}

type CreateCompanyRequest struct {
	Name      string `json:"name" binding:"required,alpha,min=3,max=20"`
	CountryId int    `json:"countryId" binding:"required"`
}

type UpdateCompanyRequest struct {
	Name      string `json:"name,omitempty" binding:"alpha,min=3,max=20"`
	CountryId int    `json:"countryId,omitempty"`
}

type CompanyResponse struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Country CountryResponse `json:"country,omitempty"`
}

// Color

type CreateColorRequest struct {
	Name    string `json:"name" binding:"alpha,min=3,max=15"`
	HexCode string `json:"hexCode" binding:"min=7,max=7"`
}

type UpdateColorRequest struct {
	Name    string `json:"name,omitempty" binding:"alpha,min=3,max=15"`
	HexCode string `json:"hexCode,omitempty" binding:"min=7,max=7"`
}

type ColorResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	HexCode string `json:"hexCode"`
}

// Year
type CreatePersianYearRequest struct {
	PersianTitle string    `json:"persianTitle" binding:"required,min=4,max=4"`
	Year         int       `json:"year" binding:"required"`
	StartAt      time.Time `json:"startAt" binding:"required"`
	EndAt        time.Time `json:"endAt" binding:"required"`
}

type UpdatePersianYearRequest struct {
	PersianTitle string    `json:"persianTitle,omitempty" binding:"min=4,max=4"`
	Year         int       `json:"year,omitempty"`
	StartAt      time.Time `json:"startAt,omitempty"`
	EndAt        time.Time `json:"endAt,omitempty"`
}

type PersianYearResponse struct {
	Id           int       `json:"id"`
	PersianTitle string    `json:"persianTitle,omitempty"`
	Year         int       `json:"year,omitempty"`
	StartAt      time.Time `json:"startAt,omitempty"`
	EndAt        time.Time `json:"endAt,omitempty"`
}

type PersianYearWithoutDateResponse struct {
	Id           int    `json:"id"`
	PersianTitle string `json:"persianTitle,omitempty"`
	Year         int    `json:"year,omitempty"`
}
