package main 

import (
  "fmt"
  "time"
  // "net/url"
  "net/http"
  "log"
  "io/ioutil"
  "crypto/md5"
  "encoding/hex"  
  "os/exec"    
  "encoding/xml"
  "github.com/qiniu/iconv"
  "strconv"
  "strings"
  "net/url"  
)

type Q_HF_Short struct{
agent_id, // 给商户分配的唯一标识 是 
version, //版本号，固定为 1 是 
user_identity, // 商户用户唯一标识  
hy_auth_uid, // 用户快捷支付的授权码  
mobile, // 用户手机号码  
device_type, //设备类型： 0=Wap 1=Web 2=AndroidApp
device_id, //设备唯一标识  
display, //终端显示样式： 0= Wap（默认） 
custom_page, //自定义支付页面： 0=否，支付页面使用汇付宝 提供（默认）
return_url, // 支付页面返回地址
notify_url, //异步通知地址  
timestamp, //时间戳  
bank_id, //签约的银行ID默认为0，用户没有签约时在汇付宝页面进行签约的情况下使用该属性可以直接跳转到指定的商家支持的银行进行签约支付。

agent_bill_id, //商户订单号，必须唯一
agent_bill_time, // 商户订单时间，格式为“yyyyMMddhhmmss”,例20130910170601
pay_amt, //支付金额，单位：元，保留二位小数
goods_name, // 商品名称 是 
goods_note,  //商品描述
goods_num, // 商品数量 
user_ip, //用户IP地址  
auth_card_type, // 银行卡类型： 0=储蓄卡（默认） 1=信用卡
ext_param1, // 商户保留，支付完成后的通知会将该参数返回给商户
ext_param2 string; //商户保留，支付完成后的通知会将该参数返回给商户

server string;
};

type R_HF_Short_Ret struct{
	Root    xml.Name  `xml:"root"`;
	Ret_code string `xml:"ret_code"`;
	Ret_msg string `xml:"ret_msg"`;
	Encrypt_data string `xml:"encrypt_data"`;
	Sign string `xml:"sign"`;
};

func (this *Q_HF_Short)CryData()(ret string){
	ret = "agent_bill_id=" + this.agent_bill_id + "&" +
	       "agent_bill_time=" + this.agent_bill_time + "&" + 
	       "agent_id=" + this.agent_id + "&" +
	       "auth_card_type=" + this.auth_card_type + "&" +
	       "custom_page=" + this.custom_page + "&" +
	       "device_id=" + this.device_id + "&" +
	       "device_type=" + this.device_type + "&" +
	       "ext_param1=" + this.ext_param1 + "&" +
	       "ext_param2=" + this.ext_param2 + "&" +
	       "goods_name=" + this.goods_name + "&" +
	       "goods_note=" + this.goods_note + "&" +
	       "goods_num=" + this.goods_num + "&" +
	       "hy_auth_uid=" + this.hy_auth_uid + "&" +
	       "mobile=" + "" + "&" +
	       "notify_url=" + this.notify_url + "&" +
	       "pay_amt=" + this.pay_amt + "&" +
	       "return_url=" + this.return_url + "&" +
	       "timestamp="+ this.timestamp + "&" +
	       "user_identity=" + this.user_identity + "&" +
	       "user_ip=" + this.user_ip + "&" +
	       "version=" + "1";   
	return;
}

func (this *Q_HF_Short)Sign(key string)(ret string){
	ret = "agent_bill_id=" + this.agent_bill_id + "&" +
	       "agent_bill_time=" + this.agent_bill_time + "&" + 
	       "agent_id=" + this.agent_id + "&" +
	       "auth_card_type=" + this.auth_card_type + "&" +
	       "custom_page=" + this.custom_page + "&" +
	       "device_id=" + this.device_id + "&" +
	       "device_type=" + this.device_type + "&" +
	       "ext_param1=" + this.ext_param1 + "&" +
	       "ext_param2=" + this.ext_param2 + "&" +
	       "goods_name=" + this.goods_name + "&" +
	       "goods_note=" + this.goods_note + "&" +
	       "goods_num=" + this.goods_num + "&" +
	       "hy_auth_uid=" + this.hy_auth_uid + "&" +
	       "key=" + key + "&" +
	       "mobile=" + "" + "&" +
	       "notify_url=" + this.notify_url + "&" +
	       "pay_amt=" + this.pay_amt + "&" +
	       "return_url=" + this.return_url + "&" +
	       "timestamp="+ this.timestamp + "&" +
	       "user_identity=" + this.user_identity + "&" +
	       "user_ip=" + this.user_ip + "&" +
	       "version=" + "1";   
	return;
}

