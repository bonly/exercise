package main

/*
#include <stdio.h>
#include <stdlib.h>

static void cprint(int argc, char* argv[]) {
	extern void goprint(int p0, char** p1);
	goprint(argc, argv);
}
*/
import "C"
import (
	"fmt"
	"os"
	"unsafe"
)

func main() {
	argc := C.int(len(os.Args))
	argv := make([]*C.char, len(os.Args))
	for i := 0; i < len(os.Args); i++ {
		argv[i] = C.CString(os.Args[i])
		defer C.free(unsafe.Pointer(argv[i]))
	}
	C.cprint(argc, (**C.char)(unsafe.Pointer(&argv[0])))
	fmt.Println("Done")
}

//export goprint
func goprint(argc C.int, argv_ **C.char) {
	argv := (*(*[1 << 30]*C.char)(unsafe.Pointer(argv_)))[:int(argc)]
	args := make([]string, int(argc))
	for i := 0; i < int(argc); i++ {
		args[i] = C.GoString(argv[i])
	}
	for i := 0; i < len(args); i++ {
		fmt.Printf("argv[%d]: %s\n", i, args[i])
	}
}
