package routes

import (
	"github.com/Rus203/shop/handler"

	"github.com/gin-gonic/gin"
)

func RegisterWebSocketRoutes(router *gin.RouterGroup, websocketHandler handlers.IWebSocketHandler) {
	router.GET("/", websocketHandler.HandleConnection)
}