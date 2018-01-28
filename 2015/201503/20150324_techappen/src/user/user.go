package user

import "C"

import (
	"proto"
	// "golang.org/x/net/context"
	"log"
	// "reflect"
	// "srv"
	// "encoding/json"
)


type Q_User struct{
	Name string `json:"Name"`;
	Passwd string `json:"Passwd"`;
	Func string `json:"Func"`;
};

func init(){
	log.Printf("User Lib init\n");

	//注册类到清单中
	proto.Proto["Q_User"] = func() interface{}{return &Q_User{}};
	
	// Fnc = reflect.ValueOf(&inf);
}

func Start(){
	// if Net == nil{
	// 	Net = proto.NewUserClient(srv.Cnn);
	// 	log.Printf("设置服务器连接成功\n");	
	// }
}

// func (this *Q_User)Decode(data []byte) error{
// 	return json.Unmarshal(data, this);
// }


// func (this *Inf)Login(data interface{}){
// 	user := data.(*Q_User);
// 	log.Printf("user: %v\n", user);
// 	_, err := Net.Login(context.Background(), &proto.ReqLogin{Name: "ok"});
// 	if err != nil{
// 		log.Fatalf("send failed: %v", err);
// 	}
// 	return;
// }

// func (this *Q_User)Decode(pack []byte)(error){
// 	return json.Unmarshal(pack, this);
// }

