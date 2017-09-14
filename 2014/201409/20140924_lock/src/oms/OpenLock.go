/*
auth: bonly
create: 2016.9.20
desc: 远程开门
*/
package oms

import (
"fmt"
"encoding/xml"
"proto"
"strconv"
"manage"
)

type OpenLock_REQ struct{
	XMLName xml.Name `xml:"Command"`;
	Name string `xml:"name,attr"`;	
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
	return xml.MarshalIndent(this, " ", " ");
}

func (this *OpenLock_REQ)Decode(pack []byte)(error){
	return xml.Unmarshal(pack, this);
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
		manage.Works().Work_map[idx].Task <- this.Tran;
		val.Result = result;
		manage.Works().Work_map[idx] = val;
		manage.Works().Conn_map[val.Conn] = &val;
		fmt.Printf("修改过的%+X\n", manage.Works().Work_map[idx]);
	}else{
		fmt.Printf("对象盒子[%s]不在线%X\n", this.MidId, idx);
		return;
	}
}

func (this *OpenLock_REQ)Tran()(pack []byte, size int){
	return	this.cmd.Buf(), this.cmd.Total_len();
}

func (this *OpenLock_REQ)Res(){
	fmt.Printf("in res\n");
}

type OpenLock_RESP struct{
	XMLName xml.Name `xml:"Command"`;
	Name string `xml:"name,attr"`;
	ResultID int;
	Description string;
};

func (this *OpenLock_RESP)New()(cmd interface{}){
	this.Name = "OpenLock_RESP";
	cmd = PMS_Command{Data:this, Name:&this.Name};
	return cmd;
}

func (this *OpenLock_RESP)Encode()([]byte, error){
	return xml.MarshalIndent(this, " ", " ");
}
