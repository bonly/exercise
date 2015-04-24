package main
import (
  "database/sql"
  "fmt"
  _ "github.com/mattn/go-sqlite3"
)

func main(){
  db, err := sql.Open("sqlite3", "./foo.db");
  if err != nil {
     panic(err);
  }
  defer db.Close();

  _, err = db.Exec("create table if not exists userinfo (username varchar(12), departname varchar(12), created varchar(22))");
  if err != nil {
    panic(err);
  }
  
  stmt, err := db.Prepare("replace into userinfo(username, departname, created) values(?,?,?)");
  if err != nil {
    panic(err);
  }

  _, err = stmt.Exec("astaxie", "研发部门", "2012-12-09");
  fmt.Println("end");
}




















