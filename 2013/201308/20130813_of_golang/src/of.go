package main

/*
#cgo CPPFLAGS: -g
#cgo LDFLAGS: -l emp -L. 
#include "main.h"
*/
import "C"

import (
"fmt"
)

func main(){
  fmt.Println("Begin");
  C.of_main();
  defer func(){fmt.Println("End");}();
}
