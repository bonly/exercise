package main

import (
	"log"
	//"html/template"
	"net/http"
	"net"
)

func main(){
	http.HandleFunc("/help", help);
	http.Handle("/", http.FileServer(http.Dir(".")));
	//http.StripPrefix()
	
	fd, err := net.Listen("unix", "/tmp/tp");
	if err != nil {
		log.Fatal(err);
	}
	err = http.Serve(fd, nil);
	if err != nil {
		log.Fatal(err);
	}
}

func help(w http.ResponseWriter, r *http.Request){
	
}