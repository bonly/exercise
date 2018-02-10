package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zserge/webview"
	"log"
	"net"
	"net/http"
)

var router *gin.Engine

func srv() string {
	addr, _ := net.Listen("tcp", "127.0.0.1:0")

	go func() {
		router = gin.Default()
		router.LoadHTMLGlob("templates/*")
		initializeRoutes()
		http.Serve(addr, router)
	}()
	return "http://" + addr.Addr().String()
}

func rpc(wb webview.WebView, data string) {
	log.Printf("get data from html: %s\n", data)
	wb.Eval(`document.getElementById("ft").innerText="文案已变"`)
}

func main() {
	url := srv()
	wb := webview.New(webview.Settings{
		Title:     "测试程序",
		Resizable: true,
		URL:       url,
		ExternalInvokeCallback: rpc,
	})
	defer wb.Exit()
	wb.Run()
}
