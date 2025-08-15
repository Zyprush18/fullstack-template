package request

type User struct {
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=12"`
	RoleId uint `json:"role_id"`	
}

type Login struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}