package middlewares

import (
	"chat_me_api/src/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientToken := ctx.Request.Header.Get("token")
		if clientToken == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"code": "NoTokenProvided", "error": "No Authorization Header Provided"})
			ctx.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)

		if err != "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"code": "InvalidToken", "error": err})
			ctx.Abort()
			return
		}

		ctx.Set("email", claims.Email)
		ctx.Set("fullName", claims.FullName)
		ctx.Set("uid", claims.Uid)
		ctx.Set("role", claims.Role)
		ctx.Next()
	}
}
