package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/Rus203/shop/handler"

)
func RegisterAppRoutes(ctx *gin.RouterGroup) {
	appHandler := handlers.NewAppHandler()

	ctx.GET("/ping", appHandler.HealCheck)
}