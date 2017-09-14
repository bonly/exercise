package oms
import (
"fmt"
"flag"
"net/http"
"log"
"io/ioutil"
"sync"
"config"
"time"
"proto"
"encoding/json"
"strings"
// "runtime/pprof"
"strconv"
"runtime"
)
import _ "net/http/pprof" 

type SRV struct{
};

func (this *SRV)Srv(){
	flag.Parse();

	http.HandleFunc("/goroutines", 
		func(w http.ResponseWriter, r *http.Request) {  
        	num := strconv.FormatInt(int64(runtime.NumGoroutine()), 10);
        	w.Write([]byte(num));
    });

	http.HandleFunc("/cmd", this.cmd);
	err := http.ListenAndServe(*config.Oms_srv, nil)
	if err != nil {
		log.Fatal(err);
	}
}

func (this *SRV)cmd(wr http.ResponseWriter, rq *http.Request){
	// pprof.WriteHeapProfile(config.MemFile);
	wr.Header().Set("Access-Control-Allow-Origin", "*");
	wr.Header().Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");
	
	var resp PMS_CMD;
	body, err := ioutil.ReadAll(rq.Body);
	if err != nil{
		fmt.Printf("%v\n", err);
		return;
	}

	if len(body) <= 0{
		fmt.Printf("消息体空\n");
		return;
	}

	var wg sync.WaitGroup;
	wg.Add(2);
	ret, cmd_name := this.process(body, &wg);
	cmd_resp_name := strings.Replace(cmd_name, "REQ", "RESP", 1);

	go func(){
		defer func(){
			if err := recover(); err != nil{
				fmt.Printf("关闭通道失败, %v\n", err);
				return;
			}		
		}();
		defer func(){
			wg.Done();
			if ret != nil{
				close(ret);
			}
		}();
		for {
			select{
				case <- time.After(5 * time.Second):
					fmt.Printf("接收结果超时\n");
					timeout := PMS_Manual{
						Name : cmd_resp_name,
						ResultID : "1",
						Description : "超时",
					};
					pack, _ := json.MarshalIndent(timeout, " ", " ");
					wr.Write(pack);
					return;
				case result := <- ret:{
					if result == nil{
						fmt.Printf("收到nil\n");
						resp = &PMS_Command{};
						resp.New();
						*(resp.(*PMS_Command).Name) = cmd_resp_name;
						pack, _ := resp.Encode();
						wr.Write(pack);					
						return;
					}
					// 解析并把转换结果回写到客户端	
					ret_dat := this.parse_resp(result);
					pack, _ := ret_dat.Encode();

					wr.Write(pack);
					return;
				}
			}
		}
	}();
	wg.Wait();
	fmt.Printf("结束处理\n");
}

func (this *SRV)process(pack []byte, wg *sync.WaitGroup)(ret chan interface{}, cmd_name string){
	var cmd PMS_CMD;
	cmd = &PMS_Command{};
	err := cmd.Decode(pack);
	if err != nil{
		fmt.Printf("json wrong %v\n", err);
		return nil,"PROTO_ERR"; //todo 返回错误到chan
	}
	cmd_name = *(cmd.(*PMS_Command).Name);
	fmt.Printf("命令: %s\n", cmd_name);	
	ret = make(chan interface{});
	go func(){
		defer wg.Done();

		switch(cmd_name){
			case "Add_OpenUser_REQ":{
				fmt.Printf("新增用户\n");
				cmd = &Add_OpenUser_REQ{};
				break;
			}
			case "Delete_OpenUser_REQ":{
				fmt.Printf("删除用户\n");
				cmd = &Delete_OpenUser_REQ{};
				break;
			}
			case "OpenLock_REQ":{
				fmt.Printf("远程开门\n");
				cmd = &OpenLock_REQ{};
				break;
			}			
			case "Time_REQ":{
				fmt.Printf("时间校正\n");
				cmd = &Time_REQ{};
				break;
			}
		}
		err = cmd.Decode(pack);
		if err != nil{
			return; //todo 返回错误到chan
		}
		cmd.Process(ret);
	}();
	return ret, cmd_name;
}

func (this *SRV)parse_resp(inf interface{})(ret PMS_CMD){			
	switch inf.(type){
		case *proto.R_Remote_open_door:{
			fmt.Printf("收到远程开门的处理结果\n");
			var org *proto.R_Remote_open_door = inf.(*proto.R_Remote_open_door);
			var retp OpenLock_RESP;
			retp.New();
			if org.Head().Type == proto.LOCK_YES{
				retp.ResultID = "0";
				retp.Description = "成功";
			}else{
				retp.ResultID = "1";
				retp.Description = "失败";
			}
			ret = &retp;
			break;
		}
		case *proto.R_User_add_passwd:{
			fmt.Printf("收到增加用户(密码)的处理结果\n");
			var org *proto.R_User_add_passwd = inf.(*proto.R_User_add_passwd);
			var retp Add_OpenUser_RESP;
			retp.New();
			if org.Head().Type == proto.LOCK_YES{
				retp.ResultID = "0";
				retp.Description = "成功";
			}else{
				retp.ResultID = "1";
				retp.Description = "失败";
			}
			ret = &retp;
			break;
		}
		case *proto.R_User_del_passwd:{
			fmt.Printf("收到删除用户(密码)的处理结果\n");
			var org *proto.R_User_del_passwd = inf.(*proto.R_User_del_passwd);
			var retp Delete_OpenUser_RESP;
			retp.New();
			if org.Head().Type == proto.LOCK_YES{
				retp.ResultID = "0";
				retp.Description = "成功";
			}else{
				retp.ResultID = "1";
				retp.Description = "失败";
			}
			ret = &retp;
			break;
		}
		case *proto.R_Setup_lock_time:{
			fmt.Printf("收到校正时间的处理结果\n");
			var org *proto.R_Setup_lock_time = inf.(*proto.R_Setup_lock_time);
			var retp Time_RESP;
			retp.New();
			if org.Head().Type == proto.LOCK_YES{
				retp.ResultID = "0";
				retp.Description = "成功";
			}else{
				retp.ResultID = "1";
				retp.Description = "失败";
			}
			ret = &retp;
			break;
		}		
		default:{
			fmt.Printf("未知的应答结果\n");
			break;
		}
	}
	return ret;
}
