package utils

import (
	"github.com/gin-gonic/gin"
)

func WriteJSON(ctx *gin.Context, body any, status int) {
	ctx.JSON(status, body)
}

func WriteErrorJSON(ctx *gin.Context, status int, err error) {
	WriteJSON(ctx, map[string]string { "error": err.Error() }, status)
}

func ParseJSON(ctx *gin.Context, payload any) error {
	return ctx.ShouldBindBodyWithJSON(payload)
}
