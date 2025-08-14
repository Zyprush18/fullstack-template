package response

import "gorm.io/gorm"

type Role struct {
	Role string `json:"role"`
	gorm.Model
}