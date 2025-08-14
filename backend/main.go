package main

import (
	"github.com/Zyprush18/fullstack-template/backend/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.NewRouteApi(router)
}
