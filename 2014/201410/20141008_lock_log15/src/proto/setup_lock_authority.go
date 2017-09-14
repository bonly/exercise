/*
auth: bonly
create: 2016.9.18
desc: 设置开门权限
*/
package proto

import(
"fmt"
)

var SETUP_LOCK_AUTHORITY_CMD_PARAM = []byte{0x90, 0x38};

const ALLOW_OPEN = 0x30; //允许开门，所有用户可以开门
const FORBIT_OPEN = 0x31; //禁止开门，管理员可以开门

type Setup_lock_authority struct{
	Command;
	Authority uint8;
}

func (this *Setup_lock_authority)New()(cmd interface{}){
	this.Command.New();
	this.head.Type = SRV_2_LOCK;
	this.head.Cmd = SETUP_LOCK_AUTHORITY_CMD_PARAM[0];
	this.head.Param = SETUP_LOCK_AUTHORITY_CMD_PARAM[1]; 
	this.Authority = ALLOW_OPEN;
	this.verify = false;
	return this;
}

func (this *Setup_lock_authority)Encode()(pack []byte, err error){
	arr := []byte{this.Authority};
	this.Command.Add_data(arr);

	return this.Command.Encode();
}

type R_Setup_lock_authority struct{
	Command;
}

func (this *R_Setup_lock_authority)Decode(pack []byte)(cmd interface{}){
	cmd = this.Command.Decode(pack);
	if cmd == nil{
		return nil;
	}
	fmt.Printf("收到设置门锁权限设置应答:\n");
	Hex_Dump(this.buf, len(this.buf));

	fmt.Printf("头数据head: %+v\n", *(this.head));
	fmt.Printf("尾数据tail: %+v\n", *(this.tail));
	return nil;
}