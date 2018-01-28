package pk

import (
	"golang.org/x/net/context"
	. "proto"
	log "golang.org/x/glog"
	"srv/db"
)

type PK struct{};

/*
匹配
*/
func (this *PK) Match (ctx context.Context, in *ReqMatch) (*RepMatch, error) {
	log.Info("recv Match: ", in.Id);
	data, err := db.Match(in.Id);
	if err != nil || len(data) <= 0{
		return &RepMatch{Sn: in.Sn, Msg: "匹配失败", Ret:"-1"}, err;
	}
	return &RepMatch{Sn: in.Sn, Msg: "匹配成功", Ret:"0", Data:data}, nil;
}

