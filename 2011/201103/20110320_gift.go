package main
import (
  "fmt"
  "html/template"
  "net/http"
  "log"
  _ "strings" 
  _ "github.com/go-sql-driver/mysql"
  "database/sql"
)

//select * from customer.customer where name = '390325630';
func get_customer (name string) (ret map[string]string){
  cust_id := "not found player ";
	realm_id := "not in space";
	ret = make(map[string]string);
	db, err := sql.Open("mysql", "mysql:6yitddfzx@tcp(117.135.154.58:3306)/customer?charset=utf8");
  if err != nil{
  	panic(err);
  }
  defer db.Close();
  
  rows, err := db.Query("select cu.customer_id,realm_id from customer.customer cu, customer.customer_realm rm where cu.customer_id=rm.customer_id and Name = '" + name +"'");
  if err != nil{
  	panic(err);
  }
  
  for rows.Next(){
  	err = rows.Scan(&cust_id, &realm_id);
  	if err != nil{
  		panic(err);
  	}
  	ret[realm_id] = cust_id;
  }
  return ret;
}

//SELECT * FROM player WHERE account_id=273010;
func get_player (realm string, account_id string) (player_id string){
	srv := "";
	switch realm {
		case "3": //xajh
		   srv = "mysql:6yitddfzx@tcp(117.135.154.58:3306)/paladin?charset=utf8";
		   break;
		case "4": //tlbb
		   srv = "mysql:2405767f@tcp(117.135.154.59:3306)/paladin?charset=utf8";
		   break;
		case "2": //sdxl
		   //srv = "mysql:4860e49a@tcp(117.135.154.120:3306)/paladin?charset=utf8";
		   srv = "mysql:4860e49a@tcp(117.135.154.59:3307)/paladin?charset=utf8";
		   break;
		default: //sdxl
			 srv = "mysql:4860e49a@tcp(117.135.154.59:3307)/paladin?charset=utf8";
		   break;
	}
	
	db, err := sql.Open("mysql", srv);
  if err != nil{
  	panic(err);
  }
  defer db.Close();
  
  rows, err := db.Query("SELECT player_id FROM player WHERE account_id = '" + account_id +"'");
  if err != nil{
  	panic(err);
  }
  
  for rows.Next(){
  	err = rows.Scan(&player_id);
  	if err != nil{
  		panic(err);
  	}
  }
  return player_id;  
}


func search(w http.ResponseWriter, r *http.Request){
	fmt.Println("method:", r.Method);
	if r.Method == "GET"{
		t, _:= template.ParseFiles("user.html");
		t.Execute(w, nil);	
	}else{ 
		username := r.FormValue("username");
		
		fmt.Fprintf(w, "<head>search for %s </head> </br>", username);
		fmt.Fprintf(w, "<body></br>");
				
		cust := get_customer (username);
		
		fmt.Fprintf(w, "<form action='/user' method='post'>");
		fmt.Fprintf(w, "<ol>");
		for key,value := range cust{
			fmt.Fprintf(w, "<li><input type='radio' name='cust_play' value='%s:%s'> %s server: ", key, value, key);
			fmt.Fprintf(w, "cust_id: %s ", value);
			fmt.Fprintf(w, "realm_id: %s ", key);
			player_id := get_player (key, value);
			fmt.Fprintf(w, "player_id: %s </li>", player_id);
		}
		fmt.Fprintf(w, "</ol>");
		fmt.Fprintf(w, "</form>");
		
		fmt.Fprintf(w, "</body>");
	}
}

func main(){
	http.HandleFunc("/user", search);
	err := http.ListenAndServe(":9090", nil);
	if err != nil{
		log.Fatal("ListenAndServe: ", err);
	}
}
