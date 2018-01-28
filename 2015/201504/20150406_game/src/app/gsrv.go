package main 

import (
	"net"
	"google.golang.org/grpc"
	log "golang.org/x/glog"
	"proto"
	"flag"
	"srv/user"
	"srv/db"
	"srv/scene"
	"srv/cards"
	"srv/pk"
)

const (
	port = ":50051";
)

func init(){
  flag.Set("alsologtostderr", "true");
  flag.Set("v", "99");
  flag.Set("log_dir", "./");  
  flag.Parse();
}

func main(){
	db.Init();

	lis, err := net.Listen("tcp", port);
	if err != nil{
		log.Fatalf("failed to listen: %v\n", err);
	}

	srv := grpc.NewServer();
	proto.RegisterUserServer(srv, &user.User{});
	proto.RegisterSceneServer(srv, &scene.Scene{});
	proto.RegisterCardsServer(srv, &cards.Cards{});
	proto.RegisterPkServer(srv, &pk.PK{});
	srv.Serve(lis);
}