package main

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"fmt"
	"log"
	// "strings"
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
  token = get_token();
  get_srv();
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

//https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=ACCESS_TOKEN

func get_srv(){
    client := &http.Client{};

	req, err := http.NewRequest("GET",
		"https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token="+token,
		nil);

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
