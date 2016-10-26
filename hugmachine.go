package main

import (
	"net/http"

	"golang.org/x/net/websocket"
	"github.com/Jack47/hugger"
)

func handler(ws *websocket.Conn) {
	ws.Write([]byte(hugger.Hug()))
}
func main() {
	http.Handle("/heartbreaker", websocket.Handler(handler))
	http.ListenAndServe(":3000", nil)
}
