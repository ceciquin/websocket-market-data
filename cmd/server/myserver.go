package main

import (
	"alpaca-market-data-api/pkg/server"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	server := server.NewServer
	http.Handle("/ws", websocket.Handler(server().HandleWS))
	http.Handle("/marketData", websocket.Handler(server().HandleWSMarketData))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
