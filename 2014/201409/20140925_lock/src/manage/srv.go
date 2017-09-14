/*
auth: bonly
create: 2016.9.19
desc: 盒子管理服务
*/
package manage

import (
"log"
"fmt"
"net"
"sync"
"time"
"reflect"
. "proto"
"config"
)

type SRV struct{
};

func (this *SRV)Srv(){
	srv, err := net.ResolveTCPAddr("tcp4", *config.Box_srv);
	if err != nil{
		log.Printf("Box服务地址解释起动失败 %v\n", err);
		return;
	}
	listener, err := net.ListenTCP("tcp", srv);
	if err != nil{
		log.Printf("Box服务地址侦听失败 %v\n", err);
		return;
	}

	go func(){
		for {
			select {
				case stat := <- config.Run_chn:{
					config.Run = stat;
				}
			}
		}
	}();
	for config.Run { //从全局程序上控制，其实有break已可以
		// listener.SetDeadline(time.Now().Add(30 * time.Second));

		conn, err := listener.Accept();
		if err != nil{
			fmt.Printf("listen : %s\n", err.Error());
			continue; //正式环境应继续
			// listener.Close();
			// break; //测试时可以结束了
		}
		fmt.Printf("收到新的连接\n");
		go this.handle_connect(conn);
	}
}

func (this *SRV)handle_connect(conn net.Conn){
	defer conn.Close();
	var wg sync.WaitGroup;
	wg.Add(3);

	//接收消息
	reg, task := this.recv_Msg(conn, &wg);

	//处理消息
	this.proc_task(conn, task, &wg);
	this.proc_reg(conn, reg, &wg);

	//应答消息

	fmt.Printf("设置Box任务处理完毕\n");
	wg.Wait();
	fmt.Printf("所有Box任务完成\n");	
	config.Run_chn <- false;
}

func (this *SRV)recv_Msg(conn net.Conn, wg *sync.WaitGroup)(reg chan interface{}, task chan interface{}){
	task = make(chan interface{});
	reg  = make(chan interface{});
	go func(){
		defer wg.Done();
		for {
			buf := make([]byte, 255);
			// if *Test_mode {
				conn.SetDeadline(time.Now().Add(60 * time.Second));
			// }
			hl_len, err := conn.Read(buf);
			if err != nil{
				fmt.Printf("读取包头失败 %s\n", err.Error());
				close(task);
				close(reg);
				Works().Lost(conn);
				return;
			}
			fmt.Printf("recv[%d]: ", hl_len);

			if hl_len == 5{  //==5的是注册盒子
				// var box *Register;
				// box = (*Register)(unsafe.Pointer(&buf[0]));
				var box Register;
				box.Decode(buf);
				fmt.Printf("收到盒子心跳: \n");
				Hex_Dump(buf[:hl_len], hl_len);
				reg <- box;
			}else{ //大于5的是协议
				var zh Command;
				cmd := zh.Decode(buf);
				if cmd == nil{
					fmt.Printf("收到错误的数据包\n");
					return;//todo 需要关闭连接？
				}
				fmt.Printf("收到协议数据包: \n");
				Hex_Dump(buf[:hl_len], hl_len);
				task <-zh;
			}
		}
	}();

	return reg, task;
}

func (this *SRV)proc_task(conn net.Conn, task chan interface{}, wg *sync.WaitGroup){
	fmt.Printf("Box任务处理器就绪\n");
	go func(){
		defer wg.Done();
		for {
			select{
				case work := <- task:{
					if work == nil{
						return;
					}
					var answer CMD;
					var resp CMD = nil;
					cmd := work.(Command);
					cmd_param := cmd.Head().Get_CMD_PARAM();
					switch {
						case reflect.DeepEqual(cmd_param, LOCK_STAT_CMD_PARAM):{
							answer = &R_Lock_stat{};
							break;
						}
						case reflect.DeepEqual(cmd_param, REMOTE_OPEN_DOOR_CMD_PARAM):{
							answer = &R_Remote_open_door{};
							break;
						}
						case reflect.DeepEqual(cmd_param, SETUP_LOCK_TIME_CMD_PARAM):{
							answer = &R_Setup_lock_time{};
							break;
						}		
						case reflect.DeepEqual(cmd_param, SETUP_LOCK_AUTHORITY_CMD_PARAM):{
							answer = &R_Setup_lock_authority{};
							break;
						}	
						case reflect.DeepEqual(cmd_param, SETUP_LOCK_CLEAN_LOG_CMD_PARAM):{
							answer = &R_Setup_lock_clean_log{};
							break;
						}		
						case reflect.DeepEqual(cmd_param, LOCK_UPLOAD_LOG_CMD_PARAM):{
							answer = &Lock_upload_log{};
							resp = &R_Lock_upload_log{};
							resp.New();
							break;
						}	
						case reflect.DeepEqual(cmd_param, USER_ADD_PASSWD_CMD_PARAM):{
							answer = &R_User_add_passwd{};
							break;
						}	
						case reflect.DeepEqual(cmd_param, USER_DEL_PASSWD_CMD_PARAM):{
							answer = &R_User_del_passwd{};
							break;
						}		
						case reflect.DeepEqual(cmd_param, USER_DEL_ALL_CMD_PARAM):{
							answer = &R_User_del_all{};
							break;
						}																																								
						default:{
							fmt.Printf("未知数据包\n");
							return;
						}
					}
					answer.Decode(cmd.Buf());
					if cnn, ok := Works().Conn_map[conn]; ok{
						if reflect.DeepEqual(cnn.Cur_task, cmd_param) && cnn.Result!= nil{
							defer func(){
								if err := recover(); err != nil{
									fmt.Printf("通道在前端已关闭, %v\n", err);
									Works().Lost(conn);
									return;
								}
							}();
							cnn.Result <- answer; //回写结果包到通道
						}
					}

					if resp != nil{
						rep := resp.(*R_Lock_upload_log);
						rep.Head().LockID = cmd.Head().LockID;
						rep.Verify(false);
						pack, err := resp.Encode();
						cnt, err := conn.Write(pack);
						if err != nil{
							fmt.Printf("send %v\n", err);
						}else{
							fmt.Printf("send len[%d]:\n", cnt);
							Hex_Dump(pack, len(pack));
						}			
					}
				}
			}
		}
	}();
}

func (this *SRV)proc_reg(conn net.Conn, reg chan interface{}, wg *sync.WaitGroup){
	fmt.Printf("Box注册处理器就绪\n");
	go func(){
		defer wg.Done();
		for {
			select{
				case work := <- reg:{
					if work == nil{
						return;
					}
					fmt.Printf("%+v\n", work);
					Works().Health(work.(Register).Ctl, conn);	
				}
			}
		}
	}();
}