package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/net/websocket"
)

func main() {
	flag.Parse()

	// Connect to the WebSocket server.
	ws, err := websocket.Dial("ws://localhost:3000/marketData", "", "http://localhost")
	if err != nil {
		fmt.Printf("Failed to connect to WebSocket server: %v\n", err)
		return
	}
	//Receive message from server using Websocket connection
	for {
		var message string
		err := websocket.Message.Receive(ws, &message)
		if err != nil {
			break
		}
		fmt.Println("Received message:", message)
	}

	// Close connection correctly on exit
	sigs := make(chan os.Signal, 1)

	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// The program will wait here until it gets the
	<-sigs
	ws.Close()
	fmt.Println("Goodbye")
}
