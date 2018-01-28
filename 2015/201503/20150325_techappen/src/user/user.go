package user

import "C"

import (
	"proto"
	"golang.org/x/net/context"
	"log"
	// "reflect"
	"srv"
	"encoding/json"
)


type Q_User struct{
	Name string `json:"Name"`;
	Passwd string `json:"Passwd"`;
};

func init(){
	log.Printf("User Lib init\n");

	//注册类到清单中
	proto.Proto["Q_User"] = func() interface{}{return &Q_User{}};
}

func (this *Q_User)Login(data interface{}){
	user := data.(*Q_User);
	log.Printf("Login: %#v\n", user);

	Net := proto.NewUserClient(srv.Cnn);
	rep, err := Net.Login(context.Background(), &proto.ReqLogin{Name: user.Name, Passwd: user.Passwd});
	if err != nil{
		log.Fatalf("send failed: %v", err);
	}

	log.Printf("Recv: %#v\n", rep);
	return;
}

func (this *Q_User)Encode(pack []byte)(error){
	return json.Unmarshal(pack, this);
}

