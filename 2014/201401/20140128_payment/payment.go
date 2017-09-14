/*
auth: bonly
create: 2015.12.18
*/

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
"github.com/jmoiron/sqlx"
_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB;

func main(){
	flag.Parse();
	defer glog.Flush();
	flag.Set("alsologtostderr", "true");
	
	var err error;
	db, err = sqlx.Open("mysql", "db_writer:XqH/a5aOnAlw@tcp(112.74.195.114:3306)/xbed_service?charset=utf8");
	if err != nil{
		glog.Info(err);
		return;
	}

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))));
	http.Handle("/Main", websocket.Handler(Main));

	go open.Run("http://127.0.0.1:9998/");

	err = http.ListenAndServe(":9998", nil);
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
		case "Query":
			ret = &R_Query{};
			Cmd_Query(ws, ret);
			break;
		}
		websocket.JSON.Send(ws, &ret);
	}
}