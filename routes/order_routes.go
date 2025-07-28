package routes

import  (
	"github.com/gin-gonic/gin"

	"github.com/Rus203/shop/service"
	"github.com/Rus203/shop/handler"
)

func RegisterOrderRoutes(router *gin.RouterGroup, messagePublisher services.IMessagePublisher) {
	orderHandler := handlers.NewOrderHandler(messagePublisher)

	router.POST("create", orderHandler.CreateOrder)
}