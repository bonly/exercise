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
/*
xbed的微信开发服务配置
*/
const (
	appid = "wx0c49d2c7f9d36648";
	secret = "910fe488f3ff205b428905e2e1733a94";
)

var token = "G2Svkwf8qy9TwoMSB8GDB5-TSiOfCcxOFyaCrQPhMJmLgcSq9n9YEdb8VJ_2XCF2I-9L3cTQfxvQpIuPiwcnnFHulbR1Dqk0Sh_aTL64TyI";

func main(){
	// get_token();
	// get_menu();
  token = get_token();
  get_media();
	// set_menu();
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

var new_menu string=`{
"button": [
    {
        "type": "view",
        "name": "预订",
        "url": "https://open.weixin.qq.com/connect/oauth2/authorize?%s",
        "sub_button": []
    },
    {
         "name":"服务",
         "sub_button":[
         {  
             "type":"view",
             "name":"办理入住",
             "url":"http://wxi.xbed.com.cn/checkin"
          },
          {
            "type": "view",
            "name": "房间服务",
            "url": "https://open.weixin.qq.com/connect/oauth2/authorize?%s",
            "sub_button": []
          },
          {
            "type": "view",
            "name": "退房",
            "url": "http://wxi.xbed.com.cn/checkout",
            "sub_button": []
          }          
          ]
    },
    {
        "name": "更多",
        "sub_button": [
        {
            "type": "view",
            "name": "直接支付",
            "url": "http://devpay.xbed.com.cn/wechat/pay.php?lodger_id=0&product=大床房&order_no=113431&fee=100&cb_cancel=http://devpay.xbed.com.cn/wechat/pay_cancel.php&cb_ok=http://devpay.xbed.com.cn/wechat/pay_ok.php",
            "sub_button": []
        },        
        {
            "type": "view",
            "name": "支付",
            "url": "http://devpay.xbed.com.cn/wechat/auth.php",
            "sub_button": []
        },         
        {
            "type": "view",
            "name": "个人中心",
            "url": "https://open.weixin.qq.com/connect/oauth2/authorize?%s",
            "sub_button": []
        }, 
        {
            "type": "view",
            "name": "房间服务",
            "url": "http://devpay.xbed.com.cn:7002/views/user_myroom.html",
            "sub_button": []
        }, 
        {
            "type": "view",
            "name": "Xbed微商城",
            "url": "http://kdt.im/rYr1FDO4r",
            "sub_button": []
        }, 
	}
            "type": "click",
            "name": "客服热线",
            "key": "call400",
            "sub_button": []
        }
        ]
  ]
}
`;
        
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

func get_media(){
    str := `{"type":"image", "offset":0, "count":20}`;
    client := &http.Client{};
    req, err := http.NewRequest("POST","https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token="+token,
      strings.NewReader(str));
    req.Header.Set("Content-Type","application/x-www-form-urlencoded");
    req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");

    resp,err := client.Do(req);
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
  /* 旧代码
  room_srv := url.Values{};
  room_srv.Add("appid",appid);
  room_srv.Add("redirect_uri","http://wxi.xbed.com.cn/app/auth");
  room_srv.Add("response_type","code");
  room_srv.Add("scope","snsapi_base");
  room_srv.Add("state","http://wxi.xbed.com.cn:9091/index.html");

	ali_pay := url.Values{};
	ali_pay.Add("appid",appid);
	ali_pay.Add("redirect_uri","http://pay.xbed.com.cn/pay");
	ali_pay.Add("response_type","code");
	ali_pay.Add("scope","snsapi_base");
	ali_pay.Add("state","http://pay.xbed.com.cn/xbed/pay.html");
  */
  // adb := fmt.Sprintf(new_menu);
	// adb := fmt.Sprintf(new_menu, room_srv.Encode()+"#wechat_redirect", ali_pay.Encode()+"#wechat_redirect");

  order := url.Values{};
  order.Add("appid",appid);
  order.Add("redirect_uri","http://devpay.xbed.com.cn/wechat/order.php");//"http://devpay.xbed.com.cn:7002/views/index.html");
  order.Add("response_type","code");
  order.Add("scope","snsapi_base");
  order.Add("state","");

  room_srv := url.Values{};
  room_srv.Add("appid",appid);
  room_srv.Add("redirect_uri","http://devpay.xbed.com.cn/wechat/myroom.php");//"http://devpay.xbed.com.cn:7002/views/index.html");
  room_srv.Add("response_type","code");
  room_srv.Add("scope","snsapi_base");
  room_srv.Add("state","");

  user_info := url.Values{};
  user_info.Add("appid",appid);
  user_info.Add("redirect_uri","http://devpay.xbed.com.cn/wechat/user_center.php");//"http://devpay.xbed.com.cn:7002/views/index.html");
  user_info.Add("response_type","code");
  user_info.Add("scope","snsapi_base");
  user_info.Add("state","");

  adb := fmt.Sprintf(new_menu, 
    order.Encode()+"#wechat_redirect",
    room_srv.Encode()+"#wechat_redirect",
    user_info.Encode()+"#wechat_redirect");
  // fmt.Println("菜单: ",adb);

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
