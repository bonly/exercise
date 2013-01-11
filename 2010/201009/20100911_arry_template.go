package main

//import _ "os";
import "os/exec"
import "bytes"
import "fmt"
import "time"
import "flag"
import "net/http"
import "html/template"
import "log"

//import "encoding/binary";
//import "strconv";
//import "strings";
import "github.com/ziutek/mymysql/mysql"
import _ "github.com/ziutek/mymysql/native"

var listenPort = flag.String("l", ":8090", "http listen port");
var cssFile = flag.String("f", "css.html", "template file");

type CountSrv struct {
	addr   *string
	user   *string
	pass   *string
	dbname *string
	Cname  *string

	TimeNow  time.Time
	D帐户数     int;
	D角色数     int;
	D活跃数     int
	D隔日流失率   int
	D登录_30   int
	D登录_60   int
	D级别   [10]int;
	D累计充值    int
}

var srv1 CountSrv

func show_count(w http.ResponseWriter, r *http.Request) {
	log.Println("process count request");
	srv1.TimeNow = time.Now()

	db1 := mysql.New("tcp", "", *(srv1.addr), *(srv1.user), *(srv1.pass), *(srv1.dbname));
	err := db1.Connect()
	defer db1.Close()
	if err != nil {
		panic(err)
	}
	_, _, _ = db1.Query("set wait_timeout = 3600")
	_, _, _ = db1.Query("set interactive_timeout = 3600")

	//	     fmt.Fprintf(w, player_page, 
	//	                 get_all_reg(&db1),
	//	                 pve_online_5(&db1),
	//	                 get_login_30(&db1)); 

	srv1.D帐户数 = get_all_reg(&db1)
	srv1.D角色数 = get_player(&db1);
	srv1.D活跃数 = pve_online_5(&db1)
	srv1.D登录_30 = get_login_30(&db1)
	srv1.D登录_60 = get_login_60(&db1)
	get_level_count(&db1, &srv1)

  log.Println("parse file: ", *cssFile);
	tmp, _ := template.ParseFiles(*cssFile)
	tmp.Execute(w, srv1);
	
  fmt.Fprintf(w, `<br>`);
	level_tmp := template.New("");
	level_tmp.Parse(`
	  <table width="307" height="414" border="1" cellspacing="0" bordercolor="#333366">
	  {{range $k,$v := .D级别}}
	  <tr>
		  <td><span class="style3">级别 {{$k}}0-{{$k}}9 级</span></td>
		  <td width="88"><span class="style5">  {{$v}} &nbsp;</span></td>
		</tr>
		{{end}}
	  </table>
	`);
	level_tmp.Execute(w, srv1);
	
	nav(w);

}

func nav(w http.ResponseWriter){
	fmt.Fprintf(w,`
<br>
<hr>
<a href="/">实时在线</a>
<a href="http://183.60.126.26:8888">项目管理</a>
<br>
  `);	
}
func exec_sql(sql string, db *(mysql.Conn)) int {
	var ret = 0
	rows, _, err := (*db).Query(sql)
	if err != nil {
		panic(err)
	}
	for _, row := range rows {
		ret = row.Int(0)
	}
	return ret
}

func get_level_count(db *(mysql.Conn), sh *CountSrv) {
	sql := `select level_class,count(level_class) from
					(
					  select floor(level/10) as level_class from player where vip>0 and level>=1
					) as tmp_tab
					group by level_class order by 1 limit 10;
  `
	rows, _, err := (*db).Query(sql)
	if err != nil {
		panic(err)
	}
	for i, row := range rows {
		sh.D级别[i] = row.Int(1);
	}
	return

}

func pve_online_5(db *(mysql.Conn)) int {
	sql := `SELECT COUNT(*) ct FROM player 
	                WHERE vip>=1 AND player_id>27048 
	                AND (time_to_sec(timediff(now(), spirit_drop_time ))<=15*60 || time_to_sec(timediff(now(), energy_drop_time ))<=30*60)`
	return exec_sql(sql, db)
}

func get_all_reg(db *(mysql.Conn)) int {
	sql := "select count(customer_id)-54 from customer.customer"
	return exec_sql(sql, db)
}

func get_login_30(db *(mysql.Conn)) int {
	sql := "SELECT count(*) ct FROM player WHERE TIMESTAMPDIFF(SECOND, NOW(), login_time) >= -1800"
	return exec_sql(sql, db)
}

func get_login_60(db *(mysql.Conn)) int {
	sql := "SELECT count(*) ct FROM player WHERE TIMESTAMPDIFF(SECOND, NOW(), login_time) >= -3600*12"
	return exec_sql(sql, db)
}

func get_player(db *(mysql.Conn)) (int){
   return exec_sql("select count(player_id) from player where vip>=1 and player_id>27048", db);
}

func get_count() {
	cmd0 := exec.Command("netstat", "-an")
	cmd1 := exec.Command("grep", "8098")
	cmd2 := exec.Command("wc", "-l")

	cmd1.Stdin, _ = cmd0.StdoutPipe()
	cmd2.Stdin, _ = cmd1.StdoutPipe()
	var output bytes.Buffer
	cmd2.Stdout = &output

	cmd2.Start()
	cmd1.Start()
	cmd0.Run()
	cmd1.Wait()
	cmd2.Wait()

	connect_count = output.String()
}

var connect_count string

func watch(w http.ResponseWriter, r *http.Request) {
	get_count()
	fmt.Fprintf(w, "<head><meta http-equiv=\"refresh\" content=\"5\"></head>")
	fmt.Fprintf(w, "当前并发的在线人数为: %s\n", connect_count)
	fmt.Fprintf(w, `<br><hr><a href="player">统计数据</a>`)
}

func main() {
	log.Println("server begin...");
	srv1.addr = flag.String("a", "117.135.154.58:3306", "mysql address")
	srv1.user = flag.String("u", "mysql", "mysql user name")
	srv1.pass = flag.String("p", "6yitddfzx", "mysql password")
	srv1.dbname = flag.String("d", "paladin", "mysql database name")
	flag.Parse();

  log.Println("define handle");
	http.HandleFunc("/", watch);
	http.HandleFunc("/player", show_count);
	
	log.Println("listen ", *listenPort);
	if err := http.ListenAndServe(*listenPort, nil); err != nil{
		panic (err);
	}
}

var player_page = `
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
`

/*
select case when level < 10 then ('1-9')
            when level >=10 and level < 20 then ('10-20')
            when level >=20 and level < 30 then ('20-30') end 级别,
count(player_id) from player where vip >0 and level >=1 
group by 
case when level < 10 then ('1-9')
            when level >=10 and level < 20 then ('10-20')
            when level >=20 and level < 30 then ('20-30')
end
*/

/*
select level_class,count(level_class) from
(
  select floor(level/10)*10 as level_class from player where vip>0 and level>=1
) as tmp_tab
group by level_class
*/

/*
select chapter_id,sub_id,count(player_id) 
from player_copy
group by chapter_id,sub_id
*/