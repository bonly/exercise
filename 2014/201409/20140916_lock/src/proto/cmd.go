/*
auth: bonly
create: 2016.9.14
desc: 紫蜂协议
*/
package proto

import (
// "encoding/binary"
// "bytes"
"unsafe"
"fmt"
)

const PRE_FIX = 0x2;
type Zigbee_head struct{
	LockID,   	//协议中没描述的前缀，门锁
	Head,  		//包头
	Len,   		//包长 = 包类型 + 命令 + 参数 + 数据
	Type,  		//包类型
	Cmd uint8;   //命令
	Param uint8; //参数	
};

const SRV_2_LOCK = 0x80;
const NEED_VERIFY = 0x03;
const NO_VERIFY = 0x0D;
const NO_VERIFY_VAL = 0x0;
type Zigbee_tail struct{
	Tail,  //包尾
	Verify uint8;//检验
};

type CMD interface{
	Decode(buf []byte)(cmd interface{});
	Encode()(pack []byte, err error);
	New()(cmd interface{});
};

type Command struct{
	buf  []byte;    //全包
	head *Zigbee_head; //头指针
	tail *Zigbee_tail; //尾指针
	data []byte;  //真实数据
	data_len int;  //真实数据大小
};

func (this *Command)New()(cmd interface{}){
	this.buf = make([]byte, 255);//制作空间
	this.head = (*Zigbee_head)(unsafe.Pointer(&this.buf[0])); //设置头指针
	this.data = ([]byte)(this.buf[unsafe.Sizeof(Zigbee_head{}):]); //设置数据指针
	this.head.Head = PRE_FIX; //协议固定值 0x02

	return this;
}

func (this *Command)Encode()(pack []byte, err error){
	//tail的开始位置需要加上data长度
	tail_begin := (int)(unsafe.Sizeof(Zigbee_head{})) + this.data_len;
	this.tail = (*Zigbee_tail)(unsafe.Pointer(&this.buf[tail_begin]));//设置尾指针
	//todo 校验
	this.tail.Tail = NO_VERIFY;
	this.tail.Verify = NO_VERIFY_VAL;

	total_len := tail_begin + 2; //全包大小
    //包长 = 全包长 - 锁号 - 包头 - 包长 - 包尾 - 校验
	this.head.Len = (uint8)(total_len - 1 - 1 -1 -1 -1);

	this.buf = append([]byte(nil), this.buf[:total_len]...); //自适应包大小
	return this.buf, nil;
}

func (this *Command)Add_data(dt interface{})(err error){
	data := dt.([]byte);
	copy(this.data[this.data_len:], data);
	this.data_len += len(data);
	return nil;
}

func (this *Command)Decode(pack []byte)(cmd interface{}){
	this.buf = append([]byte(nil), pack[:]...); //设置缓冲包数据
	this.head = (*Zigbee_head)(unsafe.Pointer(&this.buf[0]));//设置头指针
	this.data = ([]byte)(this.buf[unsafe.Sizeof(Zigbee_head{}):]); //设置数据指针

	//锁号 + 包头 + 包长 + 包长所说的大小
	tail_begin := 1 + 1 + 1 + this.head.Len; 
	if (int)(tail_begin) != len(this.buf) - 2 {
		fmt.Println("包长不正确");
		return nil;
	}
	this.tail = (*Zigbee_tail)(unsafe.Pointer(&this.buf[tail_begin]));//设置尾指针
	
	//包长所说的大小 - 包类型 - 命令 - 参数
	this.data_len = (int)(this.head.Len) - 1 - 1 - 1;
	return *this;
}