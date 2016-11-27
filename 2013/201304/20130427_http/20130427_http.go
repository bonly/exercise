package main

import (
	"log"
	//"html/template"
	"net/http"
)

func main(){
	http.HandleFunc("/help", help);
	http.Handle("/", http.FileServer(http.Dir(".")));
	//http.StripPrefix()
	

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err);
	}
}

func help(w http.ResponseWriter, r *http.Request){
	
}