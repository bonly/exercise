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

func main(){}

//export Run
func Run()(int){
	fmt.Printf("in go\n");
	return 0;
}

//export Cb
func Cb(){
	// go func(){ //报：condition `refcount' not met
		cstr := C.CString("callback");
		defer C.free(unsafe.Pointer(cstr));
		C.bridge_callback(1, cstr);
	// }();
}

//export SetCallBack
func SetCallBack(pt C.Callback){
	C.fn = pt;
}
/*
go build -buildmode=c-shared -o libtechappen.so dll.go dll_c.go
*/