func main(){
	var apply Q_HF_Short;
	//设置值
	// apply.server = "https://Pay.Heepay.com/ShortPay/SubmitOrder.aspx";
	apply.server = "http://211.103.157.45/PayHeepay/ShortPay/SubmitOrder.aspx"
	apply.agent_id = "1733154";
	apply.version = "1";
	apply.device_type = "2";
	apply.custom_page = "0";
	apply.return_url = "http://183.61.112.5:8090/huifubao";
	apply.notify_url = "http://127.0.0.112";
	apply.timestamp ="";

	apply.hy_auth_uid ="pre-b0b933b70673458a9c0a2ad0eb76ee09";  //第一次时为空

	apply.agent_bill_id = "000009";
    tn := time.Now();
    apply.agent_bill_time = fmt.Sprintf("%d%02d%02d%02d%02d%02d",tn.Year(),tn.Month(),tn.Day(), tn.Hour(), tn.Minute(), tn.Second());
    apply.timestamp = strconv.FormatInt(tn.Unix()*1000, 10);
    apply.pay_amt = "0.01";
    apply.goods_name = "test";
    apply.goods_num = "1";
    apply.user_ip = "127.0.0.1";

    apply.auth_card_type = "1"; //0:储蓄卡,1:信用卡

    //生成数据对象
	// val := url.Values{};
	// val.Add("agent_id", apply.agent_id);
	// val.Add("version", apply.version);
	// val.Add("device_type", apply.device_type);
	// val.Add("custom_page", apply.custom_page);
	// val.Add("return_url", apply.return_url);
	// val.Add("notify_url", apply.notify_url);
	// val.Add("timestamp", apply.timestamp);
	// val.Add("agent_bill_id", apply.agent_bill_id);
	// val.Add("agent_bill_time", apply.agent_bill_time);
	// val.Add("pay_amt", apply.pay_amt);
	// val.Add("goods_name", apply.goods_name);
	// val.Add("goods_num", apply.goods_num);
	// val.Add("user_ip", apply.user_ip);
    // ucode := val.Encode();

    allstr := apply.CryData();
    log.Println("cry data: ", allstr);

	cmd_argv := []string{"aes.php", "en","76NOWeXeZDTlwfwxNAOuhqDoFuGs7xnDY65u1Or5XWg=", allstr};
	ret_c := exec.Command("php", cmd_argv...);
	encrypt_data, _ := ret_c.Output();

    allstr = apply.Sign(strings.ToLower("B524CA826E5A4036A0F167AE"));
    log.Println("sign field: ", allstr);
    md := md5.New();
    md.Write([]byte(allstr));       
    strSign := hex.EncodeToString(md.Sum(nil));	    

    addr := apply.server + "?" + "agent_id=" + apply.agent_id + "&encrypt_data=" + string(encrypt_data) + "&sign=" + string(strSign);
	req, err := http.NewRequest("POST", addr, nil);
	log.Println("send: ", addr);

    cli := &http.Client{};
    resp, err := cli.Do(req);
    if err != nil {
       log.Println("handle error: ", err.Error());
    }	
    body, err := ioutil.ReadAll(resp.Body);

    // gb2312转utf8
    cd, err := iconv.Open("utf-8", "gb2312");
    if err != nil{
    	log.Println("iconv open failed! ", err.Error());
    }
    defer cd.Close();

    // strUtf8 := cd.ConvString(string(body));
    log.Println("res: ", string(body));     

    var ret R_HF_Short_Ret;
    // err = xml.Unmarshal([]byte(strUtf8), &ret);
    err = xml.Unmarshal(body, &ret);
    if err != nil {
        log.Printf("error: %v", err);
    }

    log.Println("ret_code=", ret.Ret_code);
    log.Println("ret_msg=", ret.Ret_msg);
    log.Println("encrypt_data=", ret.Encrypt_data);

	cmd_argv = []string{"aes.php", "de","76NOWeXeZDTlwfwxNAOuhqDoFuGs7xnDY65u1Or5XWg=", ret.Encrypt_data};
	ret_c = exec.Command("php", cmd_argv...);
	decrypt_data, _ := ret_c.Output();    

    //解密内容
	log.Println("decode_data=", string(decrypt_data));
	res, err := url.ParseQuery(string(decrypt_data)); //还原结构
    if err != nil{
        log.Println("err: ", err.Error());
    }	
    log.Println("hy_auth_uid=", res.Get("hy_auth_uid"));
}