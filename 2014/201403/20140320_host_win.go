/*
auth: bonly
*/

package main

import (
"fmt"
"io/ioutil"
"net/http"
"bufio"
"bytes"
"os"
)

func Get_Hosts()(body []byte, err error){
	resp, err := http.Get("https://raw.githubusercontent.com/racaljk/hosts/master/hosts");
	if err != nil {
		fmt.Println("Get Host err: ", err);
		return nil, err;
	}
	defer resp.Body.Close();

	body, err = ioutil.ReadAll(resp.Body);
	if err != nil{
		fmt.Println("Body err: ", err);
		return nil, err;
	}

	return body, nil;
}

func main(){
	fl, err := os.OpenFile("c:\\Windows\\System32\\Drivers\\etc\\hosts", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666);
	if err != nil{
		fmt.Println("打开文件失败:",err);
		return;
	}
	defer fl.Close();

	body, err := Get_Hosts();
	if err != nil{
		return;
	}

	scanner := bufio.NewScanner(bytes.NewReader(body));
	var begin_write = false;
	for scanner.Scan(){
		if scanner.Text() == "# Modified hosts start"{
			begin_write = true;
			fl.Write([]byte("127.0.0.1\tlocalhost\r\n"));
			fl.Write([]byte("::1\tlocalhost\r\n"));
		}
		if begin_write {
			fl.Write([]byte(scanner.Text()));
			fl.Write([]byte("\r\n"));
		}
	}
	if err := scanner.Err(); err != nil{
		fmt.Println(err);
	}
	// fmt.Println(string(body));
}