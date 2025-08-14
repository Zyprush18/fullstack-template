package authrepo

import (
	"github.com/Zyprush18/fullstack-template/backend/src/model/request"
	"gorm.io/gorm"
)

type AuthRepo interface {
	Registration(req *request.User) error
}

type database struct {
	db *gorm.DB	
}

func NewConnectDB(d *gorm.DB) database {
	return database{db: d}
}

func (d *database) Registration(req *request.User) error {
	if err:= d.db.Debug().Table("users").Create(&req).Error;err != nil {
		return err
	}

	return nil
}