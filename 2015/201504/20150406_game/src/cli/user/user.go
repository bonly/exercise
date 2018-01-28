package user

import (
	"proto"
	"golang.org/x/net/context"
	"log"
	// "reflect"
	srv "cli"
	"encoding/json"
	"time"
)

type Q_User struct{
	Sn uint32 `json:"Sn"`;
	Name string `json:"Name"`;
	Passwd string `json:"Passwd"`;
};

func init(){
	log.Printf("User Lib init\n");

	//注册类到清单中
	proto.Proto["Q_User"] = func() interface{}{return &Q_User{}};
	proto.Inter["Q_User"] = func() interface{}{return &Q_User{}};
}

func (this *Q_User)Login(data interface{})(ret string){
	//@todo 数据验证
	user := data.(*Q_User);
	log.Printf("Login: %#v\n", user);

	go func(){
		time.Sleep(500* time.Millisecond);
		// proto.RepCallback( 13, "okaddddd");
		return;
	}();
	// go func(){
		Net := proto.NewUserClient(srv.Cnn);
		rep, err := Net.Login(context.Background(), &proto.ReqLogin{Sn:user.Sn , Name: user.Name, Passwd: user.Passwd});
		if err != nil{
			log.Printf("send failed: %v", err);
			proto.RepCallback(user.Sn, proto.Mk_Ret(user.Sn, "-1", "send failed"));
			return;
		}

		log.Printf("Recv: %#v\n", rep);
		js, _ := json.MarshalIndent(rep, " ", " ");
		log.Printf("回应答包: %s",js);
		proto.RepCallback(user.Sn, string(js));
		return;
	// }();

	return string(proto.Mk_Ret(user.Sn, "0", "OK"));
}

func (this *Q_User)Encode(pack []byte)(error){
	return json.Unmarshal(pack, this);
}

func (this *Q_User) Register (data interface{})(ret string){
	user := data.(*Q_User);
	log.Printf("Register: %#v\n", user);

	go func(){
		Net := proto.NewUserClient(srv.Cnn);
		rep, err := Net.Register(context.Background(), &proto.ReqRegister{Sn: user.Sn, Name: user.Name, Passwd: user.Passwd});
		if err != nil{
			log.Printf("send failed: %v", err);
			proto.RepCallback(user.Sn, proto.Mk_Ret(user.Sn, "-1", "send failed"));
			return;
		}

		log.Printf("Recv: %#v\n", rep);
		js, _ := json.MarshalIndent(rep, " ", " ");
		proto.RepCallback(user.Sn, string(js));
		return;
	}();
	return string(proto.Mk_Ret(user.Sn, "0", "OK"));
}
