package main

import (
	"github.com/zserge/webview"
	"log"
	"net"
	"net/http"
)

const (
	windowWidth  = 480
	windowHeight = 320
)

var indexHTML = `
<!doctype html>
<html>
<head><title>world</title> </head>
<body>
   <button id="bttn" onclick="external.invoke('fromHtml')">click me</button>
</body>
<html>
`

func startSrv() string {
	ln, _ := net.Listen("tcp", "127.0.0.1:0")
	go func() {
		defer ln.Close()
		http.HandleFunc("/", func(wr http.ResponseWriter, rp *http.Request) {
			wr.Write([]byte(indexHTML))
		})
		log.Fatal(http.Serve(ln, nil))
	}()
	return "http://" + ln.Addr().String()
}

func rpc(wb webview.WebView, data string) {
	switch {
	case data == "fromHtml":
		log.Printf("get data from html: %s\n", data)
		wb.Eval(`document.getElementById("bttn").innerText="hello"`)
	}
}

func main() {
	log.Printf("begin\n")
	defer log.Printf("end.\n")

	url := startSrv()
	wb := webview.New(webview.Settings{
		Width:     windowWidth,
		Height:    windowHeight,
		Title:     "test app",
		Resizable: true,
		URL:       url,
		ExternalInvokeCallback: rpc,
	})
	defer wb.Exit()
	wb.Run()
}
