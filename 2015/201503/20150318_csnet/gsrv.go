package main 

import (
	"net"
	"google.golang.org/grpc"
	"log"
	"golang.org/x/net/context"
	"Proto"
)

const (
	port = ":50051";
)

type User struct{};

func (this *User) Login (ctx context.Context, in *Proto.ReqLogin) (*Proto.RepLogin, error) {
	log.Println("recv Login from ", in.Name);
	return &Proto.RepLogin{Msg: "OK", Ret:0}, nil;
}

func main(){
	lis, err := net.Listen("tcp", port);
	if err != nil{
		log.Fatalf("failed to listen: %v\n", err);
	}

	srv := grpc.NewServer();
	Proto.RegisterUserServer(srv, &User{});
	srv.Serve(lis);
}