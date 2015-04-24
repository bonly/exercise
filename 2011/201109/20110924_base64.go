package main

import (
	"encoding/base64"
        "io/ioutil"
	"os"
)

func main() {
        input, err := ioutil.ReadFile("/home/bonly/global.jpg");
        if err != nil {
            panic(err);
        }
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	encoder.Write(input)
	// Must close the encoder when finished to flush any partial blocks.
	// If you comment out the following line, the last partial block "r"
	// won't be encoded.
	encoder.Close()
}

