package main 

import "C"

import (
"fmt"
)

//export Gf
func Gf(){
	fmt.Println("hello from golang");
}

func main(){}
