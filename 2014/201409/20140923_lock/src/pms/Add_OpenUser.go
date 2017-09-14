/*
auth: bonly
create: 2016.9.19
desc: 添加开门用户
*/
package pms

import (
"fmt"
"encoding/xml"
)

type Add_OpenUser_REQ struct{
	XMLName xml.Name `xml:"Command"`;
	Name string `xml:"name,attr"`;	
	MidId string;  //盒子编号 
	LockId string;  //锁号
	CardType string;  //开门用户类型，MF卡：1；身份证Id：2；密码：3
	CardData string;  //开门数据，密码最长为10位
	BeginTime string; //有效期的开始时间
	EndTime string; //有效期的结束时间
};

func (this *Add_OpenUser_REQ)New()(cmd interface{}){
	this.Name = "Add_OpenUser_REQ";
	cmd = PMS_Command{Data:this, Name:&this.Name};
	return cmd;
}

func (this *Add_OpenUser_REQ)Encode()([]byte, error){
	return xml.MarshalIndent(this, " ", " ");
}

func (this *Add_OpenUser_REQ)Decode(pack []byte)(error){
	return xml.Unmarshal(pack, this);
}

func (this *Add_OpenUser_REQ)Process(){
	fmt.Printf("数据： %v\n", this);
	//todo 入库
	//读取盒子编号
	//读取门锁号
	//检查类型，只支持密码方式
	//密码转换
	//开始时间转换
	//结束时间转换
	//查找对应的box，把转换操作函数调用起来
}

func (this *Add_OpenUser_REQ)Tran(){
	
}

type Add_OpenUser_RESP struct{
	XMLName xml.Name `xml:"Command"`;
	Name string `xml:"name,attr"`;
	ResultID int;
	Description string;
};

func (this *Add_OpenUser_RESP)New()(cmd interface{}){
	this.Name = "Add_OpenUser_RESP";
	cmd = PMS_Command{Data:this, Name:&this.Name};
	return cmd;
}

func (this *Add_OpenUser_RESP)Encode()([]byte, error){
	return xml.MarshalIndent(this, " ", " ");
}