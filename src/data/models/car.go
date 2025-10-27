package models

import "time"

type Gearbox struct {
	BaseModel
	Name      string `gorm:"size:15;type:string;not null;unique"`
	CarModels []CarModel
}

type CarType struct {
	BaseModel
	Name      string `gorm:"size:15;type:string;not null;unique"`
	CarModels []CarModel
}

type Company struct {
	BaseModel
	Name      string  `gorm:"size:15;type:string;not null;unique"`
	Country   Country `gorm:"foreignKey:CountryId"`
	CountryId int
	CarModels []CarModel
}

type CarModel struct {
	BaseModel
	Name               string  `gorm:"size:15;type:string;not null;unique"`
	Company            Company `gorm:"foreignKey:CompanyId"`
	CompanyId          int
	CarType            CarType `gorm:"foreignKey:CarTypeId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	CarTypId           int
	Gearbox            Gearbox `gorm:"foreignKey:GearBoxId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	GearBoxId          int
	CarModelColors     []CarModelColor
	CarModelYears      []CarModelYear
	CarModelProperties []CarModelProperty
	CarModelImages     []CarModelImage
	CarModelComments   []CarModelComment
}
type CarModelColor struct {
	BaseModel
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	CarModelId int
	Color      Color `gorm:"foreignKey:ColorId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	ColorId    Color
}

type CarModelYear struct {
	BaseModel
	CarModel               CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	CarModelId             int
	PersianYear            PersianYear `gorm:"foreignKey:PersianYearId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	PersianYearId          int
	CarModelPriceHistories []CarModelPriceHistory
}

type CarModelImage struct {
	BaseModel
	CarModel    CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	CarModelId  int
	Image       File `gorm:"foreignKey:ImageId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	ImageId     Color
	IsMainImage bool
}

type CarModelPriceHistory struct {
	BaseModel
	CarModelYear   CarModelYear `gorm:"foreignKey:CarModelYearId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	CarModelYearId int
	Price          float32   `gorm:"type:decimal(10,2);not null"`
	PriceAt        time.Time `gorm:"type:TIMESTAMP with time zone;not null"`
}

type CarModelProperty struct {
	BaseModel
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	CarModelId int
	Properties Property `gorm:"foreignKey:PropertyId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	PropertyId int
	Value      string `gorm:"size:100;type:string;not null"`
}

type CarModelComment struct {
	BaseModel
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	CarModelId int
	User       User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	UserId     int
	Message    string `gorm:"size:500;type:string;not null"`
}
