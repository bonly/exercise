package main

import (
	"fmt"
)

func Add(nums... int) int {
    total := 0
    for _, v := range nums {
        total += v
    }
    return total  
}

func main() {
    fmt.Println("Hello, playground")
    fmt.Println(Add(1, 3, 4, 5,))
}