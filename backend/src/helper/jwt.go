package helper

import (
	"fmt"
	"time"

	"github.com/Zyprush18/fullstack-template/backend/src/database/postgre/entity"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Role string	`json:"role"`
	jwt.RegisteredClaims
}

func GenerateJwtToken(key string, data *entity.User) (string, error) {
	claims := CustomClaims{
		Role: data.Roles.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ID: fmt.Sprintf("%d",data.IdUser),
			Subject: data.Email,
			Issuer: "http://localhost:8080", // di ganti nanti dengan domain atau url yg dimiliki
			Audience: []string{"my-api-service"},
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24 jam
		},
	}

	jwtclaim := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	
	// return token dan error
	return jwtclaim.SignedString([]byte(key))
}