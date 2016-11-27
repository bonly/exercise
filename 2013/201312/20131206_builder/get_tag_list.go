package main 

import (
"fmt"
// "net/http"
"golang.org/x/net/websocket"
// "open"
"glog"
"os/exec"
"strings"
)

type Q_Tag_list struct{
	Git_Root string;
};
type C_Tag struct{
	Tag string;
};
type R_Tag_list struct{
	Ret;
	Ver []C_Tag;
};
func Cmd_Tag_list(ws *websocket.Conn, it interface{}){
	ret := it.(*R_Tag_list);
	ret.Ret.Cmd = "RTag_list";
	var qry Q_Tag_list;
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

	git_get_tag_list(qry, ret);
}

func git_get_tag_list(qry Q_Tag_list, ret *R_Tag_list){
	cmd_argv := []string{"tag"};
	cmd := exec.Command("git", cmd_argv...);
	cmd.Env = []string{"GIT_DIR=" + qry.Git_Root};

	out, err := cmd.Output();
	if err != nil{
		glog.Info(fmt.Printf("err: %s", err));
		ret.Ret.Ret = "2";
		ret.Ret.Msg = "Get tag failed: " + err.Error();
	}

	// rt := strings.Trim(string(out), " ");
	rt := string(out);
	lines := strings.Split(string(rt), "\n");

	for _, oneLine := range lines {
		if len(oneLine) <= 0{
			continue;
		}
		ret.Ver = append(ret.Ver, C_Tag{oneLine});
	}
}