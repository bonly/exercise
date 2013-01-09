package main;
//import _ "os";
import "os/exec";
import "fmt";
import "flag";
import "bytes";
import "net/http";
//import "encoding/binary";
//import "strconv";
//import "strings";
import "github.com/ziutek/mymysql/mysql";
import _ "github.com/ziutek/mymysql/native"

var listenPort = flag.String("l", ":8090", "http listen port");

var addr = flag.String("a", "117.135.154.58:3306", "mysql address");
var user = flag.String("u", "mysql", "mysql user name");
var pass = flag.String("p", "6yitddfzx", "mysql password");
var dbname = flag.String("d", "paladin", "mysql database name");

var db = mysql.New("tcp", "", *addr, *user, *pass, *dbname);

func show_count(w http.ResponseWriter, r *http.Request){
	   err := db.Connect();
	   defer db.Close();
	   if err != nil{
      panic(err);
     }
     _, _, _ = db.Query("set wait_timeout = 3600");
     _, _, _ = db.Query("set interactive_timeout = 3600");
	   
     //fmt.Fprintf(w, "注册玩家数为: %d\n", all);
     //fmt.Fprintf(w, "创建角色的数量为: %d\n", player);
     fmt.Fprintf(w, player_page, 
                 get_all_reg(),
                 pve_online_5(),
                 get_login_30());   
       
}

func exec_sql(sql string)(int){
	 var ret = 0;
	 rows, _, err := db.Query(sql);
   if err != nil {
      panic(err);
   }
   for _, row := range rows{
      ret = row.Int(0);
   }
   return ret;
}

func pve_online_5() (int){
	sql := `SELECT COUNT(*) ct FROM player 
	                WHERE vip>=1 AND player_id>27048 
	                AND (time_to_sec(timediff(now(), spirit_drop_time ))<=15*60 || time_to_sec(timediff(now(), energy_drop_time ))<=30*60)`
	return exec_sql(sql);
}

func get_all_reg() (int){
   sql := "select count(customer_id)-54 from customer.customer";
   return exec_sql(sql);
}

func get_login_30() (int){
	sql := "SELECT count(*) ct FROM player WHERE TIMESTAMPDIFF(SECOND, NOW(), login_time) >= -1800";
	return exec_sql(sql);
}

func get_player() (int){
   return exec_sql("select count(player_id) from player where vip>=1 and player_id>27048");
}

func get_count(){
  cmd0 := exec.Command("netstat", "-an");
  cmd1 := exec.Command("grep", "8098");
  cmd2 := exec.Command("wc", "-l");

  cmd1.Stdin, _ = cmd0.StdoutPipe();
  cmd2.Stdin, _ = cmd1.StdoutPipe();
  var output bytes.Buffer;
  cmd2.Stdout = &output;

  cmd2.Start();
  cmd1.Start();
  cmd0.Run();
  cmd1.Wait();
  cmd2.Wait();

  connect_count = output.String();
}

var connect_count string;
func watch(w http.ResponseWriter, r *http.Request){
     get_count();
     fmt.Fprintf(w, "<head><meta http-equiv=\"refresh\" content=\"5\"></head>");
     fmt.Fprintf(w, "当前并发的在线人数为: %s\n", connect_count);
}

func main(){
   http.HandleFunc("/", watch);
   http.HandleFunc("/player", show_count);
   http.ListenAndServe(*listenPort, nil);
}

var player_page =
`
<HEAD><META><meta http-equiv="Content-Type" content="text/htm; charset=utf8"></HEAD>     
<style type=\"text/css\">
<!--
.style1 {color: #FF0000}
.style2 {font-size: 18px}
.style3 {font-size: 12px}
.style5 {font-size: 12px; color: #CC0000; }
-->
</style>
<table width="615" height="414" border="1" cellspacing="0" bordercolor="#333366">
  <tr>
    <td colspan="6"><div align="center" class="style2"><span class="style1">《武林萌主》</span>当前情况<span class="style3"> <span class="style5"><?php echo $datetime; ?></span></span></div></td>
  </tr>
  <tr>
    <td width="98"><span class="style3">总注册玩家数量</span></td>
    <td width="88"><span class="style5">  %d &nbsp;</span></td>
    <td width="132"><span class="style3">PVE/PVP活跃数量</span></td>
    <td width="95"><span class="style3"><span class="style5"> %d </span></span></td>
    <td width="88"><span class="style3">隔日流失率</span></td>
    <td width="88"><span class="style3"><span class="style5">100</span></span></td>
  </tr>
  <tr>
    <td><span class="style3">级别1-5级</span></td>
    <td><span class="style5"><?php echo $total_player_1; ?></span></td>
    <td><span class="style3">30分钟登录玩家数量</span></td>
    <td><span class="style3"><span class="style5"> %d </span></span></td>
    <td><span class="style3"></span></td>
    <td><span class="style3"></span></td>
  </tr>
  <tr>
    <td><span class="style3">级别5-10级</span></td>
    <td><span class="style5"><?php echo $total_player_5; ?></span></td>
    <td><span class="style3">60分钟登录玩家数量</span></td>
    <td><span class="style3"><span class="style5"><?php echo $login_60; ?></span></span></td>
    <td><span class="style3"></span></td>
    <td><span class="style3"></span></td>
  </tr>
  <tr>
    <td><span class="style3">级别10-20级</span></td>
    <td><span class="style5"><?php echo $total_player_10; ?></span></td>
    <td><span class="style3"></span></td>
    <td><span class="style3"></span></td>
    <td><span class="style3"></span></td>
    <td><span class="style3"></span></td>
  </tr>
  <tr>
    <td><span class="style3">级别20-40级</span></td>
    <td><span class="style5"><?php echo $total_player_20; ?></span></td>
    <td><span class="style3"></span></td>
    <td><span class="style3"></span></td>
    <td><span class="style3"></span></td>
    <td><span class="style3"></span></td>
  </tr>
  <tr>
    <td><span class="style3">级别&gt;=40</span></td>
    <td><span class="style5"><?php echo $total_player_40; ?></span></td>
    <td><span class="style3"></span></td>
    <td><span class="style3"></span></td>
    <td><span class="style3"></span></td>
    <td><span class="style3"></span></td>
  </tr>
  <tr>
    <td><span class="style3"></span></td>
    <td>&nbsp;</td>
    <td><span class="style3"></span></td>
    <td><span class="style3"></span></td>
    <td><span class="style3"></span></td>
    <td><span class="style3"></span></td>
  </tr>
  <tr>
    <td><span class="style3"></span></td>
    <td>&nbsp;</td>
    <td><span class="style3"></span></td>
    <td><span class="style3"></span></td>
    <td><span class="style3"></span></td>
    <td><span class="style3"></span></td>
  </tr>
  <tr>
    <td><span class="style3">累计充值</span></td>
    <td><span class="style2"><span class="style3"><span class="style5"><?php echo $deposit; ?></span></span></span></td>
    <td><span class="style3"></span></td>
    <td><span class="style3"></span></td>
    <td><span class="style3"></span></td>
    <td><span class="style3"></span></td>
  </tr>
  </table>
`;