package gocallback

import "fmt"

/*
#include <stdio.h>
extern void ACFunction();
*/
import "C"

//export AGoFunction
func AGoFunction(){
  fmt.Println("AGoFunction()");
}

func Example(){
  C.ACFunction();
}
