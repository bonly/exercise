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
	http.HandleFunc("/x5/pay", Pay);
	http.Handle("/x5/socket", websocket.Handler(Payment));
	http.ListenAndServe(":8091", nil);
}

func Pay(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.URL);
	fmt.Fprint(w, Main_page);
}

type Price struct{
	Gold int;
	WType int;
};

func Payment(ws *websocket.Conn) {
	fmt.Println("get a WebSocket connected!");
	websocket.JSON.Send(ws, Price{100, 1});
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
 
        // fmt.Println("Received back from client: " + reply);
 
        // msg := "Received:  " + reply;
        // fmt.Println("Sending to client: " + msg);
 
        // if err = websocket.Message.Send(ws, msg); err != nil {
        //     fmt.Println("Can't send");
        //     break;
        // }
    }
}

const Main_page string=`
<!DOCTYPE html>
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1" />  

<script>
var wsUri ="ws://127.0.0.1:8091/x5/socket"; 
var output;  

function init() { 
    output = document.getElementById("output"); 
    testWebSocket(); 
}
function testWebSocket() { 
	websocket = new WebSocket(wsUri); 
    websocket.onopen = function(evt) { onOpen(evt) }; 
    websocket.onclose = function(evt) { onClose(evt) }; 
    websocket.onmessage = function(evt) { onMessage(evt) }; 
    websocket.onerror = function(evt) { onError(evt) }; 
}  
 
function onOpen(evt) {writeToScreen("CONNECTED");}
function onClose(evt) { writeToScreen("DISCONNECTED"); }
function onError(evt) {writeToScreen('<span style="color: red;">ERROR:</span> '+ evt.data); }

function onMessage(evt) { 
    writeToScreen('<span style="color: blue;">RESPONSE: '+ evt.data+'</span>'); 
    // var b='{"name":"Mike","sex":"女","age":"29"}'; //字符串转为json对象
    // var bToObj=JSON.parse(b);
    // websocket.close(); 
}  

function doSend(message) { 
    writeToScreen("SENT: " + message);  
    websocket.send(message); 
}  

function writeToScreen(message) { 
    var pre = document.createElement("p"); 
    pre.style.wordWrap = "break-word"; 
    pre.innerHTML = message; 
    output.appendChild(pre); 
}

window.addEventListener("load", init, false);  

</script>

<script>
function btnClick(){
	// doSend('{"Gold":101,"WType":2}');  //以字符串方式发送json,服务器能解释
	var a={"Gold":103,"WType":3};  //json对象,需要转为字符串后发送
	doSend(JSON.stringify(a));
}
</script>
</head>
<body>
<button style="padding-right:10" onclick="btnClick()">支付</button>

<h2>金额:0.01</h2>  
<div id="output"></div> 

</body>
</html> 
`;