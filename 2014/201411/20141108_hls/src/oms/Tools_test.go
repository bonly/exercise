package oms

import(
"net/http"
"strings"
"fmt"
"io/ioutil"
"flag"
)

var MidId = flag.String("mid", "00000001", "MidId");
var Srv = flag.String("s", "http://127.0.0.1:5010/cmd", "srv addr");

func init(){
	flag.Parse();
}

func Post(srv string, data string)(err error){
	fmt.Printf("Send: \n%s\n", data);
	request, err := http.NewRequest("POST",
		srv, strings.NewReader(data));
	if err != nil{
		fmt.Printf("REQ %s\n", err.Error());
		return err;
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");
	request.Header.Set("Content-Type", "application/json");
	request.Header.Set("Accept", "application/json");

	cli := &http.Client{};
	resp, err := cli.Do(request);
	if err != nil{
		fmt.Printf("POST %s\n", err.Error());
		return err;
	}	
	defer resp.Body.Close();

	body, err := ioutil.ReadAll(resp.Body);
	if err != nil{
		fmt.Printf("body %s\n", err.Error());
		return err;
	}
	fmt.Printf("Recv: \n%s\n", string(body));

	var ret PMS_Manual;
	err = ret.Decode(body);
	if err != nil{
		fmt.Printf("decode result %s\n", err.Error());
		return err;
	}
	if ret.ResultID != "0"{
		return fmt.Errorf("执行失败");
	}
	return nil;
}