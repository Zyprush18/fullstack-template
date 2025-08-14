package entity

import "gorm.io/gorm"

type Role struct {
	Role string `json:"role" gorm:"type:varchar(255);not null"`
	User []User `gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	gorm.Model
}