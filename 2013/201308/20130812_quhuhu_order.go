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
	"encoding/json"
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

const srv_addr string = "http://provider.beta9.qunar.com/api/";
const key string = "bc2f469e-e1fa-4c91-af3f-b502202d5ffd";

func hmac_check_lock_network_state(dt *NewOrder, vl *url.Values){
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
	action := "order/newOrder.do";
	var cn NewOrder;
	cn.version = "1.0";
	cn.SignKey = key;
	cn.hotelId = "1";

    var rm = D_RoomInfo{
    RoomId: "1",
    ChannelOrderNo: "1",
    CheckInTime: "20150723181650",
    CheckOutTime: "20150724181650",
    CustomerName: "test",
    CustomerMobile: "13336677817",
    ContactName: "test",
    ContactMobile: "13336677817",
    RoomPrice: "220",
    };

    var st StayInfo;
    st.RoomInfo = append(st.RoomInfo, rm);
    //*
    js, err := json.Marshal(st);
    if err != nil{
    	log.Println("json: ", err);
    	return;
    }
    lg := len(string(js));
    cn.stayInfo = string(js[12:lg-1]);
    //*/

    // test_str:=`[{"roomId": "1","channelOrderNo": "1","checkInTime": "20150723181650","checkOutTime": "20150724181650","customerName": "测试","customerMobile": "13336677817","contactName": "测试","contactMobile": "13336677817","roomPrice": 220}]`;
    // cn.stayInfo = test_str;

	val := url.Values{};
	val.Add("SignKey", cn.SignKey);
	val.Add("version", cn.version);
	val.Add("hotelId", cn.hotelId);
	val.Add("stayInfo", cn.stayInfo);
	
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
