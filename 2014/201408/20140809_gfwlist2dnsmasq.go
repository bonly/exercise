package main  

import (
"fmt"
"io/ioutil"
"net/http"
"bufio"
"bytes"
// "os"
"encoding/base64"
"log"
"regexp"
)

func Get_GFW_list()(body []byte, err error){
    resp, err := http.Get("https://raw.githubusercontent.com/gfwlist/gfwlist/master/gfwlist.txt");
    if err != nil {
        log.Println("Get Host err: ", err);
        return nil, err;
    }
    defer resp.Body.Close();

    body, err = ioutil.ReadAll(resp.Body);
    if err != nil{
        log.Println("Body err: ", err);
        return nil, err;
    }

    return body, nil;
}

func main() {
    body, err := Get_GFW_list();
    if err != nil{
        return;
    }

    list, err := base64.StdEncoding.DecodeString(string(body));
    if err != nil{
        log.Printf("decode err: %v", err);
        return;
    }

    // log.Printf("%s", string(list));
    scanner := bufio.NewScanner(bytes.NewReader(list));
    // var begin_write = false;

    re_comm := regexp.MustCompile(`^\!|\[|^@@|^\d+\.\d+\.\d+\.\d+`);

    re_domain := regexp.MustCompile(`([\w\-\_]+\.[\w\.\-\_]+)[\/\*]*`);
    // if err != nil{
    //     log.Println(err);
    //     return;
    // }

    for scanner.Scan(){
        if len(re_comm.FindString(scanner.Text())) > 0{
            // log.Println("注释行:");
            // fmt.Println(scanner.Text());
        }else{
            domain := re_domain.FindAllString(scanner.Text(), -1);
            // fmt.Println(domain);
            if len(domain) > 0{
                fmt.Printf("server=/.%s/%s#%s\n", domain[0], "127.0.0.1", "1053");
            }
        }
    //     if scanner.Text() == "# Modified hosts start"{
    //         begin_write = true;
    //         fl.Write([]byte("127.0.0.1\tlocalhost\r\n"));
    //         fl.Write([]byte("::1\tlocalhost\r\n"));
    //     }
    //     if begin_write {
    //         fl.Write([]byte(scanner.Text()));
    //         fl.Write([]byte("\r\n"));
    //     }
    }
    // if err := scanner.Err(); err != nil{
    //     fmt.Println(err);
    // }    
}