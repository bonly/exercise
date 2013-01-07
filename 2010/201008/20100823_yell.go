package main

import (
    _ "fmt"
    "github.com/ziutek/mymysql/mysql"
    _ "github.com/ziutek/mymysql/native" // Native engine
    "flag"
)

const user = "mysql";  
var pass = "" ;
var dbname = "paladin"; 

func yell() {
    db := mysql.New("tcp", "", "117.135.154.58:3306", user, pass, dbname);
    //db := mysql.New("tcp", "", "183.60.126.26:3306", user, pass, dbname);

    err := db.Connect();
    if err != nil {
        panic(err);
    }
    defer db.Close();

    _, _, err = db.Query("insert into messages " +
                             "(account_id, name, sex, player_id, " +
                             "  recv_player_id, content, create_at, type) " +
                             "values (0, '系统', 1, -99, 0, '%s', now(), 1)", flag.Arg(0));
    if err != nil {
        panic(err);
    }
}

func main(){
   flag.Parse();
   if flag.NArg() == 1 {
      yell();
   }
}


