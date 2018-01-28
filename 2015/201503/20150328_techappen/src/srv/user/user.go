package user

import (
	"golang.org/x/net/context"
	. "proto"
	log "golang.org/x/glog"
	"srv/db"
)

type User struct{};

func (this *User) Login (ctx context.Context, in *ReqLogin) (*RepLogin, error) {
	log.Info("recv Login: ", in.Name);
	user_id, err := db.User_Get(in.Name, in.Passwd);
	if err != nil || len(user_id) <= 0{
		return &RepLogin{Msg: "登录失败", Ret:-1}, err;
	}
	return &RepLogin{Msg: "登录成功", Ret:0}, nil;
}

func (this *User) Register (ctx context.Context, in *ReqRegister) (ret *RepRegister, err error){
	log.Info("recv register ");
	user_id, err := db.User_Add(in.Name, in.Passwd);
	if err != nil{
		log.Error("Register failed");
		ret = &RepRegister{
			Ret : -1,
		    Msg :"注册失败",
		};
		return 	ret, err;
	}
	ret = &RepRegister{
		Ret : 0,
		Msg : "注册成功",
		Id : user_id,
	};
	return ret, nil;
}