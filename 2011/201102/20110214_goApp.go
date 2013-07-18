package main
/*
#cgo LDFLAGS: -L/home/bonly/src/ -lmyl
#include "20110213_mylib.h"
*/
import "C"
import "fmt"
func main(){
  fmt.Println("1+2=", C.sum(1,2));
}