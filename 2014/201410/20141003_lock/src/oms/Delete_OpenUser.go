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
)

type Delete_OpenUser_REQ struct{
	Name string ;	
	MidId string;  //盒子编号 
	LockId string;  //锁号
	CardType string;  //开门用户类型，MF卡：1；身份证Id：2；密码：3
	CardData string;  //开门数据，密码最长为10位
	cmd proto.User_del_passwd;
};

func (this *Delete_OpenUser_REQ)New()(cmd interface{}){
	this.Name = "Delete_OpenUser_REQ";
	cmd = PMS_Command{Data:this, Name:&this.Name};
	return cmd;
}

func (this *Delete_OpenUser_REQ)Encode()([]byte, error){
	return json.MarshalIndent(this, " ", " ");
}

func (this *Delete_OpenUser_REQ)Decode(pack []byte)(error){
	return json.Unmarshal(pack, this);
}

func (this *Delete_OpenUser_REQ)Process(result chan interface{}){
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

	idx := (uint32)(box);
	//查找对应的box，把转换操作函数调用起来
	// manage.Works().Lock();
	// defer manage.Works().Unlock();	
	if val, ok := manage.Works().Work_map[idx]; ok{
		defer func(){
			if err := recover(); err != nil{
				fmt.Printf("操作盒子失败, %v\n", err);
				manage.Works().Lost(manage.Works().Work_map[idx].Conn);
			}
		}();	
		fmt.Printf("对象盒子[%s]在线\n", this.MidId);
		fmt.Printf("box: %v\n", val);
		//把调用传递进去任务中 val.Task 

		this.cmd.Encode();
		manage.Works().Work_map[idx].Task <- this.Tran;
		manage.Works().Work_map[idx].Result = result;
	}else{
		fmt.Printf("对象盒子[%s]不在线%X\n", this.MidId, idx);
		return;
	}
}

func (this *Delete_OpenUser_REQ)Tran()(pack []byte, size int, cmd_name []byte){
	return	this.cmd.Buf(), this.cmd.Total_len(), this.cmd.Head().Get_CMD_PARAM();
}

type Delete_OpenUser_RESP struct{
	Name string;
	ResultID string;
	Description string;
};

func (this *Delete_OpenUser_RESP)New()(cmd interface{}){
	this.Name = "Delete_OpenUser_RESP";
	cmd = PMS_Command{Data:this, Name:&this.Name};
	return cmd;
}

func (this *Delete_OpenUser_RESP)Encode()([]byte, error){
	return json.MarshalIndent(this, " ", " ");
}

func (this *Delete_OpenUser_RESP)Decode(pack []byte)(error){
	return fmt.Errorf("未实现");
}

func (this *Delete_OpenUser_RESP)Process(result chan interface{}){
}