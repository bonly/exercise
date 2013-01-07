package main

import (
    //"fmt"
    //"os"
    "github.com/ziutek/mymysql/mysql"
    _ "github.com/ziutek/mymysql/native" // Native engine
    // _ "github.com/ziutek/mymysql/thrsafe" // Thread safe engine
)

func main() {
    user:="mysql";
    pass:="";
    dbname:="paladin";
    db := mysql.New("tcp", "", "117.135.154.58:3306", user, pass, dbname);

    err := db.Connect();
    if err != nil {
        panic(err);
    }
    defer db.Close();

    _, _, err = db.Query("update player set spirit=190,energy=4 where player_id = %d", 27028);
    if err != nil {
        panic(err);
    }

}
