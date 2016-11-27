package main

import (
	"log"
	"os"
	"net/http"
)

func main(){
	if len(os.Args) < 2{
		log.Println("useage: ", os.Args[0], " ip:port")
		return
	}
	http.Handle("/", http.FileServer(http.Dir(".")));
	//http.StripPrefix()
	
	err := http.ListenAndServe(os.Args[1], nil)
	if err != nil {
		log.Fatal(err);
	}
}
