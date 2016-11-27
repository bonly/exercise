package main

import (
	"fmt"

	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/menu"
)

var AccessTokenServer = mp.NewDefaultAccessTokenServer("wx0c49d2c7f9d36648", 
		"910fe488f3ff205b428905e2e1733a94", nil) // 一個應用只能有一個實例
var mpClient = mp.NewClient(AccessTokenServer, nil)

func main() {
	var subButtons = make([]menu.Button, 3);
	subButtons[0].SetAsViewButton("办理入住", "http://wxi.xbed.com.cn");
	subButtons[1].SetAsViewButton("房间服务", "http://wxi.xbed.com.cn");
	subButtons[2].SetAsViewButton("退房", "http://wxi.xbed.com.cn");
	// subButtons[1].SetAsClickButton("赞一下我们", "V1001_GOOD")

	var mn menu.Menu
	mn.Buttons = make([]menu.Button, 3);
	// mn.Buttons[0].SetAsClickButton("今日歌曲", "V1001_TODAY_MUSIC")
	mn.Buttons[0].SetAsViewButton("预订", "http://wxi.xbed.com.cn");
	mn.Buttons[1].SetAsSubMenuButton("服务", subButtons);

	red := url.Values{};
	red.Add("appid",appid);
	red.Add("redirect_uri","http://wxi.xbed.com.cn/x4");
	red.Add("response_type","code");
	red.Add("scope","snsapi_base");
	red.Add("state","http://wxi.xbed.com.cn/x4/web/pay.html");	
	mn.Buttons[2].SetAsViewButton("更多", "https://open.weixin.qq.com/connect/oauth2/authorize?");


	menuClient := (*menu.Client)(mpClient)
	if err := menuClient.CreateMenu(mn); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ok")
}