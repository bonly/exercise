/*
auth: bonly
*/
package main

import (
"log"
"os/exec"
"errors"
"flag"
"fmt"
)

func Get_Hosts() (err error){
	cmd_argv := []string{"https://raw.githubusercontent.com/racaljk/hosts/master/hosts", "-o", "/tmp/fetchedhosts"};
	cmd := exec.Command("curl", cmd_argv...);
	_, err = cmd.CombinedOutput();
	// log.Println(string(out));
	if err != nil{
		log.Println("fetch file err: ", err);
		return err;
	}
	return err;
}

func Modi_area(){
	cmd_argv := []string{"-i","s/# Copyright (c) 2014.*/# Author: bonly/g", "/tmp/fetchedhosts"};
	cmd := exec.Command("sed", cmd_argv...);
	_, err := cmd.CombinedOutput();
	// log.Println(string(out));
	if err != nil{
		log.Println("modi 1 err: ", err);
	}	

	cmd_argv = []string{"-i","s/# https.*//g", "/tmp/fetchedhosts"};
	cmd = exec.Command("sed", cmd_argv...);
	_, err = cmd.CombinedOutput();
	// log.Println(string(out));
	if err != nil{
		log.Println("modi 2 err: ", err);
	}	

	cmd_argv = []string{"-i","s/# This work is licensed.*//g", "/tmp/fetchedhosts"};
	cmd = exec.Command("sed", cmd_argv...);
	_, err = cmd.CombinedOutput();
	// log.Println(string(out));
	if err != nil{
		log.Println("modi 3 err: ", err);
	}			
}

func Del_area(){
	// cmd_argv := []string{"-i", "/# Copyright (c) 2014/,/# Modified hosts end/d", "/etc/hosts"};
	cmd_argv := []string{"-i", "/# Author: bonly/,/# Modified hosts end/d", "/etc/hosts"};
	cmd := exec.Command("sed", cmd_argv...);
	_, err := cmd.CombinedOutput();
	// log.Println(string(out));
	if err != nil{
		log.Println("del err: ", err);
	}
}

func Cat_area(){
	cmd_argv := []string{"-c", "cat /tmp/fetchedhosts >> /etc/hosts"};
	cmd := exec.Command("bash", cmd_argv...);
	_, err := cmd.CombinedOutput();
	// log.Println(string(out));
	if err != nil{
		log.Println("cat err: ", err);
	}
}

func Rm_file(){
	cmd_argv := []string{"-f", "/tmp/fetchedhosts"};
	cmd := exec.Command("rm", cmd_argv...);
	_, err := cmd.CombinedOutput();
	// log.Println(string(out));
	if err != nil{
		log.Println("rm err: ", err);
	}	
}

func Chk_root()(err error){
	cmd_argv := []string{"-u"};
	cmd := exec.Command("id", cmd_argv...);
	out, err := cmd.CombinedOutput();
	ret := string(out); //无string转换时，原数据带\n
	if err != nil{
		log.Printf("操作员： %+v",ret);
		log.Println("err: ", err);
		return err;
	}	
	if ret != "0\n" {
		log.Printf("操作员： %+v",ret);
		return errors.New("无操作权限，请用root执行!");
	}
	return err;
}

func main(){
	version := flag.Bool("V", false, "Version");

	flag.Parse();

	if *version == true {
		fmt.Println("Version: 1.0\nAuthor: bonly@163.com");
		return;
	}

	err := Chk_root();
	if err != nil{
		log.Println(err);
		return;
	}
	if Get_Hosts() != nil{
		return;
	}
	Modi_area();
	Del_area();
	Cat_area();
	Rm_file();
}

/*
GOOS=darwin go build
*/