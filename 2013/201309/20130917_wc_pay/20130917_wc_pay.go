package main

import (
    "log"
    "net/http"
    "fmt"

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
    mpServer := mp.NewDefaultServer("gh_7cd2c63862d9", "wechat4xbed", "wx0c49d2c7f9d36648", aesKey, messageServeMux)

    mpServerFrontend := mp.NewServerFrontend(mpServer, mp.ErrorHandlerFunc(ErrorHandler), nil)

    // 如果你在微信后台设置的回调地址是
    //   http://xxx.yyy.zzz/wechat
    // 那么可以这么注册 http.Handler
    http.Handle("/", mpServerFrontend);
    http.HandleFunc("/x4/pay", Pay);
    http.Handle("/x4/socket", websocket.Handler(Payment));
    http.ListenAndServe(":8091", nil);
}

func Pay(w http.ResponseWriter, r *http.Request){
    fmt.Println(r.URL);
    fmt.Fprint(w, Main_page);
}

type Price struct{
    Gold int;
    WType string;
};

func Payment(ws *websocket.Conn) {
    fmt.Println("get a WebSocket connected!");
    // websocket.JSON.Send(ws, Price{100, 1});
    // var err error;
 
    for {
        // var reply string;
 
        // if err = websocket.Message.Receive(ws, &reply); err != nil {
        //     fmt.Println("Can't receive");
        //     break;
        // }

        var aPr Price;
        websocket.JSON.Receive(ws, &aPr);
        fmt.Println(aPr);
 
        client := &http.Client{};

        val := map[string]string{
            "appid":"wx0c49d2c7f9d36648",
            "mch_id":"1263833801",
            // "device_info":"WEB",
            "nonce_str":"e61463f8efa94090b1f366cccfbbb444",
            "sign":"C380BEC2BFD727A4B6845133519F3AD6",
            "body":"大床房间",
            "detail":"大床房间",
            "attach":"backlog",
            "out_trade_no":"1217752501201407033233368018",
            // "fee_type":"CNY",
            "total_fee":"1",
            "spbill_create_ip":"127.0.0.1",
            "notify_url":"wxi.xbed.com.cn",
            "trade_type":"JSAPI",
            // "trade_type":"APP",
            // "product_id":"12235413214070356458058",
            // "limit_pay":"no_credit",
            "openid":aPr.WType,
        };
        pro := mch.NewProxy("gh_7cd2c63862d9", "wechat4xbed", "wx0c49d2c7f9d36648", client);
        ab, err := pay.UnifiedOrder(pro, val);
        if err != nil{
            fmt.Println("unified err: ", err);
            // return;
        }
        fmt.Println(ab);
        // fmt.Println("Received back from client: " + reply);
 
        // msg := "Received:  " + reply;
        // fmt.Println("Sending to client: " + msg);
 
        // if err = websocket.Message.Send(ws, msg); err != nil {
        //     fmt.Println("Can't send");
        //     break;
        // }
    }
}


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

*/
