package main 

import (
// "encoding/xml"
"net/http"
"log"
"fmt"
"strings"
"io/ioutil"
// "net/url"
// "html"
)


type Command struct{
	Name string `xml:"name,attr"`;
	Sn string `xml:"sn,attr"`;
	Version string `xml:"version,attr"`;
	URL string;
};

const SoapTemplate = `<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><soap:Body><NetLockWeb xmlns="http://tempuri.org/">
    <recdatastr>%s
</recdatastr>
</NetLockWeb></soap:Body></soap:Envelope>`;

func main(){
	// var qry Command;
	// qry.Name="Registering_Callbacks_Url_REQ";
	// qry.Sn="7";
	// qry.Version="1.0.0";
	// qry.URL="http://127.0.0.1:8005/Callbacks.asmx";

	// xm, err := xml.Marshal(qry);
	// if err != nil{
	// 	log.Println("XML错误: ", err);
	// 	return;
	// }
	// qdata := string(xm);
	// log.Println("原请求数据：", qdata);

	// str := fmt.Sprintf(SoapTemplate1, url.QueryEscape(`<?xml version="1.0" encoding="utf-8"?>` + qdata));
	// str := fmt.Sprintf(SoapTemplate1, html.EscapeString(`<?xml version="1.0" encoding="utf-8"?>` + qdata));
	str := fmt.Sprintf(SoapTemplate, `&lt;?xml version="1.0" encoding="UTF-8"?&gt;&lt;Command name="QuertNetMidComStatus_REQ" sn="7" version="1.0.0"&gt;&lt;MidComNo&gt;01020304&lt;/MidComNo&gt;&lt;/Command&gt;`);
	log.Println("请求服务器: ", str);

	client := &http.Client{};
	req, err := http.NewRequest("POST",
		"http://183.63.52.28:8005/NetLockWebServer.asmx",
		strings.NewReader(str));

	req.Header.Set("Content-Type", "text/xml;charset=utf-8");
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");
	req.Header.Set("SOAPAction", "http://tempuri.org/NetLockWeb");
	// req.Header.Set("Host", "183.63.52.28:8005");
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
