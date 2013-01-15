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
import _ "github.com/Go-SQL-Driver/MySQL";
import   "database/sql";

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
	D级别   [13]int;
	D累计充值    int
}

type MissionCnt struct {
	Chapter int;
	Segment int;
	Name string;
	Count int;
}

var srv1 CountSrv

func show_count(w http.ResponseWriter, r *http.Request) {
	log.Println("process count request");
	nav(w);
	srv1.TimeNow = time.Now()

  conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", *(srv1.user),*(srv1.pass),*(srv1.addr),*(srv1.dbname));
  db1, err := sql.Open("mysql", conn);
	defer db1.Close()
	if err != nil {
		log.Fatal(err)
	}
	_, _ = db1.Query("set wait_timeout = 3600")
	_, _ = db1.Query("set interactive_timeout = 3600")

	srv1.D帐户数 = get_all_reg(db1)
	srv1.D角色数 = get_player(db1);
	srv1.D活跃数 = pve_online_5(db1)
	srv1.D登录_30 = get_login_30(db1)
	srv1.D登录_60 = get_login_60(db1)

  log.Println("parse file: ", *cssFile);
	tmp, err := template.ParseFiles(*cssFile)
	tmp.Execute(w, srv1);
	if err != nil {
		log.Fatal(err);
	}
}

func show_level(w http.ResponseWriter, r *http.Request) {
	log.Println("process level request");
	nav(w);

  conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", *(srv1.user),*(srv1.pass),*(srv1.addr),*(srv1.dbname));
  db1, err := sql.Open("mysql", conn);
  
	defer db1.Close()
	if err != nil {
		panic(err)
	}
	_, _ = db1.Query("set wait_timeout = 3600")
	_, _ = db1.Query("set interactive_timeout = 3600")
	
	get_level_count(db1, &srv1);
	
	fmt.Fprintf(w, `<br>`);
	level_tmp := template.New("");
	level_tmp.Parse(`
	  <table width="307" height="414" border="1" cellspacing="0" bordercolor="#333366">
	  <tr>
		  <td><span class="style3">等级</span></td>
		  <td width="88"><span class="style5">  人数 &nbsp;</span></td>
		</tr>
	  {{range $k,$v := .D级别}}
	  <tr>
		  <td><span class="style3">级别 {{$k}}0-{{$k}}9 级</span></td>
		  <td width="88"><span class="style5">  {{$v}} &nbsp;</span></td>
		</tr>
		{{end}}
	  </table>
	`);
	level_tmp.Execute(w, srv1);
}

func show_mission(w http.ResponseWriter, r *http.Request) {
	log.Println("process mission request");
	nav(w);
	
  conn := fmt.Sprintf("%v:%s@tcp(%s)/%s?charset=utf8", *(srv1.user),*(srv1.pass),*(srv1.addr),*(srv1.dbname));
  db1, err := sql.Open("mysql", conn);
  log.Println("conn:", conn);
	defer db1.Close()
	if err != nil {
		panic(err)
	}
	_, err = db1.Query("set wait_timeout = 3600")
	_, err = db1.Query("set interactive_timeout = 3600")
	if err != nil {
		panic(err)
	}
		
	mission := get_mission_count(db1);
	
	fmt.Fprintf(w, `<br>`);
	fmt.Fprintf(w, `<table width="307" height="414" border="1" cellspacing="0" bordercolor="#333366">`);
	fmt.Fprintf(w,`
	<td><span class="style3"> 书</span></td>
	<td><span class="style3"> 章</span></td>
	<td width="88"><span class="style5">  简述 &nbsp;</span></td>
	<td width="88"><span class="style5">  人数 &nbsp;</span></td>`);
	
	for _, v := range mission{
		fmt.Fprintf(w, `<tr>`);
		if v.Count > 0 {
			level_tmp,_ := template.New("tr").Parse(`
			   <td><span class="style3"> {{.Chapter}}</span></td>
			   <td><span class="style3"> {{.Segment}}</span></td>
			   <td width="88"><span class="style5">  {{.Name}} &nbsp;</span></td>
			   <td width="88"><span class="style5">  {{.Count}} &nbsp;</span></td>`);
			level_tmp.Execute(w, v);
			fmt.Fprintf(w, `</tr>`);
	  }
	}
	fmt.Fprintf(w, `</table>`);
}

