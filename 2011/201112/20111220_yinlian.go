package main

import (
  //"testing"
  "net/http"
  "net/url"
  "io/ioutil"
  "crypto/md5"
  "encoding/hex"
  "log"  
  "bytes"
  "fmt"
  "time"
  //"strings"
)

var (
	//timezone                    string = "Asia/Shanghai"; //时区

    version                     string = "1.0.0"; // 版本号
    charset                     string = "UTF-8"; // 字符编码
    signMethod                  string = "MD5"; // 签名方法，目前仅支持MD5
    signature                   string;
    
    transType                   string = "01"; //01:消费 02:预授权

    merId                       string = "880000000000084"; // 商户号
    security_key                string = "vWGyoD0Clm8v098ezbhmI9n4HOWBFsNE"; // 商户密钥
    backEndUrl                  string = "http://183.61.112.5:8090/pay"; // 后台通知地址
    frontEndUrl                 string = "http://183.61.112.5:8090/pay"; // 前台通知地址
 
    //backEndUrl                  string = "abcpay"; // 后台通知地址
    //frontEndUrl                 string = "acdpay"; // 前台通知地址

    //acqCode                     string = "";

    orderTime                   string ;//= "20140429143000";
    orderTimeout                string ;//= "20140429150020";
    orderNumber                 string = "10000001";
    orderAmount                 string = "100";
    orderCurrency               string = "156";
    orderDescription            string = "test";

    merReserved                 string;
    reqReserved                 string = "aaa";
    sysReserved                 string;

    upmp                        string = "http://222.66.233.198:8080/gateway/merchant/trade";
)

func main(){
    cli := &http.Client{};
    resp, err := cli.Get(upmp);

    tn := time.Now();
    orderTime = fmt.Sprintf("%d%02d%02d%02d%02d%02d",tn.Year(),tn.Month(),tn.Day(), tn.Hour(), tn.Minute(), tn.Second());
    to := time.Now().Add(1*time.Hour);
    orderTimeout = fmt.Sprintf("%d%02d%02d%02d%02d%02d",to.Year(),to.Month(),to.Day(), to.Hour(), to.Minute(), to.Second());

    val := url.Values{};
    val.Add("version", version);
    val.Add("charset", charset);
    val.Add("transType", transType);
    val.Add("merId", merId);
    val.Add("backEndUrl", backEndUrl);
    val.Add("frontEndUrl", frontEndUrl);
    //val.Add("acqCode", acqCode);
    val.Add("orderTime", orderTime);
    val.Add("orderTimeout", orderTimeout);
    val.Add("orderNumber", orderNumber);
    val.Add("orderAmount", orderAmount);
    val.Add("orderCurrency", orderCurrency);
    val.Add("orderDescription", orderDescription);
    //val.Add("merReserved", merReserved);
    val.Add("reqReserved", reqReserved);
    //val.Add("sysReserved", sysReserved);

    keymd5 := md5.New();
    keymd5.Write([]byte(security_key));
    strkey := hex.EncodeToString(keymd5.Sum(nil));

    md := md5.New();
    str := mkstr(strkey);
    md.Write([]byte(str));

    log.Println("sign: ", str);

    val.Add("signMethod", signMethod);
    val.Add("signature", hex.EncodeToString(md.Sum(nil)));

    postDataBytes := []byte(val.Encode());
    postBytesReader := bytes.NewReader(postDataBytes);

    req, err := http.NewRequest("POST", upmp, postBytesReader);
    req.Header.Set("Content-Type", "application/octet-stream");
    
    log.Println("url: ", upmp);
    log.Println("data: ", val.Encode());
    resp, err = cli.Do(req);

    if err != nil {
       log.Println("handle error");
    }

    body, err := ioutil.ReadAll(resp.Body);

    log.Println("res: ", string(body));      

    res, err := url.ParseQuery(string(body)); 
    if err != nil{
        log.Println("err: ", err);
    }

    log.Println("Msg: ", res.Get("respMsg"));
}

func mkstr(key string) (str string){
    str = "backEndUrl=" + backEndUrl + "&" +
           "charset=" + charset + "&" +
           "frontEndUrl=" + frontEndUrl + "&" +
           "merId=" + merId + "&" +
           "orderAmount=" + orderAmount + "&" +
           "orderCurrency=" + orderCurrency + "&" +
           "orderDescription=" + orderDescription + "&" +
           "orderNumber=" + orderNumber + "&" +
           "orderTime=" + orderTime + "&" +
           "orderTimeout=" + orderTimeout + "&" +
           "reqReserved=" + reqReserved + "&" +
           "transType=" + transType + "&" +
           "version=" + version + "&" + key;
    return str;
}
/*
三个备用字段中， mer和sys的内容是{}格式的
对密钥md5一次，拼在原生字段（不包括签名和签名方法字段）的后面（注意不能url.encode()的）
再把上面的串md5作为签名
*/