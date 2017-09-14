package main

import (
	// "fmt"
	"github.com/icza/gowut/gwu"
)

func main() {
	// Create and build a window
	win := gwu.NewWindow("bonly", "bonly Window"); //路径名，窗口名

	nameLabel := gwu.NewLabel("输入名字:");  //标签
	nameLabel.Style().SetColor(gwu.ClrRed); //颜色
	
	tb := gwu.NewTextBox("我叫"); //输入框
	tb.AddSyncOnETypes(gwu.ETypeKeyUp); //增加响应事件类型
	tb.AddEHandlerFunc(func(e gwu.Event) {
		nameLabel.SetText(tb.Text());
		e.MarkDirty(nameLabel);
	}, gwu.ETypeChange, gwu.ETypeKeyUp); //定义事件处理

	win.Add(nameLabel);
	win.Add(tb);

	server := gwu.NewServer("", "localhost:8081");//服务器网页根地址，侦听地址
	server.AddWin(win); //加入窗口
	server.Start("bonly"); //服务起动，等待结束,当窗口路径名为空时，打开的是列表
}