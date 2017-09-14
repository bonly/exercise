/*
auth: bonly
create: 2016.9.19
desc: 删除全部用户
*/
package proto

import(
"fmt"
)

var USER_DEL_ALL_CMD_PARAM = []byte{0x90, 0x34};

type User_del_all struct{
	Command;
};

func (this *User_del_all)New()(cmd interface{}){
	this.Command.New();
	this.head.Type = SRV_2_LOCK;
	this.head.Cmd = USER_DEL_ALL_CMD_PARAM[0];
	this.head.Param = USER_DEL_ALL_CMD_PARAM[1];

	this.verify = false;	
	return this;
}

func (this *User_del_all)Encode()(pack []byte, err error){
	return this.Command.Encode();
}

type R_User_del_all struct{
	Command;
}

func (this *R_User_del_all)Decode(pack []byte)(cmd interface{}){
	cmd = this.Command.Decode(pack);
	if cmd == nil{
		return nil;
	}
	fmt.Printf("收到删除所有用户应答:\n");
	Hex_Dump(this.buf, len(this.buf));

	fmt.Printf("头数据head: %+v\n", *(this.head));
	fmt.Printf("尾数据tail: %+v\n", *(this.tail));
	return nil;
}