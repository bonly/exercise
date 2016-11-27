package main

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"fmt"
	"log"
  "os"
)

const (
	token = "t16YiU7OMQz6RsDL2S17nCxoMm9kqfM0kygofb9RMxUfkpJR6KTBoT8-sBRhy9530zVMvbsae0DcGxxKPk--j3vwvDXah7EIgn5gIVAx3rE";
	appid = "wx100949d5a719dac2";
	secret = "10b624d9c9faa83165edfc1d4a336935";
)

func main(){
	get_openid(os.Args[1]);
}


func get_openid(cd string){
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
}


