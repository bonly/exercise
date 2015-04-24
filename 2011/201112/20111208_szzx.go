package main 

import (
  "net/http"
  "net/url"
  "log"
  "io/ioutil"
  "crypto/md5"
  "encoding/hex"
  //"crypto/des"
  //"crypto/cipher"  
  "os/exec"
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

    www := "http://pay3.shenzhoufu.com/interface/version3/serverconnszx/entry-noxml.aspx";
	resp, err := cli.Get(www);
    //resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf);
    //resp, err := http.PostForm("http://example.com/form",url.Values{"key": {"Value"}, "id": {"123"}});

    data := SZZX{"3","151525","100","103","http://183.61.112.5:8080/main", 
                 "","test User","abc@163.com","","1",
                 "0",""};
    val := url.Values{};
    val.Add("version", data.Version);
    val.Add("merId", data.MerId);
    val.Add("payMoney", data.PayMoney);
    val.Add("orderId", data.OrderId);
    val.Add("returnUrl", data.ReturnUrl);
    val.Add("cardInfo", data.CardInfo);
    val.Add("merUserName", data.MerUserName);
    val.Add("merUserMail", data.MerUserMail);
    val.Add("privateField", data.PrivateField);
    val.Add("verifyType", data.VerifyType); //固定传1
    val.Add("cardTypeCombine", data.CardTypeCombine); //0：移动；1：联通；2：电信
    val.Add("md5String", "");

    //卡信息
    card := Card{"50", "11387014148153216", "013809579253610971"};

    cmd_argv := []string{"20111206_DES.php", card.Value, card.SN, card.Passwd, desKey};
    ret_c := exec.Command("php", cmd_argv...);
    ret_out, _ := ret_c.Output();
    data.CardInfo = string(ret_out);
    val.Set("cardInfo", data.CardInfo); //设置加密后信息

	log.Printf("orgdata: %v", card);
	log.Println("crydata: ", data.CardInfo); 

    //设置md5值
    md := md5.New();
    md.Write([]byte(data.Version+data.MerId+data.PayMoney+data.OrderId+data.ReturnUrl+
    	     data.CardInfo+data.PrivateField+data.VerifyType+privateKey));

    log.Println(data.Version+data.MerId+data.PayMoney+data.OrderId+data.ReturnUrl+
    	     data.CardInfo+data.PrivateField+data.VerifyType+privateKey);
    val.Set("md5String", hex.EncodeToString(md.Sum(nil)));
    
    req, err := http.NewRequest("POST",www + "?" + val.Encode(),nil);
    
    log.Println("url:", www+"?"+val.Encode());
    resp, err = cli.Do(req);

    if err != nil {
       log.Println("handle error");
	}

	body, err := ioutil.ReadAll(resp.Body);

	log.Println("res: ", string(body));
}
