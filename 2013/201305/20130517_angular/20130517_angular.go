package main

import (
	"log"
	//"html/template"
	"net/http"
	"fmt"
	"golang.org/x/net/websocket"
	"open"
)

func main(){
	http.HandleFunc("/help", help);
	http.Handle("/", http.FileServer(http.Dir(".")));
	http.Handle("/html", http.HandlerFunc(sendHtml));
	http.Handle("/toolbar.js", http.HandlerFunc(sendJs));
	//http.StripPrefix()
	http.Handle("/ws", websocket.Handler(App));

        go open.Run("http://127.0.0.1:8888");

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err);
	}
}

func help(w http.ResponseWriter, r *http.Request){
}

func sendHtml(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, Index);
}

func sendJs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript");
	fmt.Fprint(w, Toolbar);
}

func App(ws *websocket.Conn){
	log.Println("get a connect!");
	ws.Write([]byte(`{"app":[{"x":"20","y":"20","o":"+","v":"50"},{"x":"3","y":"3","o":"+","v":"8"}]}`));
}
