package main 

/*
#cgo LDFLAGS: -pthread -fPIC
struct Data{
	int cmd;
	char msg[1024];
};
*/
import "C"

import (
	"fmt"
	"he"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	// address = "localhost:50051";
	address = "192.168.1.109:50051";
)

var cnn *grpc.ClientConn;
var cli  he.GreeterClient;

func main(){}

//export Send
func Send(msg *C.char) C.int{
	fmt.Printf("cli: %v\n", cli);
	_, err := cli.SayHello(context.Background(), &he.HelloRequest{Name: "ok"});
	if err != nil{
		log.Fatalf("send failed: %v", err);
	}
	return 0;
}

func Recv(msg *C.char) C.int{
	return 0;
}

func Print(data *C.struct_Data) *C.char{
	return C.CString("oooo");
}

//export Srv
func Srv(){
	go srv(); 
}

func srv(){
	// for{
	    // fmt.Println("loop in go");
	// }
	var err error;
	cnn, err = grpc.Dial(address, grpc.WithInsecure());
	if err != nil{
		log.Fatalf("did not connect: %v", err);
	}
	// defer cnn.Close();

	cli = he.NewGreeterClient(cnn);
	fmt.Printf("create cli: %v\n", cli);
}