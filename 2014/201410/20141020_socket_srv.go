package main

import (
"net"
"log"
)

func Srv(){
	lis, err := net.Listen("tcp", "0.0.0.0:8989");
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
	ngr_conn(conn.RemoteAddr().String());
}

func main(){
	Srv();
}

func ngr_conn(addr string){
	conn, err := net.Dial("tcp", addr);
	if err != nil{
		log.Printf("%s\n", err.Error());
		return;
	}
	log.Printf("ngr connect ok");
	conn.Write([]byte("ngr connect u ok"));
}