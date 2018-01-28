package main 

import(
"net/http"
"encoding/json"
"crypto/md5"
"encoding/hex"
"strings"
"fmt"
"io/ioutil"
"reflect"
"sort"
)

var srv = "http://link.beta.quhuhu.com";
var pms = "xbed";

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
	verify();
}

func verify(){
	var dat Search_Channel_Rate_Plan;
	dat.PmsId = pms;
	dat.Version = "1.0";

	// mp, _:= ToMap(dat, "sr");	
	// fmt.Println(mp);	
	// Display(&dat);
	mp := make(map[string]string);
	ToStrMap(&dat, "sr", &mp);
	fmt.Printf("org: %v\n", mp);
	sort.Strings(mp);
}

func send(){
	var dat Search_Channel_Rate_Plan;
	dat.PmsId = pms;
	dat.Version = "1.0";

	mp, err := ToMap(dat, "json");
	if err != nil{
		fmt.Printf("data map: %s\n", err.Error());
		return;
	}

	dat.Hmac = Gen_hmac(&mp);
	mp["hmac"] = dat.Hmac;

	js, err := json.MarshalIndent(dat, " ", " ");

	qry := string(js);
	fmt.Printf("qry:\n%s\n", qry);

	req, err := http.NewRequest("POST", srv + "/api/" + pms + "/search/searchChannelRatePlan.do",
		strings.NewReader(qry));

	req.Header.Set("Content-Type","application/x-www-form-urlencoded");
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");	

	if err != nil{
		fmt.Printf("new req: ", err);
		return;
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded");
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

func Display(in interface{}){
	reflectType := reflect.TypeOf(in).Elem(); //取类型对象
	reflectValue := reflect.ValueOf(in).Elem();  //取值对象

	for idx := 0; idx < reflectType.NumField(); idx++{  //遍历所有字段
		typeName := reflectType.Field(idx).Name;  //取每个字段的名字

		valueType := reflectValue.Field(idx).Type(); //字段类的字符串式的描述
		valueValue := reflectValue.Field(idx).Interface();  //把字段转换为接口

		switch reflectValue.Field(idx).Kind(){ //字段类型判断
	 		case reflect.String:
	            fmt.Printf("%s : %s(%s)\n", typeName, valueValue, valueType);
	            break;
	        case reflect.Int32:
	            fmt.Printf("%s : %i(%s)\n", typeName, valueValue, valueType);
	            break;
	        case reflect.Struct:
	            fmt.Printf("%s : it is %s\n", typeName, valueType);
	            val := reflectValue.Field(idx).Addr(); //取字段的地址
	            Display(val.Interface());  //深度遍历
	            break;
		}
	}
}

func ToStrMap(in interface{}, tag string, mp *(map[string]string))(err error){
	reflectType := reflect.TypeOf(in).Elem(); //取类型对象
	reflectValue := reflect.ValueOf(in).Elem();  //取值对象

	for idx := 0; idx < reflectType.NumField(); idx++{  //遍历所有字段
		typeName := reflectType.Field(idx).Name;  //取每个字段的名字

		valueType := reflectValue.Field(idx).Type(); //字段类的字符串式的描述
		valueValue := reflectValue.Field(idx).Interface();  //把字段转换为接口

		switch reflectValue.Field(idx).Kind(){ //字段类型判断
	 		case reflect.String:
	            fmt.Printf("%s : %s(%s)\n", typeName, valueValue, valueType);
	            (*mp)[typeName] = fmt.Sprintf("%s", valueValue);
	            break;
	        case reflect.Int32:
	            fmt.Printf("%s : %i(%s)\n", typeName, valueValue, valueType);
	            break;
	        case reflect.Struct:
	            fmt.Printf("%s : it is %s\n", typeName, valueType);
	            val := reflectValue.Field(idx).Addr(); //取字段的地址
	            ToStrMap(val.Interface(), tag, mp);  //深度遍历
	            break;
		}
	}
	return nil;
}

//转换数据结构到map
func ToMap(in interface{}, tag string) (map[string]interface{}, error){
    out := make(map[string]interface{});

    v := reflect.ValueOf(in);
    fmt.Printf("%v\n", v.Kind());
    if v.Kind() == reflect.Ptr {
    	fmt.Println("in ptr");
        v = v.Elem(); //取interface所指的元素
    }

    // 只处理 structs
    if v.Kind() != reflect.Struct {
        return nil, fmt.Errorf("ToMap only accepts structs; got %T", v);
    }

    typ := v.Type();
    fmt.Println("type: ", typ);
    for i := 0; i < v.NumField(); i++ {
        // 取一个字段
        fi := typ.Field(i);
        if tagv := fi.Tag.Get(tag); tagv != "" { //处理指定标签字段
        	vl := reflect.TypeOf(fi);
        	fmt.Printf("in %v\n", vl.Kind());
            // 推入map中
            out[tagv] = v.Field(i).Interface();
            // glog.Info("push: ", tagv);
        }
    }
    return out, nil;
}

//产生签名
func Gen_hmac(vl *map[string]interface{})(ret string){
	var sv []string;
	for _,v := range *vl {
		if v != nil{
			switch v.(type){
			case string:
				sv = append(sv, v.(string));
			}
		}
	}
	sort.Strings(sv);

	all_data := strings.Join(sv,"");
	fmt.Printf("md5 data: %s\n",all_data);
	md := md5.New();
	md.Write([]byte(all_data));
	return strings.ToUpper(hex.EncodeToString(md.Sum(nil)));
}