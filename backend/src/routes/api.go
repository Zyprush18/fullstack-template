package routes

import (
	"log"

	"github.com/Zyprush18/fullstack-template/backend/src/config"
	"github.com/Zyprush18/fullstack-template/backend/src/database/migration"
	"github.com/Zyprush18/fullstack-template/backend/src/handler/authhandler"
	"github.com/Zyprush18/fullstack-template/backend/src/handler/rolehandler"
	"github.com/Zyprush18/fullstack-template/backend/src/repository/authrepo"
	"github.com/Zyprush18/fullstack-template/backend/src/repository/rolerepo"
	"github.com/Zyprush18/fullstack-template/backend/src/service/authservice"
	"github.com/Zyprush18/fullstack-template/backend/src/service/roleservice"
	"github.com/gin-gonic/gin"
)

func NewRouteApi(r *gin.Engine)  {
	c :=config.NewConfig()
	initdb, err:= migration.Migrate(c.Host,c.Port,c.DBName,c.Username,c.Password)
	if err != nil {
		log.Fatalln(err.Error())
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(500, gin.H{
			"message": "pong",
		})
	})

	// authentication
	repoauth := authrepo.NewConnectDB(initdb)
	serviceauth := authservice.NewConnectRepo(&repoauth)
	handlerauth := authhandler.NewConnectService(serviceauth)

	r.POST("/api/register", handlerauth.Registration)

	// role
	reporole := rolerepo.NewConnectDB(initdb)
	servicerole := roleservice.NewConnectRepo(&reporole)
	handlerrole := rolehandler.NewConnectService(servicerole)


	r.GET("/api/role", handlerrole.GetAll)
	r.Run()
}