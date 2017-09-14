package main

import (
    "log"
    "net/smtp"
    "flag"
    "text/template"
    "bytes"
    "github.com/jmoiron/sqlx"
    _ "github.com/go-sql-driver/mysql"
    "crypto/md5"
    "encoding/hex"
    "math/rand"
    "time"
)

var db *sqlx.DB;

var tpl = `To: {{ .To }}
From: {{ .From }}>
Subject: {{ .Subject }}
Content-Type: text/html;charset=UTF-8

your new passwd: {{ .Body }}`;

type Mail struct{
	To       string;
	From     string;
	Subject  string;
	Body     string;
};

type User struct{
	ID string `db:"id"`;
	Salt string `db:"salt"`;
	Create_date string `db:"createdate"`;
};


var from_mail   = flag.String("f", "enterprise@xbed.com.cn", "mail from");
var passwd      = flag.String("p", "Xbed2015", "passwd");
var smtp_srv    = flag.String("s", "smtp.exmail.qq.com", "smtp srv addr");
var smtp_port   = flag.String("t", "25", "smtp srv port");
var db_con      = flag.String("d", "db_writer:XqH/a5aOnAlw@tcp(112.74.195.114:3306)/xbed_service?charset=utf8", "db connect str");
var username    = flag.String("n", "hejb", "user name");
var to_mail     = flag.String("m", "hejb@xbed.com.cn", "mail to");
var null_pwd2   = flag.Bool("u", true, "write password2 field as null");
var set_pwd     = flag.String("x", "", "setup password");

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ");

func init(){
	flag.Parse();
	rand.Seed(time.Now().UnixNano());
}

func main(){
	var err error;
	db, err = sqlx.Open("mysql", *db_con);
	if err != nil{
		log.Fatal(err);
	}
	defer db.Close();

	auth := smtp.PlainAuth(
		"",
		*from_mail,
		*passwd,
		*smtp_srv,
	);

	user, err := get_user(*username);
	if err != nil{
		log.Fatal(err);
	}

	pwd := "";
	if len(*set_pwd)<6 {
		pwd = RandStringRunes(6);
	}else{
		pwd = *set_pwd;
	}
	

	m1 := md5.New();
	m1.Write([]byte(user.Salt + pwd + user.Create_date));
	pwd_crypt := hex.EncodeToString(m1.Sum(nil));

	log.Println("passwd: ", pwd_crypt);

	pwd2 := pwd;
	if *null_pwd2 == true {
		pwd2 = "";
	}

	err = update_passwd(pwd_crypt, pwd2, *username);
	if err != nil{
		log.Fatal(err);
	}

	var mail Mail;
	mail.To = *to_mail;
	mail.From = "operation@xbed.com.cn";
	mail.Subject = "Xbed's oms new password";
	mail.Body = pwd;
	
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

func update_passwd(pwd_cry string, pwd string, name string)(err error){
	_, err = db.Exec(`update xbed_service.oms_new_user set password=?, password2=? where username=?`, pwd_cry, pwd, name);
	if err != nil{
		log.Println(err);
		return err;
	}
	return nil;
}


func get_user(name string)(user User, err error){
	rows, err := db.Queryx("select id, salt, createdate from xbed_service.oms_new_user where username=?", name);
	if err != nil{
		log.Println(err);
		return user, err;
	}
	for rows.Next(){
		err := rows.StructScan(&user);
		if err != nil{
			log.Println(err);
			return user, err;
		}
	}
	return user, nil;
}

func RandStringRunes(n int) string {
    b := make([]rune, n);
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))];
    }
    return string(b);
}