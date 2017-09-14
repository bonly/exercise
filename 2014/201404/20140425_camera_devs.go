package main

import (
"fmt"
"glog"
"sync"
"encoding/xml"
"encoding/json"
"flag"
"io/ioutil"
"net/http"
"strings"
)

var Version string;
var Code string;
var App_time string;
var ver *bool = flag.Bool("V", false, "Version infomation");

var Wg sync.WaitGroup;

var srv *string = flag.String("srv", "http://device.sjkd189.com:80", "service address");
var usr *string = flag.String("user", "17701906586", "user name");
var pwd *string = flag.String("pwd", "xbedadmin2016", "password");
var key *string = flag.String("auth", "0290D5FC888563EC55761A4B4FB637AC", "auth key");

func main(){
	flag.Parse();

	if *ver{
		fmt.Println("Ver: ", Version);
		fmt.Println("Code: ", Code);
		fmt.Println("Time: ", App_time);
		return;
	}
	defer glog.Flush();

	flag.Set("alsologtostderr", "true");
	glog.Info("====== Camera check begin ======");
	defer func(){
		glog.Info("====== Camera check end ======");
		glog.Flush();
	}();


	Get_devices_list();
	Wg.Wait();
}

/* 
//请求
<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <GetDevicesByUid xmlns="http://tempuri.org/">
      <data>{"UserName":"17701906586","Password":"xbedadmin2016","PageIndex":1,"PageSize":10}</data>
      <authKey>0290D5FC888563EC55761A4B4FB637AC</authKey>
    </GetDevicesByUid>
  </soap:Body>
</soap:Envelope>

//结果
<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
    <soap:Body>
        <GetDevicesByUidResponse xmlns="http://tempuri.org/">
            <GetDevicesByUidResult>{"Code":0,"Message":"成功","Data":{"PageIndex":1,"PageCount":36,"CurrentPageSize":1,"Devices":[{"Name":"海业路xbed卡片机","Id":"3HK6E510110M7ST","ServerUrl":"ddns1.sdvideo.cn","UserName":"admin","Password":"12345","LinkedType":0,"AccessCode":""}]}}</GetDevicesByUidResult>
        </GetDevicesByUidResponse>
    </soap:Body>
</soap:Envelope>
*/

type Q_Get_devices struct{
	XMLName xml.Name `xml:"soap:Envelope"`;
	Xsi string `xml:"xmlns:xsi,attr"`;
	Xsd string `xml:"xmlns:xsd,attr"`;
	Soap string `xml:"xmlns:soap,attr"`;
	Xmlns string `xml:"xmlns,attr"`;
	Data string `xml:"soap:Body>GetDevicesByUid>data"`;
	Authkey string `xml:"soap:Body>GetDevicesByUid>authKey"`;
};

type Q_Data struct{
	UserName string;
	Password string;
	PageIndex int;
	PageSize int;
};

type R_Get_devices struct{
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`;
	// Xsi string `xml:"xmlns:xsi,attr"`;
	// Xsd string `xml:"xmlns:xsd,attr"`;
	// Soap string `xml:"xmlns:soap,attr"`;
	// Xmlns string `xml:"xmlns,attr"`;	
	Body struct{
		GetDevicesByUidResponse struct{
			GetDevicesByUidResult string `xml:"GetDevicesByUidResult"`;
		}`xml:"http://tempuri.org/ GetDevicesByUidResponse"`;
	}`xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`;
};

type R_Devices struct{
	Code int;
	Message string;
	Data struct{
		PageIndex int;
		PageCount int;
		CurrentPageSize int;
		Devices []struct{
			Name string;
			Id string;
			ServerUrl string;
			UserName string;
			Password string;
			LinkedType int;
			AccessCode string;
		}
	}
};

func Get_devices_list() (devices R_Devices, err error){
	dat := Q_Data{
		UserName : *usr,
		Password : *pwd,
		PageIndex : 1,
		PageSize : 2,
	};

	dat_str, err := json.Marshal(dat);
	if err != nil{
		glog.Info(err);
		return devices, err;
	}

	qry := Q_Get_devices{
		Xsi : "http://www.w3.org/2001/XMLSchema-instance",
		Xsd : "http://www.w3.org/2001/XMLSchema",
		Soap : "http://schemas.xmlsoap.org/soap/envelope/",
		Xmlns : "http://tempuri.org/",
		Data : string(dat_str), 
		Authkey : *key,
	};
	qry_str, err := xml.MarshalIndent(qry, " ", " ");
	if err != nil{
		glog.Info(err);
		return devices, err;
	}
	glog.Info("send: ", string(qry_str));

	body, err := Post(*srv+"/SmartHomeService.asmx", string(qry_str));
	if err != nil{
		glog.Info(err);
		return devices, err;
	}
	glog.Info(string(body));

	var rep R_Get_devices;
	err = xml.Unmarshal(body, &rep);
	if err != nil{
		glog.Info(err);
		return devices, err;
	}
	glog.Info(fmt.Sprintf("ret: %+v", rep.Body.GetDevicesByUidResponse.GetDevicesByUidResult));

	err = json.Unmarshal([]byte(rep.Body.GetDevicesByUidResponse.GetDevicesByUidResult), &devices);
	if err != nil{
		glog.Info(err);
		return devices, err;
	}
	glog.Info(fmt.Sprintf("devices: %+v", devices));
	return devices, err;
}

func Post(url string, data string)(rs []byte, err error){
	cli := &http.Client{};
	req, err := http.NewRequest("POST", url, strings.NewReader(data));
	if err != nil{
		glog.Info(err);
		return rs, err;
	}
	req.Header.Set("Content-Type","text/xml");
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");
	resp, err := cli.Do(req);
	if err != nil{
		glog.Info(err);
		return rs, err;
	}
	defer resp.Body.Close();

	rs, err = ioutil.ReadAll(resp.Body);	
	return rs, err;
}