package main 

import (
"fmt"
"net/http"
"io/ioutil"
"regexp"
)

func GetLocalPublicIp() (string, error) {  
    // `nc ns1.dnspod.cn 6666`  
    res, err := http.Get("http://iframe.ip138.com/ic.asp")  
    if err != nil {  
        return "", err  
    }  
    defer res.Body.Close()  
    result, err := ioutil.ReadAll(res.Body)  
    if err != nil {  
        return "", err  
    }  
    reg := regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`)  
    return reg.FindString(string(result)), nil  
}  

func main(){
	str, _ := GetLocalPublicIp();
	fmt.Printf("%s", str);
}