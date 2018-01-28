package main 

import (
"net/http"
"flag"
"io/ioutil"
"fmt"
)

var srv = flag.String("s", "https://www.googleapis.com/customsearch/v1", "google server");
var key = flag.String("key", "AIzaSyCDRPHwkd8jJkD00u9_diE6lGkbTmtjfzE", "key");
var cx  = flag.String("cx", "003049783992381642025:l6tuuj5z7hg", "cx ID");

func init(){
	flag.Parse();
}

func main(){
	qry := *srv + "?" + "key=" + *key + "&cx=" + *cx +"&q=" + "bonly"; 
	resp, err := http.Get(qry);
	if err != nil{
		return;
	}

	defer resp.Body.Close();
	body, err := ioutil.ReadAll(resp.Body);
	if err != nil{
		return;
	}

	fmt.Println(string(body));
}

/*
key: https://console.cloud.google.com/apis/credentials?project=bonly-api
AIzaSyCDRPHwkd8jJkD00u9_diE6lGkbTmtjfzE

cx: https://cse.google.com/cse/setup/basic?cx=003049783992381642025:l6tuuj5z7hg
003049783992381642025:l6tuuj5z7hg

*/