package main 

import (
  "fmt"
  "net/http"
  "log"
  "net/url"
  "io/ioutil"
  "github.com/qiniu/iconv"
  "time"
  "crypto/md5"
  "encoding/hex"  
  "crypto/des"
  "crypto/cipher"
  "encoding/base64"
  "bytes"  
)

type Q_HF_Charge struct{
agent_id, // int 商家ID (必填) 汇付宝注册的账户七位数字ID
bill_id, // string 商家提交的唯一订单号（必填）必须唯一6到50位
bill_time, // string 商户订单时间(必填 格式为 yyyyMMddHHmmss 4位年+2位月+2位日+2位时+2位分+2位秒)
card_type, // int 卡类型（参考3.6说明）
card_data, //string 骏卡一卡通最多支持3张(其他卡类只支持1张),
//格式为：卡号1,密码1,金额1|卡号2，密码2,金额2|卡号3，密码3，金额3），
//必填,双方协商的对称加密, 使用3DES加密,合作方不能保存记录卡密, （金额为整数。1 5 10 20 30等等）(必填)
pay_amt, //订单总金额 不可为空，单位：元， 订单金额为整数
//注1：订单金额为0默认支付卡内所有金额，仅限骏卡，
//注2：其他卡种请输入和卡面额相符的金额。
qq, // string 用户的联系方式qq(可选)
email,// string 用户Email(可选)
client_ip, // string 用户来源IP(必填)
notify_url, // string 支付后返回的商户处理页面(可选),
//URL参数是以http://开头的完整URL地址(后台处理) 不填写不通知可通过查询接口确定单据状态
desc, //string 简要说明(可选)
ext_param, // string商户自定义参数或扩展参数，接口按原值返回(可选)
time_stamp, //string提交时间戳(必填格式为yyyyMMddHHmmss 4位年+2位月+2位日+2位时+2位分+2位秒，不足14位用0补齐。)(必填)
sign  string; // string数字签名（32位的md5加密,加密后转换成小写）
//需要加密的串：
//agent_id=***&bill_id=***&bill_time=***&card_type=***&card_data=***&pay_amt=***&notify_url=***&time_stamp=***|||md5Key

HuiFu string;
};

func (this *Q_HF_Charge)Sign(md5key string){
	str := "agent_id=" + this.agent_id + "&" +
	       "bill_id=" + this.bill_id + "&" + 
	       "bill_time=" + this.bill_time + "&" +
	       "card_type=" + this.card_type + "&" +
	       "card_data=" + this.card_data + "&" +
	       "pay_amt=" + this.pay_amt + "&" +
	       "notify_url=" + this.notify_url + "&" +
	       "time_stamp=" + this.time_stamp + "|||" +
	       md5key;

  log.Println("md5 data: ", str);
  md := md5.New();
  md.Write([]byte(str));       
  this.sign = hex.EncodeToString(md.Sum(nil));
}

type R_HF_Ret struct{
ret_code, // int 返回结,果代码 0=接收成功，其他值参考3.5
ret_msg, // string 返回消息
agent_id,// string 商户ID号
bill_id,// string 商户订单号
jnet_bill_no,// string 成功后在汇元网产生的单据号
bill_status,// int 单据状态：0=未知；1=成功；-1=失败
card_real_amt,// string 收到的卡的实际面值金额
card_settle_amt,// string 卡的结算金额
card_detail_data,//string卡明细信息,格式为 卡号1=金额1,卡号2=金额2,卡号3=金额3
ext_param, //string 商户自定义参数或扩展参数
sign string; // string 数字签名
//ret_code=***&agent_id=***&bill_id=***&jnet_bill_no=***&bill_status=***&card_real_amt=***&card_settle_amt&card_detail_data=***|||md5Key	
}

func (this *R_HF_Ret)Sign(md5key string)(ret string){
	str := "ret_code=" + this.ret_code + "&" +
	       "agent_id=" + this.agent_id + "&" +
	       "bill_id=" + this.bill_id + "&" +
	       "jnet_bill_no=" + this.jnet_bill_no + "&" +
	       "bill_status=" + this.bill_status + "&" +
	       "card_real_amt=" + this.card_real_amt + "&" +
	       "card_settle_amt=" + this.card_settle_amt + "&" +
	       "card_detail_data="+ this.card_detail_data + "|||" +
           md5key;
    md := md5.New();
    md.Write([]byte(str));       
    return hex.EncodeToString(md.Sum(nil));
}

//3des ECB PKCS7
func TDesEncrypt(src, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key);
	if err != nil {
		return nil, err;
	}
	bs := block.BlockSize();
	src = PKCS7Pad(src, bs);

	if len(src)%bs != 0 {
		return nil, errors.New("Need a multiple of the blocksize");
	}
	out := make([]byte, len(src));
	dst := out;
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs]);
		src = src[bs:];
		dst = dst[bs:];
	}
	return out, nil;
}
 
