/*
auth: bonly
create: 2016.9.19
desc: 添加开门用户
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

type Add_OpenUser_REQ struct{
	Name string ;	
	MidId string;  //盒子编号 
	LockId string;  //锁号
	CardType string;  //开门用户类型，MF卡：1；身份证Id：2；密码：3
	CardData string;  //开门数据，密码最长为10位
	BeginTime string; //有效期的开始时间
	EndTime string; //有效期的结束时间
	cmd proto.User_add_passwd;
};

func (this *Add_OpenUser_REQ)New()(cmd interface{}){
	this.Name = "Add_OpenUser_REQ";
	cmd = PMS_Command{Data:this, Name:&this.Name};
	return cmd;
}

func (this *Add_OpenUser_REQ)Encode()([]byte, error){
	return json.MarshalIndent(this, " ", " ");
}

func (this *Add_OpenUser_REQ)Decode(pack []byte)(error){
	return json.Unmarshal(pack, this);
}

func (this *Add_OpenUser_REQ)Process(result chan interface{}){
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

	//检查类型，只支持密码方式
	if this.CardType != "3"{
		return; 
	}

	//密码转换
	pwd_len := len(this.CardData);
	if pwd_len > 11{
		return; //最大只能是11位
	}
	for idx := 0; idx < pwd_len; idx++{
		this.cmd.Passwd[idx] = this.CardData[idx] << 4 >> 4;
	}

	//开始时间转换
	tn, err := time.Parse("2006-01-02 15:04:05", this.BeginTime);
	if err != nil{
		return;
	}
	this.cmd.Period[0] = (uint8)(tn.Year() - 2000);
	this.cmd.Period[1] = (uint8)(tn.Month());
	this.cmd.Period[2] = (uint8)(tn.Day());
	this.cmd.Period[3] = (uint8)(tn.Hour());
	this.cmd.Period[4] = (uint8)(tn.Minute());
	this.cmd.Period[5] = (uint8)(tn.Second());

	//结束时间转换
	tn_end, err := time.Parse("2006-01-02 15:04:05", this.EndTime);
	if err != nil{
		return;
	}	
	this.cmd.Period[6] = (uint8)(tn_end.Year() - 2000);
	this.cmd.Period[7] = (uint8)(tn_end.Month());
	this.cmd.Period[8] = (uint8)(tn_end.Day());
	this.cmd.Period[9] = (uint8)(tn_end.Hour());
	this.cmd.Period[10] = (uint8)(tn_end.Minute());
	this.cmd.Period[11] = (uint8)(tn_end.Second());

	idx := (uint32)(box);
	//查找对应的box，把转换操作函数调用起来
	if val, ok := manage.Works().Work_map[idx]; ok{
		fmt.Printf("对象盒子[%s]在线\n", this.MidId);
		fmt.Printf("box: %v\n", val);
		//把调用传递进去任务中 val.Task 

		this.cmd.Encode();
		manage.Works().Lock();
		manage.Works().Work_map[idx].Task <- this.Tran;
		manage.Works().Work_map[idx].Result = result;
		manage.Works().Unlock();
	}else{
		fmt.Printf("对象盒子[%s]不在线%X\n", this.MidId, idx);
		return;
	}
}

func (this *Add_OpenUser_REQ)Tran()(pack []byte, size int, cmd_name []byte){
	return	this.cmd.Buf(), this.cmd.Total_len(), this.cmd.Head().Get_CMD_PARAM();
}

type Add_OpenUser_RESP struct{
	Name string;
	ResultID string;
	Description string;
};

func (this *Add_OpenUser_RESP)New()(cmd interface{}){
	this.Name = "Add_OpenUser_RESP";
	cmd = PMS_Command{Data:this, Name:&this.Name};
	return cmd;
}

func (this *Add_OpenUser_RESP)Encode()([]byte, error){
	return json.MarshalIndent(this, " ", " ");
}

func (this *Add_OpenUser_RESP)Decode(pack []byte)(error){
	return fmt.Errorf("未实现");
}

func (this *Add_OpenUser_RESP)Process(result chan interface{}){
}