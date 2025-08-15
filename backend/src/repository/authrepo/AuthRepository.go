package authrepo

import (
	"github.com/Zyprush18/fullstack-template/backend/src/database/postgre/entity"
	"github.com/Zyprush18/fullstack-template/backend/src/model/request"
	"gorm.io/gorm"
)

type AuthRepo interface {
	Registration(req *request.User) error
	Login(req *request.Login) (*entity.User, error)
}

type database struct {
	db *gorm.DB	
}

func NewConnectDB(d *gorm.DB) database {
	return database{db: d, }
}

func (d *database) Registration(req *request.User) error {
	if err:= d.db.Table("users").Create(&req).Error;err != nil {
		return err
	}

	return nil
}

func (d *database) Login(req *request.Login) (*entity.User,error) {
	var model_user entity.User
	if err:= d.db.Model(&entity.User{}).Preload("Roles").Where("email = ?", req.Email).First(&model_user).Error;err != nil {
		return nil,err
	}

	return &model_user,nil
}