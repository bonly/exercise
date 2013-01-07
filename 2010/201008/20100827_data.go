package main;
//import _ "os";
import "os/exec";
import "fmt";
import "bytes";
import "net/http";
//import "encoding/binary";
//import "strconv";
//import "strings";
import "github.com/ziutek/mymysql/mysql";
import _ "github.com/ziutek/mymysql/native"

var connect_count string;
const user = "mysql"
var pass = "";
var dbname = "paladin";

func get_player(w http.ResponseWriter, r *http.Request){
     all, player := get_reg();
     fmt.Fprintf(w, "注册玩家数为: %d\n", all);
     fmt.Fprintf(w, "创建角色的数量为: %d\n", player);
}

func watch(w http.ResponseWriter, r *http.Request){
     get_count();
     fmt.Fprintf(w, "当前并发的在线人数为: %s\n", connect_count);
}

func get_reg() (int,int){
   db := mysql.New("tcp", "", "117.135.154.58:3306", user, pass, dbname);
   err := db.Connect();
   var ret_all = 0;
   var ret_player = 0;
   if err != nil{
      panic(err);
      return ret_all, ret_player;
   }
   defer db.Close();
   _, _, _ = db.Query("set wait_timeout = 3600");
   _, _, _ = db.Query("set interactive_timeout = 3600");

   rows, _, err := db.Query("select count(customer_id)-54 from customer.customer");
   if err != nil {
      panic(err);
   }
   for _, row := range rows{
      ret_all = row.Int(0);
   }
   rows, _, err = db.Query("select count(player_id) from player where vip>=1 and player_id>27048");
   if err != nil {
      panic(err);
   }
   for _, row := range rows{
      ret_player = row.Int(0);
   }
   return ret_all,ret_player;
}

func get_count(){
  cmd0 := exec.Command("netstat", "-an");
  cmd1 := exec.Command("grep", "8098");
  cmd2 := exec.Command("wc", "-l");

  cmd1.Stdin, _ = cmd0.StdoutPipe();
  cmd2.Stdin, _ = cmd1.StdoutPipe();
  //cmd2.Stdout = os.Stdout;
  var output bytes.Buffer;
  cmd2.Stdout = &output;

  cmd2.Start();
  cmd1.Start();
  cmd0.Run();
  cmd1.Wait();
  cmd2.Wait();

  connect_count = output.String();
  /*connect_count, err := strconv.ParseInt(strcount, 10, 0);
  //connect_count, err := strconv.Atoi(strcount);
  if err != nil{
     fmt.Println(err);
  }
  */
  //fmt.Printf("%s", connect_count);
}

func main(){
   http.HandleFunc("/", watch);
   http.HandleFunc("/player", get_player);
   http.ListenAndServe(":8090", nil);
}

