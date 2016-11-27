/*
去呼呼测试
*/
package main

import(
	"net/http"
	"strings"
	"log"
	"net/url"
	"crypto/md5"
	"encoding/hex"
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
)

type Quhuhu struct{
	SignKey string;
	version string;
	hmac string;
};

type D_RoomInfo struct{
	RoomId string `json:"roomId"`;
	ChannelOrderNo string `json:"channelOrderNo"`;
	CheckInTime string `json:"checkInTime"`;
	CheckOutTime string `json:"checkOutTime"`;
	CustomerName string `json:"customerName"`;
	CustomerMobile string `json:"customerMobile"`;
	ContactName string `json:"contactName"`;
	ContactMobile string `json:"contactMobile"`;
	RoomPrice string `json:"roomPrice"`;
};

type StayInfo struct{
    RoomInfo []D_RoomInfo `json:"stayInfo"`;
};

type NewOrder struct{
	Quhuhu;	
	hotelId string;
	stayInfo string;
};

type QryOrder struct{
	fromTime,
	toTime,
	hotelId,
	typeId string;
	Quhuhu;
};

const srv_addr string = "http://sp.qunar.com/api/";
const key string = "50cc695f-4602-4f23-b356-17bc61ed1806";

func hmac_check_lock_network_state(dt *QryOrder, vl *url.Values){
	//不是对key排序，是对值排序
	var sv []string;
	for _,v := range *vl {
		sv = append(sv, v[0]);
	}
	sort.Strings(sv);
	
	all_data := strings.Join(sv,"");
	log.Println("md5 data: ",all_data);
	md := md5.New();
	md.Write([]byte(all_data));
	(*dt).hmac = strings.ToUpper(hex.EncodeToString(md.Sum(nil)));
	(*vl).Add("hmac", (*dt).hmac);
}

func check_lock_network_state (){
	action := "order/queryOrder.do";
	var cn QryOrder;
	cn.version = "1.0";
	cn.SignKey = key;

	cn.hotelId = "1";
	cn.fromTime = "20150819";
	cn.toTime = "20150820";
	cn.typeId = "1";

	val := url.Values{};
	val.Add("SignKey", cn.SignKey);
	val.Add("version", cn.version);
	val.Add("hotelId", cn.hotelId);
	val.Add("fromTime", cn.fromTime);
	val.Add("toTime", cn.toTime);
	val.Add("type", cn.typeId);

	hmac_check_lock_network_state(&cn, &val);
	
	// val.Del("stayInfo");
	str := val.Encode();

	addr := srv_addr + action;
	
	cli := &http.Client{};
	req, err := http.NewRequest("POST",
		addr, strings.NewReader(str));
	req.Header.Set("Content-Type","application/x-www-form-urlencoded");
	resp, err := cli.Do(req);
	if err != nil{
		log.Println("send: ", err);
		return;
	}
	defer resp.Body.Close();
	
	log.Println("send to: ", addr);
	log.Println("post dt: ", str);
	
	body, err := ioutil.ReadAll(resp.Body);
	if err != nil{
		fmt.Println("recv: ", err);
		return;
	}
	fmt.Println("服务器应答：", string(body));
}

func main(){
    check_lock_network_state();
	
}
