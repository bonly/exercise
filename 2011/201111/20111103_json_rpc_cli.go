package main

import (
    "fmt"
    "log"
    "net/rpc/jsonrpc"
    "os"
)

type Args struct {
    A, B int
}

type Quotient struct {
    Quo, Rem int
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: ", os.Args[0], "server:port")
        log.Fatal(1)
    }
    service := os.Args[1]

    client, err := jsonrpc.Dial("tcp", service)
    if err != nil {
        log.Fatal("dialing:", err)
    }
    // Synchronous call
    //args := Args{17, 8}
    var reply int
    err = client.Call("Arith.ForTet", nil, &reply); //没有参数
    if err != nil {
        log.Fatal("arith error:", err)
    }
    //fmt.Printf("Register: %d*%d=%d\n", args.A, args.B, reply)
    fmt.Printf("Register return %d\n", reply)

}
