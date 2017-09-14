/*
auth: bonly
create: 2016.9.18
desc: 设置门锁时间
*/
package proto

import(
"fmt"
// "unsafe"
)

var SETUP_LOCK_TIME_CMD_PARAM = []byte{0x90, 0x35};

type Setup_lock_time struct{
	Command;
	DateTime [6]uint8;
}

func (this *Setup_lock_time)New()(cmd interface{}){
	this.Command.New();
	this.head.Type = SRV_2_LOCK;
	this.head.Cmd = SETUP_LOCK_TIME_CMD_PARAM[0];
	this.head.Param = SETUP_LOCK_TIME_CMD_PARAM[1]; 
	return this;
}

func (this *Setup_lock_time)Encode()(pack []byte, err error){
	this.Command.Add_data(this.DateTime[:]);

	return this.Command.Encode();
}

func (this *Setup_lock_time)Decode(pack []byte)(cmd interface{}){
	cmd = this.Command.Decode(pack);
	if cmd == nil{
		return nil;
	}
	Hex_Dump(this.buf, len(this.buf));
	copy(this.DateTime[:], this.data[:this.data_len]);
	return nil;
}

type R_Setup_lock_time struct{
	Command;
}

func (this *R_Setup_lock_time)Decode(pack []byte)(cmd interface{}){
	cmd = this.Command.Decode(pack);
	if cmd == nil{
		return nil;
	}
	fmt.Printf("收到设置门锁时间应答:\n");
	Hex_Dump(this.buf, len(this.buf));

	fmt.Printf("头数据head: %+v\n", *(this.head));
	fmt.Printf("尾数据tail: %+v\n", *(this.tail));
	return nil;
}