package main

import (
	_ "github.com/alexbrainman/odbc"
	"database/sql"
        "log"
        "fmt"
)

func main() {
  // Replace the DSN value with the name of your ODBC data source.
	db, err := sql.Open("odbc","DSN=FromAccess")
	if err != nil {
		log.Fatal("open: ",err)
	}

	var (
		id string
		name string
	)

  // This is a SQL Server AdventureWorks database query.
  _, err = db.Exec("create table iprom ( akey int, aname varchar(20))");
  if err != nil{
     fmt.Println("create: ", err);
  }
  
  _, err = db.Exec("insert into iprom (akey, aname) values('1', 'bonly')");
  if err != nil{
     fmt.Println("insert: ", err);
  }
	rows, err := db.Query("select akey, aname from iprom ");
//	rows, err := db.Query("select NAME from t_guest where id = 1186", 1)

	if err != nil {
		log.Fatal("qry: ", err)
	}
	defer rows.Close()


	fmt.Println("ok")
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal("scan: ", err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal("row: ", err)
	}

	_, err = db.Exec("drop table iprom");
	if err != nil{
	  fmt.Println("drop: ", err);
	}
	defer db.Close()
}