// 3DES加密 CBC pkcs5
func TripleDesEncrypt(origData , key []byte)([]byte,  error){
    block, err := des.NewTripleDESCipher(key);
    if err != nil {
        return nil, err;
    }

    origData = PKCS5Padding(origData, block.BlockSize());
    // origData = ZeroPadding(origData, block.BlockSize())
    blockMode := cipher.NewCBCEncrypter(block, key[:8]);
    crypted := make([]byte, len(origData));
    blockMode.CryptBlocks(crypted, origData);
    return crypted, nil;
}

// 3DES解密 CBC pkcs5
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

func main(){
	var _ = fmt.Sprint("");
	//设置值
	var apply Q_HF_Charge;
	apply.HuiFu = "https://pay.Heepay.com/Api/CardPaySubmitService.aspx?";
	apply.agent_id = "1733154";
	apply.bill_id = "000001"; //6-50位订单号  无此会报:1无效的订单号
    // 生成时间点  报:1无效的时间戳
    tn := time.Now();
    apply.time_stamp = fmt.Sprintf("%d%02d%02d%02d%02d%02d",tn.Year(),tn.Month(),tn.Day(), tn.Hour(), tn.Minute(), tn.Second());
    apply.bill_time = fmt.Sprintf("%d%02d%02d%02d%02d%02d",tn.Year(),tn.Month(),tn.Day(), tn.Hour(), tn.Minute(), tn.Second());
    // 卡数据 报:2无效的卡数据
	key := []byte("B524CA826E5A4036A0F167AE"); 
	crytext, err := TripleDesEncrypt([]byte("1401220375869122,3768830876522604,1"), key);
	if err != nil {
		panic(err);
	}    

    apply.card_data = string(base64.StdEncoding.EncodeToString(crytext));
    apply.card_type = "10";
    apply.pay_amt = "1";

    // 客户端IP
    apply.client_ip = "127.13.11.1";
    // 签名
    apply.Sign("B524CA826E5A4036A0F167AE");

	
	//生成数据对象
	val := url.Values{};
	val.Add("agent_id", apply.agent_id);
	val.Add("bill_id", apply.bill_id);
	val.Add("bill_time", apply.bill_time);
	val.Add("card_type", apply.card_type);
	val.Add("card_data", apply.card_data);
	val.Add("pay_amt", apply.pay_amt);
	val.Add("client_ip", apply.client_ip);
	val.Add("notify_url", apply.notify_url);
	val.Add("time_stamp", apply.time_stamp);
	val.Add("sign", apply.sign);

    srv_url := apply.HuiFu + val.Encode();
    log.Println("srv_url: ", srv_url);
	cli := &http.Client{};
    resp, err := cli.Get(srv_url);
    if err != nil {
    	log.Println(err.Error());
    }

    body, err := ioutil.ReadAll(resp.Body);

    log.Println("res: ", string(body));      

    res, err := url.ParseQuery(string(body)); 
    if err != nil{
        log.Println("err: ", err.Error());
    }

    // gb2312转utf8
    cd, err := iconv.Open("utf-8", "gb2312");
    if err != nil{
    	log.Println("iconv open failed! ", err.Error());
    }
    defer cd.Close();


    log.Println("ret_code=", res.Get("ret_code"));
    log.Println("ret_msg=", cd.ConvString(res.Get("ret_msg")));      
}


 
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
 
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
 
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}
 
func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}
 
//des ECB PKCS7
func DesEncrypt(src, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key);
	if err != nil {
		return nil, err;
	}
	bs := block.BlockSize()
	src = PKCS7Pad(src, bs)
	// src = ZeroPadding(src, bs)
	// src = PKCS5Padding(src, bs)
	if len(src)%bs != 0 {
		return nil, errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return out, nil
}
 
func DesDecrypt(src, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = ZeroUnPadding(out)
	// out = PKCS5UnPadding(out)
	return out, nil
}

func PKCS7Pad(message []byte, blocksize int) (padded []byte) {
	// block size must be bigger or equal 2
	if blocksize < 1<<1 {
		panic("block size is too small (minimum is 2 bytes)")
	}
	// block size up to 255 requires 1 byte padding
	if blocksize < 1<<8 {
		// calculate padding length
		padlen := PadLength(len(message), blocksize)

		// define PKCS7 padding block
		padding := bytes.Repeat([]byte{byte(padlen)}, padlen)

		// apply padding
		padded = append(message, padding...)
		return padded
	}
	// block size bigger or equal 256 is not currently supported
	panic("unsupported block size")
}


func PKCS7Unpad(padded []byte) (message []byte, err error) {
	// read padding length
	plen := len(padded)
	last_byte := padded[plen-1]
	padlen := int(last_byte)

	// check validity of PKCS7 padding
	for i := padlen; i > 1; i-- {
		if padded[plen-i] != last_byte {
			err = errors.New(fmt.Sprintf("Invalid padding (byte -%d: %d). Is the message supplied PKCS7 padded?", i, padded[plen-i]))
			break
		}
	}

	// remove padding
	return padded[:plen-padlen], err
}






/*
对方返回的字符集为gb2312
*/

