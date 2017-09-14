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
	Len,   		//包长 = 包类型 + 命令 + 参数 + 数据 (不包括尾部和校验位)
	Type,  		//包类型
	Cmd uint8;   //命令
	Param uint8; //参数	
};

func (this *Zigbee_head) Get_CMD_PARAM()(ret []byte){
	ret = []byte{this.Cmd, this.Param};
	return ret;
}

const CMD_MAX_LEN = 256;

const SRV_2_LOCK = 0x80;

const NEED_VERIFY = 0x03;
const NO_VERIFY = 0x0D;
const NO_VERIFY_VAL = 0x0;

const YES = 0x46;
const NO = 0x4F;

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
	total_len int; //全包大小
	verify bool; //是否校验
};

func (this *Command)Head()(*Zigbee_head){
	return this.head;
}

func (this *Command)Buf()([]byte){
	return this.buf;
}

func (this *Command)Verify(val bool){
	this.verify = val;
}

func (this *Command)New()(cmd interface{}){
	this.buf = make([]byte, CMD_MAX_LEN);//制作空间
	this.head = (*Zigbee_head)(unsafe.Pointer(&this.buf[0])); //设置头指针
	this.data = ([]byte)(this.buf[unsafe.Sizeof(Zigbee_head{}):]); //设置数据指针
	this.head.Head = PRE_FIX; //协议固定值 0x02
	this.head.LockID = 0x0;
	this.verify = true;

	return this;
}

func (this *Command)Encode()(pack []byte, err error){
	if this.head.LockID == 0x0{
		return nil, fmt.Errorf("需要设置门锁ID再压包!");
	}
	//tail的开始位置需要加上data长度
	tail_begin := (int)(unsafe.Sizeof(Zigbee_head{})) + this.data_len;
	this.tail = (*Zigbee_tail)(unsafe.Pointer(&this.buf[tail_begin]));//设置尾指针

	this.total_len = tail_begin + 1; //全包大小,未计算是否要校验位

	//校验
	if this.verify == false{
		this.tail.Tail = NO_VERIFY
		this.tail.Verify = NO_VERIFY_VAL;
		
	}else{
		this.total_len += 1; //全包大小，加上校验位
		this.Crypt();
	}

    //包长 = 数据长 + 包类型 + 命令 + 参数
	this.head.Len = (uint8)(this.data_len + 1 + 1 + 1);

	this.buf = append([]byte(nil), this.buf[:this.total_len]...); //自适应包大小
	return this.buf, nil;
}

func (this *Command)Add_data(dt interface{})(err error){
	data := dt.([]byte);
	copy(this.data[this.data_len:], data);
	this.data_len += len(data);
	return nil;
}

func (this *Command)Decode(pack []byte)(cmd interface{}){
	this.head = (*Zigbee_head)(unsafe.Pointer(&pack[0]));//设置在源数据中的头指针	
	//包长所说的大小 - 包类型 - 命令 - 参数
	this.data_len = (int)(this.head.Len) - 1 - 1 - 1;//计算真实数据长度

	tail_begin :=  (int)(unsafe.Sizeof(Zigbee_head{})) + this.data_len; //计算尾指针偏移量
	this.tail = (*Zigbee_tail)(unsafe.Pointer(&pack[tail_begin]));//设置源数据中的尾指针

	tail_len := 1; //根据是否校验来纠正尾包长度
	if this.tail.Tail == NEED_VERIFY {
		tail_len = 2;
	}

	//包长所说的大小 + 尾长 + 锁号 + 包头 + 包长
	this.total_len = (int)(this.head.Len) + tail_len + 1 + 1 + 1;	 //全包大小
	this.buf = append([]byte(nil), pack[:this.total_len]...); //设置缓冲包数据	

	this.head = (*Zigbee_head)(unsafe.Pointer(&this.buf[0]));//设置头指针
	this.data = ([]byte)(this.buf[unsafe.Sizeof(Zigbee_head{}):]); //设置数据指针
	this.tail = (*Zigbee_tail)(unsafe.Pointer(&this.buf[tail_begin]));//设置正确的尾指针	

	return *this;
}

func (this *Command)Crypt(){
	this.tail.Tail = NEED_VERIFY;
	this.tail.Verify = this.buf[1];  //不包括锁地址
	for idx := 2; idx < this.total_len - 2; idx++{ //不包括校验位本身
		this.tail.Verify = this.tail.Verify ^ this.buf[idx]; 
	}
}

