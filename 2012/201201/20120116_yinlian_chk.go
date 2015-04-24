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
  // "fmt"
  // "time"
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

    orderTime                   string = "20140523161521"; //关键数据，申请流水时的字段
    orderNumber                 string = "000000000072";  //关键数据

    merReserved                 string;
    reqReserved                 string = "aaa";
    sysReserved                 string;

    upmp                        string = "http://222.66.233.198:8080/gateway/merchant/query"
)

func main(){
    cli := &http.Client{};
    resp, err := cli.Get(upmp);

    // orderTime = "20140523150727";

    val := url.Values{};
    val.Add("version", version);
    val.Add("charset", charset);
    val.Add("transType", transType);
    val.Add("merId", merId);
    val.Add("orderTime", orderTime);
    val.Add("orderNumber", orderNumber);
    //val.Add("merReserved", merReserved);
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
    log.Println("Stat: ", res.Get("transStatus"));
}

func mkstr(key string) (str string){
    str =  "charset=" + charset + "&" +
           "merId=" + merId + "&" +
           "orderNumber=" + orderNumber + "&" +
           "orderTime=" + orderTime + "&" +
           "transType=" + transType + "&" +
           "version=" + version + "&" + key;
    return str;
}
/*
三个备用字段中， mer和sys的内容是{}格式的
对密钥md5一次，拼在原生字段（不包括签名和签名方法字段）的后面（注意不能url.encode()的）
再把上面的串md5作为签名
*/