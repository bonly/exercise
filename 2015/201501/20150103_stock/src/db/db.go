package db

import (
log "glog"
_ "github.com/mattn/go-sqlite3"
"github.com/jmoiron/sqlx"
"flag"
. "stock"
)

var db *sqlx.DB;

func init(){
	flag.Parse();
	flag.Set("alsologtostderr", "true");
	open();
}


func open(){
	var err error;
	// db, err = sqlx.Open("sqlite3", ":memory:");
	db, err = sqlx.Open("sqlite3", "history");
	if err != nil{
		log.Info(err);
	}
	
	err = db.Ping();
	if err != nil{
		log.Info(err);
	}
}

func create_table(){
	_, err := db.Exec(`
		create table if not exists history(
			PK INTEGER PRIMARY KEY,
			Code varchar(22),
			Dt   int(22),
			Open float(22),
			High float(22),
			Low  float(22),
			Close float(22),
			Volume float(22)
		)`);
	if err != nil{
		log.Info(err);
	}
}


func select_table()(ret []Day, err error){
	his := Day{};
	// rows, err := db.Queryx("select PK,Code,Dt,Open,High,Low,Close,Volume,Adj_Close from history  limit 5");
	rows, err := db.Queryx("select * from history");
	if err != nil{
		log.Info(err);
		return ret, err;
	}
	for rows.Next(){
		// log.Infof("%#v\n", rows);
		err := rows.StructScan(&his);
		if err != nil{
			log.Info(err);
		}
		// log.Infof("%#v\n", his);
		ret = append(ret, his);
	}
	return ret, nil;
}