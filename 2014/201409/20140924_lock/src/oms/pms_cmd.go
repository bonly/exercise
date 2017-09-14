/*
auth: bonly
create: 2016.9.19
desc: PMS端的接口
*/
package oms 

import(
"encoding/xml"
"fmt"
)

type PMS_CMD interface{
	Decode(pack []byte)(error);
	Encode()(pack []byte, err error);
	New()(cmd interface{});
	Process(chan interface{});
};

type PMS_Command struct{
	XMLName *xml.Name `xml:"Command"`;
	Name *string `xml:"name,attr"`;		
	Data interface{};	
}

func (this *PMS_Command)Decode(pack []byte)(error){
	return xml.Unmarshal(pack, this);
}

func (this *PMS_Command)Encode()(pack []byte, err error){
	return nil, fmt.Errorf("未实现\n");
}

func (this *PMS_Command)New()(cmd interface{}){
	var name string;
	this.Name = &name;
	return this;
}

func (this *PMS_Command)Process(chan interface{}){
	fmt.Println("没啥好处理的\n");
}