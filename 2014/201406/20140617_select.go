package main

import (
"fmt"
"sync"
)
var wg sync.WaitGroup;

func main() {
    messages := make(chan string)
    signals := make(chan bool)
    
    wg.Add(1);
    go func(){
      defer wg.Done();
      for{
        select {
        case msg := <-messages:
            fmt.Println("received message", msg)
        default:
            // fmt.Println("no message received")
        }
      }
    }();

    msg := "hi"
    go func(){
      defer wg.Done();
      for{
        select {
        case messages <-msg:
            fmt.Println("sent message", msg)
        default:
            // fmt.Println("no message sent")
        }
      }
    }();

    messages<-msg
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
    wg.Wait();
}