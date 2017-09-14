package main

import (
    "log"
    "net/smtp"
    "flag"
    tp "text/template"
    "html/template" 
    "bytes"
    "github.com/jmoiron/sqlx"
    _ "github.com/go-sql-driver/mysql"
    "crypto/md5"
    "encoding/hex"
    "math/rand"
    "time"
    "net/http"
    "fmt"
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
	User_name string `db:"username"`;
	Create_date string `db:"createdate"`;
};


var from_mail   = flag.String("f", "enterprise@xbed.com.cn", "mail from");
var passwd      = flag.String("p", "Xbed2015", "passwd");
var smtp_srv    = flag.String("s", "smtp.exmail.qq.com", "smtp srv addr");
var smtp_port   = flag.String("t", "25", "smtp srv port");
var db_con      = flag.String("d", "db_writer:XqH/a5aOnAlw@tcp(112.74.195.114:3306)/xbed_service?charset=utf8", "db connect str");
var username    = flag.String("n", "hejb", "user name");
var to_mail     = flag.String("m", "hejb@xbed.com.cn", "mail to");

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ");

func init(){
	flag.Parse();
	rand.Seed(time.Now().UnixNano());
}

func chg_pwd(){
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

	pwd  := RandStringRunes(6);

	m1 := md5.New();
	m1.Write([]byte(user.Salt + pwd + user.Create_date));
	pwd_crypt := hex.EncodeToString(m1.Sum(nil));

	log.Println("passwd: ", pwd_crypt);

	err = update_passwd(pwd_crypt, pwd, *username);
	if err != nil{
		log.Fatal(err);
	}

	var mail Mail;
	mail.To = *to_mail;
	mail.From = "operation@xbed.com.cn";
	mail.Subject = "oms new password";
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

func list_user()(users []User, err error){
	rows, err := db.Queryx("select id, username, salt, createdate from xbed_service.oms_new_user");
	if err != nil{
		log.Println(err);
		return users, err;
	}
	for rows.Next(){
		var user User;
		err := rows.StructScan(&user);
		if err != nil{
			log.Println(err);
			return users, err;
		}
		users = append(users, user);
	}
	return users, nil;
}

func RandStringRunes(n int) string {
    b := make([]rune, n);
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))];
    }
    return string(b);
}

func index(wr http.ResponseWriter, req *http.Request){
	var err error;
	// users, err := list_user();
	// if err != nil{
	// 	fmt.Println(err);
	// 	return;
	// }
	users := []User{
		{"1","salt","username","create_date"},
		{"2", "s2", "un", "20161102"},
	};
	
	fmt.Println(users);
	// buf := new(bytes.Buffer);
	var buf bytes.Buffer;
	rows := template.Must(template.ParseFiles("tab_rows.html"));
	err = rows.ExecuteTemplate(&buf, "tab_rows.html", users);
	if err != nil{
		fmt.Println(err);
		return;
	}

	tpl, err := tp.ParseFiles("index.html", "body.html");
	if err != nil{
		fmt.Println(err);
		return;
	}

	err = tpl.ExecuteTemplate(wr, "index.html", map[string]string{"Rows":buf.String()});
	if err != nil{
		fmt.Println(err);
		return;
	}
	
	fmt.Println(buf.String());
}


func main(){
	var err error;
	// db, err = sqlx.Open("mysql", *db_con);
	// if err != nil{
	// 	log.Fatal(err);
	// }
	// defer db.Close();	

	http.HandleFunc("/", index);

	err = http.ListenAndServe("0.0.0.0:9997", nil);
	if err != nil{
		fmt.Println(err);
	}
}