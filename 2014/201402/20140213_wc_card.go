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
  token = get_token();
  batch, err := get_card();
  if err != nil{
    return;
  }

  for _, ab := range batch.Card_id_list {
    log.Println("==============","get card info: ", ab, "====================");
    Get_card(ab);
    log.Println("==============", "end card info: ", ab, "==================="); 
  }
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


type Batchget struct{
  Offset int `json:"offset"`;
  Count int `json:"count"`;
  Status_list []string  `json:"status_list"`;
};

type RetBatchget struct{
  ErrCode int `json:"errcode"`;
  ErrMsg string `json:"errmsg"`;
  Card_id_list []string `json:"card_id_list"`;
};

func get_card()(ret RetBatchget, err error){
  ps := Batchget{
    Offset:0,
    Count:10,
  };

  ps.Status_list = append(ps.Status_list, "CARD_STATUS_VERIFY_OK");
  ps.Status_list = append(ps.Status_list, "CARD_STATUS_DISPATCH");
  js, _ := json.Marshal(ps);

  client := &http.Client{};

	req, err := http.NewRequest("POST",
		"https://api.weixin.qq.com/card/batchget?access_token="+token,
		strings.NewReader(string(js)));

	req.Header.Set("Content-Type","application/x-www-form-urlencoded");
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");
	resp, err := client.Do(req);
	if err != nil{
		log.Println("send: ", err);
		return ret, err;
	}	
  defer resp.Body.Close();
    
	body, err := ioutil.ReadAll(resp.Body);
  if err != nil {
      log.Println("body: ", err);
      return ret, err;
  }
 
  fmt.Println("recv: ", string(body));	
  err = json.Unmarshal(body, &ret);
  if err != nil{
    log.Println("ret err: ", err);
    return ret, err;
  }
  return ret, nil;
}

func Get_card(card string){
  ps := fmt.Sprintf(`{"card_id":"%s"}`, card);

  client := &http.Client{};

  req, err := http.NewRequest("POST",
    "https://api.weixin.qq.com/card/get?access_token="+token,
    strings.NewReader(ps));

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
 
  fmt.Println("recv: ", string(body));  
}