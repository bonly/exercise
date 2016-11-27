package main 

import (
"log"
"fmt"
"net/http"
"golang.org/x/net/websocket"

)

func main(){
	http.Handle("/", http.StripPrefix("/",http.FileServer(http.Dir("."))));
	http.Handle("/Main", websocket.Handler(Main));
	err := http.ListenAndServe(":8080", nil);
	if err != nil{
		panic(err.Error());
	}
}

type Head struct{
	Cmd string;
};

func Main(ws *websocket.Conn){
	log.Println(fmt.Sprintf("Get a connected from %s", "a"));

	for {
		//取消息头
		var head Head;
		err := websocket.JSON.Receive(ws, &head);
		log.Println(fmt.Sprintf("Head: %s", head));
		if err != nil{
			log.Println(err);
			return;
		}

		switch(head.Cmd){

		}
	}
}