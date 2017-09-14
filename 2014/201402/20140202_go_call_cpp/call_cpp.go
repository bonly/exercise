package main
// #cgo LDFLAGS: -L . -lc_test -lstdc++
// #cgo CFLAGS: -I ./
// #include "c.h"
import "C"

func main(){
    C.test();
}