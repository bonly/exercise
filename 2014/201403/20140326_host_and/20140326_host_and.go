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
"io"
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
	fl, err := os.OpenFile("/etc/hosts", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666);
	if err != nil{
		fmt.Println("打开文件失败: ", err);
		return;
	}
	defer fl.Close();

	body, err := Get_Hosts();
	if err != nil{
		fmt.Println("获取文件失败: ", err);
		return;
	}

	var begin_write = false;
	br := bufio.NewReader(bytes.NewReader(body));
	bw := bufio.NewWriter(fl);

	for{
		line, err := br.ReadString('\n');
		if err == nil{
			fmt.Println(line);
			if line == "# Modified hosts start\n"{
				begin_write = true;
				fl.Write([]byte("127.0.0.1\tlocalhost\n"));
				fl.Write([]byte("::1\tlocalhost\n"));
			}
			if begin_write {
				// fl.Write([]byte(line));
				fmt.Fprintln(bw, line); //最快
				// bw.WriteString(line); //比上面的慢些
			}			
		}else{
			if err == io.EOF{
				// fmt.Fprintln(bw, line);
				break;
			}
			fmt.Println("other:", err);
			break;
		}
	}
	defer bw.Flush();

	fmt.Println("更新完成!");
	return;
}
