package proto

import (
"testing"
"log"
"fmt"
"unsafe"
"flag"
"net"
"sync"
)

func Test_pack_recv(ts *testing.T){
	var Srv_addr = flag.String("s", "0.0.0.0:5020", "服务器地址及端口");

	flag.Parse();

	srv, err := net.ResolveTCPAddr("tcp4", *Srv_addr);
	if err != nil{
		log.Printf("srv start failed %v\n", err);
		return;
	}
	listener, err := net.ListenTCP("tcp", srv);
	if err != nil{
		log.Printf("srv listen failed %v\n", err);
		return;
	}
	for {
		conn, err := listener.Accept();
		if err != nil{
			continue;
		}
		go handleDoor(conn);
	}
}

func handleDoor(conn net.Conn){
	defer conn.Close();
	var wg sync.WaitGroup;
	wg.Add(2);

	//接收消息
	reg, task := recvMsg(conn, &wg);

	//处理消息
	proc_task(conn, task, &wg);

	proc_reg(conn, reg, &wg);

	//应答消息

	fmt.Printf("设置任务处理完毕\n");
	wg.Wait();
	fmt.Printf("所有任务完成\n");	
}

func recvMsg(conn net.Conn, wg *sync.WaitGroup)(reg chan interface{}, task chan interface{}){
	task = make(chan interface{});
	reg  = make(chan interface{});
	go func(){
		defer wg.Done();
		for {
			buf := make([]byte, 255);
			hl_len, err := conn.Read(buf);
			if err != nil{
				fmt.Printf("read head failed %s\n", err.Error());
				close(task);
				// continue;
			}
			fmt.Printf("recv[%d]: %X\n", hl_len, buf[:hl_len]);

			if hl_len == 5{  //==5的是注册盒子
				var box *Register;
				box = (*Register)(unsafe.Pointer(&buf[0]));
				fmt.Printf("盒子注册: %v\n", *box);
				reg <- box;
			}else{ //大于5的是协议
				var zh Command;
				cmd := zh.Decode(buf);
				fmt.Printf("协议数据包: %v\n", cmd);
				task <-zh;
			}
		}
	}();

	return reg, task;
}

func proc_task(conn net.Conn, task chan interface{}, wg *sync.WaitGroup){
	fmt.Printf("任务处理器就绪\n");
	go func(){
		defer wg.Done();
		for {
			select{
				case work := <- task:{
					fmt.Printf("%+v\n", work);	
				}
			}
		}
	}();
}

func proc_reg(conn net.Conn, reg chan interface{}, wg *sync.WaitGroup){
	fmt.Printf("注册处理器就绪\n");
	go func(){
		defer wg.Done();
		for {
			select{
				case work := <- reg:{
					fmt.Printf("%+v\n", work);	
					var rop Remote_open_door;
					rop.New();
					rop.Passwd = []byte{0x06, 0x01, 0x08, 0x05, 0x03, 0x09};
					pack, _ := rop.Encode();
					cnt, err := conn.Write(pack);
					if err != nil{
						fmt.Printf("send %v\n", err);
					}else{
						fmt.Printf("send len[%d]: %v\n", cnt, pack);
					}				
					// conn.Write(pack);
					// conn.Write(pack);
					// conn.Write(pack);
					// conn.Write(pack);	
				}
			}
		}
	}();
}