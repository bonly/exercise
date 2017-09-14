/*
auth: bonly
create: 2016.9.15
desc: 门锁状态查询
*/
package proto

import(
"fmt"
"unsafe"
);

var LOCK_STAT_CMD_PARAM = []byte{0x91, 0x30};

type Lock_stat struct{
	Command;
};

func (this *Lock_stat)New()(cmd interface{}){
	this.Command.New();
	this.head.Type = SRV_2_LOCK;
	this.head.Cmd = LOCK_STAT_CMD_PARAM[0];
	this.head.Param = LOCK_STAT_CMD_PARAM[1];

	return this;
}

func (this *Lock_stat)Encode()(pack []byte, err error){
	return this.Command.Encode();;
}

func (this *Lock_stat)Decode(pack []byte)(cmd interface{}){
	return nil;
}

/*
1个字节.
BIT2:门磁信号 =0 关门；=1 开门
BIT1:反锁信号：=0 反锁；=1 未反锁
BIT0:锁舌信号：=0 开锁；=1 上锁
*/

type R_Lock_stat_data struct{
	Card_user_count uint8;  //卡用户数
	Passwd_user_count uint8; //密码用户数
	Power [2]uint8; //电池电压
	Now_date [3]uint8; //门锁日期  年-月-日(10进制) 
	Now_time [3]uint8; //门锁时间 时-分-秒(BCD码)
	Sensor uint8; //传感器状态  
	Authority uint8; //权限  30:开门有效 31：禁止一切开门方式
};

type R_Lock_stat struct{
	Command;
	Data *R_Lock_stat_data;
};

func (this *R_Lock_stat)Encode()(pack []byte, err error){
	return nil, fmt.Errorf("未实现");
}

func (this *R_Lock_stat)Decode(pack []byte)(cmd interface{}){
	cmd = this.Command.Decode(pack);
	if cmd == nil{
		return nil;
	}
	fmt.Printf("收到门锁状态应答:\n");
	Hex_Dump(this.buf, len(this.buf));

	this.Data = (*R_Lock_stat_data)(unsafe.Pointer(&this.data[0]));
	fmt.Printf("头数据head: %+v\n", *(this.head));
	fmt.Printf("体数据body: %+v\n", *(this.Data));
	fmt.Printf("尾数据tail: %+v\n", *(this.tail));
	return *this;
}