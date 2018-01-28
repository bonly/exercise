package scene

import (
	"golang.org/x/net/context"
	. "proto"
	log "golang.org/x/glog"
	"srv/db"
)

type Scene struct{};

/*
取场景布局数据
*/
func (this *Scene) Get (ctx context.Context, in *ReqKeyValue) (*RepKeyValue, error) {
	log.Info("recv Scene Get: ", in.Key);
	data, err := db.Scene_Get(in.Key);
	if err != nil{
		return &RepKeyValue{Ret:"-1", Msg:"数据错误", Data: ""}, err;
	}

	return &RepKeyValue{Ret: "0", Msg:"成功", Data: data}, nil;
}

/*
保存场景数据
*/
func (this *Scene) Set (ctx context.Context, in *ReqKeyValue) (*Rep, error) {
	log.Info("recv Scene Set: ", in.Key);
	err := db.Scene_Save(in.Key, in.Data);
	if err != nil{
		return &Rep{Ret:"-1", Msg:"数据错误"}, err;
	}

	return &Rep{Ret: "0", Msg:"成功"}, nil;
}