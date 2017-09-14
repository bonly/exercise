package main  

import (
"fmt"
// "io/ioutil"
// "net/http"
"bufio"
// "bytes"
"os"
// "encoding/base64"
// "log"
"regexp"
"strings"
"flag"
"golang.org/x/crypto/ssh"
)

var Pac = flag.String("p", "/tmp/OmegaProfile.pac", "read pac file from");
var Dns = flag.String("d", "/tmp/dnsmasq.conf", "write dnsmasq.conf file to");

func main() {
    flag.Parse();

    infile, err := os.OpenFile(*Pac, os.O_RDONLY, 0666);
    if err != nil{
        fmt.Printf("打开PAC[%s]失败: %s\n", Pac, err);
        return;
    }
    defer infile.Close();

    //建ssh 到路由器上
    config := &ssh.ClientConfig{ //配置
        User: "root",
        Auth: []ssh.AuthMethod{
            ssh.Password("hayes"),
        },
    };
    client, err := ssh.Dial("tcp", "192.168.1.1:22", config); //连接
    if err != nil {
        panic("Failed to dial: " + err.Error());
    };
    session, err := client.NewSession();//创建会话
    if err != nil {
        panic("Failed to create session: " + err.Error());
    }
    defer session.Close();
    of, _ := session.StdinPipe();
    defer of.Close();

    go func() {
        scanner := bufio.NewScanner(infile);

        re_redirct   := regexp.MustCompile(`if \((.*)\) return \"\+GoAgent\";`);
        re_domain    := regexp.MustCompile(`.*url.indexOf\(\"(.*)"\).*`);
        re_del_ext   := regexp.MustCompile(`\/\(\?\:\^\|\\\.\)|\\|\$|\/.*`);

        fmt.Fprintln(of, "C0644", 0, "/tmp/testfile1"); //创建文件
        fmt.Fprint(of, "# auth: bonly \n");
        domain_list := make([]string, 0);
        for scanner.Scan(){
            text := scanner.Text();
            if re_redirct.MatchString(text){
                text_Redirct := re_redirct.ReplaceAllString(text, "$1"); //换成只剩$1
                text_domain := re_domain.ReplaceAllString(text_Redirct, "$1"); //换成只剩indexOf($1)中的变量
                text_host := re_del_ext.ReplaceAllString(text_domain, "");
                domainSpec := strings.TrimSpace(text_host);
                if len(domainSpec) > 0{
                    // fmt.Println(domainSpec);
                    xd := -1;
                    for idx, val := range domain_list{
                        if domainSpec == val{ //found
                            xd = idx;
                            break;
                        }
                    }
                    if xd == -1 { //没找到的插入数据
                        fmt.Fprint(of, fmt.Sprintf("server=/%s/%s#%s\n", domainSpec, "127.0.0.1", "1053"));
                        fmt.Fprint(of, fmt.Sprintf("ipset=/%s/gfwlist\n", domainSpec));
                        domain_list = append(domain_list, domainSpec);
                    }
                }
            }
        }
        fmt.Fprint(of, "\x00"); //结束 scp的输入
        if err := scanner.Err(); err != nil{
            fmt.Println(err);
        }    
    }();
    if err := session.Run("/usr/bin/scp -tr ./"); err != nil { //开终端接收信息处理
        panic("Failed to run: " + err.Error())
    }
}