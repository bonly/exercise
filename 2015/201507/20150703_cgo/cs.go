package main

/*
#include <stdlib.h>
typedef void (*Callback) (unsigned int sn, char *buf);
extern Callback fn;
*/
import "C"

import (
	"fmt"
)

//export Tabc
func Tabc(){
	fmt.Printf("pok\n");
}

func main(){}

