package main

import (
        "fmt"
        "crypto/des"
        "encoding/hex"
)

func main() {
        key, _ := hex.DecodeString("0000000000000000")
        block, _ := des.NewCipher(key)
        data := make([]byte, 8)
        block.Encrypt(data, key)
        fmt.Println(hex.EncodeToString(data))
}
