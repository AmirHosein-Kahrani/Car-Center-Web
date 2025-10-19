package models

type UserRole struct {
	BaseModel
	UserId int
	RoleId int
	User   User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
	Role   Role `gorm:"foreignKey:RoleId;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION"`
}
