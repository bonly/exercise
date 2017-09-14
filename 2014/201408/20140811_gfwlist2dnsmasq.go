package main  

import (
"fmt"
"io/ioutil"
"net/http"
"bufio"
"bytes"
"os"
"encoding/base64"
"log"
"regexp"
"strings"
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
    fl, err := os.OpenFile("/tmp/dnsmasq.conf", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666);
    if err != nil{
        fmt.Println("打开文件失败:",err);
        return;
    }
    defer fl.Close();

    body, err := Get_GFW_list();
    if err != nil{
        return;
    }

    list, err := base64.StdEncoding.DecodeString(string(body));
    if err != nil{
        log.Printf("decode err: %v", err);
        return;
    }

    fmt.Printf("%s\n", string(list));
    scanner := bufio.NewScanner(bytes.NewReader(list));

    re_comm := regexp.MustCompile(`^\!|\[|^@@|^\d+\.\d+\.\d+\.\d+`);
    // re_comm := regexp.MustCompile("(^!|^@@|^\\[|!--)[\\s\\w-:,+\\.\\]\\*#\\(\\)\\|/@]*");
    re_domain := regexp.MustCompile(`([\w\-\_]+\.[\w\.\-\_]+)[\/\*]*`);
    // re_domain := regexp.MustCompile("^\\|\\||^http(s?)://|^\\|http(s?)://|^\\.|^\\w+\\*\\.|(\\.?)[/*]\\S+|^%\\S+|/|^\\w+$");

    fl.Write([]byte("# auth: bonly \n"));
    domain_list := make([]string, 0);
    for scanner.Scan(){
        if len(re_comm.FindString(scanner.Text())) > 0{
            // log.Println("注释行:");
            // fmt.Println(scanner.Text());
        }else{
            domain := re_domain.FindAllString(scanner.Text(), -1);
            // fmt.Println(domain);            
            if len(domain) > 0{
                cur_domain := strings.Replace(domain[0], "*", "", -1);//去掉后面多余的*
                cur_domain = strings.Replace(cur_domain, "/", "", -1);//去掉后面多余的/
                xd := -1;
                for idx, val := range domain_list{
                    // fmt.Printf("%s [%s]\n", cur_domain, val);
                    if cur_domain == val{ //found
                        xd = idx;
                        break;
                    }
                }
                if xd == -1 { //没找到的插入数据
                    fl.Write([]byte(fmt.Sprintf("server=/.%s/%s#%s\n", cur_domain, "127.0.0.1", "1053")));
                    fl.Write([]byte(fmt.Sprintf("ipset=/.%s/gfwlist\n", cur_domain)));
                    domain_list = append(domain_list, cur_domain);
                } 
            }
        }
    }
    if err := scanner.Err(); err != nil{
        fmt.Println(err);
    }    
}