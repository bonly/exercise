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
// "syscall"
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

func Exec_remount()(ret string, err error){
	cmd_argv := []string{"-c", "mount -o remount,rw /system"};
	cmd := exec.Command("/system/xbin/su", cmd_argv...);
	// cmd.SysProcAttr = &syscall.SysProcAttr{};
	// cmd.SysProcAttr.Credential = &syscall.Credential{Uid: 0, Gid: 0};
	_, err = cmd.CombinedOutput();
	if err != nil{
		ret = fmt.Sprintf("exec err: %v", err);
		return;
	}

	// cmd_argv = []string{"-c", "chmod a+rw /system/etc/hosts"};
	// cmd = exec.Command("/system/xbin/su", cmd_argv...);
	// _, err = cmd.CombinedOutput();
	// if err != nil{
	// 	ret = fmt.Sprintf("exec err: %v", err);
	// 	return;
	// }	
	return;
}

func Exec_mv(file string)(ret string, err error){
	cmd_argv := []string{"-c", "mv " + file + " /etc/hosts"};
	cmd := exec.Command("/system/xbin/su", cmd_argv...);
	// cmd.SysProcAttr = &syscall.SysProcAttr{};
	// cmd.SysProcAttr.Credential = &syscall.Credential{Uid: 0, Gid: 0};
	_, err = cmd.CombinedOutput();
	if err != nil{
		ret = fmt.Sprintf("mv hosts err: %v", err);
		return;
	}
	return;
}

func mk_resolv(pre string)(ret string){
	fl, err := os.OpenFile(pre+"/resolv.conf", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666);
	if err != nil{
		ret = fmt.Sprintf("打开文件失败: %v", err);
		return;
	}
	defer fl.Close();
	fl.Write([]byte("nameserver 114.114.114.114\n"));
	return;
}

func mv_resolv(file string)(ret string, err error){
	cmd_argv := []string{"-c", "mv " + file + " /etc/resolv.conf"};
	cmd := exec.Command("/system/xbin/su", cmd_argv...);
	_, err = cmd.CombinedOutput();
	if err != nil{
		ret = fmt.Sprintf("mv resolv.conf err: %v", err);
		return;
	}
	return;
}

func rm_resolv()(ret string){
	cmd_argv := []string{"-c", "rm -rf /etc/resolv.conf"};
	cmd := exec.Command("/system/xbin/su", cmd_argv...);
	_, err := cmd.CombinedOutput();
	if err != nil{
		ret = fmt.Sprintf("rm err: %v", err);
		return;
	}
	return;
}

func Update(pre string)(ret string){
	ret, err := Exec_remount(); //重挂载
	if err != nil{
		return;
	}

	fl, err := os.OpenFile(pre+"/hosts", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666);
	if err != nil{ //打开临时hosts
		ret = fmt.Sprintf("打开文件失败: %v", err);
		return;
	}
	defer fl.Close();

	mk_resolv(pre); //建临时resolv.conf
	mv_resolv(pre+"/resolv.conf"); //移动到系统目录

	body, err := Get_Hosts(); //修改文件
	if err != nil{
		ret = fmt.Sprintf("获取文件失败: %v", err);
		return;
	}

	rm_resolv(); //删除resolv.conf

	var begin_write = false; //写入临时host
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

	Exec_mv(pre+"/hosts"); //转换hosts为正式
	ret = fmt.Sprintf("更新完成!");
	fmt.Println(ret);
	return;
}
