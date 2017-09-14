/*
auth: bonly
create: 2016.9.21
desc: 删除开门用户
*/
package oms

import (
"fmt"
"encoding/json"
"proto"
"strconv"
"manage"
"time"
)

type Time_REQ struct{
	Name string ;	
	MidId string;  //盒子编号 
	LockId string;  //锁号
	Time string; //时间
	cmd proto.Setup_lock_time;
};

func (this *Time_REQ)New()(cmd interface{}){
	this.Name = "Time_REQ";
	cmd = PMS_Command{Data:this, Name:&this.Name};
	return cmd;
}

func (this *Time_REQ)Encode()([]byte, error){
	return json.MarshalIndent(this, " ", " ");
}

func (this *Time_REQ)Decode(pack []byte)(error){
	return json.Unmarshal(pack, this);
}

func (this *Time_REQ)Process(result chan interface{}){
	fmt.Printf("数据： %v\n", this);
	//todo 入库
	this.cmd.New();

	//参数检查
	//读取门锁号	
	lock, err := strconv.ParseUint(this.LockId, 10, 8);
	if err != nil{
		return; //todo 错误处理
	}
	this.cmd.Command.Head().LockID = (uint8)(lock);	

	//读取盒子编号	
	box, err := strconv.ParseUint(this.MidId, 16, 32);
	if err != nil{
		return; //todo 错误处理
	}	

	//时间转换
	var tn time.Time;
	if len(this.Time) > 0{
		tn, err = time.Parse("2006-01-02 15:04:05", this.Time);
		if err != nil{
			return;
		}
	}else{
		tn = time.Now();
	}

	this.cmd.DateTime[0] = (uint8)(tn.Year() - 2000);
	this.cmd.DateTime[1] = (uint8)(tn.Month());
	this.cmd.DateTime[2] = (uint8)(tn.Day());
	this.cmd.DateTime[3] = (uint8)(tn.Hour());
	this.cmd.DateTime[4] = (uint8)(tn.Minute());
	this.cmd.DateTime[5] = (uint8)(tn.Second());

	idx := (uint32)(box);
	//查找对应的box，把转换操作函数调用起来
	if val, ok := manage.Works().Work_map[idx]; ok{
		fmt.Printf("对象盒子[%s]在线\n", this.MidId);
		fmt.Printf("box: %v\n", val);
		//把调用传递进去任务中 val.Task 

		this.cmd.Encode();
		manage.Works().Work_map[idx].Task <- this.Tran;
		// if manage.Works().Work_map[idx].Result != nil{
		// 	close(manage.Works().Work_map[idx].Result); //关闭旧通道
		// } //to think: 新旧两个指令，无法判断，由发起方来关闭通道
		manage.Works().Work_map[idx].Result = result;
	}else{
		fmt.Printf("对象盒子[%s]不在线%X\n", this.MidId, idx);
		return;
	}
}

func (this *Time_REQ)Tran()(pack []byte, size int, cmd_name []byte){
	return	this.cmd.Buf(), this.cmd.Total_len(), this.cmd.Head().Get_CMD_PARAM();
}

type Time_RESP struct{
	Name string;
	ResultID string;
	Description string;
};

func (this *Time_RESP)New()(cmd interface{}){
	this.Name = "Time_RESP";
	cmd = PMS_Command{Data:this, Name:&this.Name};
	return cmd;
}

func (this *Time_RESP)Encode()([]byte, error){
	return json.MarshalIndent(this, " ", " ");
}

func (this *Time_RESP)Decode(pack []byte)(error){
	return fmt.Errorf("未实现");
}

func (this *Time_RESP)Process(result chan interface{}){
}