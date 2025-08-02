package utils

import 	"github.com/gorilla/websocket"

func WriteWebSocketMessage(conn *websocket.Conn, message string) error {
	return  conn.WriteMessage(websocket.TextMessage, []byte(message))	
}