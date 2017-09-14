package main

import (
	_ "github.com/alexbrainman/odbc"
	"database/sql"
        "log"
        "fmt"
)

func main() {
  // Replace the DSN value with the name of your ODBC data source.
	db, err := sql.Open("odbc","DSN=testdb");
	if err != nil {
		log.Fatal("open: ",err);
	}
	defer db.Close();

	var (
		id string;
		akey string;
		atext string;
		ablob []byte;
	);

  // _, err = db.Exec("create table iprom ( akey int, aname varchar(20))");
  // if err != nil{
     // fmt.Println("create: ", err);
  // }

	var blob = []byte{0x31, 0x32, 0x43};
	_, err = db.Exec("insert into mytbl(id, akey, atext, ablob) values('1', '12', 'test', ?)", blob);
		if err != nil{
		fmt.Println("insert: ", err);
	}
	rows, err := db.Query("select id, akey, atext, ablob from mytbl ");

	if err != nil {
		log.Fatal("qry: ", err);
	}
	defer rows.Close();


	for rows.Next() {
		err := rows.Scan(&id, &akey, &atext, &ablob);
		if err != nil {
			log.Fatal("scan: ", err);
		}
		log.Printf("%s %s %s %x len: %d\n", id, akey, atext, ablob, len(ablob));
	}
	err = rows.Err();
	if err != nil {
		log.Fatal("row: ", err);
	}

	_, err = db.Exec("delete from mytbl where id=1");
	if err != nil{
		fmt.Println("del: ", err);
	}
	// _, err = db.Exec("drop table iprom");
	// if err != nil{
	  // fmt.Println("drop: ", err);
	// }
}