package main 

import(
"net/http"
// "encoding/json"
"crypto/md5"
"encoding/hex"
"strings"
"fmt"
"io/ioutil"
"reflect"
"sort"
"flag"
// "os"
)

var srv = flag.String("srv", "http://link.beta.quhuhu.com", "service addr");
var pms = flag.String("pms", "xbed", "pms id");
var token = flag.String("token", "cf2CQAjcSlpKeXeVwEBpvNQ773p110bp", "token");
var channelCode = flag.String("c", "QUNAR", "渠道ID");
var hotelNo = flag.String("h", "1", "酒店ID");
var userIp = flag.String("ip", "120.25.106.243", "接口请求IP");
var userAccount = flag.String("acc", "13693583818", "ebooking中的帐号");
var verifyCode = flag.String("verify", "493702", "手机验证码")
var cmd = flag.String("cmd", "41", `command api NO:
		21 查询渠道酒店
		41 渠道价格计划查询
		12 开通渠道直连
		11 发送手机动态码
`);

func init(){
	flag.Parse();
}

type Must struct{
	Hmac string `json:"hmac"`;
	Version string `sr:"version",json:"version"`;
	PmsId string `sr:"pmsId",json:"pmsId"`;
};

type Search_Channel_Rate_Plan struct{
	Must `sr:"must"`;
	ChannelCode string `sr:"channelCode",json:"channelCode"`;
	HotelNo string `sr:"hotelNo",json:"hotelNo"`;
};

func main(){
	switch(*cmd){
		case "41":
			search_channel_rate_plan();
			break;
		case "21":
			search_account_and_hotel_docking();
			break;
		case "12":
			docking_account();
			break;
		case "11":
			send_verification_code();
			break;
		default:
			fmt.Println("此接口未定义");
			break;
	}
}

func search_channel_rate_plan(){
	var dat Search_Channel_Rate_Plan;
	dat.PmsId = *pms;
	dat.Version = "1.0";
	dat.HotelNo = *hotelNo;
	dat.ChannelCode = *channelCode;

	qry, _ := Marsh(dat);
	send("/search/searchChannelRatePlan.do", qry);
}

type Search_Account_And_Hotel_Docking struct{
	Must `sr:"must"`;
	ChannelCode string `sr:"channelCode",json:"channelCode"`;
	HotelNo string `sr:"hotelNo",json:"hotelNo"`;
};

func search_account_and_hotel_docking(){
	var dat Search_Account_And_Hotel_Docking;
	dat.PmsId = *pms;
	dat.Version = "1.0";
	dat.HotelNo = *hotelNo;
	dat.ChannelCode = *channelCode;

	qry, _ := Marsh(dat);
	send("/search/searchAccountAndHotelDocking.do", qry);
}

type Docking_Account struct{
	Must `sr:"must"`;
	ChannelCode string `sr:"channelCode",json:"channelCode"`;
	HotelNo string `sr:"hotelNo",json:"hotelNo"`;
	UserAccount string `sr:"userAccount"`;
	VerificationCode string `sr:"verificationCode"`;
	UserIp string `sr:"userIp"`;
	OperatorGuid string `sr:"operatorGuid"`;
	OperatorName string `sr:"operatorName"`;
};

func docking_account(){
	var dat Docking_Account;
	dat.PmsId = *pms;
	dat.Version = "1.0";
	dat.HotelNo = *hotelNo;
	dat.ChannelCode = *channelCode;
	dat.UserAccount = *userAccount;
	dat.VerificationCode = *verifyCode;
	dat.UserIp = *userIp;
	dat.OperatorGuid = "0000";
	dat.OperatorName = "SYSTEM";

	qry, _ := Marsh(dat);
	send("/docking/dockingAccount.do", qry);
}


type Send_Verification_Code struct{
	Must `sr:"must"`;
	ChannelCode string `sr:"channelCode"`;
	Mobile string `sr:"mobile"`;
	UserIp string `sr:"userIp"`;
	OperatorGuid string `sr:"operatorGuid"`;
	OperatorName string `sr:"operatorName"`;
};

func send_verification_code(){
	var dat Send_Verification_Code;
	dat.PmsId = *pms;
	dat.Version = "1.0";
	dat.ChannelCode = *channelCode;
	dat.Mobile = *userAccount;
	dat.UserIp = *userIp;
	dat.OperatorGuid = "0000";
	dat.OperatorName = "SYSTEM";

	qry, _ := Marsh(dat);
	send("/docking/sendVerificationCode.do", qry);
}

