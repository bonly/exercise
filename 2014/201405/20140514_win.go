package main

import (
	// "fmt"
	"github.com/icza/gowut/gwu"
)

func main() {
	// Create and build a window
	win := gwu.NewWindow("bonly", "bonly Window"); //路径名，窗口名

	server := gwu.NewServer("", "localhost:8081");//服务器网页根地址，侦听地址
	server.AddWin(win); //加入窗口
	server.Start("bonly"); //服务起动，等待结束
}