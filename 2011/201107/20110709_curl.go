package main

import (
    "fmt"
    curl "github.com/andelf/go-curl"
)

func main() {
  //for i:=0; i<10; i++{
    easy := curl.EasyInit()
    defer easy.Cleanup()

    easy.Setopt(curl.OPT_URL, "http://127.0.0.1:9001")

    // make a callback function
    fooTest := func (buf []byte, userdata interface{}) bool {
        println("DEBUG: size=>", len(buf))
        println("DEBUG: content=>", string(buf))
        return true
    }

    easy.Setopt(curl.OPT_WRITEFUNCTION, fooTest)
    easy.Setopt(curl.OPT_POSTFIELDS, "{\"user_id\":134}");

    
    if err := easy.Perform(); err != nil {
        fmt.Printf("ERROR: %v\n", err)
    }
  //}
}

//curl -d "this is body" -u "user:pass" "http://localhost/?ss=ss&qq=11"
//curl -d '{"user_id":33}' "http://127.0.0.1/8000'
