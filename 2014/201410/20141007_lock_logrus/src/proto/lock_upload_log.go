/*
auth: bonly
create: 2016.9.18
desc: 锁主动上传开门记录
*/
package proto

import(
"fmt"
"unsafe"
)

var LOCK_UPLOAD_LOG_CMD_PARAM = []byte{0xB0, 0x30};

type Lock_log struct{
	Open_type uint8; //开锁类型
	SN_len uint8; //密码或卡号的长度
	SN [11]uint8; //密码或卡号
	DateTime [6]uint8; //开锁时间
};

type Lock_upload_log struct{
	Command;
	Data *Lock_log;
};

func (this *Lock_upload_log)New()(cmd interface{}){
	return nil;
}

func (this *Lock_upload_log)Encode()(pack []byte, err error){
	return nil, fmt.Errorf("未实现协议");
}

func (this *Lock_upload_log)Decode(pack []byte)(cmd interface{}){
	cmd = this.Command.Decode(pack);
	if cmd == nil{
		return nil;
	}
	fmt.Printf("收到门锁开门日志:\n");
	Hex_Dump(this.buf, len(this.buf));

	this.Data = (*Lock_log)(unsafe.Pointer(&this.data[0]));
	fmt.Printf("头数据head: %+v\n", *(this.head));
	fmt.Printf("体数据body: %+v\n", *(this.Data));
	fmt.Printf("尾数据tail: %+v\n", *(this.tail));
	return *this;
}

type R_Lock_upload_log struct{
	Command;
}

func (this *R_Lock_upload_log)New()(cmd interface{}){
	this.Command.New();
	this.Command.verify = false;
	this.head.Type = YES;
	this.head.Cmd = LOCK_UPLOAD_LOG_CMD_PARAM[0];
	this.head.Param = LOCK_UPLOAD_LOG_CMD_PARAM[1];	
	return this;
}

func (this *R_Lock_upload_log)Encode()(pack []byte, err error){
	return this.Command.Encode();
}

func (this *R_Lock_upload_log)Decode(pack []byte)(cmd interface{}){
	return nil;
}