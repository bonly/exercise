package cards

import (
	"golang.org/x/net/context"
	. "proto"
	log "golang.org/x/glog"
	"srv/db"
)

type Cards struct{};

/*
取卡牌数据
*/
func (this *Cards) Get (ctx context.Context, in *ReqKeyValue) (*RepKeyValue, error) {
	log.Info("recv Cards Get: ", in.Key);
	data, err := db.Cards_Get(in.Key);
	if err != nil{
		return &RepKeyValue{Ret:"-1", Msg:"数据错误", Data: ""}, err;
	}

	return &RepKeyValue{Ret: "0", Msg:"成功", Data: data}, nil;
}

/*
保存卡牌数据
*/
func (this *Cards) Set (ctx context.Context, in *ReqKeyValue) (*Rep, error) {
	log.Info("recv Cards Set: ", in.Key);
	err := db.Cards_Save(in.Key, in.Data);
	if err != nil{
		return &Rep{Ret:"-1", Msg:"数据错误"}, err;
	}

	return &Rep{Ret: "0", Msg:"成功"}, nil;
}