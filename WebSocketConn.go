package conns

import (
	"fmt"
	"golang.org/x/net/websocket"
)

// WebSocket
type WebSocketConn struct {
	Conn *websocket.Conn
}

func (webSock *WebSocketConn) Send(msg string) error {
	if len(msg) < 1 {
		return nil
	}
	return websocket.Message.Send(webSock.Conn, msg)
}

func (webSock *WebSocketConn) Read() (string, error) {
	var err error
	var str string
	err = websocket.Message.Receive(webSock.Conn, &str)
	return str, err
}

func (webSock *WebSocketConn) Close() error {
	return webSock.Conn.Close()
}

func (webSock *WebSocketConn) GetIPInfo() string {
	r := webSock.Conn.Request()
	return fmt.Sprint(webSock.Conn.RemoteAddr().Network(), "|", r.RemoteAddr)
}
