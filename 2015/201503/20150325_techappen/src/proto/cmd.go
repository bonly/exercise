package proto

import (
	"encoding/json"
	"reflect"
	// "log"
)

var Proto = make(map[string]func() interface{});

type Inf struct{
	Fnc  reflect.Value;
	Net  UserClient;
	Sc   string; //Scope
	Func string;
};

type ICmd interface{
	Decode(data []byte) error;
	Encode() ([]byte, error);
};

type Cmd struct{
	Key  Inf `json:"Key,omitempty"`;
	Data interface{} `json:"Data,omitempty"`;
};

	// Fnc = reflect.ValueOf(&inf);
/*
name: 结构名
data: 数据
pt:   对应的interface指针 Inf/Req/Rep
*/
func (this *Cmd)Decode(data []byte)(err error){
	var raw json.RawMessage;
	this.Data = &raw;  //把对应的成员变成暂存裸数据

	if err = json.Unmarshal(data, this); err != nil{  //取出裸数据
		return err;
	}

	obj := Proto[this.Key.Sc](); //根据类型名新建相应的对象 @todo 检查类型是否正确
	if err = json.Unmarshal(raw, &obj); err != nil{ //裸数据解释到新对象中
		return err;
	}

	this.Data = obj; //把新对象替换原协议对象
	this.Key.Fnc = reflect.ValueOf(obj);
	return nil;
}

func (this *Cmd)Encode()(ret []byte, err error){
	return json.MarshalIndent(this, " ", " ");
}

