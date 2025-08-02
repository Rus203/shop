package routes

import (
	"github.com/Rus203/shop/service"

	"github.com/Rus203/shop/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(r *gin.Engine, messagePublisher services.IMessagePublisher, webSockerHandler handlers.IWebSocketHandler) {
	router := r.Group("/")

	wsr := router.Group("/ws") 
	RegisterWebSocketRoutes(wsr, webSockerHandler)

	or := router.Group("/orders")
	RegisterOrderRoutes(or, messagePublisher)

	ar := router.Group("/")
	RegisterAppRoutes(ar)
}
