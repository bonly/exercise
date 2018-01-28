package main

/*
#include <stdlib.h> // for C.free
typedef void (*Callback)(unsigned int sn, char *buffer);
extern Callback fn; //extern的东西需要把实体放到另一个文件中实现
extern void bridge_callback(unsigned int sn, char *buffer);
*/
import "C"

import (
	"fmt"
	"C"  //用c-shared必须要有这个
	"unsafe"
)

var sock chan string;

func main(){}

var quit = make(chan int);

//export Init 
func Init()(int){
	fmt.Printf("init go\n");
	sock = make (chan string);
	return 0;
}

//export Run
func Run()(int){
	for {
		select {
			case str := <- sock: {
			  fmt.Printf("get str %s\n", str);
			  //call cs
			  cstr := C.CString("call back from go");
			  defer C.free(unsafe.Pointer(cstr));
			  C.bridge_callback(C.uint(18), cstr);
			  fmt.Printf("回调CS成功\n");
			  break;
			}
			case <- quit:
			  fmt.Printf("get stop sign\n");
			  return 0;
		}
	}
	return 0;
}

//export Put_pack
func Put_pack()(int){
	fmt.Printf("put pack\n");
	sock <- "this is a test";
	quit <- 0;
	return 0;
}

//export SetCallBack
func SetCallBack(pt C.Callback){
	C.fn = pt;
}

/*
go build -buildmode=c-shared -o libtechappen.so dll.go dll_c.go
*/