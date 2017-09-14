/*
auth: bonly
create: 2016-5-4
*/
package main

import (
  "log"
  "sync"
  "github.com/bitly/go-nsq"
  "os/signal"
  "os"
  "syscall"
  "encoding/xml"
  "time"
  "io/ioutil"
  "strings"
  "net/http"
  "fmt"
  "flag"
  "net/url"
  // "bytes"
  // "mime/multipart"
)

var Wg sync.WaitGroup;
var Run bool;

var Card_srv *string = flag.String("p", "http://120.25.106.243:6871", "coupon card srv addr");
var Card_sec *string = flag.String("s", "02a1a682e1f448hhbcd2cdc5682baaab", "coupon card srv secret");

func main() {
  flag.Parse();

  Run = true;

  /// 建立要收集的信号
  sigs := make(chan os.Signal, 1);
  signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM);

  go func(){
    Wg.Add(1);
    defer Wg.Done();
    for{
      sig := <-sigs;
      process_sig(sig);
      if Run == false {
        break;
      }
    }
  }();

  User_get_card();
  User_del_card();
  Wg.Wait();
}

func User_get_card(){
  config := nsq.NewConfig();
  q, _ := nsq.NewConsumer("wechat.event.user_get_card", "notify_card", config);
  q.AddHandler(nsq.HandlerFunc(notify_add_card));
  
  err := q.ConnectToNSQD("devpay.xbed.com.cn:4150")
  //err := q.ConnectToNSQLookupd("192.168.1.10:4161");
  if err != nil {
      log.Panic("Could not connect");
  }
}

func User_del_card(){
  config := nsq.NewConfig();
  q, _ := nsq.NewConsumer("wechat.event.user_del_card", "notify_card", config);
  q.AddHandler(nsq.HandlerFunc(notify_del_card));
  
  err := q.ConnectToNSQD("devpay.xbed.com.cn:4150")
  //err := q.ConnectToNSQLookupd("192.168.1.10:4161");
  if err != nil {
      log.Panic("Could not connect");
  }
}
/*
<xml><ToUserName><![CDATA[gh_64de4351ce17]]></ToUserName>
<FromUserName><![CDATA[oM9dAwgu7wAR92d40chOcBesHlNQ]]></FromUserName>
<CreateTime>1462349553</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[user_get_card]]></Event>
<CardId><![CDATA[pM9dAwjvZriZH65LXTUHkc_bPve0]]></CardId>
<IsGiveByFriend>0</IsGiveByFriend>
<UserCardCode><![CDATA[148822443028]]></UserCardCode>
<FriendUserName><![CDATA[]]></FriendUserName>
<OuterId>0</OuterId>
<OldUserCardCode><![CDATA[]]></OldUserCardCode>
<IsRestoreMemberCard>0</IsRestoreMemberCard>
</xml>
*/
type WcBody struct {
  XMLName      xml.Name `xml:"xml"`;
  ToUserName   string;
  FromUserName string;
  CreateTime   time.Duration;
  MsgType      string;
  Event        string;
  CardId       string;
  IsGiveByFriend int;
  UserCardCode string;
  FriendUserName string;
  OuterId int;
  OldUserCardCode string;
  IsRestoreMemberCard string;
};

func Get_wc_msg(body *string) *WcBody {
  requestBody := &WcBody{};
  xml.Unmarshal([]byte(*body), requestBody);
  return requestBody;
}

