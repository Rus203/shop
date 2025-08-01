package handlers

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/Rus203/shop/logger"
	"github.com/Rus203/shop/service"
	utils "github.com/Rus203/shop/util"
)

type IWebSocketHandler interface {
	HandleConnection(ctx *gin.Context)
	GetConnectionMap() *map[string]services.IWebSocketConnection
}

type WebSocketHandler struct {
	connection *map[string]services.IWebSocketConnection
	mutex      sync.Mutex
	upgrader   websocket.Upgrader
}

func (wh *WebSocketHandler) HandleConnection(ctx *gin.Context) {
	conn, err := wh.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		logger.Log("Failed to upgrade connection")
	}

	defer conn.Close()

	utils.WriteWebSocketMessage(conn, "Started taking order")

	connection := services.NewWebSocketConnection(conn)
	wh.addConnection("pizza", connection)		// todo: implement multiply conevtion storing

	for {
		logger.Log("no message is coming from client")
		time.Sleep(time.Second)
	}
}

func (wh *WebSocketHandler) addConnection(clientId string, connection services.IWebSocketConnection) {
	wh.mutex.Lock()
	defer wh.mutex.Unlock()

	(*wh.connection)[clientId] = connection
}

func (wh *WebSocketHandler) GetConnectionMap() *map[string]services.IWebSocketConnection {
	return wh.connection
}


func NewWebSocketHandler() *WebSocketHandler {
	connection := make(map[string]services.IWebSocketConnection)
	return &WebSocketHandler{
		connection: &connection,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true }, // todo: finish it later
		},
	}
}
