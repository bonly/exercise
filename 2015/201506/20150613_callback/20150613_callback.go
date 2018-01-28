package main

/*
#include <stddef.h>
typedef void (*GOLOG_PROC)(char* szDscr);
GOLOG_PROC g_logcallback = NULL;


void SetLogCallback(GOLOG_PROC logcallback) {
	 g_logcallback = logcallback;
 }

*/
import "C"

import (
	"fmt"
)

func logcallback(dscr *C.char) {
	fmt.Println("logcallback: ")
}

var Myfunc = logcallback

func main() {
	fn := C.GOLOG_PROC(logcallback)

	C.SetLogCallback(Myfunc)
}
