package main

import(
_ "net/http/pprof" 
"log"
"net/http"
"flag"
"time"
)

var t = flag.Int("t", 5, "sleep time");

func worker(c <-chan int) {
    var i int

    for {
        i += <-c
    }
}

func wrapper() {
    c := make(chan int)
    defer close(c) // fix here

    go worker(c)

    for i := 0; i < 0xff; i++ {
        c <- i
    }
}

func main() {
    flag.Parse();
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil)) 
    }();

    for {
        time.Sleep((time.Duration)(*t)*time.Second);
        wrapper()
    }
}