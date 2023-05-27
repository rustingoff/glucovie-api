package middleware

import (
	jwtoken "glucovie/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	token := ctx.Request.Header.Get("authorization")
	if token == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	jwt := strings.Split(token, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	_, err := jwtoken.ParseToken(jwt[1])
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	details, err := jwtoken.ExtractTokenMetadata(ctx.Request)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("user_id", details.UserID)
	ctx.Next()
}
