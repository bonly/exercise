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
)

var srv = flag.String("srv", "http://link.beta.quhuhu.com", "service addr");
var pms = flag.String("pms", "xbed", "pms id");
var token = flag.String("token", "cf2CQAjcSlpKeXeVwEBpvNQ773p110bp", "token");
var channelCode = flag.String("c", "QUNAR", "channel code");
var hotelNo = flag.String("h", "1", "hotel NO");

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
	// verify();
	send();
}

func verify(){
	var dat Search_Channel_Rate_Plan;
	dat.PmsId = *pms;
	dat.Version = "1.0";
	dat.HotelNo = "1";
	dat.ChannelCode = "quhuhu";

	str, _ := Marsh(dat);
	fmt.Printf("ret: %s\n", str);
}

func send(){
	var dat Search_Channel_Rate_Plan;
	dat.PmsId = *pms;
	dat.Version = "1.0";
	dat.HotelNo = *hotelNo;
	dat.ChannelCode = *channelCode;

	qry, _ := Marsh(dat);

	req, err := http.NewRequest("POST", *srv + "/api/" + *pms + "/search/searchChannelRatePlan.do",
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
	dat := data.(Search_Channel_Rate_Plan);
	
	To_Str_Arry(&dat, "sr", &mp);
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
	
	dat.Hmac = hmac;
	// fmt.Printf("%+v\n", dat);	

	ret = fmt.Sprintf("%s&hmac=%s", pich_data, hmac);
	return ret, nil; 
}