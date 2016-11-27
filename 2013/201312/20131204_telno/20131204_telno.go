/*
http://tcc.taobao.com/cc/json/mobile_tel_segment.htm?tel=15850781443
http://virtual.paipai.com/extinfo/GetMobileProductInfo?mobile=15850781443&amount=10000&callname=getPhoneNumInfoExtCallback
http://life.tenpay.com/cgi-bin/mobile/MobileQueryAttribution.cgi?chgmobile=15850781443
*/


package main 

import (
"log"
"net/http"
"io/ioutil"
"fmt"
"encoding/json"
"strings"
// "iconv"
"encoding/csv"
"os"
)

type GetZoneResult struct{
	Mts string `json:"mts"`;
	Province string `json:"province"`;
	CatName string `json:"catName"`;
	TelString string `json:"telString"`;
	AreaVid string `json:"areaVid"`;
	IspVid string `json:"ispVid"`;
	Carrier string `json:"carrier"`;
};

func main(){
	zone, err := Get_zone("17701906025");
	if err != nil{
		fmt.Println(zone);
		return;
	}

	fs, err := os.Create("tel.csv");
	if err != nil{
		fmt.Println(err);
		return;
	}
	defer fs.Close();

	// fs.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	ws := csv.NewWriter(fs);
	// ws.Write([]string{"电话", "归属地", "运营商"});
	ws.Write([]string{"17701906025", zone.Province, zone.Carrier});
	ws.Flush();
}


func Get_zone(key string) (zone GetZoneResult, err error){
    resp, err := http.Get("http://tcc.taobao.com/cc/json/mobile_tel_segment.htm?tel="+key);
    if err != nil {
        log.Println("http get: ",err);
        return zone, err;
    }
    defer resp.Body.Close();

    // cd, err := iconv.Open("gbk", "utf-8");
    // if err != nil{
    // 	log.Println("iconv open fail: ", err.Error());
    // 	return zone, err;
    // }	
    // defer cd.Close();

	body, err := ioutil.ReadAll(resp.Body);
    if err != nil {
        log.Println(fmt.Printf("body: %s", err.Error()));
        return zone, err;
    }
 
 	// bt := make([]byte, 255);
  //   bt, _, err = cd.Conv(body, bt);
  //   if err != nil{
  //   	fmt.Println(err);
  //   }
 	ret := string(body);
    // fmt.Println(gret);

	idx := strings.Index(ret, "{");
	ret = ret[idx:];

	ret = strings.Replace(ret, "mts:", "\"mts\":", 1);
	ret = strings.Replace(ret, "province:", "\"province\":",1);
	ret = strings.Replace(ret, "catName:", "\"catName\":", 1);
	ret = strings.Replace(ret, "telString:", "\"telString\":", 1);
	ret = strings.Replace(ret, "areaVid:", "\"areaVid\":", 1);
	ret = strings.Replace(ret, "ispVid:", "\"ispVid\":", 1);
	ret = strings.Replace(ret, "carrier:", "\"carrier\":", 1);
	ret = strings.Replace(ret, "'", "\"", -1);
	fmt.Println(ret);

    if err = json.Unmarshal([]byte(ret), &zone); err != nil{
    	log.Println("解包错误: ", err);
    	return zone, err
    }

    // zone.Province = cd.ConvString(zone.Province);  
    // zone.CatName = cd.ConvString(zone.CatName);  
    // zone.Carrier = cd.ConvString(zone.Carrier);  
	return zone, nil;	
}