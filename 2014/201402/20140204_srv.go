package main

import (
	"log"
	"os"
	"net/http"
)

func main(){
	if len(os.Args) < 2{
		log.Println("useage: ", os.Args[0], " ip:port ", "dir")
		return
	}
	http.Handle("/", http.FileServer(http.Dir(os.Args[2])));
	//http.StripPrefix()
	
	err := http.ListenAndServe(os.Args[1], nil)
	if err != nil {
		log.Fatal(err);
	}
}
