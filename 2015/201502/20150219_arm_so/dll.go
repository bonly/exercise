package main

/*
#cgo LDFLAGS: -pthread -fPIC 
*/
import "C"

func main(){}

//export Hello
func Hello() *C.char{
    println("test hello");
    return C.CString("ok hehe");
}

