package main

import (
    _ "github.com/Go-SQL-Driver/MySQL"
    "database/sql"
    "fmt"
    //"time"
)

func main() {
    db, err := sql.Open("mysql", "bonly:1234@tcp(183.60.126.26:3306)/paladin?charset=utf8")
    checkErr(err)

    //插入数据
    stmt, err := db.Prepare("INSERT copy SET copy_id=?,level=?,sub_id=?")
    checkErr(err)

    res, err := stmt.Exec("110", "10", "113")
    checkErr(err)

    id, err := res.LastInsertId()
    checkErr(err)

    fmt.Println(id)

    //更新数据
    stmt, err = db.Prepare("update copy set copy_id=? where copy_id=?")
    checkErr(err)

    res, err = stmt.Exec("111", 110)
    checkErr(err)

    affect, err = res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)

    
    //查询数据
    rows, err := db.Query("SELECT copy_id, level, chapter_id, sub_id FROM copy")
    checkErr(err)

    for rows.Next() {
        var uid int
        var username string
        var department string
        var created string
        err = rows.Scan(&uid, &username, &department, &created)
        checkErr(err)
        fmt.Println(uid)
        fmt.Println(username)
        fmt.Println(department)
        fmt.Println(created)
    }
    
    //删除数据
    stmt, err = db.Prepare("delete from copy where copy_id=?")
    checkErr(err)

    res, err = stmt.Exec(111)
    checkErr(err)

    affect, err = res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)

    db.Close()
    

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}