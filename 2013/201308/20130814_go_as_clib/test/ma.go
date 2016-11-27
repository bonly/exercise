package main

/*
#cgo LDFLAGS: -l 20130814_go_as_clib -L.
*/
import "C"

import "fmt"

func main(){
  fmt.Println(Foo());
}

