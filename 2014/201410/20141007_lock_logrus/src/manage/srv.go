/*
auth: bonly
create: 2016.9.19
desc: 盒子管理服务
*/
package manage

import (
// "log"
// "fmt"
"net"
"sync"
"time"
"reflect"
. "proto"
"config"
log "logrus"
)

type SRV struct{
};

func (this *SRV)Srv(){
	srv, err := net.ResolveTCPAddr("tcp4", *config.Box_srv);
	if err != nil{
		log.Warnf("Box服务地址解释起动失败 %v\n", err);
		return;
	}
	listener, err := net.ListenTCP("tcp", srv);
	if err != nil{
		log.Warnf("Box服务地址侦听失败 %v\n", err);
		return;
	}

	for config.Run { //从全局程序上控制，其实有break已可以
		log.WithFields(log.Fields{"IP_PORT":*config.Box_srv}).Info("盒子管理服务就绪");
		// listener.SetDeadline(time.Now().Add(30 * time.Second));

		conn, err := listener.Accept();
		if err != nil{
			log.Warnf("listen : %s\n", err.Error());
			continue; //正式环境应继续
			// listener.Close();
			// break; //测试时可以结束了
		}
		log.Infof("收到新的连接\n");
		var connect Connect;
		connect.Conn = conn;
		go connect.Handle_connect();
	}
}

type Connect struct{
	Conn net.Conn;
	task chan interface{};
	// reg  chan interface{};	
	run bool;
};

func (this *Connect)close(){
	log.Infof("退出清理开始\n");
	this.run = false;
	var wg sync.WaitGroup;
	wg.Add(2);	
	go func(){
		defer func(){
			wg.Done();
			return;
		}();
		this.Conn.Close();
		log.Infof("关闭socket连接\n");
	}();
	go func(){
		defer func(){
			if err := recover(); err != nil{
				log.Warnf("关闭连接的任务通道失败, %v\n", err);
			}	
			wg.Done();
			return;
		}();
		close(this.task); //关闭任务通道
		log.Infof("关闭连接的任务通道\n");
	}();
	// go func(){
	// 	defer func(){
	// 		if err := recover(); err != nil{
	// 			log.Warnf("关闭连接的注册通道失败, %v\n", err);
	// 		}	
	// 		wg.Done();
	// 		return;
	// 	}();
	// 	close(this.reg);//关闭注册通道
	// 	log.Warnf("关闭连接的注册通道\n");
	// }();
	wg.Wait();
	Works().Lost(this.Conn);
	log.Infof("退出清理结束\n");	
}

func (this *Connect)Handle_connect(){
	defer func(){
		if err := recover(); err != nil{
			log.Warnf("处理连接失败, %v\n", err);
			return;
		}
	}();	
	this.task = make(chan interface{});
	// this.reg  = make(chan interface{});
	this.run = true;
	defer this.close();

	var wg sync.WaitGroup;
	wg.Add(3);

	//接收消息
	err := this.recv_Msg(this.Conn, &wg);
	if err != nil{
		return;
	}

	//处理消息
	this.proc_task(this.Conn, &wg);
	// this.proc_reg(this.Conn, &wg);

	//应答消息

	log.Infof("设置Box任务处理完毕\n");
	wg.Wait();
	log.Infof("所有Box任务完成\n");	
	
}

func (this *Connect)recv_Msg(conn net.Conn, wg *sync.WaitGroup)(err error){
	go func(){
		defer this.close();
		defer func(){
			if err := recover(); err != nil{
				log.Warnf("处理消息失败, %v\n", err);
				return;
			}
		}();		
		defer wg.Done();
		for this.run{
			buf := make([]byte, 255);
			// if *Test_mode {
				conn.SetDeadline(time.Now().Add(60 * time.Second));
			// }
			hl_len, err := conn.Read(buf);
			if err != nil{
				log.Warnf("读取包头失败 %s\n", err.Error());
				return;
			}
			log.WithFields(log.Fields{"Len":hl_len}).
				Infof("收到数据包");

			if hl_len == 5{  //==5的是注册盒子
				var box Register;
				box.Decode(buf);
				log.Debugf("收到盒子心跳: \n");
				Hex_Dump(buf[:hl_len], hl_len);
				// this.reg <- box;
				Works().Health(box.Ctl, conn);	
			}else{ //大于5的是协议
				var zh Command;
				cmd := zh.Decode(buf);
				if cmd == nil{
					log.Warnf("收到错误的数据包\n");
					return;
				}
				log.Infof("收到协议数据包: \n");
				Hex_Dump(buf[:hl_len], hl_len);
				this.task <-zh;
			}
		}
	}();

	return nil;
}

func (this *Connect)proc_task(conn net.Conn, wg *sync.WaitGroup){
	log.Infof("Box任务处理器就绪\n");
	go func(){
		defer this.close();
		defer func(){
			if err := recover(); err != nil{
				log.Warnf("处理任务消息失败, %v\n", err);
			}
		    wg.Done();	
		    log.Infof("Box处理任务器结束\n");		
		    return;
		}();		

		for this.run{
			select{
				case work, ok := <- this.task:{
					if !ok {
						log.Warnf("接收任务消息失败,准备退出\n");
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
							log.Warnf("未知数据包\n");
							this.run = false;
							return;
						}
					}
					answer.Decode(cmd.Buf());
					if cnn, ok := Works().Conn_map[conn]; ok{
						if reflect.DeepEqual(cnn.Cur_task, cmd_param) && cnn.Result!= nil{
							defer func(){
								if err := recover(); err != nil{
									log.Warnf("通道在前端已关闭, %v\n", err);
									this.run = false;
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
							log.Warnf("send %v\n", err);
							this.run = false;
							return;
						}else{
							log.Infof("send len[%d]:\n", cnt);
							Hex_Dump(pack, len(pack));
						}			
					}
				}
			}
		}
	}();
}

// func (this *Connect)proc_reg(conn net.Conn, wg *sync.WaitGroup){
// 	log.Infof("Box注册处理器就绪\n");
// 	go func(){
// 		defer this.close();
// 		defer func(){
// 			if err := recover(); err != nil{
// 				log.Warnf("处理注册消息失败, %v\n", err);
// 			}
// 			wg.Done();
// 			log.Infof("Box注册处理器结束\n");
// 			return;
// 		}();		
// 		for this.run{
// 			select{
// 				case work, ok := <- this.reg:{
// 					if !ok {
// 						log.Warnf("接收注册消息失败,准备退出\n");
// 						return;
// 					}
// 					log.Infof("接收到的注册: %+v\n", work);
// 					Works().Health(work.(Register).Ctl, conn);	
// 				}
// 			}
// 		}
// 	}();
// }
