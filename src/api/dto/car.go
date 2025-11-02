package dto

import "time"

// CarType
type CreateCarTypeRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=15"`
}

type UpdateCarTypeRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=15"`
}

type CarTypeResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// GearBox
type CreateGearboxRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=15"`
}

type UpdateGearboxRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=15"`
}

type GearboxResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// Car Model Request

type CreateCarModelRequest struct {
	Name      string `json:"name" binding:"required,min=3,max=15"`
	CompanyId int    `json:"companyId" binding:"required"`
	GearboxId int    `json:"gearboxId" binding:"required"`
	CarTypeId int    `json:"carTypeId" binding:"required"`
}

type UpdateCarModelRequest struct {
	Name      string `json:"name,omitempty"`
	CompanyId int    `json:"companyId,omitempty"`
	GearboxId int    `json:"gearboxId,omitempty"`
	CarTypeId int    `json:"carTypeId,omitempty"`
}
type CarModelResponse struct {
	Id                 int                        `json:"id"`
	Name               string                     `json:"name"`
	CarType            CarTypeResponse            `json:"carType"`
	Company            CompanyResponse            `json:"company"`
	Gearbox            GearboxResponse            `json:"gearbox"`
	CarModelColors     []CarModelColorResponse    `json:"carModelColors,omitempty"`
	CarModelYears      []CarModelYearResponse     `json:"carModelyears,omitempty"`
	CarModelImages     []CarModelImageResponse    `json:"carModelImages,omitempty"`
	CarModelProperties []CarModelPropertyResponse `json:"carModelProperties,omitempty"`
	CarModelComments   []CarModelCommentResponse  `json:"carModelComments,omitempty"`
}

// Car Model Color

type CreateCarModelColorRequest struct {
	CarModelId int `json:"carModelId" binding:"required"`
	ColorId    int `json:"colorId" binding:"required"`
}

type UpdateCarModelColorRequest struct {
	CarModelId int `json:"carModelId,omitempty"`
	ColorId    int `json:"colorId,omitempty"`
}

type CarModelColorResponse struct {
	Id    int           `json:"id"`
	Color ColorResponse `json:"color"`
}

// Car Model Year
type CreateCarModelYearRequest struct {
	CarModelId    int `json:"carModelId" binding:"required"`
	PersianYearId int `json:"persianYearId" binding:"required"`
}

type UpdateCarModelYearRequest struct {
	CarModelId    int `json:"carModelId,omitempty"`
	PersianYearId int `json:"persianYearId,omitempty"`
}

type CarModelYearResponse struct {
	Id                     int                            `json:"id"`
	PersianYear            PersianYearWithoutDateResponse `json:"persianYear,omitempty"`
	CarModelId             int                            `json:"carModelId,omitempty"`
	CarModelPriceHistories []CarModelPriceHistoryResponse `json:"carModelPriceHistories,omitempty"`
}

type CreateCarModelPriceHistoryRequest struct {
	CarModelYearId int       `json:"carModelYearId" binding:"required"`
	PriceAt        time.Time `json:"priceAt" binding:"required"`
	Price          float32   `json:"price" binding:"required"`
}

type UpdateCarModelPriceHistoryRequest struct {
	PriceAt time.Time `json:"priceAt,omitempty"`
	Price   float32   `json:"price,omitempty"`
}

type CarModelPriceHistoryResponse struct {
	Id             int       `json:"id"`
	CarModelYearId int       `json:"carModelYearId,omitempty"`
	PriceAt        time.Time `json:"priceAt,omitempty"`
	Price          float64   `json:"price,omitempty"`
}

type CreateCarModelImageRequest struct {
	CarModelId  int  `json:"carModelId" binding:"required"`
	ImageId     int  `json:"imageId" binding:"required"`
	IsMainImage bool `json:"isMainImage"`
}

type UpdateCarModelImageRequest struct {
	IsMainImage bool `json:"isMainImage,omitempty"`
}

type CarModelImageResponse struct {
	Id          int          `json:"id"`
	CarModelId  int          `json:"carModelId,omitempty"`
	Image       FileResponse `json:"image,omitempty"`
	IsMainImage bool         `json:"isMainImage"`
}

type CreateCarModelPropertyRequest struct {
	CarModelId int    `json:"carModelId" binding:"required"`
	ProperyId  int    `json:"propertyId" binding:"required"`
	Value      string `json:"value" binding:"required,max=100"`
}

type UpdateCarModelPropertyRequest struct {
	Value string `json:"value" binding:"required,max=100"`
}

type CarModelPropertyResponse struct {
	Id         int              `json:"id"`
	CarModelId int              `json:"carModelId,omitempty"`
	Property   PropertyResponse `json:"property,omitempty"`
	Value      string           `json:"value"`
}

type CreateCarModelCommentRequest struct {
	CarModelId int    `json:"carModelId" binding:"required"`
	UserId     int    `json:"userId" binding:"required"`
	Message    string `json:"Message" binding:"required,max=500"`
}

type UpdateCarModelCommentRequest struct {
	Message string `json:"Message" binding:"required,max=500"`
}

type CarModelCommentResponse struct {
	Id         int          `json:"id"`
	CarModelId int          `json:"carModelId"`
	User       UserResponse `json:"user"`
	Message    string       `json:"Message" binding:"required,max=500"`
}

type UserResponse struct {
	Id        int    `json:"id"`
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Message   string `json:"Message"`
}
