package main

import (
	"flag"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/tomasen/realip"
)

var serverAddr = flag.String("a", "localhost:8000", "The address for the server ot listen on")
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func echoIP(w http.ResponseWriter, r *http.Request) {
	// Upgrade the client's connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}

	// Don't do anything else — keep the connection allive

	// Echo the client's IP address
	ip := realip.FromRequest(r)
	conn.WriteMessage(websocket.TextMessage, []byte(ip))
}

func main() {
	// Parse the flags
	flag.Parse()

	// Start the server
	http.ListenAndServe(*serverAddr, http.HandlerFunc(echoIP))
}
