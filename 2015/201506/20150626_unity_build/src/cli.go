package main

import (
	"log"
	"net/rpc"
	"shared"
	"time"

	"github.com/gopherjs/websocket"
	"honnef.co/go/js/dom"
)

var connectDelay = time.Second
var connectDelayMax = time.Second * 30

var document = dom.GetWindow().Document().(dom.HTMLDocument)

func main() {
	document.AddEventListener("DOMContentLoaded", false, func(_ dom.Event) {
		go ready()
	})
}

func ready() {
	rpc.Register(&RPC{})
	for {
		log.Println("Connecting...")
		connectAndServe()
		log.Println("Disconnected.")
		time.Sleep(connectDelay)
	}
}

func connectAndServe() {
	ws, err := websocket.Dial("ws://localhost:5454/ws-rpc")
	if err != nil {
		log.Println(err.Error())
		connectDelay *= 2
		if connectDelay > connectDelayMax {
			connectDelay = connectDelayMax
		}
		return
	}
	defer ws.Close()
	connectDelay = time.Second
	rpc.ServeConn(ws)
}

type RPC struct{}

func (this RPC) Message(input *shared.ChatMessage, output *shared.ChatMessage) error {
	log.Println(input)
	return nil

}
