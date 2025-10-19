package models

type City struct {
	BaseModel
	Name      string `gorm:"size:10;type:string;not null"`
	CountryID int
	Country   Country `gorm:"foreignKey:CountryID"`
}

// type City struct {
// 	BaseModel
// 	Name      string  `gorm:"size:10;not null"`
// 	CountryID uint
// 	Country   Country
// }
