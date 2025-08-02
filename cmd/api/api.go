package api

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/Rus203/shop/config"
	"github.com/Rus203/shop/constants"
	"github.com/Rus203/shop/handler"
	"github.com/Rus203/shop/logger"
	"github.com/Rus203/shop/routes"
	"github.com/Rus203/shop/service"
)

type APIServer struct {
	// db *sql.DB	// todo: implement it later
}



func (as *APIServer) Run() {
	app := gin.Default()

	app.Use(gin.Recovery())	// todo: add custoom error handling

	messagePublisher := services.NewMessagePublisher()
	webSocketHandler := handlers.NewWebSocketHandler()
	messageProcessor := services.NewMessageProcessor(messagePublisher, webSocketHandler.GetConnectionMap())

	messageConsumer := services.NewMessageConsumer()

	go func() {

		
		err := messageConsumer.ConsumeEventProcess(constants.KITCHEN_ORDER_QUEUE, messageProcessor)

		if err != nil {
			logger.Log(fmt.Sprintf("failed to consume events: %v", err))
		}
	}()

	routes.RegisterRouters(app, messagePublisher, webSocketHandler)
	port :=  fmt.Sprintf(":%d", configs.Env.Port) 
	logger.Log(fmt.Sprintf("Pizza shop started successfully on port : %s", port))
	app.Run(port)
}

func NewApiServer() *APIServer {
	return &APIServer{}
}