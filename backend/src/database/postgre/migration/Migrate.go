package migration

import (
	"fmt"
	"log"

	"github.com/Zyprush18/fullstack-template/backend/src/database/postgre/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate(host, port, dbname, username, pass string) (DB *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, username, pass, dbname, port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	if err := DB.AutoMigrate(&entity.Role{}, &entity.User{});err != nil {
		return nil, err
	}
	return DB,nil
}
