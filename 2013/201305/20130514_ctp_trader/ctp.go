package main

/*
#cgo CPPFLAGS: -g
#cgo LDFLAGS: -lbonly -lthostmduserapi -lthosttraderapi -L.
#include "trader.h"
*/
import "C"
import "fmt"

func main(){
	fmt.Println("begin");
	C.cnt();
	fmt.Println("end");
}

/*
go build -v -o gs ctp.go
*/