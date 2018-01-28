package main

import "C"
import "fmt"

//export Foo
func Foo(){
  fmt.Println("he");
}

//export GetString
func GetString() *C.char{
     return C.CString("hello from go");
}

//export MergeString
func MergeString(left *C.char, right *C.char) *C.char{
    merge := C.GoString(left) + C.GoString(right);
    return C.CString(merge);
}

func main(){}

/*
go build -buildmode=plugin -o libfc.so mylib.go // 不行，格式是go的
go build -o libfc.so -buildmode=c-shared mylib.go
*/
