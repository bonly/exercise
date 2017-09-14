package main 

import (
"net/http"
"log"
"io/ioutil"
"encoding/base64"
"encoding/json"
"crypto/des"
"bytes"
"errors"
"fmt"
)

var srv_addr = "http://120.25.160.52";
var srv_key = "nXV3Xbhx";

type Data struct {
	Name string `json:"name"`;
	Idnum string `json:"idnum"`;
	Pic string `json:"pic"`;
};

type Ret struct{
	Code int `json:"code"`;
	Msg interface{} `json:"msg"`;
};

type R_residuenum struct{
	Num int `json:"num"`;
};

func residue_num(){
	resp, err := http.Get(srv_addr + "/joggle/identity/residuenum");
	if err != nil{
		log.Printf("http get: %v\n", err);
		return;
	}
	defer resp.Body.Close();

	body, err := ioutil.ReadAll(resp.Body);
	if err != nil{
		log.Printf("body: %v\n", err);
		return;
	} 

	log.Printf("recv org: %s", string(body));
	
	buf := make([]byte, 512);
	lng, err := base64.StdEncoding.Decode(buf, body); //先base64解码
	if err != nil{
		log.Printf("base64 decode: %v\n", err);
		return;
	}
	// log.Printf("len: %d %d\n", lng, len(string(buf[:lng])));

	ret, err := DesDecrypt([]byte(string(buf[:lng])), []byte(srv_key)); //再des解密
	if err != nil{
		log.Printf("des decode: %v\n", err);
		return;
	}
	log.Printf("recv: %+v", string(ret));

	var ret_data Ret;
	var ret_msg R_residuenum;
	ret_data.Msg = &ret_msg;
	// if err = json.Unmarshal([]byte(`{"code":0,"msg":{"num":5}}`), &ret_data); err != nil{
	if err = json.Unmarshal(ret, &ret_data); err != nil{
		log.Printf("json data decode: %v\n", err);
		return;
	}

	log.Println(fmt.Sprintf("code=%d num=%+v\n", ret_data.Code, ret_msg.Num));
}

func main(){
	residue_num();
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize;
	padtext := bytes.Repeat([]byte{byte(padding)}, padding);
	return append(ciphertext, padtext...);
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData);
	unpadding := int(origData[length-1]);
	return origData[:(length - unpadding)];
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize;
	padtext := bytes.Repeat([]byte{0}, padding);
	return append(ciphertext, padtext...);
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0);
		})
}

func DesEncrypt(src, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key);
	if err != nil {
		return nil, err;
	}
	bs := block.BlockSize();
	// src = ZeroPadding(src, bs);
	src = PKCS5Padding(src, bs);
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
func DesDecrypt(src, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key);
	if err != nil {
		return nil, err;
	}
	out := make([]byte, len(src));
	dst := out;
	bs := block.BlockSize();
	if len(src)%bs != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks");
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs]);
		src = src[bs:];
		dst = dst[bs:];
	}
	// out = ZeroUnPadding(out);
	out = PKCS5UnPadding(out);
	return out, nil;
}