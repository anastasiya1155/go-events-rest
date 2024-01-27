package middleware

import (
	"net/http"

	"github.com/anastasiya1155/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func CheckAuth(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	context.Set("userId", userId)
	context.Next()
}
