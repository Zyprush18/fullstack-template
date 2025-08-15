package entity

type User struct {
	IdUser uint `json:"id_user" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"type:varchar(255);not null"`
	Email string `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
	RoleId uint `json:"role_id"`
	Roles Role  `gorm:"foreignKey:RoleId"`

	Model
}