func notify_add_card(message *nsq.Message) error{
  log.Println("============= Begin user add card ============");
  defer log.Println("============= End user add card ============");

  body := string(message.Body);
  log.Printf("User get card: %v", body);

  //检查配置
  if len(*Card_srv) <= 0 || len(*Card_sec) <= 0 {
    log.Println("缺少卡券配置[Card_srv]||[Card_sec]!");
    return fmt.Errorf("缺少卡券配置");
  }

  //取数据
  msg := Get_wc_msg(&body);
  if (msg == nil){
    log.Println("取wc数据错误");
  }

  log.Printf("%+v \n", msg);

  qry := *Card_srv + "/api/couponCard/userCard/focus";

  param := url.Values{};
  param.Add("cardId", msg.CardId);
  param.Add("code", msg.UserCardCode);
  param.Add("openId", msg.FromUserName);
  param.Add("isGiveByFriend", fmt.Sprintf("%d", msg.IsGiveByFriend));
  param.Add("outerId", fmt.Sprintf("%d", msg.OuterId));
  
  // dat := url.QueryEscape(param.Encode());
  dat := param.Encode();
  
  log.Println("send card srv: ", qry);
  log.Println("send card body: ", dat);
  
  cli := &http.Client{};

  // bd := &bytes.Buffer{};
  // writer := multipart.NewWriter(bd);
  // writer.WriteField("cardId", msg.CardId);
  // writer.WriteField("code", msg.UserCardCode);
  // writer.WriteField("openId", msg.FromUserName);
  // writer.WriteField("isGiveByFriend", fmt.Sprintf("%d", msg.IsGiveByFriend));
  // writer.WriteField("outerId", fmt.Sprintf("%d", msg.OuterId));
    
  // req, err := http.NewRequest("POST", qry, bd);
  req, err := http.NewRequest("POST", qry, strings.NewReader(dat));
  if err != nil{
    log.Println("http: ", err);
    return err;
  }
  // req.Header.Set("Content-Type", writer.FormDataContentType());
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded");
  req.Header.Set("User-Agent", "curl/7.48.0");
  req.Header.Set("secret", *Card_sec);
  req.Header.Set("Accept", "*/*");
  // req.Header.Set("Accept-Encoding", "gzip, deflate");
  // req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8");
  // req.Header.Set("Cache-Control", "no-cache");
  // req.Header.Set("Connection", "keep-alive");

  resp, err := cli.Do(req);
  if err != nil{
    log.Println("http get: ", err);
    return err;
  }
  defer resp.Body.Close();

  ret_cr, err := ioutil.ReadAll(resp.Body);
  if err != nil{
    log.Println("body: ", err);
    return err;
  }

  log.Println("card srv back: ", string(ret_cr));

  return nil;
}


func notify_del_card(message *nsq.Message) error{
  log.Println("============= Begin user del card ============");
  defer log.Println("============= End user del card ============");

  body := string(message.Body);
  log.Printf("User get card: %v", body);

  //检查配置
  if len(*Card_srv) <= 0 || len(*Card_sec) <= 0 {
    log.Println("缺少卡券配置[Card_srv]||[Card_sec]!");
    return fmt.Errorf("缺少卡券配置");
  }

  //取数据
  msg := Get_wc_msg(&body);
  if (msg == nil){
    log.Println("取wc数据错误");
  }

  log.Printf("%+v \n", msg);

  qry := *Card_srv + "/api/couponCard/userCard/blur";

  param := url.Values{};
  param.Add("cardId", msg.CardId);
  param.Add("code", msg.UserCardCode);
  
  dat := param.Encode();
  qry = qry + "?" + dat;

  log.Println("send card srv: ", qry);
  
  cli := &http.Client{};
    
  req, err := http.NewRequest("PUT", qry , nil);
  if err != nil{
    log.Println("http: ", err);
    return err;
  }
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded");
  req.Header.Set("User-Agent", "curl/7.48.0");
  req.Header.Set("secret", *Card_sec);
  req.Header.Set("Accept", "*/*");

  resp, err := cli.Do(req);
  if err != nil{
    log.Println("http get: ", err);
    return err;
  }
  defer resp.Body.Close();

  ret_cr, err := ioutil.ReadAll(resp.Body);
  if err != nil{
    log.Println("body: ", err);
    return err;
  }

  log.Println("card srv back: ", string(ret_cr));

  return nil;
}

/// 处理系统信号
func process_sig(sig os.Signal){
    switch(sig){
        case syscall.SIGINT:
           log.Println("Get SigInt");
           Run = false; 
           break;
        case syscall.SIGTERM:
           log.Println("Get SigTerm");
           Run = false;
           break;
        case syscall.SIGUSR1:
            log.Println("Get SIGUSR1.");
            Run = true;
            break;
        default:
           log.Println("Get Sig: ", sig);
           Run = true;
           break;
    }
}