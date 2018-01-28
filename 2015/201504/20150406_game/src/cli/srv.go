package cli

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
		log.Printf("服务器连接失败: %v", err);
		return;
	}
	log.Printf("初始化连接成功[%s]\n", address);
}

func End(){
	Cnn.Close();
}