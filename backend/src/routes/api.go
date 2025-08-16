package routes

import (
	"log"

	"github.com/Zyprush18/fullstack-template/backend/src/config"
	"github.com/Zyprush18/fullstack-template/backend/src/database/postgre/migration"
	"github.com/Zyprush18/fullstack-template/backend/src/database/redisdb"
	"github.com/Zyprush18/fullstack-template/backend/src/handler/authhandler"
	"github.com/Zyprush18/fullstack-template/backend/src/handler/rolehandler"
	"github.com/Zyprush18/fullstack-template/backend/src/middleware"
	"github.com/Zyprush18/fullstack-template/backend/src/repository/authrepo"
	"github.com/Zyprush18/fullstack-template/backend/src/repository/rolerepo"
	"github.com/Zyprush18/fullstack-template/backend/src/service/authservice"
	"github.com/Zyprush18/fullstack-template/backend/src/service/roleservice"
	"github.com/gin-gonic/gin"
)

func NewRouteApi(r *gin.Engine)  {
	c :=config.NewConfig()
	// database
	initdb, err:= migration.Migrate(c.Host,c.Port,c.DBName,c.Username,c.Password)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// redis
	initrdb := redisdb.ConnectRedis(c.RHost,c.RPort,c.RUsername,c.RPassword)

	// midldleware
	initmdlwr := middleware.NewMiddleware(initrdb,c.JwtKey)

	// api
	api := r.Group("/api")

	// authentication
	repoauth := authrepo.NewConnectDB(initdb)
	serviceauth := authservice.NewConnectRepo(&repoauth,initrdb,c.JwtKey)
	handlerauth := authhandler.NewConnectService(serviceauth)

	api.POST("/register", handlerauth.Registration)
	api.POST("/login", handlerauth.Login)

	
	// admin
	admin := api.Group("/admin") 
	admin.Use(initmdlwr.MiddlewareAuth())

	// logout
	admin.GET("/logout", handlerauth.Logout)

	// role
	reporole := rolerepo.NewConnectDB(initdb)
	servicerole := roleservice.NewConnectRepo(&reporole)
	handlerrole := rolehandler.NewConnectService(servicerole)


	admin.GET("/role", handlerrole.GetAll)
	r.Run()
}