package main

import (
    "log"
    "net/smtp"
    "flag"
    "text/template"
    "bytes"
)

var tpl = `To: {{ .To }}
From: {{ .From }}>
Subject: {{ .Subject }}
Content-Type: text/html;charset=UTF-8

{{ .Body }}`;

type Mail struct{
	To       string;
	From     string;
	Subject  string;
	Body     string;
};

var from_mail   = flag.String("f", "hejb@xbed.com.cn", "mail from");
var passwd      = flag.String("p", "Xbed116", "passwd");
var smtp_srv    = flag.String("s", "smtp.exmail.qq.com", "smtp srv addr");
var smtp_port   = flag.String("t", "25", "smtp srv port");

func init(){
	flag.Parse();
}

func main(){
	auth := smtp.PlainAuth(
		"",
		*from_mail,
		*passwd,
		*smtp_srv,
	);

	var mail Mail;
	mail.To = "hejb@xbed.com.cn";
	mail.From = "abc@xbed.com.cn";
	mail.Subject = "new passwd";
	mail.Body = "测试";
	
	tmpl, err := template.New("mail").Parse(tpl);
	if err != nil {
		log.Fatal(err);
	}

	var buf bytes.Buffer;
	err = tmpl.Execute(&buf, mail);
	if err != nil{
		log.Fatal(err);
	}
	log.Println(buf.String());

	err = smtp.SendMail(
		*smtp_srv + ":" + *smtp_port,
		auth,
		*from_mail,
		[]string{mail.To},
		buf.Bytes(),
	);
	if err != nil{
		log.Fatal(err);
	}
}