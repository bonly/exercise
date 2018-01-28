package proto

/*
#cgo LDFLAGS: -pthread -fPIC  -lgate
#include "../app/gate.h"
#include <stdlib.h>
// void 
// bridge_callback(Callback fn, unsigned int sn, char *buffer){
//     return fn(sn, buffer);
// }
*/
import "C"

import (
	"encoding/json"
	"reflect"
	"log"	
	"fmt"
	"unsafe"
)
// type Callback C.Callback;

var Proto = make(map[string]func() interface{});
var Inter = make(map[string]func() interface{});

type Inf struct{
	Fnc  reflect.Value;
	Net  UserClient;
	Scope string; //Scope
	Argv string; //Param
	Func string; //Func Name
};

type ICmd interface{
	Decode(data []byte) error;
	Encode() ([]byte, error);
};

type Cmd struct{
	Key  Inf `json:"Key,omitempty"`;
	Data interface{} `json:"Data,omitempty"`;
};

type Ret struct{
	Sn  uint32;
	Ret string;
	Msg string;
};

/*
name: 结构名
data: 数据
pt:   对应的interface指针 Inf/Req/Rep
*/
func (this *Cmd)Decode(data []byte)(err error){
	defer func(){
		if err := recover(); err != nil{
			fmt.Errorf("接口未定义, %v\n", err);
		}	
	}();
	var raw json.RawMessage;
	this.Data = &raw;  //把对应的成员变成暂存裸数据

	if err = json.Unmarshal(data, this); err != nil{  //取出裸数据
		return err;
	}

	if len(this.Key.Argv) <=0 || len(this.Key.Scope) <= 0{
		return fmt.Errorf("协议不正确");
	}
	obj := Proto[this.Key.Argv](); //根据类型名新建相应的参数对象 @todo 检查类型是否正确
	if err = json.Unmarshal(raw, &obj); err != nil{ //裸数据解释到新对象中
		return err;
	}

	this.Data = obj; //把新对象替换原协议对象

	sc := Inter[this.Key.Scope](); //根据范围名新建对应的接口 @todo 检查类型是否正确
	this.Key.Fnc = reflect.ValueOf(sc);
	return nil;
}

func (this *Cmd)Encode()(ret []byte, err error){
	return json.MarshalIndent(this, " ", " ");
}

func Mk_Ret(sn uint32, code string, msg string)(ret string){
	obj := &Ret{Sn:sn, Ret:code, Msg:msg};
	js, _ := json.MarshalIndent(obj, " ", " ");
	log.Printf("make ret: %#v\n", string(js));
	return string(js);
}

func RepCallback(sn uint32, str string){
	cstr := C.CString(str);
	defer C.free(unsafe.Pointer(cstr));
	C.bridge_callback(C.uint(sn), cstr);
	log.Printf("回调CS成功");
	return;
}

// func SetCallBack(pt Callback){
// 	log.Printf("注册回调函数成功");
// 	C.fn = (C.Callback)(pt);
// }