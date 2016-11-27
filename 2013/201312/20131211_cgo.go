package main 

/*
#include <malloc.h>

#define SIZE(array) (int)((sizeof array) /(sizeof *array))

int IntArraySize(const int *array) {
        return SIZE(array);
}

int *NewIntArray() {
        int *a;
        a = malloc(2);
        a[0] = 1;
        a[1] = 2;
        return a;
}
*/
import "C"

import (
        "fmt"
        "unsafe"
)

func main() {
        cArray := C.NewIntArray()
        defer C.free(unsafe.Pointer(cArray))

        fmt.Printf("C array %#v\n", cArray)
        fmt.Printf("size of C array %d\n", int(C.IntArraySize(cArray)))

        arraySize := int(C.IntArraySize(cArray))
        goSlice := (*[1 << 30]C.int)(unsafe.Pointer(cArray))[:arraySize:arraySize]
        fmt.Printf("turn C array into Go slice %#v\n", goSlice)
}