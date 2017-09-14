package main

import (
    "log"
    "net/http"
    "flag"
    "fmt"
    // "io/ioutil"
)

func main(){
    var srv *string = flag.String("s", ":9997", "service address for Listen");
    var dir *string = flag.String("d", ".", "file dir");

    flag.Parse();

    http.HandleFunc("/view", view);
    http.Handle("/", http.FileServer(http.Dir(*dir)));
    //http.StripPrefix()
    
    err := http.ListenAndServe(*srv, nil)
    if err != nil {
        log.Fatal(err);
    }
}

func view(rw http.ResponseWriter, qry *http.Request){
    log.Printf("=========== %s ===========\n", "Begin View Request");
    defer func(){
        log.Printf("=========== %s ===========\n", "End View Request");
    }();

    fmt.Println(qry.URL.RawQuery);

    rw.Header().Set("Access-Control-Allow-Origin", "*");
    // rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE");
    // rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token");
    // rw.Header().Set("Access-Control-Allow-Credentials", "true");
    http.Redirect(rw, qry, "/?" + qry.URL.RawQuery, 301);
}

