package client

import (
	"context"
	"log"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

// WebSocketClient return websocket client connection
type WebSocketClient struct {
	configStr string
	sendBuf   chan []byte
	ctx       context.Context
	ctxCancel context.CancelFunc

	mu     sync.RWMutex
	wsconn *websocket.Conn
}

// function to connect with the websocket client
func (conn *WebSocketClient) Connect() *websocket.Conn {
	origin := "http://localhost/"
	url := "ws://localhost:3000/ws"
	conn.mu.Lock()
	defer conn.mu.Unlock()
	if conn.wsconn != nil {
		return conn.wsconn
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		select {
		case <-conn.ctx.Done():
			return nil
		default:
			ws, err := websocket.Dial(url, "", origin)
			if err != nil {
				log.Fatal("Cannot connect to websocket client : ", err)
				continue
			}
			conn.wsconn = ws
			return conn.wsconn
		}
	}
}
