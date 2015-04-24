package main 

import (
  "net/http"
  "net/url"
  "log"
  "io/ioutil"
  "crypto/md5"
  "encoding/hex"
  "crypto/des"
  "crypto/cipher"  
)

type SZZX struct{
	Version, MerId, PayMoney, OrderId, ReturnUrl, 
	CardInfo, MerUserName, MerUserMail, PrivateField, VerifyType,
	CardTypeCombine, Md5String, PrivateKey string;
};

type Card struct{
    Value, SN, Passwd string;
};

/*
直接默认发送
*/
func main(){
	cli := &http.Client{};

    www := "http://pay3.shenzhoufu.com/interface/version3/serverconnszx/entry-noxml.aspx";
	resp, err := cli.Get(www);
    //resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf);
    //resp, err := http.PostForm("http://example.com/form",url.Values{"key": {"Value"}, "id": {"123"}});

    data := SZZX{"3","151525","100","100","127.0.0.1", "","test User","abc@163.com","","0","0","","123456"};
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
    card := Card{"100", "3314999333", "cardpasswd"};
    cardstr := card.Value + "@" + card.SN + "@" + card.Passwd;

    key_text := "fNCrhSyn";  //密钥
    var iv = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07};   //向量

    cyp, err := des.NewCipher([]byte(key_text)); //创建加密算法
    if err != nil {
        log.Printf("Error: NewCipher(%d bytes) = %s", len(key_text), err)
        return;
    }    
    cfb := cipher.NewCFBEncrypter(cyp, iv);
    //cfb := cipher.NewCFBEncrypter(cyp, byte[](key_text[0:8]));
    ciphertext := make([]byte, len(cardstr));
    cfb.XORKeyStream(ciphertext, []byte(cardstr));  //加密
    //log.Printf("%s=>%x\n", cardstr, ciphertext);    

    val.Set("cardInfo", string(ciphertext)); //设置加密后信息

    //设置md5值
    md := md5.New();
    md.Write([]byte(data.Version+data.MerId+data.PayMoney+data.OrderId+data.ReturnUrl+
    	     data.CardInfo+data.PrivateField+data.VerifyType+data.PrivateKey));

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
