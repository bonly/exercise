package db

import (
	log "golang.org/x/glog"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"fmt"
)

var db *sqlx.DB;

func Init(){
	var err error;
	db, err = sqlx.Open("mysql", "root:techappen@tcp(192.168.1.104:3306)/techappen?charset=utf8");
	if err != nil{
		log.Fatalf("db connect: %v\n", err);
	}	
}

func User_Add(name string, passwd string)(ret string, err error){
	qry := fmt.Sprintf("Insert into techappen.user(user_name, user_passwd, create_time) values(?, ?, ?)");

	tn := time.Now();
	nw := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", tn.Year(),tn.Month(),tn.Day(),tn.Hour(),tn.Minute(),tn.Second());

	res, err := db.Exec(qry, name, passwd, nw);
	if err != nil{
		log.Error("ins user failed: %v", err);
		return;
	}
	lsid, _ := res.LastInsertId();
	ret = fmt.Sprintf("%d", lsid);

	return ret, nil;
}

func User_Get(name string, passwd string)(ret string, err error){
	qry := fmt.Sprintf("Select user_id from techappen.user where user_nameji = ? and user_passwd = ?");

	res, err := db.Query(qry, name, passwd);
	if err != nil{
		log.Error("get user failed: ", err);
		return "", err;
	}
	for res.Next() {
        err := res.Scan(&ret);
        if err != nil {
            log.Error(err);
			return "", err;
        } 
    }

	return ret, nil;
}