/*
package main;

import (
    "fmt"
    "database/sql"
    _"github.com/weigj/odbc/driver"
)

func main(){
    // conn,err := sql.Open("odbc","driver={Microsoft Access Driver (*.mdb)};dbq=/home/bonly/hotel/Hotel.mdb");
    // conn,err := sql.Open("odbc","odbc:Driver={MDBTools};DBQ=/home/bonly/hotel/Hotel.mdb;");
    conn,err := sql.Open("odbc","DSN=FromAccess");
    if(err!=nil){
        fmt.Println("Connecting Error", err);
        return;
    }
    defer conn.Close();
    stmt,err := conn.Prepare("select * from t_room");
    if(err!=nil){
        fmt.Println("Query Error: ", err);
        return;
    }
    defer stmt.Close();
    row,err := stmt.Query();
    if err!=nil {
        fmt.Println("Query Error: ", err);
        return;
    }
    defer row.Close();
    for row.Next() {
        var id int;
        var name string;
        if err := row.Scan(&id,&name);err==nil {
            fmt.Println(id,name);
        }
    }
    fmt.Printf("%s\n","finish");
    return;
}
*/
package main

import (
    _ "github.com/alexbrainman/odbc"
    "database/sql"
    "log"
    "fmt"
)

func main() {
    // Replace the DSN value with the name of your ODBC data source.
    // db, err := sql.Open("odbc","DSN=FromAccess;Exclusive=1;Uid=admin;Pwd=;");
    db, err := sql.Open("odbc","DSN=FromAccess");
    if err != nil {
        log.Fatal("open: ", err);
    }

    var (
        id string;
        guestid string;
        name string;
    );

    _, err = db.Exec("create table iprom(akey int, aname varchar(20))");
    if err != nil{
        fmt.Println("create: ", err);
    }
    
    rows, err := db.Query("select id,guestid,name from t_guest");
    if err != nil {
        log.Fatal("qry: ", err);
    }
    defer rows.Close()
    for rows.Next() {
        err := rows.Scan(&id, &guestid, &name)
        if err != nil {
            log.Fatal(err)
        }
        log.Println(id, guestid, name)
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()

   // stmt,err := db.Prepare("select * from t_room");
   //  if(err!=nil){
   //      fmt.Println("Query Error ", err);
   //      return;
   //  }
   //  defer stmt.Close();
   //  row,err := stmt.Query();
   //  if err!=nil {
   //      fmt.Println("Query Error");
   //      return;
   //  }
   //  defer row.Close();
   //  for row.Next() {
   //      var id int;
   //      var name string;
   //      if err := row.Scan(&id,&name);err==nil {
   //          fmt.Println(id,name);
   //      }
   //  }
    fmt.Printf("%s\n","finish");
    return;    
}
