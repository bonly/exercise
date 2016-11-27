package main

import (
        "fmt"
        "encoding/hex"
)

func main(){
    a, _ := hex.DecodeString("0101A8C0") 

    fmt.Printf("%v.%v.%v.%v", a[3], a[2], a[1], a[0])
}
