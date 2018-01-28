package main 

import (
	"net"
	"google.golang.org/grpc"
	"log"
	"golang.org/x/net/context"
	pb "He"
)

const (
	port = ":50051";
)

type server struct{};

func (this *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("recv sayhello");
	return &pb.HelloReply{Message: "Hello" + in.Name}, nil;
}

func main(){
	lis, err := net.Listen("tcp", port);
	if err != nil{
		log.Fatalf("failed to listen: %v\n", err);
	}

	srv := grpc.NewServer();
	pb.RegisterGreeterServer(srv, &server{});
	srv.Serve(lis);
}