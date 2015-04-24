package main 

import (
  "fmt"
  // "time"
  "reflect"
  "net/url"
  // "net/http"
  "log"
  // "io/ioutil"
  "crypto/md5"
  "encoding/hex"  
  "os/exec"    
  "encoding/xml"
  // "github.com/qiniu/iconv"
  // "strconv"
  // "strings"
  "sort"
)

type Q_CFT struct{
Sign_type string `xml:"sign_type"`; //签名类型，取值：MD5、RSA，默认：MD5
Service_version string `xml:"service_version"`;//版本号，默认为1.0
Input_charset string `xml:"input_charset"`; //字符集 字符编码,取值：GBK、UTF-8
Sign string `xml:"sign"`; //签名
Sign_key_index string `xml:"sign_key_index"`;//密钥序号

Bank_type string `xml:"bank_type"`;//银行类型，默认为“DEFAULT”跳转到财付通支付中心。银行直连编码及额度请与技术支持联系
Body string `xml:"body"`; //商品描述
Attach string `xml:"attach"`; //商品描述,附加数据，原样返回
Return_url string `xml:"return_url"`; //返回URL:通过该路径直接将支付结果以Get的方式返回
Notify_url string `xml:"notify_url"`; //通知URL
Buyer_id string `xml:"buyer_id"`; //买方财付通账号,买方的财付通账户(QQ 或EMAIL)。若商户没有传该参数，则在财付通支付页面，买家需要输入其财付通账户。
Partner string `xml:"partner"`; //商户号,由财付通统一分配的10位正整数(120XXXXXXX)号
Out_trade_no string `xml:"out_trade_no"`; //商户系统内部的订单号,32个字符内、可包含字母,确保在商户系统唯一
Total_fee string `xml:"total_fee"`; //订单总金额，单位为分
Fee_type string `xml:"fee_type"`; //币种:现金支付币种,取值：1（人民币）,默认值是1，暂只支持1
Spbill_create_ip string `xml:"spbill_create_ip"`; //用户IP:订单生成的机器IP，指用户浏览器端IP，不是商户服务器IP
Time_start string `xml:"time_start"`; //交易起始时间
//订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。时区为GMT+8 beijing。该时间取自商户服务器
Time_expire string `xml:"time_expire"`; //交易结束时间
//订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。时区为GMT+8 beijing。该时间取自商户服务器
Transport_fee string `xml:"transport_fee"`; //物流费用，单位为分。如果有值，必须保证transport_fee + product_fee=total_fee
Product_fee string `xml:"product_fee"`; // 商品费用，单位为分。如果有值，必须保证transport_fee + product_fee=total_fee
Goods_tag  string `xml:"goods_tag"`; //商品标记，优惠券时可能用到

server string;
};

func (this *Q_CFT) Proc_Field(key string)(ret string){
	//数据结构转成map
    mp, err := ToMap(*this, "xml");
    if err != nil{
    	log.Println("err=", err.Error());
    }
    
    //提取key字段, 排序
    keys := make([]string, 0, len(mp));
    for skey := range mp {
        keys = append(keys, skey);
    }
    sort.Strings(keys);

    //拼接字段内容,并给对象赋值
    var str string;
    val := url.Values{};
    for _, kname := range keys{
    	if len(mp[kname].(string)) > 0 && !(kname == "sign"){
    		str += kname + "=" + mp[kname].(string) + "&";
            val.Add(kname, mp[kname].(string));  
        }  	
    }
    str += "key="+key;
    log.Println("org str: ", str);

    //生成MD5
	md := md5.New();
	md.Write([]byte(str));       
	// this.Sign = strings.ToUpper(hex.EncodeToString(md.Sum(nil)));
	this.Sign = hex.EncodeToString(md.Sum(nil));
	val.Add("sign", this.Sign);

	//生成编码后的url
	return val.Encode();
}

type R_CFT_Ret struct{
	Root    xml.Name  `xml:"root"`;
	Ret_code string `xml:"ret_code"`;
	Ret_msg string `xml:"ret_msg"`;
	Encrypt_data string `xml:"encrypt_data"`;
	Sign string `xml:"sign"`;
};


func main(){
	var apply Q_CFT;
	//设置值
	apply.server = "https://gw.tenpay.com/gateway/pay.htm";
	// apply.server = "https://wap.tenpay.com/cgi-bin/wappayv2.0/wappay_init.cgi";
	apply.Sign_type = "MD5";
	apply.Service_version = "1.0";
	apply.Input_charset = "GBK";
	apply.Sign_key_index = "1";

	apply.Bank_type = "DEFAULT";
	apply.Body = "test";
	apply.Attach = "";
	// apply.Return_url = "http://183.61.112.5:8090/caifutong";
	// apply.Notify_url = "http://183.61.112.5:8090/caifutong";
	apply.Return_url = "http://localhost:8080/payReturnUrl.php";
	apply.Notify_url = "http://localhost:8080/payNotifyUrl.php";
	//apply.Buyer_id ="";
	apply.Partner = "1900000109";
	apply.Out_trade_no ="201407010611199805";  //流水号
	apply.Total_fee = "2"; //以分为单位
	apply.Fee_type = "1";
	apply.Spbill_create_ip = "127.0.0.1"; //客户端IP

    apply.Time_start = "20140701061119";
    // tn := time.Now();
    // apply.Time_start = fmt.Sprintf("%d%02d%02d%02d%02d%02d",tn.Year(),tn.Month(),tn.Day(), tn.Hour(), tn.Minute(), tn.Second());
    // to := time.Now().Add(1*time.Hour);
    // apply.Time_expire = fmt.Sprintf("%d%02d%02d%02d%02d%02d",to.Year(),to.Month(),to.Day(), to.Hour(), to.Minute(), to.Second());
    
    // apply.Transport_fee = "0";
    // apply.Product_fee = "1";
    //apply.Goods_tag = "1";

    field := apply.Proc_Field("8934e7d15453e97507ef794cf7b0519d");

    addr := apply.server + "?" + field;
	log.Println("send: ", addr);

	cmd_argv := []string{addr};
	ret_c := exec.Command("chr", cmd_argv...);
	op, _ := ret_c.Output();
	log.Println(op);
}


func ToMap(in interface{}, tag string) (map[string]interface{}, error) {
	out := make(map[string]interface{});

	v := reflect.ValueOf(in);
	if v.Kind() == reflect.Ptr {
		v = v.Elem();
	}

	// 只处理结构体
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts structs; got %T", v);
	}

	typ := v.Type();
	for i := 0; i < v.NumField(); i++ {
		// 转换为字段
		fi := typ.Field(i);
		if tagv := fi.Tag.Get(tag); tagv != "" {
			out[tagv] = v.Field(i).Interface();
		}
	}
	return out, nil;
}