package main

import (
	// "bytes"
	"fmt"
	// "io"
	// "time"

	"github.com/gopherjs/gopherjs/js"
	// "github.com/gopherjs/websocket"
	"github.com/gopherjs/websocket/websocketjs"
	"github.com/oskca/gopherjs-vue"
)

type Model struct{
	*js.Object;
};

var ws *websocketjs.WebSocket;

func init(){
   fmt.Printf("init func\n");
}

func Connect(){
	wsBaseURL := getWSBaseURL();
	var err error;
	
	ws, err = websocketjs.New(wsBaseURL + "bus");
	fmt.Printf("get: %s\n", wsBaseURL + "bus" );
	if err != nil{
		fmt.Printf("create: %#v\n", err);
		return;
	}	
}

func (this *Model) Cnt(){
	Connect();
	ws.AddEventListener("open", false, func(ev *js.Object) {
		fmt.Printf( "WebSocket opened\n");
	})

	ws.AddEventListener("close", false, func(ev *js.Object) {
		fmt.Printf("websocket close\n");
	});

    ws.AddEventListener("message", false, func(ev *js.Object){
		fmt.Printf("get a data\n");
	  });	
}

func (this *Model) Discnt(){
	ws.Close();
}

func (this *Model) Send(){
	err := ws.Send("OK");
	if err != nil{
		fmt.Printf("send failed: %#v\n", err);
	}
}

func getWSBaseURL() string {
	document := js.Global.Get("window").Get("document");
	location := document.Get("location");

	wsProtocol := "ws";
	if location.Get("protocol").String() == "https:" {
		wsProtocol = "wss";
	}

	return fmt.Sprintf("%s://%s:%s/ws/", wsProtocol, location.Get("hostname"), location.Get("port"));
}

func main() {
	m := &Model{
		Object: js.Global.Get("Object").New(),
	};

	vue.New("#app", m);

}
