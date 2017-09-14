package main

import (
    "log"
    "net/smtp"
)

func main() {
    // Set up authentication information.
    // auth := smtp.PlainAuth(
    //     "",
    //     "user@example.com",
    //     "password",
    //     "mail.example.com",
    // )
    auth := smtp.PlainAuth(
        "",
        "hejb@xbed.com.cn",
        "Xbed116",
        "smtp.exmail.qq.com",
    )    
    // Connect to the server, authenticate, set the sender and recipient,
    // and send the email all in one step.
    // err := smtp.SendMail(
    //     "mail.example.com:25",
    //     auth,
    //     "sender@example.org",
    //     []string{"recipient@example.net"},
    //     []byte("This is the email body."),
    // )

    msg := []byte("To: " + "hejb@xbed.com.cn" + "\r\nFrom: " + "hejb@xbed.com.cn" + ">\r\nSubject: " + "标题" + "\r\n"  + "\r\n\r\n" + "this is body")
    err := smtp.SendMail(
        "smtp.exmail.qq.com:25",
        auth,
        "hejb@xbed.com.cn",
        []string{"hejb@xbed.com.cn"},
        msg,
    )    
    if err != nil {
        log.Fatal(err)
    }
}