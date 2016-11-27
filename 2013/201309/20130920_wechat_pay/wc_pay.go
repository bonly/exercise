package main

import (
	"log"
	"net/http"
	"fmt"
	"net/url"
	"encoding/json"
	"io/ioutil"
	"time"

	"golang.org/x/net/websocket"

	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/message/request"
	"github.com/chanxuehong/wechat/mp/message/response"
	"github.com/chanxuehong/wechat/util"

	"github.com/chanxuehong/wechat/mch"
	"github.com/chanxuehong/wechat/mch/pay"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	log.Println(err.Error())
}

// 文本消息的 Handler
func TextMessageHandler(w http.ResponseWriter, r *mp.Request) {
	// 简单起见，把用户发送过来的文本原样回复过去
	text := request.GetText(r.MixedMsg) // 可以省略, 直接从 r.MixedMsg 取值
	resp := response.NewText(text.FromUserName, text.ToUserName, text.CreateTime, text.Content)
	mp.WriteRawResponse(w, r, resp) // 明文模式
	// mp.WriteAESResponse(w, r, resp) // 安全模式
}

func main() {
	
	aesKey, err := util.AESKeyDecode("teWHTRTT12gjyEkfgI91CzK05IwWID6UNOLgKoBGU0a") // 这里 encodedAESKey 改成你自己的参数
	if err != nil {
		panic(err)
	}

	messageServeMux := mp.NewMessageServeMux()
	messageServeMux.MessageHandleFunc(request.MsgTypeText, TextMessageHandler) // 注册文本处理 Handler

	// 下面函数的几个参数设置成你自己的参数: oriId, token, appId
	mpServer := mp.NewDefaultServer("gh_7cd2c63862d9", token, appid, aesKey, messageServeMux)

	mpServerFrontend := mp.NewServerFrontend(mpServer, mp.ErrorHandlerFunc(ErrorHandler), nil)
    
	// 如果你在微信后台设置的回调地址是
	//   http://xxx.yyy.zzz/wechat
	// 那么可以这么注册 http.Handler
	http.Handle("/wc", mpServerFrontend);
	http.HandleFunc("/pay", Pay);
	http.Handle("/x4/socket", websocket.Handler(Payment));
    // http.Handle("/x4/", http.FileServer(http.Dir("./web/")));
    http.Handle("/example/", http.StripPrefix("/example/",http.FileServer(http.Dir("web"))));
	http.ListenAndServe(":80", nil);
}
/*
跳转到加了参数的页面
*/
func Pay(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.URL);
    uri, _ := url.Parse(r.URL.String());
	vl := uri.Query();

	code := vl.Get("code");

	fmt.Println("code: ",code);
	fmt.Println("state: ", vl.Get("state"));

	tk := get_openid(code); //通过code取得openid

	val := url.Values{};
	val.Add("access_token", tk.Access_token);
	val.Add("openid", tk.Openid);
	ret := fmt.Sprintf(Jump_page, vl.Get("state"), val.Encode());	
	fmt.Fprint(w, ret);
}

type Head struct{
	Cmd string;
}

type R_Token struct {
Access_token string `json:"access_token"`;
Expires_in int `json:"expires_in"`;
Refresh_token string `json:"refresh_token"`;
Openid string `json:"openid"`;
Scope string `json:"scope"`;
};

const (
	token = "wechat4xbed";
	appid = "wx0c49d2c7f9d36648";
	secret = "910fe488f3ff205b428905e2e1733a94";	
	mch_id = "1263833801";
)

func get_openid(cd string)(ret R_Token){
	val := url.Values{};
        val.Add("grant_type", "authorization_code");
	val.Add("code", cd);
	val.Add("appid", appid);
	val.Add("secret", secret);
	addr := "https://api.weixin.qq.com/sns/oauth2/access_token?"+val.Encode();

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

	if err := json.Unmarshal(body,&ret); err != nil{
		log.Println("token: ", err);
	}    
    return ret;
}

type Price struct{
	Appid string `json:"appid"`;
	Mch_id string `json:"mch_id"`;
	Nonce_str string `json:"nonce_str"`;
	Body string `json:"body"`;
	Detail string `json:"detail"`;
	Attach string `json:"attach"`;
	Out_trade_no string `json:"out_trade_no"`;
	Total_fee string `json:"total_fee"`;
	Spbill_create_ip string `json:"spbill_create_ip"`;
	Notify_url string `json:"notify_url"`;
	Trade_type string `json:"trade_type"`;
	Openid string `json:"openid"`;
};

