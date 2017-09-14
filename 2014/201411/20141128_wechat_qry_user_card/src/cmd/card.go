package main

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"fmt"
	log "glog"
	"strings"
  "encoding/json"
  "flag"
  "os"
  "bufio"
)

var token = "G2Svkwf8qy9TwoMSB8GDB5-TSiOfCcxOFyaCrQPhMJmLgcSq9n9YEdb8VJ_2XCF2I-9L3cTQfxvQpIuPiwcnnFHulbR1Dqk0Sh_aTL64TyI";

func init(){

}

/*
select open_id, card_id from xbed_marketing.xb_user_coupon_cards where coupon_card_source='wechat';
*/
var file_name = flag.String("in", "user.txt", "客户的openid");
var sql_name = flag.String("out", "card.sql", "输出sql");
var appid = flag.String("a", "wx100949d5a719dac2", "appid");
var secret = flag.String("s", "10b624d9c9faa83165edfc1d4a336935", "secret");

func main(){
  flag.Parse();
  flag.Set("alsologtostderr", "true");  
  flag.Set("v", "99");  

  fl, err := os.Open(*file_name);
  if err != nil{
    log.Info(err);
    return;
  }
  defer fl.Close();

  fo, err := os.OpenFile(*sql_name, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666);
  if err != nil{
    log.Info(err);
    return;
  }
  defer fo.Close();

  token = get_token();

  scanner := bufio.NewScanner(fl);
  for scanner.Scan(){
    var openid, card_id string;
    _, err = fmt.Sscanf(scanner.Text(), "%s %s", &openid, &card_id);
    if err != nil{
      continue;
    }
    err, _ := get_card(openid, card_id);
    if err == nil{
      // if err, valid, desc := process_card(&card, code); err == nil && valid == false{
      //   dec := fmt.Sprintf("-- %s\r\n", desc);
      //   sql := fmt.Sprintf("update xbed_marketing.xb_user_coupon_cards set stat=1 where code='%s';\r\n\r\n", code);
      //   fo.WriteString(dec);
      //   fo.WriteString(sql);
      // }
    }
  }
  if err := scanner.Err(); err != nil{
    log.Info(err);
    return;
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
	val.Add("appid", *appid);
	val.Add("secret", *secret);
	addr := "https://api.weixin.qq.com/cgi-bin/token?"+val.Encode();

  resp, err := http.Get(addr);
  if err != nil {
      log.Info("http get: ",err);
  }
  defer resp.Body.Close();
    
	body, err := ioutil.ReadAll(resp.Body);
  if err != nil {
      log.Info("body: ", err);
  }
 	log.V(99).Info("send: ", addr);
  log.V(99).Info("recv: ", string(body));	

  var token R_Token;
  if err := json.Unmarshal(body,&token); err != nil{
    log.Info("token: ", err);
  }   
  return token.Access_token;
}

type T_Card struct{
  Card_id string `json:"card_id"`;
  Code string `json:"code"`;
};

type R_Card struct{
  Errcode int `json:"errcode"`;
  Errmsg string `json:"errmsg"`;
  Cards []T_Card `json:"card_list"`;
};

func get_card(openid string, card_id string) (err error, ret R_Card){
  var qry_card string;
  if len(card_id) != '0' {
	  qry_card = fmt.Sprintf(`
	  {
	    "openid":"%s",
	    "card_id":"%s"
	  }
	  `, openid, card_id);
  }else{
	  qry_card = fmt.Sprintf(`
	  {
	    "openid":"%s"
	  }
	  `, openid);  	
  }
  client := &http.Client{};

	req, err := http.NewRequest("POST",
		"https://api.weixin.qq.com/card/user/getcardlist?access_token="+token,
		strings.NewReader(qry_card));

	req.Header.Set("Content-Type","application/x-www-form-urlencoded");
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");
	resp, err := client.Do(req);
	if err != nil{
		log.Info("send: ", err);
		return err, ret;
	}	
  defer resp.Body.Close();
    
	body, err := ioutil.ReadAll(resp.Body);
    if err != nil {
        log.Info("body: ", err);
        return err, ret;
  }
 
  log.V(99).Info(string(body));	
  if err = json.Unmarshal(body, &ret);err != nil{
    return err, ret;
  } 
  return nil, ret; 	
}

// func process_card(card *R_Card, code string)(err error, valid bool, desc string){
//   if card.Errcode != 0{ //不成功的直接结束
//     desc = fmt.Sprintf("卡券[%s]在微信检查失败: %#v", code, card);
//     log.Info(desc);
//     return nil, false, desc;
//   }
//   switch {
//     case card.Can_consume == false:{
//       desc = fmt.Sprintf("卡券已不可核销 card=%s code=%s status=%s", card.Card.Card_id, card.Card.Code, card.User_card_status);
//       log.Infof(desc);
//       valid = false;
//       break;
//     }
//     case card.User_card_status == "NORMAL":{
//       valid = true;
//       break;
//     }
//     case card.User_card_status == "GIFT_TIMEOUT":{
//       valid = true;
//       break;
//     }
//     default:{
//       desc = fmt.Sprintf("卡券已不可用 card=%s code=%s status=%s", card.Card.Card_id, card.Card.Code, card.User_card_status);
//       log.Infof(desc);
//       valid = false;
//       break;
//     }
//   }
//   return nil, valid, desc;
// }
