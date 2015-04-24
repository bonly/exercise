package main
/*
#ifndef PRINTS_HEAD
#define PRINTS_HEAD
void Prints(char* str);
#endif

#include <stdio.h>
#include <stdlib.h>
void Prints(char* str){
   printf("%s\n", str);
}
*/
import "C"
import "unsafe"

func main(){
  cs := C.CString("hi world");
  
  C.Prints(cs);
  
  C.free(unsafe.Pointer(cs));
}
