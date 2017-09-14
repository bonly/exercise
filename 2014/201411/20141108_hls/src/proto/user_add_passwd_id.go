/*
auth: bonly
create: 2016.9.27
desc: 添加用户密码ID
*/
package proto

import(
"fmt"
)

var USER_ADD_PASSWD_ID_CMD_PARAM = []byte{0x90, 0x51};

type User_add_passwd_id struct{
	Command;
	Passwd_len [1]uint8; //密码长度
	Passwd [11]uint8; //密码 11
	Period [12]uint8; //时段 12
};

func (this *User_add_passwd_id)New()(cmd interface{}){
	this.Command.New();
	this.head.Type = SRV_2_LOCK;
	this.head.Cmd = USER_ADD_PASSWD_ID_CMD_PARAM[0];
	this.head.Param = USER_ADD_PASSWD_ID_CMD_PARAM[1];

	this.verify = false;	
	return this;
}

func (this *User_add_passwd_id)Encode()(pack []byte, err error){
	this.Command.Add_data(this.Passwd_len[:]);
	this.Command.Add_data(this.Passwd[:]);
	this.Command.Add_data(this.Period[:]);
	return this.Command.Encode();
}

type R_User_add_passwd_id struct{
	Command;
	Data [2]byte;
}

func (this *R_User_add_passwd_id)Decode(pack []byte)(cmd interface{}){
	cmd = this.Command.Decode(pack);
	if cmd == nil{
		return nil;
	}
	fmt.Printf("收到增加用户(密码)应答:\n");
	Hex_Dump(this.buf, len(this.buf));

	fmt.Printf("头数据head: %+v\n", *(this.head));
	fmt.Printf("尾数据tail: %+v\n", *(this.tail));
	return nil;
}


