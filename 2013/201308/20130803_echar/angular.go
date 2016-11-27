package main

import (
"fmt"
"net/http"
"syscall"
"os"
"os/signal"
"log"
)

func main(){
   sigs := make(chan os.Signal, 1);
   signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM);
   
   go srv();
   
   <-sigs;
   fmt.Println("\n recv end sigs\n");
}

func srv(){
    http.Handle("/",http.FileServer(http.Dir(".")));
    err := http.ListenAndServe(":8888", nil);
    if err != nil{
        log.Fatal(err);
    }
}

