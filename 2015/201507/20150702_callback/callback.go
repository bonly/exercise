package main

/*
#include <stddef.h>
#include <stdlib.h>
typedef void (*GOLOG_PROC) (char *szDscr);
void bridge(char *szDscr);
extern GOLOG_PROC fn;

extern void go_logcallback(char*);

extern void set_callback(void *cb);
*/
import "C"

import(
	"fmt"
	"unsafe"
)

//export go_logcallback
func go_logcallback(dscr *C.char){
	fmt.Println("go log callback")
}

func main(){
	str := C.CString("abc")
	defer C.free(unsafe.Pointer(str))

	// C.fn = (*[0]byte)(C.go_logcallback)
	C.set_callback(C.go_logcallback)
	C.bridge(str)
}

