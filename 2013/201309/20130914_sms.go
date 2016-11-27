/*
调用玄武的短信接口
*/

package main

import (
	"net/http"
	"net/url"
	"log"
	"fmt"
	"io/ioutil"
	"os"
)

type Sms struct{
	username,
	password,
	to,
	text,
	subid,
	msgtype string;
};

const srv_addr string = "http://211.147.239.62:9050/cgi-bin/sendsms?";
// const srv_addr string = "http://211.147.239.62:9080/cgi-bin/sendsms?";

func main(){
	param := url.Values{};
	// param.Add("username", "jacy@bjhg");
	// param.Add("password", "23!xwwx8");
	// param.Add("to", "15360534225");
	param.Add("to", os.Args[1]);	
	// param.Add("text", "test from bonly");
	param.Add("text", "room key:" + os.Args[2]);
	param.Add("msgtype", "1");
	
	//str := srv_addr+param.Encode()+"&username="+"jacy@bjhg"+"&password="+"23!xwwx8";
	str := srv_addr+param.Encode()+"&username="+"abd@abd"+"&password="+"3Xbed2048_";
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
	
	fmt.Println("return: ", string(body));
}