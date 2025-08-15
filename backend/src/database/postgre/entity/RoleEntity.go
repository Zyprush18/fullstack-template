package entity


type Role struct {
	IdRole uint `json:"id_role" gorm:"primaryKey;autoIncrement"`
	Role string `json:"role" gorm:"type:varchar(255);not null"`
	User []User `gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Model
}