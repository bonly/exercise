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
	kcp "github.com/xtaci/kcp-go"
	"golang.org/x/crypto/pbkdf2"
	"crypto/sha1"
)


func main(){}

var quit bool = false;

const addr = "127.0.0.1:4242";
var key = []byte("testkey");
var salt = []byte("techappen");

var sess *kcp.UDPSession = nil;


//export Init 
func Init()(int){
	fmt.Printf("init go\n");
	return Connect();
}

//export Run
func Run()(int){
	for ;quit == false; {
		Pull();
	}
	return 0;
}

func Pull(){
	cnt, err := Get_pack();
	if err == nil && cnt > 0{
		cstr := C.CString("call back from go");
		defer C.free(unsafe.Pointer(cstr));
		C.bridge_callback(C.uint(18), cstr);
		fmt.Printf("回调CS成功\n");
	}
}

func Connect()int{
	var err error;

	pass := pbkdf2.Key(key, salt, 4096, 32, sha1.New);	
	block, _ := kcp.NewSalsa20BlockCrypt(pass);
	sess, err = kcp.DialWithOptions(addr, block, 10, 3)
	if err != nil{
		fmt.Printf("%v\n", err);
		return -1;
	}

	sess.SetStreamMode(true)
	sess.SetStreamMode(false)
	sess.SetStreamMode(true)
	sess.SetWindowSize(4096, 4096)
	sess.SetReadBuffer(4 * 1024 * 1024)
	sess.SetWriteBuffer(4 * 1024 * 1024)
	sess.SetStreamMode(true)
	sess.SetNoDelay(1, 10, 2, 1)
	sess.SetMtu(1400)
	sess.SetMtu(1600)
	sess.SetMtu(1400)
	sess.SetACKNoDelay(true)
	sess.SetDeadline(time.Now().Add(5 * time.Second))	

	sess.SetWriteDelay(true);
	sess.SetDUP(1);

	return 0;
}

func Get_pack()(cnt int, err error){
    buf := make([]byte, 255);
	cnt, err = sess.Read(buf);
	if err != nil{
		// fmt.Printf("read err: %v\n", err);
		return;
	}
	fmt.Printf("get data %d: %s\n", cnt, string(buf));
	return;
}

//export Stop
func Stop(){
	fmt.Printf("default1\n");
	quit = true;
	if sess != nil{
		sess.Close();
		sess = nil;
	}
	fmt.Printf("default2\n");
}

//export Put_pack
func Put_pack()(int){
	fmt.Printf("put pack\n");
	nc, err := sess.Write([]byte("bonly"));
	if err != nil{
		// fmt.Printf("send err: %v\n", err);
		return -1;
	}
	fmt.Printf("send %d success\n", nc);
	return 0;
}

//export SetCallBack
func SetCallBack(pt C.Callback){
	C.fn = pt;
}

/*
go build -buildmode=c-shared -o libtechappen.so dll.go dll_c.go
*/