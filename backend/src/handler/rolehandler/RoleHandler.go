package rolehandler

import (
	"net/http"

	"github.com/Zyprush18/fullstack-template/backend/src/service/roleservice"
	"github.com/gin-gonic/gin"
)

type HandlerRole struct {
	services roleservice.RoleService
}

func NewConnectService(s roleservice.RoleService) HandlerRole  {
	return HandlerRole{services: s}
}

func (h *HandlerRole) GetAll(c *gin.Context) {
	
	// from service
	resp,err := h.services.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something Went Wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"success",
		"data": resp,
	})
}