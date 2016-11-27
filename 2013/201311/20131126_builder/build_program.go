package main 

import (
"fmt"
// "net/http"
"golang.org/x/net/websocket"
// "open"
"glog"
"os/exec"
// "strings"
"io/ioutil"
)

type Q_Build_program struct{
	Git_Root string;
	Git_Work_Tree string;
	Tag string;
};
type R_Build_program struct{
	Ret;
};
func Cmd_Build_program(ws *websocket.Conn, it interface{}){
	ret := it.(*R_Build_program);
	ret.Ret.Cmd = "RBuild_program";
	var qry Q_Build_program;
	err := websocket.JSON.Receive(ws, &qry);
	if err != nil{
		glog.Info("qry body error");
		ret.Ret.Ret = "1";
		ret.Ret.Msg = "body error";
		return;
	}
	glog.Info(qry);

	ret.Ret.Ret = "0";
	ret.Ret.Msg = "OK";

	if err = Build_program(&qry); err != nil{
		ret.Ret.Ret = "1";
		ret.Ret.Msg = "Failed";
	}
}

func Build_program(qry *Q_Build_program) error{
	//建目录
	cmd_argv := []string{"-p", qry.Git_Work_Tree};
	cmd := exec.Command("mkdir", cmd_argv...);
	out, err := cmd.CombinedOutput();
	glog.Info(string(out));
	if err != nil{
		glog.Info("err: ", err);
		return err;
	}

	//取出代码
	if err = git_checkout(qry); err != nil{
		return err;
	}

	//设置版本
	if set_version(qry); err != nil{
		return err;
	}

	//编译
	cmd = exec.Command("make");
	cmd.Dir = qry.Git_Work_Tree;
	out, err = cmd.CombinedOutput();
	glog.Info(string(out));
	if err != nil{
		glog.Info("err: ", err);	
		return err;
	}	

	return nil;
}

func git_checkout(qry *Q_Build_program) error{
	cmd_argv := []string{"checkout", "-f", qry.Tag};
	cmd := exec.Command("git", cmd_argv...);
	cmd.Env = []string{"GIT_DIR=" + qry.Git_Root,
					   "GIT_WORK_TREE=" + qry.Git_Work_Tree};

	out, err := cmd.CombinedOutput();
	
	glog.Info(cmd);

	glog.Info(string(out));

	if err != nil{
		glog.Info("err: ", err);
		return err;	
	}	
	return nil;
}

func set_version(qry *Q_Build_program) error{
	//取版本
	//git describe --abbrev=0 --tags
	cmd_argv := []string{"describe", "--abbrev=0", "--tags"};
	cmd := exec.Command("git", cmd_argv...);
	cmd.Env = []string{"GIT_DIR=" + qry.Git_Root,
					   "GIT_WORK_TREE=" + qry.Git_Work_Tree};

	out, err := cmd.Output();
	if err != nil{
		glog.Info("err: ", err);
		return err;	
	}	
	strver := "V"+string(out);
	glog.Info("build ver: ", strver);

	//git describe --always
	cmd_argv = []string{"describe", "--always"};
	cmd = exec.Command("git", cmd_argv...);
	cmd.Env = []string{"GIT_DIR=" + qry.Git_Root,
					   "GIT_WORK_TREE=" + qry.Git_Work_Tree};

	out, err = cmd.Output();
	if err != nil{
		glog.Info("err: ", err);
		return err;	
	}	
	strcode := "C"+string(out);
	glog.Info("build code: ", strcode);

    //git log --pretty=format:%cd --date=short -n1
	cmd_argv = []string{"log", "--pretty=format:%cd", "--date=short", "-n1"};
	cmd = exec.Command("git", cmd_argv...);
	cmd.Env = []string{"GIT_DIR=" + qry.Git_Root,
					   "GIT_WORK_TREE=" + qry.Git_Work_Tree};

	out, err = cmd.Output();
	if err != nil{
		glog.Info("err: ", err);
		return err;	
	}	
	strtimes := "T"+string(out);
	glog.Info("build times: ", strtimes);

	//git describe --tags
	cmd_argv = []string{"describe", "--tags"};
	cmd = exec.Command("git", cmd_argv...);
	cmd.Env = []string{"GIT_DIR=" + qry.Git_Root,
					   "GIT_WORK_TREE=" + qry.Git_Work_Tree};

	out, err = cmd.Output();
	if err != nil{
		glog.Info("err: ", err);
		return err;	
	}	
	strtag := string(out);
	glog.Info("build times: ", strtag);


	//输出版本
	ver := fmt.Sprintf("APP_VERSION:=%s\nAPP_CODE:=%s\nAPP_TIME:=%s\nAPP_TAG:=%s\n", 
		strver, strcode, strtimes, strtag);
	ioutil.WriteFile(qry.Git_Work_Tree + "/makefile.init", []byte(ver), 644);

	return nil;	
}