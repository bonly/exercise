package main

import (
	"log"
	//"html/template"
	"net/http"
	"golang.org/x/net/websocket"
)

func main(){
	http.HandleFunc("/help", help);
	http.Handle("/", http.FileServer(http.Dir(".")));
	//http.StripPrefix()
	http.Handle("/ws", websocket.Handler(App));

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err);
	}
}

func help(w http.ResponseWriter, r *http.Request){
	
}

func App(ws *websocket.Conn){
	log.Println("get a connect!");
}