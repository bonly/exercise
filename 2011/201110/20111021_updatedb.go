package main
 
import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
    "github.com/qiniu/iconv"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "flag"
    )

var tb = map[string]string{"装备定义表":"equipment_template", "卡牌定义表":"mech_template", "道具定义表":"item_template", "商店定义表":"mall", "角色定义表":"player_setup", "卡牌升级经验":"mech_upgrade","装备升级经验":"equipment_upgrade"};


var from *string = flag.String("f","http://14.18.207.155:8000/dev/raw-attachment/wiki/策划/数值数据/","from wiki dir");
var to *string = flag.String("t","mech:2405767f@tcp(14.18.207.155:3306)/mech", "to database");

func update(db *(sql.DB), wiki_name string, table_name string){
    response, err := http.Get(*from + wiki_name + ".csv");
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
        
        cd, err := iconv.Open("utf-8","gbk");
        if err != nil{
          fmt.Printf("iconv open failed!");
          return;
        }
        defer cd.Close();
        
        fout, err := os.Create("./tmp.csv");
        defer fout.Close();
        if err != nil{
          fmt.Println("创建临时文件失败: ", err);
          return;
        }
        fout.WriteString(cd.ConvString(string(contents)));     
        
        load2db(db, table_name);
   }      
}

func load2db(db *(sql.DB), table_name string){
  _, err := db.Exec("delete from " + table_name);
  if err != nil {
    fmt.Println(err);
  }
  
  _, err = db.Exec("LOAD DATA Local INFILE './tmp.csv' INTO TABLE " + table_name + " fields terminated by ',' IGNORE 1 LINES;");
  if err != nil {
    fmt.Println(err);
  }  
}

func main(){
  flag.Parse();
  
  db, err := sql.Open("mysql",*to+"?allowAllFiles=true");
  if err != nil{
    fmt.Println("Open database failed: ", err);
    return;
  }
  defer db.Close();
  
  err = db.Ping();
  if err != nil{
    fmt.Println(err);
  }
  
  for key,value := range tb{
    update(db, key, value);
  }
}
    
