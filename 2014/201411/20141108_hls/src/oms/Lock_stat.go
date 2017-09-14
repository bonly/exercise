/*
auth: bonly
create: 2016.9.28
desc: 门锁状态查询
*/
package oms

import (
"fmt"
"encoding/json"
"proto"
"strconv"
"manage"
log "glog"
)

type Lock_stat_REQ struct{
	Name string;	
	MidId string;  //盒子编号 
	LockId string;  //锁号
	cmd proto.Lock_stat;
};

func (this *Lock_stat_REQ)New()(cmd interface{}){
	this.Name = "Lock_stat_REQ";
	cmd = PMS_Command{Data:this, Name:&this.Name};
	return cmd;
}

func (this *Lock_stat_REQ)Encode()([]byte, error){
	return json.MarshalIndent(this, " ", " ");
}

func (this *Lock_stat_REQ)Decode(pack []byte)(error){
	return json.Unmarshal(pack, this);
}

func (this *Lock_stat_REQ)Process(result chan interface{}){
	log.V(99).Infof("数据： %v\n", this);
	
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
	// manage.Works().Lock();
	// defer manage.Works().Unlock();	
	if val, ok := manage.Works().Work_map[idx]; ok{
		defer func(){
			if err := recover(); err != nil{
				log.Warningf("操作盒子失败, %v\n", err);
				manage.Works().Lost(manage.Works().Work_map[idx].Conn);
			}
		}();		
		log.Infof("对象盒子[%s]在线\n", this.MidId);
		log.V(99).Infof("box: %v\n", val);
		//把调用传递进去任务中 val.Task 

		this.cmd.Encode();
		manage.Works().Work_map[idx].Task <- this.Tran;
		manage.Works().Work_map[idx].Result = result;
	}else{
		log.Warningf("对象盒子[%s]不在线%X\n", this.MidId, idx);
		return;
	}
}

func (this *Lock_stat_REQ)Tran()(pack []byte, size int, cmd_name []byte){
	return	this.cmd.Buf(), this.cmd.Total_len(), this.cmd.Head().Get_CMD_PARAM();
}

type Lock_stat_RESP struct{
	Name string;
	ResultID string;
	Description string;
	Card_user_count string;
	Passwd_user_count string;
	Power string;
	Now string;
	Authority string;
	cmd proto.R_Lock_stat;
};

func (this *Lock_stat_RESP)New()(cmd interface{}){
	this.Name = "Lock_stat_RESP";
	cmd = PMS_Command{Data:this, Name:&this.Name};
	return cmd;
}

func (this *Lock_stat_RESP)Encode()([]byte, error){
	return json.MarshalIndent(this, " ", " ");
}

func (this *Lock_stat_RESP)Decode(pack []byte)(error){
	return fmt.Errorf("未实现");
}

func (this *Lock_stat_RESP)Process(result chan interface{}){
}