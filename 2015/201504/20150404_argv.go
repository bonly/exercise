package main 

import "C"
import "unsafe"

func GoStrings(argc C.int, argv **C.char) []string {

    length := int(argc)
    tmpslice := (*[1 << 30]*C.char)(unsafe.Pointer(argv))[:length:length]
    gostrings := make([]string, length)
    for i, s := range tmpslice {
        gostrings[i] = C.GoString(s)
    }
    return gostrings
}

func main(){

}
/*
The easiest and safest way is to copy it to a slice, not specifically to [1024]byte

mySlice := C.GoBytes(unsafe.Pointer(&C.my_buff), C.BUFF_SIZE)
To use the memory directly without a copy, you can "cast" it through an unsafe.Pointer.

mySlice := (*[1 << 30]byte)(unsafe.Pointer(&C.my_buf))[:int(C.BUFF_SIZE):int(C.BUFF_SIZE)]
// or for an array if BUFF_SIZE is a constant
myArray := *(*[C.BUFF_SIZE]byte)(unsafe.Pointer(&C.my_buf))

https://stackoverflow.com/questions/27532523/how-to-convert-1024c-char-to-1024byte
https://github.com/golang/go/wiki/cgo#Turning_C_arrays_into_Go_slices
https://my.oschina.net/chai2010/blog/168484

			ptr := unsafe.Pointer(&ret);
			arr := *((*[]byte)(ptr));
			sz = copy(arr, []byte("OK"));
*/