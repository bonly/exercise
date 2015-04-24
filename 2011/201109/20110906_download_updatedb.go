package main
 
import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
    "github.com/qiniu/iconv"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    )
 
func main() {
    response, err := http.Get("http://14.18.207.155:8000/dev/raw-attachment/wiki/策划/数值数据/装备定义表.csv");
    if err != nil {
        fmt.Printf("%s", err);
        os.Exit(1);
    } else {
        defer response.Body.Close();
        contents, err := ioutil.ReadAll(response.Body);
        if err != nil {
            fmt.Printf("%s", err);
            os.Exit(1);
        }
        //fmt.Printf("%s\n", string(contents));
        
        cd, err := iconv.Open("utf-8","gbk");
        if err != nil{
          fmt.Printf("iconv open failed!");
          return;
        }
        defer cd.Close();
        
        //gbk := cd.ConvString(string(contents));
        //fmt.Printf("%s\n", gbk);
        
        fout, err := os.Create("./tmp.csv");
        defer fout.Close();
        if err != nil{
          fmt.Println("创建临时文件失败: ", err);
          return;
        }
        fout.WriteString(cd.ConvString(string(contents)));     
        
        dump2sql();
    }
}

func dump2sql(){
  db, err := sql.Open("mysql","mech:2405767f@tcp(14.18.207.155:3306)/mech?allowAllFiles=true");
  if err != nil{
    fmt.Println("Open database failed: ", err);
    return;
  }
  defer db.Close();
  
  err = db.Ping();
  if err != nil{
    fmt.Println(err);
  }
  
  _, err = db.Exec("delete from equipment_template");
  if err != nil {
    fmt.Println(err);
  }
  
  _, err = db.Exec("LOAD DATA Local INFILE './tmp.csv' INTO TABLE equipment_template  fields terminated by ',' IGNORE 1 LINES;");
  if err != nil {
    fmt.Println(err);
  }
}