package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Zyprush18/fullstack-template/backend/src/helper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type Midleware struct {
	rdb *redis.Client
	jwtKey string
}

func NewMiddleware(r *redis.Client, j string) Midleware {
	return Midleware{rdb: r, jwtKey: j}
}

func (m *Midleware) MiddlewareAuth() gin.HandlerFunc {
	return  func(ctx *gin.Context) {
		claimtoken := ctx.GetHeader("Authorization")
		// fmt.Println(claimtoken)
		if !strings.HasPrefix(claimtoken, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token Is Missing",
			})
			ctx.Abort()
			return 
		}

		// cek token jwt
		token := strings.Split(claimtoken, " ")
		claim, err := helper.DecryptJwtToken(token[1],m.jwtKey)
		if err != nil {
			if errors.Is(err, jwt.ErrInvalidKey) || errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrSignatureInvalid) {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthorized",
				})
				ctx.Abort()
				return 
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Something Went Wrong",
			})

			ctx.Abort()
			return 
		}

		// cek di redis
		val, err:= m.rdb.Get(context.Background(), claim.Subject).Result()
		fmt.Println()
		if !strings.EqualFold(claim.Role, "admin")|| val != token[1] || err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})

			ctx.Abort()
			return 
		}

		ctx.Set("mail", claim.Subject)


		ctx.Next()
	}
}