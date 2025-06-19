package service

import "github.com/gorilla/websocket"

type DanmakuFetcher interface {
	Start()
	Close() error
	AddClientConn(ws *websocket.Conn)
	RemoveClientConn(ws *websocket.Conn)
}
