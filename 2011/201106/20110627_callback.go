package gocallback

import (
        "unsafe"
        "fmt"
)

/*
extern void go_callback_int(void* foo, int p1);

void CallMyFunction(void* pfoo) {
        go_callback_int(pfoo, 5);
}
*/
import "C"

//export go_callback_int
func go_callback_int(pfoo unsafe.Pointer, p1 C.int) {
        foo := *(*func(C.int))(pfoo)
        foo(p1)
}

func MyCallback(x C.int) {
        fmt.Println("callback with", x)
}
//we store it in a global variable so that the garbage collector doesn't clean up the memory for any temporary variables created.
var MyCallbackFunc = MyCallback

func Example() {
        C.CallMyFunction(unsafe.Pointer(&MyCallbackFunc))
}


//https://code.google.com/p/go-wiki/wiki/cgo


