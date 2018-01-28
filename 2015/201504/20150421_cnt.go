package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"
    "strings"

    _ "github.com/go-sql-driver/mysql"
)

type tude struct {
    Lat float32 `json:lat`
    Lng float32 `json:lng`
}

type region struct {
    Id       string   `json:id`
    Name     string   `json:name,omitempty`
    Fullname string   `json:fullname`
    Location tude     `json:location`
    Cidx     []int    `json:cidx,omitempty`
    Pinyin   []string `json:pinyin,omitempty`
}

type response struct {
    Status  int        `json:status`
    Message string     `json:message`
    Result  [][]region `json:result`
}

func main() {
    result := getRegionData()
    db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/DBName")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    insertIntoDB(db, result) // 插入获取到的地址数据

    fixData(db, "[1-9][0-9]0000", "", 0, 1)                // 处理省级：只需要将省级 level 更新为 1
    fixData(db, "[1-9][0-9]{3}00", "[0-9]{2}00", 10000, 2) // 处理地级市：更新市级 level 为 2、parent 为其上级 region_id
    fixData(db, "[1-9][0-9]{3}00", "[0-9]{2}", 100, 3)     // 处理区县：更新区县 level 为 3、parent 为其上级 region_id
    fixData(db, "[1-9][0-9]{3}00", "[0-9]{4}", 10000, 2)   // 处理县级市：更新市级 level 为 2、parent 为其上级 region_id
}

/**
* 整理数据层级
 */
func fixData(db *sql.DB, regexp, regexp2 string, offset, level int) {
    stmt, err := db.Prepare("UPDATE `lc_tencent_regions` SET level=?,parent=? WHERE region_code REGEXP ? AND parent=0 AND level=0")
    defer stmt.Close()
    if err != nil {
        log.Fatal(err)
    }
    if regexp2 == "" {
        _, err := stmt.Exec(level, 0, regexp)
        if err != nil {
            log.Fatal(err)
        }
    } else {
        rows, err := db.Query("SELECT region_id,region_code FROM `lc_tencent_regions` WHERE region_code REGEXP ?", regexp)
        defer rows.Close()
        if err != nil {
            log.Fatal(err)
        }
        var (
            region_id   int
            region_code int
        )
        for rows.Next() {
            rows.Scan(&region_id, &region_code)
            regexp := strconv.Itoa(region_code/offset) + regexp2
            _, err := stmt.Exec(level, region_id, regexp)
            if err != nil {
                log.Fatal(err)
            }
        }
    }
}

func getRegionData() [][]region {
    resp, err := http.Get("http://apis.map.qq.com/ws/district/v1/list?key=腾讯LBS的key，可以在 lbs.qq.com 上免费获得")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    var rg response
    err = json.Unmarshal(body, &rg)
    if err != nil {
        log.Fatal(err)
    }
    if rg.Status != 0 {
        log.Fatal(rg.Message)
    }

    return rg.Result
}

/**
* 数据入库
 */
func insertIntoDB(db *sql.DB, result [][]region) {
    sqlStr := "INSERT INTO `lc_tencent_regions`(region_code,name,region_name,lat,lng,pinyin,cidx)VALUES"
    vals := []interface{}{}
    for _, value := range result {
        for _, v := range value {
            cidx := int2Str(v.Cidx)
            pinyin := strings.Join(v.Pinyin, ",")
            sqlStr += "(?,?,?,?,?,?,?),"
            vals = append(vals, v.Id, v.Name, v.Fullname, fmt.Sprintf("%.5f", v.Location.Lat), fmt.Sprintf("%.5f", v.Location.Lng), pinyin, cidx)
            fmt.Println(v)
        }
    }
    stmt, err := db.Prepare(sqlStr[:len(sqlStr)-1])
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()
    res, err := stmt.Exec(vals...)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(res.RowsAffected())
}

/**
* 数字转字符串
 */
func int2Str(v []int) string {
    var str []string
    for _, i := range v {
        str = append(str, strconv.Itoa(i))
    }
    return strings.Join(str, ",")
}