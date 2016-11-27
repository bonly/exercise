package main

/*
#include <stdio.h>
#include <stdlib.h>
void put_strings(char *sarr[], int n)
{
        int i;
        for (i = 0; i < n; ++i) {
                puts(sarr[i]);
        }
}

char *sarr[] = {
        "THIS", "IS", "A", "TEST",
};

char **parr = sarr;
int parr_size = sizeof(sarr)/sizeof(sarr[0]);
*/
import "C"

import (
        "fmt"
        "unsafe"
)

func main() {
        // Go 转 C
        gostrarr := []string{
                "this", "is", "a", "test",
        }
        ptrarr := make([]*C.char, len(gostrarr))
        for i := range gostrarr {
                ptrarr[i] = C.CString(gostrarr[i])
                defer C.free(unsafe.Pointer(ptrarr[i]))
        }
        C.put_strings((**C.char)(unsafe.Pointer(&ptrarr[0])), C.int(len(ptrarr)))
        // C 数组转 Go
        gosarr := make([]string, unsafe.Sizeof(C.sarr)/unsafe.Sizeof(C.sarr[0]))
        for i := range gosarr {
                gosarr[i] = C.GoString(C.sarr[i])
        }
        put_strings(gosarr)
        // C 指针转 Go，需要手工做指针算术，很恶心
        gosarr2 := make([]string, C.parr_size)
        for i := range gosarr2 {
                p := uintptr(unsafe.Pointer(C.parr)) + uintptr(i)*unsafe.Sizeof(*C.parr)
                q := *(**C.char)(unsafe.Pointer(p))
                gosarr2[i] = C.GoString(q)
        }
        put_strings(gosarr2)
}

func put_strings(sarr []string) {
        for _, s := range sarr {
                fmt.Println(s)
        }
}