package pomelo

/*
#cgo CFLAGS: -g -I/home/bonly/libpomelo/include -I/home/bonly/libpomelo -I/home/bonly/libpomelo/deps/uv/include -I/home/bonly/libpomelo/deps/jansson/src
#cgo linux CFLAGS: -DLINUX=1
#cgo LDFLAGS: -L/home/bonly/libpomelo -lpomelo -L/home/bonly/libpomelo/deps/jansson/src/.libs/ -ljansson -L/home/bonly/libpomelo/deps/uv -luv
#include "libpomelo.h" 
*/
import "C"

import (
  _ "fmt"
  _ "reflect"
)

var Client *_Ctype_pc_client_t;

func Connect(ip string, port int) (*_Ctype_pc_client_t){
   cip := C.CString(ip);
   ret := C.Connect (cip, C.int(port));
   return ret;
}

func Notify(cli *_Ctype_pc_client_t){
   js := C.Add(C.CString("content"), C.CString("hi all"));
   C.Notify(cli, C.CString("chat.chatHandler.notifyall"), js);
}

func WaitJoin(cli *_Ctype_pc_client_t){
   C.wait_join(cli);
}

