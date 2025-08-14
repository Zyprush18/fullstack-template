package response

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	RoleId uint `json:"role_id"`
	
}