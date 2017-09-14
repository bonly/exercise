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
"os/exec"
"golang.org/x/crypto/ssh"
)

var Pac = flag.String("p", "/home/opt/Downloads/OmegaProfile_Go_Auto.pac", "read pac file from");
var Dns = flag.String("d", "/tmp/dnsmasq.conf", "write dnsmasq.conf file to");
var Gfw = flag.String("g", "/overlay/openwrt/etc/dnsmasq.d/gfwlist.conf", "dnsmasq file in router");
var RPasswd = flag.String("r", "", "Password of router");
var Port = flag.String("P", "23", "ssh port");

func main() {
    flag.Parse();

    infile, err := os.OpenFile(*Pac, os.O_RDONLY, 0666);
    if err != nil{
        fmt.Printf("打开PAC[%s]失败: %v\n", *Pac, err);
        return;
    }
    defer infile.Close();

    outfile, err := os.OpenFile(*Dns, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666);
    if err != nil{
        fmt.Println("打开文件[%s]失败: %s\n", *Dns, err);
        return;
    }
    defer outfile.Close();

    scanner := bufio.NewScanner(infile);

    re_redirct   := regexp.MustCompile(`if \((.*)\) return \"\+GoAgent\";`);
    re_domain    := regexp.MustCompile(`.*url.indexOf\(\"(.*)"\).*`);
    re_del_ext   := regexp.MustCompile(`\/\(\?\:\^\|\\\.\)|\\|\$|\/.*`);

    outfile.Write([]byte("# auth: bonly \n"));
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
                    outfile.Write([]byte(fmt.Sprintf("server=/%s/%s#%s\n", domainSpec, "127.0.0.1", "1053")));
                    outfile.Write([]byte(fmt.Sprintf("ipset=/%s/gfwlist\n", domainSpec)));
                    domain_list = append(domain_list, domainSpec);
                }
            }
        }
    }
    if err := scanner.Err(); err != nil{
        fmt.Println(err);
        return;
    }

    cmd_argv := []string{"-P", *Port, *Dns, "root@192.168.1.1:" + *Gfw};
    cmd := exec.Command("scp", cmd_argv...);
    out, err := cmd.Output();
    if err != nil{
        fmt.Println("scp: ", err);
        return;
    }
    fmt.Println("scp: ", out);

    if len(*RPasswd) > 0{
        reboot();
    }
}

/*
ssh root@192.168.1.1 "tee -a /etc/dropbear/authorized_keys" < ~/.ssh/id_rsa.pub
*/

func reboot(){
    config := &ssh.ClientConfig{
        User: "root",
        Auth: []ssh.AuthMethod{
            ssh.Password(*RPasswd),
        },
    };
    client, err := ssh.Dial("tcp", "192.168.1.1:" + *Port, config);
    if err != nil {
        panic("Failed to dial: " + err.Error());
    }    
    session, err := client.NewSession();
    if err != nil {
        panic("Failed to create session: " + err.Error());
    }
    defer session.Close();

    if err := session.Run("reboot"); err != nil{
        panic("filed to reboot: " + err.Error());
    }
}