func nav(w http.ResponseWriter){
	fmt.Fprintf(w,`
<head>
<br>
<hr>
<a href="http://183.60.126.26:8888">项目管理</a>
<a href="/">实时在线</a>
<a href="/player">用户统计</a>
<a href="/level">级别统计</a>
<a href="/mission">江湖统计</a>
<br>
<hr>
</head>
  `);	
}

func exec_sql(sql string, db *(sql.DB)) int {
	var ret = 0;
	rows := (*db).QueryRow(sql);
	err := rows.Scan(&ret);
	if err != nil {
		panic(err);
	}
	return ret;
}

func get_level_count(db *(sql.DB), sh *CountSrv) {
	sql := `select level_class,count(level_class) from
					(
					  select floor(level/10) as level_class from player where vip>0 and level>=1
					) as tmp_tab
					group by level_class order by 1 limit 10;
  `
	rows, err := (*db).Query(sql)
	if err != nil {
		panic(err)
	}
	var tmp int;
	for i:=0;rows.Next();i++ {
		rows.Scan(&tmp, &sh.D级别[i]);
	}
	return
}

func get_mission_count(db *(sql.DB))(sh [500]MissionCnt) {
	var mission [500]MissionCnt;
	sql := `select m.chapter_id,m.segment_id,m.name,count(c.player_id)
				  from player_mission c, mission m
					where c.mission_id=m.mission_id 
					group by m.chapter_id,m.segment_id
  `;
	rows, err := (*db).Query(sql)
	if err != nil {
		panic(err)
	}
	for i:=0;rows.Next();i++ {
	  rows.Scan(&mission[i].Chapter, &mission[i].Segment, &mission[i].Name, &mission[i].Count);
	}
	return mission;
}

func pve_online_5(db *(sql.DB)) int {
	sql := `SELECT COUNT(*) ct FROM player 
	                WHERE vip>=1 AND player_id>27048 
	                AND (time_to_sec(timediff(now(), spirit_drop_time ))<=15*60 || time_to_sec(timediff(now(), energy_drop_time ))<=30*60)`
	return exec_sql(sql, db)
}

func get_all_reg(db *(sql.DB)) int {
	sql := "select count(customer_id)-54 from customer.customer"
	return exec_sql(sql, db)
}

func get_login_30(db *(sql.DB)) int {
	sql := "SELECT count(*) ct FROM player WHERE TIMESTAMPDIFF(SECOND, NOW(), login_time) >= -1800"
	return exec_sql(sql, db)
}

func get_login_60(db *(sql.DB)) int {
	sql := "SELECT count(*) ct FROM player WHERE TIMESTAMPDIFF(SECOND, NOW(), login_time) >= -3600"
	return exec_sql(sql, db)
}

func get_player(db *(sql.DB)) (int){
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
	nav(w);
	get_count()
	fmt.Fprintf(w, "<head><meta http-equiv=\"refresh\" content=\"5\"></head>")
	fmt.Fprintf(w, "当前并发的在线人数为: %s\n", connect_count)
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
	http.HandleFunc("/level", show_level);
	http.HandleFunc("/mission", show_mission);
	
	log.Println("listen ", *listenPort);
	if err := http.ListenAndServe(*listenPort, nil); err != nil{
		panic (err);
	}
}

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

/*
select c.mission_id,m.name,count(player_id) 
from player_mission c, mission m
where c.mission_id=m.mission_id 
group by c.mission_id;
*/

/*
select m.chapter_id,m.segment_id,m.name,count(c.player_id)
from player_mission c, mission m
where c.mission_id=m.mission_id 
group by m.chapter_id,m.segment_id
*/
