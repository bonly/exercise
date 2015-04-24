package main
 
import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/smtp"
    "os"
    "strings"
)
 
type cfgmail struct {
    Username string
    Password string
    Smtphost string
}
 
type cfg struct {
    Name, Text string
}
 
func main() {
 
    //从json文件中读取发送邮件服务器配置信息
    cfgjson := getConf()
 
    var cfg cfgmail
    dec := json.NewDecoder(strings.NewReader(cfgjson))
    for {
 
        if err := dec.Decode(&cfg); err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }
 
        //fmt.Printf("%s\n%s\n%s\n", cfg.Username, cfg.Password, cfg.Smtphost)
 
    }
 
    username := cfg.Username
    password := cfg.Password
    host := cfg.Smtphost
 
    to := "bonly@163.com"
 
    fmt.Printf("============")
    fmt.Println(username)
    subject := "能否收到邮件哟？Test send email by golang"
 
    body := `
    <html>
    <body>
    <h3>
    "Test send email by bonly，来个测试试一下"
    </h3>
    </body>
    </html>
    `
 
    err := SendMail(username, password, host, to, subject, body, "html")
    if err != nil {
        fmt.Println("send mail error!")
        fmt.Println(err)
    } else {
        fmt.Println("send mail success!")
    }
 
}
 
func SendMail(user, password, host, to, subject, body, mailtype string) error {
    hp := strings.Split(host, ":")
    auth := smtp.PlainAuth("", user, password, hp[0])
    var content_type string
    if mailtype == "html" {
        content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
    } else {
        content_type = "Content-Type: text/plain" + "; charset=UTF-8"
    }
 
    msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
    send_to := strings.Split(to, ";")
    err := smtp.SendMail(host, auth, user, send_to, msg)
    return err
}
 
func getConf() string {
    filename := "conf.json"
    file, err := os.Open(filename)
 
    defer file.Close()
    if err != nil {
        fmt.Println("read conf file error")
        log.Fatal(err)
    }
 
    buf := make([]byte, 512)
    var str1 string
    for {
        n, _ := file.Read(buf)
        if 0 == n {
            break
        }
        //os.Stdout.Write(buf[:n])
 
        str := string(buf[:n])
 
        str1 = str1 + str
    }
    return str1
}

/*
{
"Username": "sunxxg@163.com",
"Password": "123456",
"Smtphost":"smtp.163.com:25"
}
*/
