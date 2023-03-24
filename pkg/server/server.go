package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"golang.org/x/net/websocket"
)

// This structs is to handdle multiple websocket connection
type Server struct {
	conns map[*websocket.Conn]bool
}

// initialize a pointer to a new server connection and a channel to anable concurrency to handle  multiple websocket connections
func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}
func (s *Server) HandleWSMarketData(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client to orderbook feed:", ws.RemoteAddr())

	for i := 0; i < 5; i++ {
		myJsonData, err := json.Marshal(GenerateMarketData())
		if err != nil {
			panic(err)
		}
		fmt.Println(string(myJsonData))

		websocket.Message.Send(ws, myJsonData)
	}

}

// check connection with client and their address
func (s *Server) HandleWS(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client:", ws.RemoteAddr())

	s.conns[ws] = true
	s.readLoop(ws)
}

// private function to read incoming messages from client
func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error:", err)
			continue
		}
		msg := buf[:n]
		fmt.Println(string(msg))
		s.broadcast(msg)
	}
}

// functon that helps to send the message to the client and perform I/O operation concurrently and avoid any task for being block
func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				log.Fatal("Error during broadcast of message: ", err)
			}
		}(ws)
	}
}
