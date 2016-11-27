package main 

import (
"fmt"
"net/http"
"golang.org/x/net/websocket"
"open"
"glog"
// "os/exec"
// "strings"
"flag"
)

func main(){
	flag.Parse();
	defer glog.Flush();
	flag.Set("alsologtostderr", "true");
	
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))));
	http.Handle("/Main", websocket.Handler(Main));

	go open.Run("http://127.0.0.1:9998/");

	err := http.ListenAndServe(":9998", nil);
	if err != nil{
		panic(err.Error());
	}
}

type Head struct{
	Cmd string;
};

type Ret struct{
	Cmd string;
	Ret string;
	Msg string;
};

func Main(ws *websocket.Conn){
	glog.Info(fmt.Sprintf("Get a connect from %s", ws.RemoteAddr().String()));

	for{
		var ret interface{};

		//取消息头
		var head Head;
		err := websocket.JSON.Receive(ws, &head);
		glog.Info(fmt.Sprintf("Head: %s", head));
		if err != nil{
			glog.Info(err);
			ret = Ret{"","-1","command not found"};
			websocket.JSON.Send(ws, &ret);
			return;
		}

		fmt.Println("recv qry ");
		switch(head.Cmd){
		case "Tag_list":
			ret = &R_Tag_list{};
			Cmd_Tag_list(ws, ret);
			break;
		case "Build_program":
			ret = &R_Build_program{};
			Cmd_Build_program(ws, ret);
			break;
		case "Turn_QA":
			ret = &R_Turn_QA{};
			Cmd_Turn_qa(ws, ret);
			break;
		case "Restart_TomCat":
			ret = &R_Restart_TomCat{};
			Cmd_Restart_TomCat(ws, ret);
		}
		websocket.JSON.Send(ws, &ret);
	}
}




