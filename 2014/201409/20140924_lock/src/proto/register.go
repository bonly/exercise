/*
auth: bonly
create: 2016.9.14
desc: 注册
*/
package proto

import(
"encoding/binary"
// "unsafe"
"bytes"
)

type Register struct{
	Ctl uint32; //控制编号
	Num uint8;  //编号
};

func (this *Register)Decode(pack []byte)(cmd interface{}){
	binary.Read(bytes.NewBuffer(pack), binary.BigEndian, this);
	return this;
}