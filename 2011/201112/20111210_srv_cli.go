/*
支付平台的处理后，模拟其返回结果给请求方
*/
package main 

import (
  "net/http"
  //"net/url"
  "log"
  "strings"
  "io/ioutil"
  //"crypto/md5"
  //"encoding/hex"
  //"crypto/des"
  //"crypto/cipher"  
  //"os/exec"
  //"bytes"
)

type SZZX struct{
	Version, MerId, PayMoney, OrderId, ReturnUrl, 
	CardInfo, MerUserName, MerUserMail, PrivateField, VerifyType,
	CardTypeCombine, Md5String  string;
};

type Card struct{
    Value, SN, Passwd string;
};

const (
    //base64Table=`ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/`;
    desKey="fNCrhSynUm4=";
    privateKey="123456";
)
/*
直接默认发送
*/
func main(){
	cli := &http.Client{};

  www := "http://127.0.0.1:8080/main";
	resp, err := cli.Get(www);
  
  req, err := http.NewRequest("POST",www, 
    strings.NewReader(`cardInfo=&md5String=0a7c488a0fb191d14d3a55132939ddd5&merId=151525&merUserMail=abc%40163.com&merUserName=test+User&payDetails=&endTime=20140409100726&errcode=201&version=3&payMoney=100&privateField=&payResult=0&signString=mLiEyMX75spAdMSUDOp1EmhlFVcGr7UB1mV0WKBsJa5woJttakALgf9Q%2BUnP98Zsx4dQ6xrC4xdtAAq7noaZRHppJMi3Q51vcPPv7hCID2sraAONyKEmIYLJjpHQ1n%2BprHe1ngCMRB1SdReRlruWr2tx2PkFPzWwe2cVQnZslqk%3D&szfOrderNo=82891369&orderId=102`));
  
  resp, err = cli.Do(req);

  if err != nil {
     log.Println("handle error");
  }

  body, err := ioutil.ReadAll(resp.Body);

	log.Println("res: ", string(body));
}
