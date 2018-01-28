package proto

import (
	"encoding/json"
	"reflect"
	"log"
)

// var Proto = make(map[string]reflect.Type); //@note 这里两种方式的含义不一样？
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

func (this *Cmd)Decode(name string, data []byte, pt *interface{})(err error){
	var raw json.RawMessage;
	*pt = &raw;

	log.Printf("aaa: %#v\n", this);

	if err = json.Unmarshal(data, this); err != nil{
		return err;
	}

	log.Printf("bbbb: %#v\n", this);

	// obj := reflect.New(Proto[name]).Elem().Interface(); // @todo 检查是否有此类型
	obj := Proto[name]();
	// log.Printf("eeee: %#v\n", reflect.TypeOf(obj).Name());
	log.Printf("eeee: %#v\n", obj);
	// if err = obj.Decode(raw); err != nil{
	if err = json.Unmarshal(raw, &obj); err != nil{
		return err;
	}

	*pt = obj;
	log.Printf("cccc: %#v\n", this);
	return nil;
}

func (this *Cmd)Encode()(ret []byte, err error){
	return json.MarshalIndent(this, " ", " ");
}

// func (this *Cmd)Decode(name string, data []byte, pt *interface{})(err error){
// 	var raw json.RawMessage;
// 	// *pt = &raw;

// 	// log.Printf("aaa: %#v\n", this);
// 	field := Cmd{
// 		Req: &raw,
// 	};

// 	if err = json.Unmarshal(data, &field); err != nil{
// 		return err;
// 	}

// 	log.Printf("bbbb: %#v\n", field);

// 	obj := reflect.New(Proto[name]).Elem().Interface(); // @todo 检查是否有此类型
// 	log.Printf("eeee: %#v\n", obj);
// 	if err = json.Unmarshal(raw, &obj); err != nil{
// 		return err;
// 	}

// 	// *pt = obj
// 	log.Printf("cccc: %#v\n", obj);
// 	return nil;
// }