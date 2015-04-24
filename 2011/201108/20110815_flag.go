package main

import (
    "flag"
    "fmt"
)

func main() {
    flag.Parse()
    for i, v := range flag.Args() {
        fmt.Println(i, v)
    }
}