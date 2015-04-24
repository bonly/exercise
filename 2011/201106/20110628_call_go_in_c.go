package main

/*
#ifndef PRINTS_HEAD
#define PRINTS_HEAD
void Acfun();
#endif

#ifndef MYCPP
#define MYCPP
#include <stdio.h>
#include <stdlib.h>
extern void Agfun();
void Acfun(){
   printf("Acfun");
   //Agfun();
}
#endif
//@note import "C",
*/
import "C"
import "fmt"

//Acfun //export Agfun  
func Agfun(){
  fmt.Println("in Agfun");
}

func main(){
  Agfun();
  C.Acfun(); //,ifdef,cgo
}
