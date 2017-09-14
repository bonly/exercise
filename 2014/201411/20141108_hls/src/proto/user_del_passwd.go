/*
auth: bonly
create: 2016.9.19
desc: 删除用户密码
*/
package proto

import(
"fmt"
)

var USER_DEL_PASSWD_CMD_PARAM = []byte{0x90, 0x32};

type User_del_passwd struct{
	Command;
	Passwd_len [1]uint8; //密码长度
	Passwd [11]uint8; //密码 11
};

func (this *User_del_passwd)New()(cmd interface{}){
	this.Command.New();
	this.head.Type = SRV_2_LOCK;
	this.head.Cmd = USER_DEL_PASSWD_CMD_PARAM[0];
	this.head.Param = USER_DEL_PASSWD_CMD_PARAM[1];

	this.verify = false;	
	return this;
}

func (this *User_del_passwd)Encode()(pack []byte, err error){
	this.Command.Add_data(this.Passwd_len[:]);
	this.Command.Add_data(this.Passwd[:]);
	return this.Command.Encode();
}

type R_User_del_passwd struct{
	Command;
}

func (this *R_User_del_passwd)Decode(pack []byte)(cmd interface{}){
	cmd = this.Command.Decode(pack);
	if cmd == nil{
		return nil;
	}
	fmt.Printf("收到删除用户(密码)应答:\n");
	Hex_Dump(this.buf, len(this.buf));

	fmt.Printf("头数据head: %+v\n", *(this.head));
	fmt.Printf("尾数据tail: %+v\n", *(this.tail));
	return nil;
}