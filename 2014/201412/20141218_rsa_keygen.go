package main

import (
    // "crypto"
    "crypto/rand"
    "crypto/rsa"
    // "crypto/sha256"
    "fmt"
"crypto/x509"
"encoding/pem"      
    "os"
    "bytes"
)

func main() {
    // Gen_key();
    Gen();
}


func Gen_key()(err error){
   private_key, err := rsa.GenerateKey(rand.Reader, 2048);

    if err != nil {
        return err;
    }

    derStream := x509.MarshalPKCS1PrivateKey(private_key);

    block := &pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: derStream,
    }
    file, err := os.Create("private.pem")
    if err != nil {
        return err
    }
    err = pem.Encode(file, block)
    if err != nil {
        return err
    }

    // 生成公钥文件
    public_key := &private_key.PublicKey;
    derPkix, err := x509.MarshalPKIXPublicKey(public_key);
    if err != nil {
        return err
    }
    block = &pem.Block{
        Type:  "PUBLIC KEY",
        Bytes: derPkix,
    }
    file, err = os.Create("public.pem")
    if err != nil {
        return err
    }
    err = pem.Encode(file, block)
    if err != nil {
        return err
    }    
    return nil;
}

func Gen()(err error){
   private_key, err := rsa.GenerateKey(rand.Reader, 2048);

    if err != nil {
        return err;
    }

    derStream := x509.MarshalPKCS1PrivateKey(private_key);

    block := &pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: derStream,
    }
    
    buf := new(bytes.Buffer);
    err = pem.Encode(buf, block)
    if err != nil {
        return err
    }
    fmt.Println(buf.String());


    // 生成公钥文件
    public_key := &private_key.PublicKey;
    derPkix, err := x509.MarshalPKIXPublicKey(public_key);
    if err != nil {
        return err
    }
    block = &pem.Block{
        Type:  "PUBLIC KEY",
        Bytes: derPkix,
    }
    
    buff := new(bytes.Buffer);
    err = pem.Encode(buff, block)
    if err != nil {
        return err
    }    
    fmt.Println(buff.String());
    return nil;
}