package main

/*
#cgo CPPFLAGS: -g -lstdc++
#include "20110704_clog.hpp"

*/
import "C"

import "reflect"
import "fmt"

func main(){
  mylog := C.get();
  
  fmt.Println(reflect.TypeOf(&mylog));
  
}
