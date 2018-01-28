package main

/*
typedef void (*Callback) (char* buffer);
Callback fn;

void bridge_callback(Callback fn, char *buffer);
*/
import "C"
import "fmt"
import (
  "os"
)

var fl *os.File;

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

//export SetCallBack
func SetCallBack(pt C.Callback) *C.char{
  C.fn = pt;
  f := C.Callback(pt);
  C.bridge_callback(f, C.CString("i am a msg, hahahahahah"));
  fl.WriteString("in set call back");
  return C.CString("set call back ok");
}

func main(){
  fl, _ := os.Create("/tmp/bonly.log");
  defer fl.Close();
  fl.WriteString("in main\n");
}

/*
go build -buildmode=plugin -o libfc.so mylib.go // 不行，格式是go的
go build -o libfc.so -buildmode=c-shared mylib.go
*/
