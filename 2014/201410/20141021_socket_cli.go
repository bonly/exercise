package main

import (
"net"
"log"
)

func Cli(){
	conn, err := net.Dial("tcp", "0.0.0.0:8989");
	if err != nil{
		log.Printf("%s\n", err.Error());
		return;
	}
	log.Printf("connect ok");
	conn.Write([]byte("org client send to srv ok"));

	tcpConn, ok := conn.(*net.TCPConn);
	if !ok {
		log.Printf("转换为tcp connect失败");
	}	
	tcpConn.SetNoDelay(true);
	//默认已经是SO_REUSEADDR

	// ngr_lis(conn.LocalAddr().String()); //反向侦听，失败，已占用
	// for {
	// 	handle(conn); //反向接收，只能是数据，不能接受
	// }
}

func main(){
	Cli();
}

func ngr_lis(addr string){
	lis, err := net.Listen("tcp", addr);
	if err != nil{
		log.Printf("%s\n", err.Error());
		return;
	}
	for {
		conn, err := lis.Accept();
		if err != nil{
			log.Printf("%s\n", err.Error());
			continue;
		}
		go handle(conn);
	}	
}

func handle(conn net.Conn){
	var buf [255]byte;
	var err error;
	lng := 0;
	if lng, err = conn.Read(buf[:]); err != nil{
		log.Printf("%s\n", err.Error());
		conn.Close();
		return;
	}
	log.Printf("get a connect: %v\n", conn.RemoteAddr());
	log.Printf("recv [%d]: %s\n", lng, string(buf[:]));
}