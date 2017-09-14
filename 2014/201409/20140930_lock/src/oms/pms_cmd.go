/*
auth: bonly
create: 2016.9.19
desc: PMS端的接口
*/
package oms 

import(
"fmt"
"encoding/json"
"proto"
)

type PMS_CMD interface{
	Decode(pack []byte)(error);
	Encode()(pack []byte, err error);
	New()(cmd interface{});
	Process(chan interface{});
};

type PMS_Command struct{
	Name *string;		
	Data interface{};
	Cmd *proto.CMD;
}

func (this *PMS_Command)Decode(pack []byte)(error){
	return json.Unmarshal(pack, this);
}

func (this *PMS_Command)Encode()(pack []byte, err error){
	return json.MarshalIndent(this, " ", " ");
}

func (this *PMS_Command)New()(cmd interface{}){
	var name string;
	this.Name = &name;
	return this;
}

func (this *PMS_Command)Process(chan interface{}){
	fmt.Println("没啥好处理的\n");
}

type PMS_Manual struct{
	Name string;
	ResultID string;
	Description string;
};