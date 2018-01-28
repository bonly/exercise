package main

import (
      "fmt"

      "github.com/gopherjs/gopherjs/js"
)

func main() {
      ws := js.Global.Get("WebSocket").New("wss://echo.websocket.org")

      ws.Call("addEventListener", "open", func(e *js.Object) {
              fmt.Println("Connection open ...")
              ws.Call("send", "Hello WebSockets!")
      })

      ws.Call("addEventListener", "message", func(e *js.Object) {
              fmt.Println("Received Message: " + e.Get("data").String())
              ws.Call("close")
      })

      ws.Call("addEventListener", "close", func(e *js.Object) {
              fmt.Println("Connection closed.")
      })
}