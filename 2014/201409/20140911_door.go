package main 

/*
#include <stdint.h>

#pragma pack(push, 1)
typedef struct{
	uint8_t Head;  //包头
	uint8_t Len;   //包长 = 命令 + 包类型
	uint8_t Type;  //包类型
	uint8_t Cmd;   //命令
	uint8_t Param; //参数	
}Zigbee_head;

typedef struct{
	Zigbee_head head;
	uint8_t Data;  //数据
	uint8_t Tail;  //包尾
	uint8_t Verify;//检验
}Zigbee;

#pragma pack(pop)
*/
import "C"
import (
// "bytes"
// "encoding/binary"
"log"
"fmt"
"unsafe"
"flag"
"net"
)

type Zigbee_head struct{
	Head ,  //包头
	Len,   //包长 = 命令 + 包类型
	Type,  //包类型
	Cmd,   //命令
	Param uint8; //参数	
};

type Zigbee_tail struct{
	Tail,  //包尾
	Verify uint8;//检验
};

type Zigbee struct{
	Zigbee_head;
	Data uint8;  //数据
	Zigbee_tail;
};

type Box struct{
	Ctl uint32; //控制编号
	Lock uint8;   //门锁编号
};

var Srv_addr = flag.String("s", "0.0.0.0:5020", "服务器地址及端口");

func main(){
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

	for {
		//接收消息
		recvMsg(conn);

		//处理消息

		//应答消息
	}
}


func recvMsg(conn net.Conn){
	// buf := &bytes.Buffer{};
	// hl_len, err := conn.Read(buf.Bytes());
	buf := make([]byte, 255);
	hl_len, err := conn.Read(buf);
	if err != nil{
		fmt.Printf("read head failed %s\n", err.Error());
		return;
	}
	fmt.Printf("recv[%d]: %X\n", hl_len, buf[:hl_len]);

	if hl_len == 5{  //==5的是注册盒子
		var box *Box;
		box = (*Box)(unsafe.Pointer(&buf[0]));
		// err = binary.Read(bytes.NewReader(buf), binary.BigEndian, box);
		// if err != nil{
				// fmt.Printf("turn head failed %s\n", err.Error());
				// return;
		// }
		fmt.Printf("盒子注册: %v\n", *box);
	}else{ //大于5的是协议
		var zh *Zigbee_head;
		zh = (*Zigbee_head)(unsafe.Pointer(&buf[0]));
		fmt.Printf("协议数据包: %v\n", zh);
	}	
}