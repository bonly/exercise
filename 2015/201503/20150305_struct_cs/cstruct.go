package main 

/*
#cgo LDFLAGS: -pthread -fPIC
struct Data{
	int cmd;
	char msg[1024];
};
*/
import "C"

func main(){}

func Send(msg *C.char) C.int{
	return 0;
}

func Recv(msg *C.char) C.int{
	return 0;
}

func Print(data *C.struct_Data) *C.char{
	return C.CString("oooo");
}