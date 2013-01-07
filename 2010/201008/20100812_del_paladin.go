package main

import (
    "fmt"
    "github.com/ziutek/mymysql/mysql"
    _ "github.com/ziutek/mymysql/native" // Native engine
)

const user = "bonly";  //全局常量不要用 :=
var pass = "" ;
var dbname = "paladin_test";  //全局变量不要写 :=

func get_id() {
    db := mysql.New("tcp", "", "183.60.126.26:3306", user, pass, dbname);
    db1 := mysql.New("tcp", "", "183.60.126.26:3306", user, pass, dbname);

    err := db.Connect();
    if err != nil {
        panic(err);
    }
    defer db.Close();

    err = db1.Connect();
    defer db1.Close();

    rows, _, err := db.Query("select bag_id,stuff_id from player_bag b, paladin p where player_id = %d and b.stuff_id=p.paladin_id and is_active=0", 26559);
    if err != nil {
        panic(err);
    }

    for _, row := range rows {
        for _, col := range row {
            if col == nil {
                // col has NULL value
            } else {
                // Do something with text in col (type []byte)
            }
        }

        bag_id := row.Int(0);
        stuff_id := row.Int(1);
        fmt.Printf("%d, %d\n", bag_id, stuff_id)

        //del(db, bag_id, stuff_id);
    }
}

func del(db mysql.Conn, bag_id int, stuff_id int){
   _, _, err :=  db.Query("delete from player_bag where bag_id=%d", bag_id);
   if err != nil{
        panic(err);
   }
   _, _, err = db.Query("delete from paladin where paladin_id=%d", stuff_id);
   if err != nil{
        panic(err);
   }
   fmt.Printf("delete %d, %d\n", bag_id, stuff_id);
}

func main(){
   get_id();
}


