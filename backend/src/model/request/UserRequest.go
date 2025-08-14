package request

type User struct {
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	RoleId uint `json:"role_id" binding:"required"`
}