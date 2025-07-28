package services

import (
	"sync"

	"github.com/gorilla/websocket"
)

type IWebSocketConnection interface {
	SendMessage(message []byte) error
	ReceiveMessage() ([]byte, error)
	Close() error
}

type WebSocketConnection struct {
	conn  *websocket.Conn
	mutex sync.Mutex
}

func (wc *WebSocketConnection) SendMessage(message []byte) error {
	wc.mutex.Lock()
	defer wc.mutex.Unlock()

	return wc.conn.WriteMessage(0, message)
}

func (wc *WebSocketConnection) ReceiveMessage() ([]byte, error) {
	wc.mutex.Lock()
	defer wc.mutex.Unlock()

	_, message, err := wc.conn.ReadMessage()

	return message, err
}

func (wc *WebSocketConnection) Close() error {
	wc.mutex.Lock()
	defer wc.mutex.Unlock()

	return wc.conn.Close()
}

func NewWebSocketConnection(conn  *websocket.Conn) *WebSocketConnection {
	return &WebSocketConnection{
		conn: conn,
	}
}

