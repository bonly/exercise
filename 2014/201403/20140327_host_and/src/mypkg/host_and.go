/*
auth: bonly
*/

package mypkg

import (
"fmt"
"io/ioutil"
"net/http"
"bufio"
"bytes"
"os"
"io"
"os/exec"
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

func Exec_Cmd()(ret string, err error){
	cmd_argv := []string{"-c", "mount -o remount,rw /system"};
	cmd := exec.Command("/system/xbin/su", cmd_argv...);
	_, err = cmd.CombinedOutput();
	if err != nil{
		ret = fmt.Sprintf("exec err: %v", err);
		return;
	}
	return;
}

func Update() (ret string){
	ret, err := Exec_Cmd();
	if err != nil{
		return;
	}

	// ret = fmt.Sprintf("hi,中文！");
	fl, err := os.OpenFile("/sdcard/hosts", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666);
	if err != nil{
		ret = fmt.Sprintf("打开文件失败: %v", err);
		return ret;
	}
	defer fl.Close();

	body, err := Get_Hosts();
	if err != nil{
		ret = fmt.Sprintf("获取文件失败: %v", err);
		return ret;
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

	ret = fmt.Sprintf("更新完成!");
	return ret;
}
