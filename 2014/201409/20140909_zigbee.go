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
"encoding/binary"
"log"
"fmt"
// "unsafe"
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

type Zigbee struct{
	Head Zigbee_head;
	Data,  //数据
	Tail,  //包尾
	Verify uint8;//检验
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
		// head := make([]byte, unsafe.Sizeof(Zigbee_head));
		var zh Zigbee_head;
		head := make([]byte, binary.Size(zh));
		hl_len, err := conn.Read(head);
		if err != nil{
			fmt.Printf("head and len read err %v\n", err);
			continue;
		}
		fmt.Printf("recv %d: %X\n", hl_len, head);

		wr, err := conn.Write(head);
		if err != nil{
			fmt.Printf("write back %v\n", err);
			continue;
		}
		fmt.Printf("send %d: %v\n", wr, head);

		hl_len, err = conn.Read(head);
		if err != nil{
			fmt.Printf("head and len read err %v\n", err);
			continue;
		}
		fmt.Printf("recv %d: %v\n", hl_len, head);
	}
}
