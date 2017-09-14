/*
auth: bonly
*/

package Bonly_Host

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

func Update()(ret string){
	fl, err := os.OpenFile("/etc/hosts", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666);
	if err != nil{
		ret = fmt.Sprintf("打开文件失败: %v", err);
		return ret;
	}
	defer fl.Close();

	body, err := Get_Hosts();
	if err != nil{
		return fmt.Sprintf("获取文件失败: %v", err);
	}

	scanner := bufio.NewScanner(bytes.NewReader(body));
	var begin_write = false;
	for scanner.Scan(){
		if scanner.Text() == "# Modified hosts start"{
			begin_write = true;
			fl.Write([]byte("127.0.0.1\tlocalhost\n"));
			fl.Write([]byte("::1\tlocalhost\n"));
		}
		if begin_write {
			fl.Write([]byte(scanner.Text()));
			fl.Write([]byte("\n"));
		}
	}
	if err := scanner.Err(); err != nil{
		fmt.Println(err);
		return fmt.Sprintf("更新文件失败: %v", err);
	}
	// fmt.Println(string(body));
	return fmt.Sprintf("更新成功");
}