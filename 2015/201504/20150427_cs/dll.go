package main

/*
#include <stdlib.h> // for C.free
typedef void (*Callback)(unsigned int sn, char *buffer);
extern Callback fn; //extern的东西需要把实体放到另一个文件中实现
extern void bridge_callback(unsigned int sn, char *buffer);
*/
import "C"

import (
	"fmt"
	"C"  //用c-shared必须要有这个
	"unsafe"
	"time"
	// "crypto/tls"
	quic "github.com/lucas-clemente/quic-go"
)


func main(){}

var sock chan string = nil;
var quit chan int = nil;
const addr = "127.0.0.1:4242";
var session quic.Session = nil;
var stream quic.Stream = nil;

//export Init 
func Init()(int){
	fmt.Printf("init go\n");
	quit = make (chan int);
	sock = make (chan string);
	return Connect();
}

//export Run
func Run()(int){
	for {
		select {
			case str := <- sock: {
			  fmt.Printf("get str %s\n", str);
			  //call cs
			  cstr := C.CString("call back from go");
			  defer C.free(unsafe.Pointer(cstr));
			  C.bridge_callback(C.uint(18), cstr);
			  fmt.Printf("回调CS成功\n");
			  break;
			}
			case <- quit:
			  fmt.Printf("get stop sign\n");
			  return 0;
			default:
			//   Get_pack();
			  break;
		}
	}
	return 0;
}

func Connect()int{
	// var cfg quic.Config;
	var err error;
	// session, err = quic.DialAddr(addr, &tls.Config{InsecureSkipVerify: true}, &quic.Config{});//&tls.Config{InsecureSkipVerify: true}, nil); //&cfg);
	session, err = quic.DialAddr(addr, &quic.Config{});//&tls.Config{InsecureSkipVerify: true}, nil); //&cfg);
	if err != nil{
		fmt.Printf("%v\n", err);
		return -1;
	}

	stream, err = session.OpenStreamSync();
	if err != nil{
		fmt.Printf("%v\n", err);
		return -2;
	}	
	fmt.Printf("stream: %#v\n", stream);
	return 0;
}

func Get_pack(){
	// fmt.Printf("default\n");
}

//export Stop
func Stop(){
	fmt.Printf("default1\n");
	if quit != nil{
		select {
    		case <- time.After(time.Second *2):
        		println("write channel timeout");
    		case quit <- 0:
        		println("write ok");
		}
	    close(quit);
	}
	fmt.Printf("default2\n");
	if sock != nil{
		close(sock);
	}
	if stream != nil{
		stream.Close();
		stream = nil;
	}
}

//export Put_pack
func Put_pack()(int){
	fmt.Printf("put pack\n");
	stream.Write([]byte("bonly"));
	// sock <- "this is a test";
	return 0;
}

//export SetCallBack
func SetCallBack(pt C.Callback){
	C.fn = pt;
}

/*
go build -buildmode=c-shared -o libtechappen.so dll.go dll_c.go
*/