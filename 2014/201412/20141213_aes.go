package main  

import(
"fmt"
"flag"
"crypto/aes"
"crypto/cipher"
"crypto/rand"
"encoding/hex"
"errors"
"io"

)

var name = flag.String("n", "bonly", "user name");
var pwd  = flag.String("p", "xbed111", "password");
var key  = flag.String("k", "123456", "key");

func main(){
	encodeBytes, err := Encrypt([]byte("weloveyouchinese"));
	if err != nil{
		fmt.Println("Encrypt: ", err);
		return;
	}

	fmt.Printf("encrypt code: %x\n", string(encodeBytes));
}

// AES加密
func Encrypt(src []byte) ([]byte, error) {
	var block cipher.Block;

    // 验证输入参数
    // 必须为aes.Blocksize的倍数
    if len(src)%aes.BlockSize != 0 {
        return nil, errors.New("crypto/cipher: input not full blocks");
    }
 
    encryptText := make([]byte, aes.BlockSize+len(src));
 
    iv := encryptText[:aes.BlockSize];
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return nil, err;
    }
 
    mode := cipher.NewCBCEncrypter(block, iv);
 
    mode.CryptBlocks(encryptText[aes.BlockSize:], src);
 
    return encryptText, nil;
}
 
// AES解密
func Decrypt(src []byte) ([]byte, error) {
	var block cipher.Block;

    // hex
    decryptText, err := hex.DecodeString(fmt.Sprintf("%x", string(src)));
    if err != nil {
        return nil, err;
    }
 
    // 长度不能小于aes.Blocksize
    if len(decryptText) < aes.BlockSize {
        return nil, errors.New("crypto/cipher: ciphertext too short");
    }
 
    iv := decryptText[:aes.BlockSize];
    decryptText = decryptText[aes.BlockSize:];
 
    // 验证输入参数
    // 必须为aes.Blocksize的倍数
    if len(decryptText)%aes.BlockSize != 0 {
        return nil, errors.New("crypto/cipher: ciphertext is not a multiple of the block size");
    }
 
    mode := cipher.NewCBCDecrypter(block, iv);
 
    mode.CryptBlocks(decryptText, decryptText);
 
    return decryptText, nil;
}