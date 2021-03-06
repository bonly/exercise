package Srv

import (
	"log"
	"google.golang.org/grpc"
)

var Cnn *grpc.ClientConn;


const (
	address = "192.168.1.109:50051";
)

func Start(){
	var err error;
	Cnn, err = grpc.Dial(address, grpc.WithInsecure());
	if err != nil{
		log.Fatalf("服务器连接失败: %v", err);
	}
}

func End(){
	Cnn.Close();
}