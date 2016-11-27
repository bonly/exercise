/*
调用助通的短信接口
http://www.ztsms.cn/sendXSms.do?username=haha&password=888888&mobile=13900000000&content=test&dstime=&productid=61341&xh=
*/

package main

import (
	"net/http"
	"net/url"
	"log"
	"fmt"
	"io/ioutil"
	"os"
	"io"
)

type Sms struct{
	username,
	password,
	to,
	text,
	subid,
	msgtype string;
};

const srv_addr string = "http://www.ztsms.cn/sendXSms.do?";

func main(){
	param := url.Values{};
	param.Add("username", "aibaidekeji");
	param.Add("password", "Abd12345");
	// param.Add("to", "15360534225");
	param.Add("mobile", os.Args[1]);	
	// param.Add("text", "test from bonly");
	param.Add("content", "【Xbed】互联网住宿:尊敬的张三客户，欢迎加入八桂快捷酒店会员，初始密码：33134，请尽快修改你的密码。");
	param.Add("dstime", "");
	param.Add("productid","456456"); //676767:优质验证码专用 887362:优持通知专用5 435227：电商会员营销4 456456:优质固定签名验证通知并发
	param.Add("xh","5501");
	
	//str := srv_addr+param.Encode()+"&username="+"jacy@bjhg"+"&password="+"23!xwwx8";
	str := srv_addr+param.Encode();
	log.Println("send: ", str);
	
	resp, err := http.Get(str);
	if err != nil{
		log.Println("http get: ", err);
		return;
	}
	defer resp.Body.Close();
	
	body, err := ioutil.ReadAll(resp.Body);
	if err != nil{
		log.Println("body: ", err);
		return;
	}
	
	fmt.Println(string(body));
	var ret_code int;
	var ret_msg string;
	_, err = fmt.Sscanf(string(body), "%d,%s", &ret_code, &ret_msg);
	if err != nil && err != io.EOF{
		log.Println(err);
		//return;
	}
	log.Println("ret: ", ret_code);
	log.Println("msg: ", ret_msg);
}