func send(cmd string, qry string){
	fmt.Printf("send:\n %s\n", qry);
	req, err := http.NewRequest("POST", *srv + "/api/" + *pms + cmd,
		strings.NewReader(qry));

	req.Header.Set("Content-Type","application/x-www-form-urlencoded");
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");	

	if err != nil{
		fmt.Printf("new req: ", err);
		return;
	}

	cli := &http.Client{};
	resp, err := cli.Do(req);
	if err != nil{
		fmt.Printf("do: ", err);
		return;
	}
	defer resp.Body.Close();

	body, err := ioutil.ReadAll(resp.Body);
	if err != nil{
		fmt.Printf("body: ", err );
		return;
	}

	fmt.Printf("recv: \n %s\n", string(body));
}

func To_Str_Arry(in interface{}, tag string, mp *([]string))(err error){
	reflectType := reflect.TypeOf(in).Elem(); //取类型对象
	reflectValue := reflect.ValueOf(in).Elem();  //取值对象

	for idx := 0; idx < reflectType.NumField(); idx++{  //遍历所有字段
		typeName := reflectType.Field(idx).Name;  //取每个字段的名字

		valueType := reflectValue.Field(idx).Type(); //字段类的字符串式的描述
		valueValue := reflectValue.Field(idx).Interface();  //把字段转换为接口

		switch reflectValue.Field(idx).Kind(){ //字段类型判断
	 		case reflect.String:
	            // fmt.Printf("%s : %s(%s)\n", typeName, valueValue, valueType);
	            if tag := reflectType.Field(idx).Tag.Get(tag); tag != ""{ //只抽取用tag标识了的字段
	            	(*mp) = append(*mp, fmt.Sprintf("%s=%s", tag, valueValue));
	        	}
	            break;
	        case reflect.Int32:
	            fmt.Printf("%s : %i(%s)\n", typeName, valueValue, valueType);
	            break;
	        case reflect.Struct:
	            // fmt.Printf("%s : it is %s\n", typeName, valueType);
	            val := reflectValue.Field(idx).Addr(); //取字段的地址
	            To_Str_Arry(val.Interface(), tag, mp);  //深度遍历
	            break;
		}
	}
	return nil;
}


func Marsh(data interface{})(ret string, err error){
	var mp []string;

	// inf := reflect.ValueOf(data);
	// dat := inf.Convert(inf.Type());
	// if dat.Type() != reflect.TypeOf(Search_Account_And_Hotel_Docking{}) {
	// 	fmt.Println("not eq");
	// 	fmt.Printf("%#v\n", dat.Type());
	// 	os.Exit(1);
	// }
	// dat := data.(Search_Account_And_Hotel_Docking);
	
	// inf := reflect.ValueOf(data);
	// dat := inf.Convert(reflect.TypeOf(data));	
	// fmt.Printf("%#v\n", dat);

	//方法1：
	switch (*cmd){
	case "21":{
		dat := data.(Search_Account_And_Hotel_Docking);
		To_Str_Arry(&dat, "sr", &mp);
		break;
	}
	case "41":{
		dat := data.(Search_Channel_Rate_Plan);
		To_Str_Arry(&dat, "sr", &mp);
		break;		
	}
	case "12":{
		dat := data.(Docking_Account);
		To_Str_Arry(&dat, "sr", &mp);
		break;				
	}
	case "11":{
		dat := data.(Send_Verification_Code);
		To_Str_Arry(&dat, "sr", &mp);
		break;
	}
	default:
		return "", fmt.Errorf("无此命令");
	}
	
	//方法2:
	// switch data.(type){
	// 	case Search_Account_And_Hotel_Docking:
	// 		dat := data.(Search_Account_And_Hotel_Docking);
	// 		To_Str_Arry(&dat, "sr", &mp);
	// 		break;
	// 	default:
	// 		return "", fmt.Errorf("未定义");
	// }		
	// return "", nil;


	// fmt.Printf("org: %v\n", mp);
	
	sort.Strings(mp);
	// fmt.Printf("sort: %v\n", mp);

	pich_data := strings.Join(mp,"&");
	// fmt.Printf("join: %v\n", pich_data);

	all_data := string(pich_data[:12]) + *token + pich_data[12:];
	// fmt.Printf("token: %v\n", all_data);

	md := md5.New();
	md.Write([]byte(all_data));
	hmac := strings.ToUpper(hex.EncodeToString(md.Sum(nil)));
	// fmt.Printf("md5: %s\n", hmac);
	
	// dat.Hmac = hmac;
	// fmt.Printf("%+v\n", dat);	

	ret = fmt.Sprintf("%s&hmac=%s", pich_data, hmac);
	return ret, nil; 
}