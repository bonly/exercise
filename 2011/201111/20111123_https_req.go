package main 

import (
  "net/http"
  "crypto/tls"
)

/*
使用https,暂未调试
*/
func main(){
	tr := &http.Transport{
	     //TLSClientConfig:    &tls.Config{RootCAs: pool},
	     DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://example.com")
}
