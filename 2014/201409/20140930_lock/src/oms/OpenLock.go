/*
auth: bonly
create: 2016.9.20
desc: 远程开门
*/
package oms

import (
"fmt"
"encoding/json"
"proto"
"strconv"
"manage"
)

type OpenLock_REQ struct{
	Name string;	
	MidId string;  //盒子编号 
	LockId string;  //锁号
	cmd proto.Remote_open_door;
};

func (this *OpenLock_REQ)New()(cmd interface{}){
	this.Name = "OpenLock_REQ";
	cmd = PMS_Command{Data:this, Name:&this.Name};
	return cmd;
}

func (this *OpenLock_REQ)Encode()([]byte, error){
	return json.MarshalIndent(this, " ", " ");
}

func (this *OpenLock_REQ)Decode(pack []byte)(error){
	return json.Unmarshal(pack, this);
}

func (this *OpenLock_REQ)Process(result chan interface{}){
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

func (this *OpenLock_REQ)Tran()(pack []byte, size int, cmd_name []byte){
	return	this.cmd.Buf(), this.cmd.Total_len(), this.cmd.Head().Get_CMD_PARAM();
}

type OpenLock_RESP struct{
	Name string;
	ResultID string;
	Description string;
	cmd proto.R_Remote_open_door;
};

func (this *OpenLock_RESP)New()(cmd interface{}){
	this.Name = "OpenLock_RESP";
	cmd = PMS_Command{Data:this, Name:&this.Name};
	return cmd;
}

func (this *OpenLock_RESP)Encode()([]byte, error){
	return json.MarshalIndent(this, " ", " ");
}

func (this *OpenLock_RESP)Decode(pack []byte)(error){
	return fmt.Errorf("未实现");
}

func (this *OpenLock_RESP)Process(result chan interface{}){
}