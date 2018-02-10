package main

/*
#include <stddef.h>
#include <stdlib.h>
typedef void (*GOLOG_PROC) (char *szDscr);
extern GOLOG_PROC g_logcallback;

extern void test();

extern void setcallback(GOLOG_PROC cb);
extern void bridge(char*);
*/
import "C"

import(
	"fmt"
	"unsafe"
)

//export Go_logcallback
func Go_logcallback(dscr *C.char){
	fmt.Println("go log callback")
}

var MyFunc = Go_logcallback;

func main(){
	// C.g_logcallback = (C.GOLOG_PROC)(unsafe.Pointer(&MyFunc))
	// C.g_logcallback = (*[0]byte)(unsafe.Pointer(&MyFunc))

	// cb := (C.GOLOG_PROC)(unsafe.Pointer(&MyFunc))
	// C.setcallback(cb)

	// C.setcallback((C.GOLOG_PROC)(unsafe.Pointer(&MyFunc)))
	// C.setcallback((*[0]byte)(unsafe.Pointer(&MyFunc)))

	fmt.Printf("%#v\n", &MyFunc)
	fmt.Printf("%#v\n", &C.g_logcallback)

	C.test()

	str := C.CString("abc")
	defer C.free(unsafe.Pointer(str))
	// MyFunc(str)

	// C.bridge(str)
}

