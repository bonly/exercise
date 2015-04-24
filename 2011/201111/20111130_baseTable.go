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
  "encoding/base64"
  "bytes"
)

type SZZX struct{
	Version, MerId, PayMoney, OrderId, ReturnUrl, 
	CardInfo, MerUserName, MerUserMail, PrivateField, VerifyType,
	CardTypeCombine, Md5String, PrivateKey string;
};

type Card struct{
    Value, SN, Passwd string;
};

const (
    base64Table=`ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/`;
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

    data := SZZX{"3","151525","100","100","127.0.0.1", 
                 "","test User","abc@163.com","","1",
                 "0","","123456"};
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

    //解base64密钥SGVsbG8sIHBsYXlncm91bmQ=
    //
    enbyte, err := base64.NewEncoding(base64Table).DecodeString("fNCrhSynUm4=");
    if err != nil {  
        log.Println(err.Error());
    }  
    //log.Printf("密钥: %s", string(enbyte));
    log.Printf("密钥: %v", enbyte);

    //des加密
	//key := []byte("sfe023f_sefiel#fi32lf3e!");  
	//key := []byte(enbyte); 
	//key := []byte("sfe023f_"); 
	crytext, err := DesEncrypt([]byte(cardstr), enbyte);
	//crytext, err := TripleDesEncrypt([]byte(cardstr), key);
	if err != nil {
		panic(err);
	}
	data.CardInfo=string(base64.NewEncoding(base64Table).EncodeToString(crytext));
	log.Println("orgdata: ", cardstr);
	log.Println("crydata: ", data.CardInfo);
//	origData, err := DesDecrypt(result, key)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(origData))

    val.Set("cardInfo", data.CardInfo); //设置加密后信息

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


func DesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	//origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func DesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	//origData = ZeroUnPadding(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}


// 3DES加密
func TripleDesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// 3DES解密
func TripleDesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}