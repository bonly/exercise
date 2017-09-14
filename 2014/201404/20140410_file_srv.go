package main

import (
	"log"
	"net/http"
	"flag"
)

func main(){
	var srv *string = flag.String("s", ":9997", "service address for Listen");
	var dir *string = flag.String("d", ".", "file dir");

	flag.Parse();

	http.Handle("/", http.FileServer(http.Dir(*dir)));
	//http.StripPrefix()
	
	err := http.ListenAndServe(*srv, nil)
	if err != nil {
		log.Fatal(err);
	}
}
