package authhandler

import (
	"errors"
	"net/http"

	"github.com/Zyprush18/fullstack-template/backend/src/helper"
	"github.com/Zyprush18/fullstack-template/backend/src/model/request"
	"github.com/Zyprush18/fullstack-template/backend/src/service/authservice"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func (h *HandlerAuth) Login(c *gin.Context) {
	req := new(request.Login)

	// cek body request
	if err:= c.ShouldBindJSON(&req);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Request Body Is Missing Or Validation Failed",
			"errors": err.Error(),	
		})
		return
	}


	tokens,err:= h.services.Login(req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || err.Error() == "invalid password" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid Credential",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"Something Went Wrong",
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"token": tokens,
	})
}

func (h *HandlerAuth) Logout(c *gin.Context) {
	// ambil dari context request
	email := c.MustGet("mail").(string)
	
	if err:= h.services.Logout(email);err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message":"Invalid Credential",
			})
			return
		}


		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"Something Wnet Wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"Success logout",
	})

}