func Cmd_price(ws *websocket.Conn){
    var aPr Price;
    err := websocket.JSON.Receive(ws, &aPr);
    if err != nil{
    	fmt.Println(err);
    	return;
    }

    aPr.Nonce_str = NewV4().String()[:32];
    aPr.Appid = appid;
    aPr.Mch_id = mch_id;
    aPr.Attach = "backlog";
    aPr.Notify_url = "wxi.xbed.com.cn";
    aPr.Trade_type = "JSAPI";

	tn := time.Now();
	strTn := fmt.Sprintf("%d%d%d",tn.Year(),tn.Month(),tn.Day());
	aPr.Out_trade_no = fmt.Sprintf("%s-%s", strTn, NewV4().String()[:8]);    
	
	fmt.Println(aPr);

	client := &http.Client{};

	val := map[string]string{
		"appid":aPr.Appid,
		"mch_id":aPr.Mch_id,
		// "device_info":"WEB",
		"nonce_str":aPr.Nonce_str,
		// "sign":"C380BEC2BFD727A4B6845133519F3AD6",
		"body":aPr.Body,
		"detail":aPr.Detail,
		"attach":aPr.Attach,
		"out_trade_no":aPr.Out_trade_no,
		// "fee_type":"CNY",
		"total_fee":aPr.Total_fee,
		"spbill_create_ip":aPr.Spbill_create_ip,
		"notify_url":aPr.Notify_url,
		"trade_type":aPr.Trade_type,
		// "trade_type":"APP",
		// "product_id":"12235413214070356458058",
		// "limit_pay":"no_credit",
		"openid":aPr.Openid,
	};
	val["sign"]=mch.Sign(val, secret, nil);
	fmt.Println("sign: ", val["sign"]);
	pro := mch.NewProxy(appid, mch_id, secret, client);
	ab, err := pay.UnifiedOrder(pro, val);
	if err != nil{
		fmt.Println("unified err: ", err);
	// return;
	}
	fmt.Println(ab);

	websocket.JSON.Send(ws, &ab); //把结束回html    
}

type CPay struct{
	AppId string `json:"appId"`;
	TimeStamp string `json:"timeStamp"`;
	NonceStr string `json:"nonceStr"`;
	Package string `json:"package"`;
	SignType string `json:"signType"`;
	PaySign string `json:"paySign"`;
	Cmd string `json:"Cmd"`;
};

func Cmd_pay(ws *websocket.Conn){
	var aPay CPay;
    err := websocket.JSON.Receive(ws, &aPay);
    if err != nil{
    	fmt.Println(err);
    	return;
    }

    utime := time.Now().Unix();
    aPay.TimeStamp = fmt.Sprintf("%d",  utime);
	aPay.AppId = appid;
	aPay.NonceStr = NewV4().String()[:32];

	val := map[string]string{
		"appId":aPay.AppId,
		"timeStamp":aPay.TimeStamp,
		"nonceStr":aPay.NonceStr,
		"package":aPay.Package,
		"signType":aPay.SignType,
	};
	val["paySign"]=mch.Sign(val, secret, nil);
	aPay.PaySign = val["paySign"];
	aPay.Cmd = "R_pay";
	fmt.Println("return to html pay:", aPay);		

	fmt.Println("nonce_str============",aPay.PaySign);
	websocket.JSON.Send(ws, &aPay); //把结束回html    	
}

//协议解释入口
func Payment(ws *websocket.Conn) {
	fmt.Println("get a WebSocket connected!");
	// websocket.JSON.Send(ws, Price{100, 1});
    // var err error;
 
    for {
    	//取得头
    	var head Head;
        err := websocket.JSON.Receive(ws, &head);
        fmt.Println("Head: ", head);
        if err != nil{
        	fmt.Println(err);
        	return;
        }
        switch(head.Cmd){
        case "price":
        	Cmd_price(ws);
        	break;
        case "pay":
        	Cmd_pay(ws);
        	break;
        default:
        	break;
        }
    }
}

const Jump_page string=`
<html>
<head>
<meta http-equiv="Refresh" content="0,url=%s?%s" />
</head>
<body >
</body>
</html>
`;

/*
{
    "appid": "wx138f189741870ffc",
    "bank_type": "CMB_DEBIT",
    "cash_fee": "1",
    "fee_type": "CNY",
    "is_subscribe": "N",
    "mch_id": "1264010901",
    "nonce_str": "dg9g2mxvw4gmmysnj78p040cy7ddnyuh",
    "openid": "oNXecwrddUtNtXIDCz6Cjb_wp67E",
    "out_trade_no": "126401090120150821145254",
    "result_code": "SUCCESS",
    "return_code": "SUCCESS",
    "sign": "4D51319A59CD417173D945B69C97FA07",
    "time_end": "20150821150242",
    "total_fee": "1",
    "trade_type": "APP",
    "transaction_id": "1001020871201508210673470246"
}
        // var reply string;
 
        // if err = websocket.Message.Receive(ws, &reply); err != nil {
        //     fmt.Println("Can't receive");
        //     break;
        // }

        // fmt.Println("Received back from client: " + reply);
 
        // msg := "Received:  " + reply;
        // fmt.Println("Sending to client: " + msg);
 
        // if err = websocket.Message.Send(ws, msg); err != nil {
        //     fmt.Println("Can't send");
        //     break;
        // }        
*/
