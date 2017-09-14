/*
auth: bonly
*/
package main
import (
    "io/ioutil"
    "net/http"
    // "net/url"
    "fmt"
    "encoding/json"
    "encoding/xml"
    "encoding/base64"
    "time"
    "encoding/hex"
    "crypto/md5"
    "strings"
    "flag"
)
  
//----------------------------------
// 在线接口文档：http://docs.yuntongxun.com/index.php/%E8%AF%AD%E9%9F%B3%E9%AA%8C%E8%AF%81%E7%A0%81
//http://docs.yuntongxun.com/index.php/Rest%E4%BB%8B%E7%BB%8D#2_.E7.BB.9F.E4.B8.80.E8.AF.B7.E6.B1.82.E5.8C.85.E5.A4.B4
//----------------------------------
const AccountSid = "8a48b55151e82a680151e8875e6700f5"; // Account Sid
const AuthToken = "53f5c3599ceb4488a090a74366c6790f";  // Auth Token
const SubAccountSid = "";  //子帐号
const SubToken = ""; //子帐号Token
const AppId = "aaf98f8954939ed50154bd30b9eb2af2";
const Srv = "https://app.cloopen.com:8883";
// const Srv = "https://sandboxapp.cloopen.com:8883";
// const Ver = "/2013-03-22";
const Ver = "/2013-12-26";

var tel *string = flag.String("t", "17701906025", "电话号码");
var sn *string = flag.String("c", "123456", "验证码");

func main(){
    flag.Parse();
  
    //1.发送语音验证码
    Request();
  
}
  
func Sign() (sign string, auth string){
    tn := time.Now();
    nw := fmt.Sprintf("%d%02d%02d%02d%02d%02d", 
        tn.Year(), tn.Month(), tn.Day(), tn.Hour(), tn.Minute(), tn.Second());

    md := md5.New();
    md.Write([]byte(AccountSid+AuthToken+nw));
    sign = strings.ToUpper(hex.EncodeToString(md.Sum(nil)));

    auth = base64.StdEncoding.EncodeToString([]byte(AccountSid+":"+nw));
    fmt.Println("sign: ", sign);
    fmt.Println("auth: ", auth);
    return sign, auth;
}

type Param struct{
    AppId string `json:"appId"`;
    VerifyCode string `json:"verifyCode"`;
    To string `json:"to"`;
    DisplayNum string `json:"displayNum"`;
    PlayTimes string `json:"playTimes"`;
    // RespUrl string `json:"respUrl"`;
    // Lang string `json:"lang"`;
    // UserData string `json:"userData"`;
    // WelcomePrompt string `json:"welcomePrompt"`;
    // PlayVerifyCode string `json:"playVerifyCode"`;
    // MaxCallTime string `json:"maxCallTime"`;
};

type Return struct{
    XMLName xml.Name `xml:"Response"`;
    StatusCode string `xml:"statusCode", json:"statusCode"`;
    StatusMsg string `xml:"statusMsg", json:"statusMsg"`;
};

//1.发送语音验证码
func Request(){
    sign, auth := Sign();
    //请求地址
    var SrvURL string = Srv + Ver + "/Accounts/" +
                  AccountSid + "/Calls/VoiceVerify" +
                  "?sig=" + sign;
  
    var param Param;
    param.AppId = AppId;
    param.To = *tel;
    param.VerifyCode = *sn;
    param.PlayTimes = "3";
    param.DisplayNum = "4006099222";
    // param.Lang = "zh";

    dat, err := json.Marshal(param);
    if err != nil{
        fmt.Println("json: ", err);
    }
  
    // //发送请求
    fmt.Println("请求：", SrvURL);
    fmt.Println("数据：", string(dat));
    data, err:=Post(SrvURL, string(dat), auth);
    fmt.Println("返回:", string(data));
    if err!=nil{
        fmt.Errorf("请求失败,错误信息:\r\n%v",err);
    }else{
        var netReturn Return;
        // xml.Unmarshal(data, &netReturn);
        json.Unmarshal(data, &netReturn);
        fmt.Printf("接口返回result字段是:\r\n%v\n", netReturn.StatusCode);
    }
}
  
// post 网络请求 ,params 是url.Values类型
func Post(apiURL string, params string, auth string)(rs[]byte,err error){
    cli := &http.Client{};
    req, err := http.NewRequest("POST", apiURL, strings.NewReader(params));
    req.Header.Set("Accept","application/json");
    req.Header.Set("Content-Type","application/json;charset=utf-8");
    req.Header.Set("Authorization", auth);
    resp,err:=cli.Do(req);
    if err!=nil{
        return nil ,err;
    }
    defer resp.Body.Close();
    return ioutil.ReadAll(resp.Body);
}
