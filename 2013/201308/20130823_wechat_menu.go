package main

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"fmt"
	"log"
	"strings"
  "encoding/json"
)

const (
	appid = "wx100949d5a719dac2";
	secret = "10b624d9c9faa83165edfc1d4a336935";
)

var token = "G2Svkwf8qy9TwoMSB8GDB5-TSiOfCcxOFyaCrQPhMJmLgcSq9n9YEdb8VJ_2XCF2I-9L3cTQfxvQpIuPiwcnnFHulbR1Dqk0Sh_aTL64TyI";

func main(){
	// get_token();
	// get_menu();
  token = get_token();
	set_menu();
}

type R_Token struct {
Access_token string `json:"access_token"`;
Expires_in int `json:"expires_in"`;
Refresh_token string `json:"refresh_token"`;
Openid string `json:"openid"`;
Scope string `json:"scope"`;
};


func get_token()(ret string){
	val := url.Values{};
	val.Add("grant_type", "client_credential");
	val.Add("appid", appid);
	val.Add("secret", secret);
	addr := "https://api.weixin.qq.com/cgi-bin/token?"+val.Encode();

  resp, err := http.Get(addr);
  if err != nil {
      log.Println("http get: ",err);
  }
  defer resp.Body.Close();
    
	body, err := ioutil.ReadAll(resp.Body);
  if err != nil {
      log.Println("body: ", err);
  }
 	fmt.Println("send: ", addr);
  fmt.Println("recv: ", string(body));	

  var token R_Token;
  if err := json.Unmarshal(body,&token); err != nil{
    log.Println("token: ", err);
  }   
  return token.Access_token;
}

const org_menu=`{
  "button": [
    {
      "type": "view",
      "name": "Xbed",
      "url": "http://wx.Xbed.com.cn",
      "sub_button": []
    },
    {
      "name": "预约",
      "sub_button": [
        {
          "type": "view",
          "name": "支付订单",
          "url": "http://wx.Xbed.com.cn/order",
          "sub_button": []
        },
        {
          "type": "view",
          "name": "自助入住",
          "url": "http://wx.Xbed.com.cn/checkin",
          "sub_button": []
        },
        {
          "type": "view",
          "name": "预约离店",
          "url": "http://wx.Xbed.com.cn/checkout",
          "sub_button": []
        }
      ]
    }
  ]
}`;

var new_menu string=`{
  "button": [
    {
      "name":"我的房间",
      "sub_button":[
        {
          "type": "view",
          "name": "自助入住",
          "url": "http://wx.Xbed.com.cn/checkin",
          "sub_button": []
        },
        {
          "type": "view",
          "name": "办理退房",
          "url": "http://wx.Xbed.com.cn/checkout",
          "sub_button": []
        }        
      ]
    },
    {
      "name": "客房服务",
      "sub_button": [
        {
          "type": "view",
          "name": "我要打扫房间",
          "url": "https://open.weixin.qq.com/connect/oauth2/authorize?%s",
          "sub_button": []
        },
        {
          "type": "click",
          "name": "客服热线",
          "key": "call400",
          "sub_button": []
        }
      ]
    },
    {
      "type": "view",
      "name": "关于Xbed",
      "url": "http://wx.xbed.com.cn:9091/aboutus.html",
      "sub_button": []
    }    
  ]
}`;
/*
,
        {
          "type": "view",
          "name": "意见反馈",
          "url": "http://wx.Xbed.com.cn/checkout",
          "sub_button": []
        }
*/
        
func get_menu(){
    resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/menu/get?access_token="+token);
    if err != nil {
        log.Println("http get: ",err);
    }
    defer resp.Body.Close();
    
	body, err := ioutil.ReadAll(resp.Body);
    if err != nil {
        log.Println("body: ", err);
    }
 
    fmt.Println(string(body));		
}

func set_menu(){
	red := url.Values{};
	red.Add("appid",appid);
	red.Add("redirect_uri","http://wx.xbed.com.cn/app/auth");
	red.Add("response_type","code");
	red.Add("scope","snsapi_base");
	red.Add("state","http://wx.xbed.com.cn:9091/index.html");
	adb := fmt.Sprintf(new_menu, red.Encode()+"#wechat_redirect");

    client := &http.Client{};

	req, err := http.NewRequest("POST",
		"https://api.weixin.qq.com/cgi-bin/menu/create?access_token="+token,
		strings.NewReader(adb));

	req.Header.Set("Content-Type","application/x-www-form-urlencoded");
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");
	resp, err := client.Do(req);
	if err != nil{
		log.Println("send: ", err);
		return;
	}	
    defer resp.Body.Close();
    
	body, err := ioutil.ReadAll(resp.Body);
    if err != nil {
        log.Println("body: ", err);
        return;
    }
 
    fmt.Println(string(body));		
}