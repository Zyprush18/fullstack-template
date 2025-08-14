package authhandler

import (
	"net/http"

	"github.com/Zyprush18/fullstack-template/backend/src/helper"
	"github.com/Zyprush18/fullstack-template/backend/src/model/request"
	"github.com/Zyprush18/fullstack-template/backend/src/service/authservice"
	"github.com/gin-gonic/gin"
)

type HandlerAuth struct {
	services authservice.AuthService
}

func NewConnectService(s authservice.AuthService) HandlerAuth  {
	return HandlerAuth{services: s}
}

func (h *HandlerAuth) Registration(c *gin.Context) {
	req := new(request.User)
	// cek body request
	if err:= c.ShouldBindJSON(&req);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Request Body Is Missing Or Validation Failed",
			"errors": err.Error(),	
		})
		return
	}

	if err:= h.services.Regist(req);err != nil {
		if helper.IsDuplicate(err) {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Email Is Already Exists",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"Something Went Wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Register Successfully",
	})

}