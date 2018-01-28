package User

import "C"

import (
	"Proto"
	"golang.org/x/net/context"
	"log"
	"reflect"
	"Srv"
	"encoding/json"
)

type Inf struct{};

var inf Inf;
var Fnc reflect.Value;
var Net Proto.UserClient;

func init(){
	log.Printf("User Lib init\n");
	
	Fnc = reflect.ValueOf(&inf);
}

func Start(){
	if Net == nil{
		Net = Proto.NewUserClient(Srv.Cnn);
		log.Printf("服务器连接成功\n");	
	}
}


type Q_User struct{
	Name string;
	Passwd string;
};

func (this *Inf)Login(data interface{}){
	user := data.(*Q_User);
	log.Printf("user: %v\n", user);
	_, err := Net.Login(context.Background(), &Proto.ReqLogin{Name: "ok"});
	if err != nil{
		log.Fatalf("send failed: %v", err);
	}
	return;
}

func (this *Q_User)Decode(pack []byte)(error){
	return json.Unmarshal(pack, this);
}

