package main

//#cgo CFLAGS: -DPNG_DEBUG=1
//#cgo linux CFLAGS: -DLINUX=1
//#cgo LDFLAGS: -L/home/bonly/src/ -lmgc 
//#include "20110211_gc.h"
import "C"

func Prints(s string){
    p := C.CString(s);
    C.mprints(p);
}

func main(){
   Prints("this is a test");
}

/*
没有main是令go build不生成任何文件
*/
