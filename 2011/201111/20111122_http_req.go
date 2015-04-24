package main 

import (
  "net/http"
  "log"
  "io/ioutil"
)

/*
@note 控制发送参数
*/
func main(){
	client := &http.Client{
       //CheckRedirect: redirectPolicyFunc, 重定向函数
	}
	resp, err := client.Get("http://example.com")
	req, err := http.NewRequest("GET", "http://example.com", nil)
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err = client.Do(req)

    if err != nil {
       log.Println("handle error");
	}
	defer resp.Body.Close();
	body, err := ioutil.ReadAll(resp.Body);

	log.Println(string(body));	
}
