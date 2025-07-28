package utils

import 	"github.com/gorilla/websocket"

func WriteWebSocketMessage(conn *websocket.Conn, message string) {
	conn.WriteMessage(websocket.TextMessage, []byte(message))	
}