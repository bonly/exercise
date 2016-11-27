package main

/*
#cgo CPPFLAGS: -g
#cgo LDFLAGS: -lbonly -lthostmduserapi -lthosttraderapi -L.
#include "acc.h"
*/
import "C"
import "fmt"

func main(){
	fmt.Println("begin");
	C.cnt();
	fmt.Println("end");
}