package main 

import (
"encoding/xml"
"net/http"
"log"
"fmt"
"strings"
"io/ioutil"
)


type Command struct{
	Name string `xml:"name,attr"`;
	Sn string `xml:"sn,attr"`;
	Version string `xml:"version,attr"`;
	URL string;
};

const soap string = `<?xml version="1.0" encoding="UTF-8" ?>`;

func main(){
	var qry Command;
	qry.Name="Registering_Callbacks_Url_REQ";
	qry.Sn="1";
	qry.Version="1.0.0";
	qry.URL="http://127.0.0.1:8005/Callbacks.asmx";

	xm, err := xml.Marshal(qry);
	if err != nil{
		log.Println("XML错误: ", err);
		return;
	}

	str := fmt.Sprint(soap, string(xm));
	log.Println("请求服务器: ", str);

	client := &http.Client{};
	req, err := http.NewRequest("POST",
		"http://183.63.52.28:8005/NetLockWebServer.asmx",
		strings.NewReader(str));

	req.Header.Set("Content-Type", "text/xml;charset=utf-8");
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");
	// req.Header.Set("SOAPAction", "www.mingyansoft.com/SetRoomIsDiry");
	resp, err := client.Do(req);
	if err != nil{
		log.Println("send: ", err);
		return;
	}	
	defer resp.Body.Close();

	body, err := ioutil.ReadAll(resp.Body);
	if err != nil {
	  	log.Println("recv: ", err);
	  	return;
	}
	
	log.Println("服务器应答：", string(body));		
}
