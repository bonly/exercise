package main 

import (
	"net"
	"google.golang.org/grpc"
	log "golang.org/x/glog"
	"golang.org/x/net/context"
	"Proto"
	// "database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

const (
	port = ":50051";
)

type User struct{};

func (this *User) Login (ctx context.Context, in *Proto.ReqLogin) (*Proto.RepLogin, error) {
	log.Info("recv Login from ", in.Name);
	return &Proto.RepLogin{Msg: "OK", Ret:0}, nil;
}

var db *sqlx.DB;

func main(){
	var err error;
	db, err = sqlx.Open("mysql", "root:techappen@tcp(192.168.1.104:3306)/techappen?charset=utf8");
	if err != nil{
		log.Fatalf("db connect: %v\n", err);
	}

	lis, err := net.Listen("tcp", port);
	if err != nil{
		log.Fatalf("failed to listen: %v\n", err);
	}

	srv := grpc.NewServer();
	Proto.RegisterUserServer(srv, &User{});
	srv.Serve(lis);
}