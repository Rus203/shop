package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8100")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true" )
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Accept")
	ctx.Writer.Header().Set("Access-Control-Allow-Meyhods", "Get,Post")

	if ctx.Request.Method == "Options" {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}

	ctx.Next()
}