package main 

import (
  "net/http"
  "log"
  "io/ioutil"
)

/*
直接默认发送
*/
func main(){
	resp, err := http.Get("http://pay3.shenzhoufu.com/interface/version3/serverconnszx/entry-noxml.aspx");
    //resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf);
    //resp, err := http.PostForm("http://example.com/form",url.Values{"key": {"Value"}, "id": {"123"}});

    if err != nil {
       log.Println("handle error");
	}
	defer resp.Body.Close();
	body, err := ioutil.ReadAll(resp.Body);

	log.Println(string(body));
}
