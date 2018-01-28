package main 

/*
// #cgo LDFLAGS: -pthread -fPIC 
#include "gate.h"
*/
import "C"

import (  //注册需用到的模块
	srv "cli"
	_ "cli/user"
	_ "cli/scene"	
	_ "cli/cards"
	_ "cli/pk"
)

import (
	"log"
	"reflect"
	"proto"
	"fmt"
)

func init(){
	log.Printf("核心加载\n");
}

func main(){
}

//export Run
func Run(){
	srv.Start();
}

//export Exit
func Exit(){
	srv.End();
}

//export Proc
func Proc(c_json *C.char)(ret int){
	defer func(){
		if err := recover(); err != nil{
			fmt.Errorf("接口未定义, %v\n", err);
			ret = -88;
		}	
	}();	
	data := C.GoString(c_json);

	var cmd proto.ICmd; //解包处理
	cmd = &proto.Cmd{};
	err := cmd.Decode([]byte(data));
	if err != nil{
		log.Printf("数据错误: %v\n", err);
		return -99;
	}

	log.Printf("call: %v \n", cmd.(*proto.Cmd).Key);

	method := cmd.(*proto.Cmd).Key.Fnc.MethodByName(cmd.(*proto.Cmd).Key.Func);
	var rcv string;
	if method.IsValid(){
		vdata := reflect.ValueOf(cmd.(*proto.Cmd).Data); //Data结构转换成参数
		res := method.Call([]reflect.Value{0:vdata});  // return []reflect.Value
		val := reflect.ValueOf(res).Interface();
		rcv = val.([]reflect.Value)[0].String();
		log.Printf("调用返回: %v\n", rcv);
		ret = 0;
	}else{
		log.Printf("无此方法: %v\n", cmd.(*proto.Cmd).Key.Func);
		ret = -2;
	}
	return ret;
}

////ddddexport SetCallBack
// func SetCallBack(pt C.Callback){
// 	proto.SetCallBack ((proto.Callback)(pt));
// }

// func Copy(dst []byte, src []byte, size int){
// 	C.memcpy(unsafe.Pointer(&dst[0]), unsafe.Pointer(&src[0]), C.size_t(size));
// }