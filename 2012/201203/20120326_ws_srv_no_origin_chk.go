package main

import (
 "fmt"
 "net/http"
 "os"
 // "io"
 "code.google.com/p/go.net/websocket"
)

func Echo(ws *websocket.Conn) {
 fmt.Println("Echoing")
 for n := 0; n < 10; n++ {
	 msg := "Hello " + string(n+48)
	 fmt.Println("Sending to client: " + msg)
	 err := websocket.Message.Send(ws, msg)
	 if err != nil {
		 fmt.Println("Can't send")
		 break
	 }
	 var reply string
	 err = websocket.Message.Receive(ws, &reply)
	 if err != nil {
		 fmt.Println("Can't receive")
		 break
	 }
	 fmt.Println("Received back from client: " + reply)
 }
}

func main() {
 // http.HandleFunc("/", websocket.Handler(Echo))
 // go func（）{  //让srv能同时处理http:80
 //    err := http.ListenAndServe(":80", nil)
 //    if err != nil {
 //        panic("ListenAndServe: " + err.Error())
 //    }
 // }
 err := http.ListenAndServe(":80", websocket.Server{websocket.Config{}, nil, Echo}) //忽略origin字段内容.
 checkError(err)
}


func checkError(err error) {
 if err != nil {
	 fmt.Println("Fatal error ", err.Error())
	 os.Exit(1)
 }
}