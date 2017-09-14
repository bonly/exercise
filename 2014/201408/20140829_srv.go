package main

import (
	"flag"
	"log"
	"net/http"
)

var port = flag.String("s", ":9997", "server port");
var dir  = flag.String("d", ".", "file dir");

func main() {
	flag.Parse();

	fs := http.StripPrefix("/", http.FileServer(http.Dir(*dir)));

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "private");
		// Needed for local proxy to Kubernetes API server to work.
		w.Header().Set("Access-Control-Allow-Origin", "*");
		w.Header().Set("Access-Control-Allow-Credentials", "true");
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS");
		w.Header().Set("Access-Control-Allow-Headers", "DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,Cache-Control,Content-Type");
		// Disable If-Modified-Since so update-demo isn't broken by 304s
		r.Header.Del("If-Modified-Since");
		fs.ServeHTTP(w, r);
	})

	go log.Fatal(http.ListenAndServe(*port, nil));

	select {};
}