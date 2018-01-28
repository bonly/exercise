package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpGet() {
	resp, err := http.Get("http://gd.189.cn/uniteWeb/J/J50010.j?a.c=0&a.u=user&a.p=pass&a.s=ECSS&jsoncallback=jQuery111305524063522834817_1513076241758&d.d01=200&d.d02=1&d.d03=100&d.d04=&d.d05=ANY&d.d06=DQGX_NUMPOOL&d.d07=&d.d11=&d.d12=&d.d16=PTK00026&d.d25=&d.d26=&d.d28=&d.d29=&_=1513076241761")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func main() {
	httpGet()
}
