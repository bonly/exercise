package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	// "encoding/hex"
	"net/url"
	"flag"
	"fmt"
)

var (
	hexKey = flag.String("key", "", "the key")
	src    = flag.String("src", "", "the content")
)

func main() {
	flag.Parse()

	var hasError = false

	if hexKey == nil || src == nil {
		fmt.Println("args error")
		hasError = true
	}

	// key, err := hex.DecodeString(*hexKey)
	// if err != nil {
	// 	fmt.Println("key error, please input hex type aes key")
	// 	hasError = true
	// }

	key := []byte(*hexKey);

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error: ", err);
		hasError = true
	}

	if *src == "" {
		fmt.Println("plain content empty")
		hasError = true
	}

	if hasError {
		flag.Usage()
		return
	}

	ecb := NewECBEncrypter(block)
	content := []byte(*src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	//fmt.Println(base64.StdEncoding.EncodeToString(crypted))
	fmt.Println(url.QueryEscape(base64.StdEncoding.EncodeToString(crypted)));
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

func (x *ecbDecrypter) BlockSize() int { return x.blockSize }

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
