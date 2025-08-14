package rolerepo

import (
	"github.com/Zyprush18/fullstack-template/backend/src/model/response"
	"gorm.io/gorm"
)

type RolesRepo interface {
	GetAll() ([]response.Role, error)
}
type database struct {
	db *gorm.DB
}

func NewConnectDB(d *gorm.DB) database {
	return database{db: d}
}

func (d *database) GetAll() ([]response.Role, error) {
	var model_role  []response.Role
	if err:= d.db.Table("roles").Find(&model_role).Error;err != nil {
		return nil, err
	}

	return model_role,nil
}