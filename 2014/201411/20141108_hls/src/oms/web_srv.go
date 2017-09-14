package oms
import (
"fmt"
"flag"
"net/http"
// "log"
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
log "glog"
"bcd"
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
		log.Error(err.Error());
	}
}

func (this *SRV)cmd(wr http.ResponseWriter, rq *http.Request){
	//defer rq.Body.Close();
	// pprof.WriteHeapProfile(config.MemFile);
	wr.Header().Set("Access-Control-Allow-Origin", "*");
	wr.Header().Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");
	
	var resp PMS_CMD;
	body, err := ioutil.ReadAll(rq.Body);
	if err != nil{
		log.Warningf("%v\n", err);
		return;
	}

	if len(body) <= 0{
		log.Warningf("消息体空\n");
		return;
	}

	var wg sync.WaitGroup;
	ret, cmd_name := this.process(body, &wg);
	cmd_resp_name := strings.Replace(cmd_name, "REQ", "RESP", 1);

	go func(){
		wg.Add(1);
		defer func(){
			if err := recover(); err != nil{
				log.Warningf("关闭通道失败, %v\n", err);
				return;
			}		
		}();
		defer func(){
			wg.Done();
			if ret != nil{
				log.Infof("关闭结果通道\n");
				close(ret);
			}
		}();
		for {
			select{
				case <- time.After(time.Duration(*config.Timeout) * time.Second):
					log.Warningf("接收结果超时\n");
					timeout := PMS_Manual{
						Name : cmd_resp_name,
						ResultID : "1",
						Description : "超时",
					};
					pack, _ := json.MarshalIndent(timeout, " ", " ");
					wr.Write(pack);
					return;
				case result, ok := <- ret:{
					if result == nil || !ok{
						log.Warningf("收到nil\n");
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
	log.Infof("结束OMS请求处理\n");
}

func (this *SRV)process(pack []byte, wg *sync.WaitGroup)(ret chan interface{}, cmd_name string){
	wg.Add(1);
	var cmd PMS_CMD;
	cmd = &PMS_Command{};
	err := cmd.Decode(pack);
	if err != nil{
		log.Warningf("json wrong %v\n", err);
		ret <- nil;
		return nil,"PROTO_ERR"; //todo 返回错误到chan
	}
	cmd_name = *(cmd.(*PMS_Command).Name);
	log.Infof("命令: %s\n", cmd_name);	
	ret = make(chan interface{});
	go func(){
		defer wg.Done();

		switch(cmd_name){
			case "Add_OpenUser_REQ":{
				log.Infof("新增用户\n");
				cmd = &Add_OpenUser_REQ{};
				break;
			}
			case "Delete_OpenUser_REQ":{
				log.Infof("删除用户\n");
				cmd = &Delete_OpenUser_REQ{};
				break;
			}
			case "User_Clean_REQ":{
				log.Infof("删除所有用户\n");
				cmd = &User_Clean_REQ{};
				break;
			}
			case "OpenLock_REQ":{
				log.Infof("远程开门\n");
				cmd = &OpenLock_REQ{};
				break;
			}			
			case "Time_REQ":{
				log.Infof("时间校正\n");
				cmd = &Time_REQ{};
				break;
			}
			case "Lock_stat_REQ":{
				log.Infof("门锁状态\n");
				cmd = &Lock_stat_REQ{};
				break;
			}		
			case "Setup_Auth_REQ":{
				log.Infof("设置门禁\n");
				cmd = &Setup_Auth_REQ{};
				break;
			}		
			case "Clean_Log_REQ":{
				log.Infof("清空开门记录\n");
				cmd = &Clean_Log_REQ{};
				break;	
			}				
		}
		err = cmd.Decode(pack);
		if err != nil{
			ret <- nil;
			return; //todo 返回错误到chan
		}
		cmd.Process(ret);
	}();
	return ret, cmd_name;
}

func (this *SRV)parse_resp(inf interface{})(ret PMS_CMD){			
	switch inf.(type){
		case *proto.R_Remote_open_door:{
			log.Infof("收到远程开门的处理结果");
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
			log.Infof("[%s]\n", retp.Description);
			ret = &retp;
			break;
		}
		case *proto.R_User_add_passwd:{
			log.Infof("收到增加用户(密码)的处理结果");
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
			log.Infof("[%s]\n", retp.Description);
			ret = &retp;
			break;
		}
		case *proto.R_User_add_passwd_id:{
			log.Infof("收到增加用户ID(密码)的处理结果");
			var org *proto.R_User_add_passwd_id = inf.(*proto.R_User_add_passwd_id);
			var retp Add_OpenUser_RESP;
			retp.New();
			if org.Head().Type == proto.LOCK_YES{
				retp.ResultID = "0";
				retp.Description = "成功";
			}else{
				retp.ResultID = "1";
				retp.Description = "失败";
			}
			log.Infof("[%s]\n", retp.Description);
			ret = &retp;
			break;
		}
		case *proto.R_User_del_passwd:{
			log.Infof("收到删除用户(密码)的处理结果");
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
			log.Infof("[%s]\n", retp.Description);
			ret = &retp;
			break;
		}
		case *proto.R_User_del_all:{
			log.Infof("收到删除所有用户的处理结果");
			var org *proto.R_User_del_all = inf.(*proto.R_User_del_all);
			var retp User_Clean_RESP;
			retp.New();
			if org.Head().Type == proto.LOCK_YES{
				retp.ResultID = "0";
				retp.Description = "成功";
			}else{
				retp.ResultID = "1";
				retp.Description = "失败";
			}
			log.Infof("[%s]\n", retp.Description);
			ret = &retp;
			break;
		}		
		case *proto.R_Setup_lock_time:{
			log.Infof("收到校正时间的处理结果");
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
			log.Infof("[%s]\n", retp.Description);
			ret = &retp;
			break;
		}	
		case *proto.R_Setup_lock_authority:{
			log.Infof("收到设置门禁的处理结果");
			var org *proto.R_Setup_lock_authority = inf.(*proto.R_Setup_lock_authority);
			var retp Setup_Auth_RESP;
			retp.New();
			if org.Head().Type == proto.LOCK_YES{
				retp.ResultID = "0";
				retp.Description = "成功";
			}else{
				retp.ResultID = "1";
				retp.Description = "失败";
			}
			log.Infof("[%s]\n", retp.Description);
			ret = &retp;
			break;
		}			
		case *proto.R_Setup_lock_clean_log:{
			log.Infof("收到清空开门记录的处理结果");
			var org *proto.R_Setup_lock_clean_log = inf.(*proto.R_Setup_lock_clean_log);
			var retp Clean_Log_RESP;
			retp.New();
			if org.Head().Type == proto.LOCK_YES{
				retp.ResultID = "0";
				retp.Description = "成功";
			}else{
				retp.ResultID = "1";
				retp.Description = "失败";
			}
			log.Infof("[%s]\n", retp.Description);
			ret = &retp;
			break;
		}				
		case *proto.R_Lock_stat:{
			log.Infof("收到门锁状态的处理结果");
			var org *proto.R_Lock_stat = inf.(*proto.R_Lock_stat);
			var retp Lock_stat_RESP;
			retp.New();
			if org.Head().Type == proto.LOCK_YES{
				retp.ResultID = "0";
				retp.Description = "成功";
				retp.Card_user_count = fmt.Sprintf("%d", (int)(org.Data.Card_user_count));
				retp.Passwd_user_count = fmt.Sprintf("%d", (int)(org.Data.Passwd_user_count));
				retp.Power = (string)(org.Data.Power[0]) + (string)(org.Data.Power[1]);
				loc, _ := time.LoadLocation("Local");
				tn := time.Date( (int)(org.Data.Now_date[0])+2000, //年
				             (time.Month)(org.Data.Now_date[1]), //月
				             (int)(org.Data.Now_date[2]), //日
				             (int)(bcd.BcdToInt((int)(org.Data.Now_time[0]))), //时
				             (int)(bcd.BcdToInt((int)(org.Data.Now_time[1]))), //分
				             (int)(bcd.BcdToInt((int)(org.Data.Now_time[2]))), //秒
				             0, loc);
				retp.Now = tn.Format("2006-01-02 15:04:05");
				if org.Data.Authority == (uint8)(proto.ALLOW_OPEN) {
					retp.Authority = "Allow Open";
				}else{
					retp.Authority = "Forbit Open";
				}
			}else{
				retp.ResultID = "1";
				retp.Description = "失败";
			}
			log.Infof("[%s]\n", retp.Description);
			ret = &retp;
			break;
		}					
		default:{
			log.Warningf("未知的应答结果\n");
			break;
		}
	}
	return ret;
}
