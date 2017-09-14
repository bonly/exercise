/*
auth: bonly
create: 2016.9.27
desc: 门锁主动上传的开门记录
*/
package oms

import (
"fmt"
"encoding/json"
"proto"
// "strconv"
// "manage"
// "time"
// log "glog"
)

type Lock_log_REQ struct{
	Name string ;	
	MidId string;    //盒子编号 
	LockId string;   //锁号
	Param string;    //参数
	DateTime string; //操作时间
	Cmd proto.Lock_upload_log `json:"-"`;
};

func (this *Lock_log_REQ)New()(cmd interface{}){
	this.Name = "Lock_log_REQ";
	cmd = PMS_Command{Data:this, Name:&this.Name};
	return cmd;
}

func (this *Lock_log_REQ)Encode()([]byte, error){
	return json.MarshalIndent(this, " ", " ");
}

func (this *Lock_log_REQ)Decode(pack []byte)(error){
	return json.Unmarshal(pack, this);
}

func (this *Lock_log_REQ)Process(result chan interface{}){
}

// func (this *Lock_log_REQ)Tran()(pack []byte, size int, cmd_name []byte){
// 	return	this.cmd.Buf(), this.cmd.Total_len(), this.cmd.Head().Get_CMD_PARAM();
// }

type Lock_log_RESP struct{
	Name string;
	ResultID string;
	Description string;
};

func (this *Lock_log_RESP)New()(cmd interface{}){
	this.Name = "Lock_log_RESP";
	cmd = PMS_Command{Data:this, Name:&this.Name};
	return cmd;
}

func (this *Lock_log_RESP)Encode()([]byte, error){
	return json.MarshalIndent(this, " ", " ");
}

func (this *Lock_log_RESP)Decode(pack []byte)(error){
	return fmt.Errorf("未实现");
}

func (this *Lock_log_RESP)Process(result chan interface{}){
}