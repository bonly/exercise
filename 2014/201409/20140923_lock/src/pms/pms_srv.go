package pms
import (
"testing"
"log"
"fmt"
// // "unsafe"
"flag"
"net"
"sync"
"time"
// "reflect"
// . "proto"
// . "manage"
)

type PMS struct{
};

var PMS_addr = flag.String("pms_srv", "0.0.0.0:5021", "作为PMS服务器地址及端口");

func (this *PMS)srv(ts *testing.T){
	flag.Parse();

	srv, err := net.ResolveTCPAddr("tcp4", *PMS_addr);
	if err != nil{
		log.Printf("srv start failed %v\n", err);
		return;
	}
	listener, err := net.ListenTCP("tcp", srv);
	if err != nil{
		log.Printf("srv listen failed %v\n", err);
		return;
	}

	for run { //从全局程序上控制，其实有break已可以
		if *Test_mode {
			// listener.SetDeadline(time.Now().Add(30 * time.Second));
		}
		conn, err := listener.Accept();
		if err != nil{
			fmt.Printf("listen : %s\n", err.Error());
			// continue; //正式环境应继续
			listener.Close();
			break; //测试时可以结束了
		}
		fmt.Printf("收到PMS新的连接\n");
		go this.handle_connect(conn);
	}
}

func (this *PMS)handle_connect(conn net.Conn){
	defer conn.Close();
	var wg sync.WaitGroup;

	//接收消息
	wg.Add(1);
	// task := this.recv_Msg(conn, &wg);
	this.recv_Msg(conn, &wg);

	//处理消息
	// wg.Add(1);
	// this.proc_task(conn, task, &wg);

	//应答消息

	fmt.Printf("设置PMS任务处理完毕\n");
	wg.Wait();
	fmt.Printf("所有PMS任务完成\n");	
	run_chn <- false;
}

func (this *PMS)recv_Msg(conn net.Conn, wg *sync.WaitGroup)(task chan interface{}){
	task = make(chan interface{});
	go func(){
		defer wg.Done();
		for {
			buf := make([]byte, 255);
			if *Test_mode {
				conn.SetDeadline(time.Now().Add(60 * time.Second));
			}
			hl_len, err := conn.Read(buf);
			if err != nil{
				fmt.Printf("read head failed %s\n", err.Error());
				close(task);
				conn.Close();
				return;
			}
			fmt.Printf("recv[%d]: ", hl_len);

			fmt.Printf("%s", string(buf));
			// task <-zh;
		}
	}();

	return task;
}

// func (this *PMS)proc_task(conn net.Conn, task chan interface{}, wg *sync.WaitGroup){
// 	fmt.Printf("任务处理器就绪\n");
// 	go func(){
// 		defer wg.Done();
// 		for {
// 			select{
// 				case work := <- task:{
// 					if work == nil{
// 						return;
// 					}
// 					var answer CMD;
// 					var resp CMD = nil;
// 					cmd := work.(Command);
// 					cmd_param := cmd.Head().Get_CMD_PARAM();
// 					switch {
// 						case reflect.DeepEqual(cmd_param, LOCK_STAT_CMD_PARAM):{
// 							answer = &R_Lock_stat{};
// 							break;
// 						}
// 						case reflect.DeepEqual(cmd_param, REMOTE_OPEN_DOOR_CMD_PARAM):{
// 							answer = &R_Remote_open_door{};
// 							break;
// 						}
// 						case reflect.DeepEqual(cmd_param, SETUP_LOCK_TIME_CMD_PARAM):{
// 							answer = &R_Setup_lock_time{};
// 							break;
// 						}		
// 						case reflect.DeepEqual(cmd_param, SETUP_LOCK_AUTHORITY_CMD_PARAM):{
// 							answer = &R_Setup_lock_authority{};
// 							break;
// 						}	
// 						case reflect.DeepEqual(cmd_param, SETUP_LOCK_CLEAN_LOG_CMD_PARAM):{
// 							answer = &R_Setup_lock_clean_log{};
// 							break;
// 						}		
// 						case reflect.DeepEqual(cmd_param, LOCK_UPLOAD_LOG_CMD_PARAM):{
// 							answer = &Lock_upload_log{};
// 							resp = &R_Lock_upload_log{};
// 							resp.New();
// 							break;
// 						}	
// 						case reflect.DeepEqual(cmd_param, USER_ADD_PASSWD_CMD_PARAM):{
// 							answer = &R_User_add_passwd{};
// 							break;
// 						}	
// 						case reflect.DeepEqual(cmd_param, USER_DEL_PASSWD_CMD_PARAM):{
// 							answer = &R_User_del_passwd{};
// 							break;
// 						}		
// 						case reflect.DeepEqual(cmd_param, USER_DEL_ALL_CMD_PARAM):{
// 							answer = &R_User_del_all{};
// 							break;
// 						}																																								
// 						default:{
// 							fmt.Printf("未知数据包\n");
// 							return;
// 						}
// 					}
// 					answer.Decode(cmd.Buf());
// 					if resp != nil{
// 						rep := resp.(*R_Lock_upload_log);
// 						rep.Head().LockID = cmd.Head().LockID;
// 						rep.Verify(false);
// 						pack, err := resp.Encode();
// 						cnt, err := conn.Write(pack);
// 						if err != nil{
// 							fmt.Printf("send %v\n", err);
// 						}else{
// 							fmt.Printf("send len[%d]:\n", cnt);
// 							Hex_Dump(pack, len(pack));
// 						}			
// 					}
// 				}
// 			}
// 		}
// 	}();
// }
