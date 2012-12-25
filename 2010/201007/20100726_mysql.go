package main

import (
    "fmt"
    "os"
    "github.com/ziutek/mymysql/mysql"
    _ "github.com/ziutek/mymysql/native" // Native engine
    // _ "github.com/ziutek/mymysql/thrsafe" // Thread safe engine
)

func main() {
    user:="bonly";
    pass:="";
    dbname:="paladin_test"
    db := mysql.New("tcp", "", "183.60.126.26:3306", user, pass, dbname)

    err := db.Connect()
    if err != nil {
        panic(err)
    }

    //rows, res, err := db.Query("select player_id from player where player_id > %d", 20)
    rows, _, err := db.Query("select player_id,name from player where player_id > %d", 20)
    if err != nil {
        panic(err)
    }

    for _, row := range rows {
        for _, col := range row {
            if col == nil {
                // col has NULL value
            } else {
                // Do something with text in col (type []byte)
            }
        }
        // You can get specific value from a row
        val1 := row[1].([]byte)

        // You can use it directly if conversion isn't needed
        os.Stdout.Write(val1)

        // You can get converted value
        number := row.Int(0)      // Zero value
        str    := row.Str(1)      // First value
        //bignum := row.MustUint(2) // Second value

        // You may get values by column name
        //first := res.Map("FirstColumn")
        //second := res.Map("SecondColumn")
        //val1, val2 := row.Int(first), row.Str(second)
        fmt.Printf("%d, %s\n", number, str)
    }
}
