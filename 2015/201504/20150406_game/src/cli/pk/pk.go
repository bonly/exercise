package scene

import "C"

import (
	"proto"
	"golang.org/x/net/context"
	"log"
	// "reflect"
	srv "cli"
	"encoding/json"
)

type Q_PK struct{
	Sn uint32;
	Id string;
};

func init(){
	log.Printf("PK Lib init\n");

	//注册类到清单中
	proto.Proto["Q_PK"] = func() interface{}{return &Q_PK{}};
	proto.Inter["PK"] = func() interface{}{return &Q_PK{}};
}

func (this *Q_PK) Get (data interface{})(ret string){
	kv := data.(*Q_PK);
	log.Printf("PK Match: %#v\n", kv.Id);

	go func(){
		Net := proto.NewPkClient(srv.Cnn);
		rep, err := Net.Match(context.Background(), &proto.ReqMatch{Sn: kv.Sn, Id: kv.Id});
		if err != nil{
			log.Printf("send failed: %v", err);
			proto.RepCallback (kv.Sn, proto.Mk_Ret(kv.Sn, "-1", "send failed"));
			return;
		}

		log.Printf("Recv: %#v\n", rep);
		js, _ := json.MarshalIndent(rep, " ", " ");
		proto.RepCallback(kv.Sn, string(js));
		return;
	}();
	return string(proto.Mk_Ret(kv.Sn, "0", "OK"));
}
