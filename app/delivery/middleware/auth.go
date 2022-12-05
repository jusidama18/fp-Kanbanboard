package middleware

import (
	"net/http"

	"Kanbanboard/app/helper"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerToken := ctx.GetHeader("Authorization")
		verifyToken, err := helper.VerifyToken(headerToken)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.Set("user", verifyToken)
		ctx.Next()
	}
}

func Authorization(roles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := getToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "you don't have access",
			})
			return
		}

		userRole := token["role"]
		for _, role := range roles {
			if role == userRole {
				ctx.Set("userID", token["id"])
				ctx.Next()
				return
			}
		}

		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "you don't have access",
		})
	}
}

func getToken(ctx *gin.Context) (jwt.MapClaims, error) {
	headerToken := ctx.GetHeader("Authorization")
	token, err := helper.VerifyToken(headerToken)
	if err != nil {
		return nil, err
	}

	return token, nil
}
