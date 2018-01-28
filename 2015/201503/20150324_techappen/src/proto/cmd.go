package proto

import (
	"encoding/json"
	"reflect"
	// "log"
)

var Proto = make(map[string]func() interface{});

type Inf struct{
	Fnc reflect.Value;
	Net UserClient;
};

type ICmd interface{
	Decode(name string, data []byte, pt *interface{}) error;
	Encode() ([]byte, error);
};

type Cmd struct{
	Inf interface{} `json:"Inf,omitempty"`;
	Req interface{} `json:"Req,omitempty"`;
	Rep interface{} `json:"Rep,omitempty"`;
};

/*
name: 结构名
data: 数据
pt:   对应的interface指针 Inf/Req/Rep
*/
func (this *Cmd)Decode(name string, data []byte, pt *interface{})(err error){
	var raw json.RawMessage;
	*pt = &raw;  //把对应的成员变成暂存裸数据

	if err = json.Unmarshal(data, this); err != nil{  //取出裸数据
		return err;
	}

	obj := Proto[name](); //根据类型名新建相应的对象
	if err = json.Unmarshal(raw, &obj); err != nil{ //裸数据解释到新对象中
		return err;
	}

	*pt = obj; //把新对象替换原协议对象
	return nil;
}

func (this *Cmd)Encode()(ret []byte, err error){
	return json.MarshalIndent(this, " ", " ");
}

