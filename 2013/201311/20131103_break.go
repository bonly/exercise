package main 

import (
"fmt"
"time"
)
func main() {

L:
    for {
        select {
        case <-time.After(time.Second):
            fmt.Println("hello")
        default:
            break L
        }
    }

    fmt.Println("ending")
}