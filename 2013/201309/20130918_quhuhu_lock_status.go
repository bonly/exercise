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
	"fmt"
	"io/ioutil"
	"sort"
)

type Quhuhu struct{
	SignKey string;
	version string;
	hmac string;
};

type CheckLockNetworkState struct{
	Quhuhu;	
	hotelId string;
	roomId string;
};

const srv_addr string = "http://sp.qunar.com/api/";
const key string = "50cc695f-4602-4f23-b356-17bc61ed1806";

func hmac_check_lock_network_state(dt *CheckLockNetworkState, vl *url.Values){
	/* //不是对key排序，是对值排序
    sort := fmt.Sprintf("hotelId=%s,roomId=%s,SignKey=%s,version=%s", 
					dt.hotelId,dt.roomId,dt.SignKey,dt.version);
	
	md := md5.New();
	md.Write([]byte(sort));
	(*dt).hmac = strings.ToUpper(hex.EncodeToString(md.Sum(nil)));
	
	(*vl).Add("hmac", (*dt).hmac);
	
	sort = fmt.Sprintf("hmac=%s,hotelId=%s,roomId=%s,SignKey=%s,version=%s", 
					dt.hmac,dt.hotelId,dt.roomId,dt.SignKey,dt.version);
	return sort;
	*/
	
	var sv []string;
	for _,v := range *vl {
		sv = append(sv, v[0]);
	}
	sort.Strings(sv);
	
	all_data := strings.Join(sv,"");
	// log.Println(all_data);
	md := md5.New();
	md.Write([]byte(all_data));
	(*dt).hmac = strings.ToUpper(hex.EncodeToString(md.Sum(nil)));
	(*vl).Add("hmac", (*dt).hmac);
}

func check_lock_network_state (){
	action := "lock/checkLockNetworkState.do";
	var cn CheckLockNetworkState;
	cn.version = "1.0";
	cn.SignKey = key;
	// cn.SignKey =strings.ToUpper(key);
	cn.hotelId = "1";
	cn.roomId = "18A-3316";
	
	val := url.Values{};
	val.Add("SignKey", cn.SignKey);
	val.Add("version", cn.version);
	val.Add("hotelId", cn.hotelId);
	val.Add("roomId", cn.roomId);
	
	hmac_check_lock_network_state(&cn, &val);
	
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
