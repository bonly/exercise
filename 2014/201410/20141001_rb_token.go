package main 

import (
"log"
"net/http"
"strings"
// "encoding/json"
"io/ioutil"
)

func main(){
	data := `{"username":"hbonly","password":"hayes11hjb"}`;
	request, err := http.NewRequest("POST",
		"https://api.robinhood.com/api-token-auth/", strings.NewReader(data));

	if err != nil{
		log.Printf("new %s\n", err.Error());
		return;
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");
	request.Header.Set("Content-Type", "application/json");
	request.Header.Set("Accept", "application/json");

	// fmt.Println("request:", request);
	cli := &http.Client{};
	resp, err := cli.Do(request);
	if err != nil{
		log.Printf("post %s\n", err.Error());
		return;
	}
	defer resp.Body.Close();

	body, err := ioutil.ReadAll(resp.Body);
	if err != nil{
		log.Printf("body %s\n", err.Error());
		return;
	}

	log.Printf("recv:\n %s\n", string(body));
}