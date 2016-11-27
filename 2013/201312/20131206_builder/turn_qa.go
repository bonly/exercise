package main 

import (
"fmt"
// "net/http"
"golang.org/x/net/websocket"
// "open"
"glog"
"os/exec"
// "strings"
// "io/ioutil"
// "os"
// "path/filepath"
)


type Q_Turn_QA struct{
	Program_name string;
};
type R_Turn_QA struct{
	Ret;
};

func Cmd_Turn_qa(ws *websocket.Conn, it interface{}){
	ret := it.(*R_Turn_QA);
	ret.Ret.Cmd = "RTurn_QA";
	var qry Q_Turn_QA;
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

	if err = Turn_qa(qry.Program_name); err != nil{
		ret.Ret.Ret = "1";
		ret.Ret.Msg = "Failed";
	}
}

func Turn_qa(pro string) error{
	// file, _ := exec.LookPath("/apps/sh/rsync_test.sh");
    // path, _ := filepath.Abs(file);
    // glog.Info(path);

	cmd_argv := []string{"/apps/sh/rsync_test.sh", pro};
	cmd := exec.Command("sudo", cmd_argv...);
	cmd.Dir = "/apps/sh";
	// cmd.Path = "/apps/sh";

	out, err := cmd.CombinedOutput();
	
	// glog.Info(cmd);

	glog.Info(string(out));

	if err != nil{
		glog.Info(fmt.Printf("err: %s", err));
		return err;	
	}	
	return nil;
}

type Q_Restart_TomCat struct{
	Srv string;
	Sh string;	
};
type R_Restart_TomCat struct{
	Ret;
};
func Cmd_Restart_TomCat(ws *websocket.Conn, it interface{}){
	ret := it.(*R_Restart_TomCat);
	ret.Ret.Cmd = "RRestart_TomCat";
	var qry Q_Restart_TomCat;
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

	if err = Restart_tomcat(qry.Srv, qry.Sh); err != nil{
		ret.Ret.Ret = "1";
		ret.Ret.Msg = "Failed";
	}
}

func Restart_tomcat(srv string, sh string) error{
	cmd_argv := []string{srv, sh};
	cmd := exec.Command("ssh", cmd_argv...);

	out, err := cmd.CombinedOutput();
	
	// glog.Info(cmd);

	glog.Info(string(out));

	if err != nil{
		glog.Info(fmt.Printf("err: %s", err));
		return err;	
	}	
	return nil;
}