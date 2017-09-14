package main 

import (
"crypto/des"
"os"
"fmt"
"bytes"
"crypto/cipher"
// "encoding/hex"
)

func main(){
	if len(os.Args) != 3{
		fmt.Println("args: ", os.Args[0], " key data");
		return;
	}

	data := []byte(os.Args[2]);
	key := []byte(os.Args[1]);
	
	ret, err := DesEncrypt(data, key);
	if err != nil{
		fmt.Println("des err: ", err);
		return;
	}
	fmt.Println(string(ret));
	return;
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
     padding := blockSize - len(ciphertext)%blockSize
     padtext := bytes.Repeat([]byte{byte(padding)}, padding)
     return append(ciphertext, padtext...);
}

func DesEncrypt(origData, key []byte) ([]byte, error) {
     block, err := des.NewCipher(key);
     if err != nil {
          return nil, err;
     }
     origData = PKCS5Padding(origData, block.BlockSize());
     blockMode := cipher.NewCBCEncrypter(block, key);
     crypted := make([]byte, len(origData));
      // 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
     // crypted := origData
     blockMode.CryptBlocks(crypted, origData);
     return crypted, nil;
}
