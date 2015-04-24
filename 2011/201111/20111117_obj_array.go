package main 
 import (
  _ "github.com/go-sql-driver/mysql"
  "database/sql"
  "net/http"
  _"fmt"
  "log"
  "html/template"
)

type Tbl struct{
    User_id string;
    Acc_id string;
}

func main(){
    http.HandleFunc("/", mylist);
    err := http.ListenAndServe(":9090", nil);
    if err != nil{
        log.Fatal("ListenAndServe: ", err);
    }
}


func mylist(w http.ResponseWriter, r *http.Request){
        db, err := sql.Open("mysql", "bonly@tcp(127.0.0.1:3306)/moudao?charset=utf8");
        if err != nil{
            panic(err);
        }
        defer db.Close();

        rows, err := db.Query("select user_id,acc_id from users");
        if err != nil{
            panic(err);
        }
        var lst []Tbl;
        for rows.Next(){
            var user_id string;
            var acc_id string;
            err = rows.Scan(&user_id, &acc_id);
            if err != nil{
                  panic(err);
            }
            var tbItems Tbl;
            tbItems.User_id = user_id;
            tbItems.Acc_id = acc_id;
            lst = append(lst, tbItems); //存入的是实例的复制品？
        } 
        //index, _ := template.ParseFiles("/tmp/untitled.html");
        index, _ := template.New("tmp").Parse(tpl);
        index.Execute(w, lst);      
  
}


var tpl string =`
<!DOCTYPE html>
<html>
    <head>
        <title> Test </title>
    </head>
    <body>
        <section id="contents">
        {{range .}}
        <p> {{.User_id}} {{.Acc_id}} </p>
        {{end}}
       </section>
    </body>
</html>
`;

//        {{range .}}    //没有数组的不能用range
//        {{end}}