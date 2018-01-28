package main

/*
#cgo LDFLAGS: -fPIC
*/
import "C"

func main() {
}

//export Hello
func Hello() {
	println("test hello")
}
