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
	he "Proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"reflect"
)

const (
	// address = "localhost:50051";
	address = "192.168.1.109:50051";
)


var cnn *grpc.ClientConn;
var cli  he.UserClient;
var fn   reflect.Value;

type Inf struct{
}

func init(){
	log.Printf("so being init\n");
	inf := &Inf{};
	fn = reflect.ValueOf(inf);
}

func main(){}

//export Send
func Send(msg *C.char) C.int{
	fmt.Printf("cli: %v\n", cli);
	_, err := cli.Login(context.Background(), &he.ReqLogin{Name: "ok"});
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

func srv (){
	// for{
	    // fmt.Println("loop in go");
	// }
	var err error;
	cnn, err = grpc.Dial(address, grpc.WithInsecure());
	if err != nil{
		log.Fatalf("did not connect: %v", err);
	}
	// defer cnn.Close();

	cli = he.NewUserClient(cnn);
	fmt.Printf("create cli: %v\n", cli);
}

func (this *Inf)Login(data string){
	log.Printf("login: %s", data);
	_, err := cli.Login(context.Background(), &he.ReqLogin{Name: "ok"});
	if err != nil{
		log.Fatalf("send failed: %v", err);
	}
	return;
}

//export Proc
func Proc(ccmd *C.char, cdata *C.char){
	cmd := C.GoString(ccmd);
	data := C.GoString(cdata);

	// log.Printf("call %v\n", C.GoString(cmd));

	method := fn.MethodByName(cmd);
	if method.IsValid(){
		vdata := reflect.ValueOf(data);
		method.Call([]reflect.Value{0:vdata});
		//reflect.Zero(reflect.TypeOf(&Thing{}))}
	}else{
		log.Printf("无此方法: %v\n", cmd);
	}
}