package main

import (
    //"fmt"
    //"os"
    "github.com/ziutek/mymysql/mysql"
    _ "github.com/ziutek/mymysql/native" // Native engine
    // _ "github.com/ziutek/mymysql/thrsafe" // Thread safe engine
)

func main() {
    user:="bonly";
    pass:="";
    dbname:="paladin_test";
    db := mysql.New("tcp", "", "183.60.126.26:3306", user, pass, dbname);

    err := db.Connect();
    if err != nil {
        panic(err);
    }
    defer db.Close();

    //rows, res, err := db.Query("select player_id from player where player_id > %d", 20);
    _, _, err = db.Query("update paladin_test.player set spirit=100,energy=4 where player_id = %d", 26585);
    if err != nil {
        panic(err);
    }

}
