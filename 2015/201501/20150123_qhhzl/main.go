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

func (this *Search_Account_And_Hotel_Docking) Cls() reflect.Type{
	fmt.Println("call in class");
	return reflect.TypeOf(this);
}

func search_account_and_hotel_docking(){
	var dat Search_Account_And_Hotel_Docking;
	dat.PmsId = *pms;
	dat.Version = "1.0";
	dat.HotelNo = *hotelNo;
	dat.ChannelCode = *channelCode;

	qry, err := Marsh(dat);
	if err != nil{
		return;
	}
	send("/search/searchAccountAndHotelDocking.do", qry);
}

type Docking_Account struct{
	Must;
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
	typ := reflect.TypeOf(in);
	val := reflect.ValueOf(in);

	if typ.Kind() == reflect.Ptr || typ.Kind() == reflect.Interface{
		typ = typ.Elem();  //如果是指针，解指针
		val = val.Elem();
		// val = reflect.Indirect(val);

		obj := val.Interface();  //还原对象至新对象，根据新对象重新反射
		val = reflect.ValueOf(obj);
		typ = reflect.TypeOf(obj);
		// fmt.Printf("入参是引用，解引用得 %s 和 %s\n", typ.Kind(), val.Type());
	}

	// fmt.Printf("得 %s 和 %s\n", typ.Kind(), val.Type());
	if typ.Kind() != reflect.Struct{
		fmt.Println("类型不正确");
		return fmt.Errorf("%v 类型不正确\n", typ.Kind()); //确保是结构
	}

	// fmt.Printf("numField: %d\n", typ.NumField());
	for idx := 0; idx < typ.NumField(); idx++{
		typ_fld := typ.Field(idx);
		val_fld := val.Field(idx).Interface();
		
		if typ_fld.Anonymous{
			// fmt.Println("子结构", typ_fld.Name);
			To_Str_Arry(&val_fld, tag, mp);
			continue;
		}

		if tag := typ.Field(idx).Tag.Get(tag); tag != ""{
			// fmt.Println("找到字段", typ.Field(idx).Name);
			(*mp) = append(*mp, fmt.Sprintf("%s=%s", tag, val_fld)); //val.Field(idx).Interface()));
		}
	}
	return nil;
}


func Marsh(data interface{})(ret string, err error){
	var mp []string;

	dat := reflect.ValueOf(data).Convert(reflect.TypeOf(data));	//转成Value是指定类型的反射
	dt  := dat.Interface(); //value转成真正的类型实例

	// reflect_struct_info(&data);
	err = To_Str_Arry(&dt, "sr", &mp);
	if err != nil{
		return "", err;
	}
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

	ret = fmt.Sprintf("%s&hmac=%s", pich_data, hmac);
	return ret, nil; 
}

//暂无作用的方法
func reflect_struct_info(it interface{}) {
    t := reflect.TypeOf(it)
    fmt.Printf("interface info:%s %s %s %s\n", t.Kind(), t.PkgPath(), t.Name(), t)
    if t.Kind() == reflect.Ptr { //if it is pointer, get it element type
        tt := t.Elem()
        if t.Kind() == reflect.Interface {
            fmt.Println(t.PkgPath(), t.Name())
            for i := 0; i < tt.NumMethod(); i++ {
                f := tt.Method(i)
                fmt.Println(i, f)
            }
        }
    }
    v := reflect.ValueOf(it)
    k := t.Kind()
    if k == reflect.Ptr {
        v = v.Elem() //指针转换为对应的结构
        t = v.Type()
        k = t.Kind()
    }
    fmt.Printf("value type info:%s %s %s\n", t.Kind(), t.PkgPath(), t.Name())
    if k == reflect.Struct { //反射结构体成员信息
        for i := 0; i < t.NumField(); i++ {
            f := t.Field(i)
            fmt.Printf("%s %v\n", i, f)
        }
        for i := 0; i < t.NumMethod(); i++ {
            f := t.Method(i)
            fmt.Println(i, f)
        }
        fmt.Printf("Fileds:\n")
        f := v.MethodByName("func_name")
        if f.IsValid() { //执行某个成员函数
            arg := []reflect.Value{reflect.ValueOf(int(2))}
            f.Call(arg)
        }
        for i := 0; i < v.NumField(); i++ {
            f := v.Field(i)
            if !f.CanInterface() {
                fmt.Printf("%d:[%s] %v\n", i, t.Field(i), f.Type())
                continue
            }
            val := f.Interface()
            fmt.Printf("%d:[%s] %v %v\n", i, t.Field(i), f.Type(), val)
        }
        fmt.Printf("Methods:\n")
        for i := 0; i < v.NumMethod(); i++ {
            m := v.Method(i)
            fmt.Printf("%d:[%v] %v\n", i, t.Method(i), m)
        }
    }
}