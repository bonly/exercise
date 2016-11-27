package main

import (
"fmt"
"net/http"
"net/url"
"log"
"io/ioutil"
)

const appid = "wx0c49d2c7f9d36648";
const appsecret = "910fe488f3ff205b428905e2e1733a94";

type CToken struct {
	Access_token string `json:"access_token"`;
	Expires_in string `json:"expires_in"`;
};

/*
获取Token
*/
func token_qry(appid string, appsecret string)(ret string){
	str := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appid, appsecret);
    resp, err := http.Get(str);
    if err != nil {
        log.Println("http get: ",err);
    }
    defer resp.Body.Close();
    
	body, err := ioutil.ReadAll(resp.Body);
    if err != nil {
        log.Println("body: ", err);
    }
 
 	ret = string(body);
    log.Println(ret);	
    return ret;
}

/*
获取Token服务
*/
func Token(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL);
	fmt.Fprintf(w, token_qry(appid, appsecret));
}

/*
获取OpenId
*/
func openid_qry(appid string, redirect string){
	str := fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=SCOPE&state=STATE#wechat_redirect", appid, url.QueryEscape(redirect));
    resp, err := http.Get(str);
    if err != nil {
        log.Println("http get: ",err);
    }
    defer resp.Body.Close();
    
	body, err := ioutil.ReadAll(resp.Body);
    if err != nil {
        log.Println("body: ", err);
    }
 
 	ret := string(body);
    log.Println(ret);	
}

/*
获取OpenId服务
*/
func Openid(w http.ResponseWriter, r *http.Request) {
	openid_qry(appid, "http://wxi.xbed.com.cn");
}

func main(){
	http.HandleFunc("/token", Token);
	http.HandleFunc("/openid", Openid);
	http.ListenAndServe(":8989", nil);
}
