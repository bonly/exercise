package cards

import "C"

import (
	"proto"
	"golang.org/x/net/context"
	"log"
	// "reflect"
	srv "cli"
	"encoding/json"
)

type Q_Card struct{
	Sn uint32;
	Key string `json:"Key"`;
	Data string `json:"Data"`;
};

func init(){
	log.Printf("Cards Lib init\n");

	//注册类到清单中
	proto.Proto["Q_Card"] = func() interface{}{return &Q_Card{}};
	proto.Inter["Cards"] = func() interface{}{return &Q_Card{}};
}

func (this *Q_Card) Get (data interface{})(ret string){
	kv := data.(*Q_Card);
	log.Printf("Cards Get: %#v\n", kv.Key);

	go func(){
		Net := proto.NewCardsClient(srv.Cnn);
		rep, err := Net.Get(context.Background(), &proto.ReqKeyValue{Sn: kv.Sn, Key: kv.Key});
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

func (this *Q_Card) Set (data interface{})(ret string){
	kv := data.(*Q_Card);
	log.Printf("Cards Set: %#v\n", kv.Key);

	go func(){
		Net := proto.NewCardsClient(srv.Cnn);
		rep, err := Net.Set(context.Background(), &proto.ReqKeyValue{Key: kv.Key, Data: kv.Data});
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