/*
auth: bonly
create: 2016.9.18
desc: 设置开门权限
*/
package proto

import(
"fmt"
)

var SETUP_LOCK_CLEAN_LOG_CMD_PARAM = []byte{0x90, 0x39};

type Setup_lock_clean_log struct{
	Command;
}

func (this *Setup_lock_clean_log)New()(cmd interface{}){
	this.Command.New();
	this.head.Type = SRV_2_LOCK;
	this.head.Cmd = SETUP_LOCK_CLEAN_LOG_CMD_PARAM[0];
	this.head.Param = SETUP_LOCK_CLEAN_LOG_CMD_PARAM[1]; 
	// this.verify = false;
	return this;
}

func (this *Setup_lock_clean_log)Encode()(pack []byte, err error){
	return this.Command.Encode();
}

type R_Setup_lock_clean_log struct{
	Command;
}

func (this *R_Setup_lock_clean_log)Decode(pack []byte)(cmd interface{}){
	cmd = this.Command.Decode(pack);
	if cmd == nil{
		return nil;
	}
	fmt.Printf("收到设置门锁日志清空应答:\n");
	Hex_Dump(this.buf, len(this.buf));

	fmt.Printf("头数据head: %+v\n", *(this.head));
	fmt.Printf("尾数据tail: %+v\n", *(this.tail));
	return nil;
}