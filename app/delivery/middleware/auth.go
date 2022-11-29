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
		_ = verifyToken

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
		userAuth := ctx.MustGet("user").(jwt.MapClaims)
		userRole := userAuth["role"]
		for _, role := range roles {
			if role == userRole {
				ctx.Next()
				return
			}
		}
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "you don't have access",
		})
	}
}

func getRole(ctx *gin.Context) (any, error) {
	headerToken := ctx.GetHeader("Authorization")
	token, err := helper.VerifyToken(headerToken)
	if err != nil {
		return nil, err
	}

	userRole := token["role"]
	return userRole, nil
}

func AuthorizeAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, err := getRole(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "you don't have access",
			})
			return
		}

		if role != "admin" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "you don't have access",
			})
			return
		}

		ctx.Next()
	}
}

func AuthorizeMember() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, err := getRole(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "you don't have access",
			})
			return
		}

		if role != "member" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "you don't have access",
			})
			return
		}

		ctx.Next()
	}
}
