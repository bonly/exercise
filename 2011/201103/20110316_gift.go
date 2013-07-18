package main
import (
  "fmt"
  _ "html/template"
  _ "log"
  _ "strings" 
  _ "github.com/go-sql-driver/mysql"
  "database/sql"
)

//select * from customer.customer where name = '390325630';
func get_customer (name string) (cr map[string]string){
	var cust_id string;
	var realm_id string
	cr = make(map[string]string);
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
  	cr[cust_id]=realm_id;
  }
  return cr;  
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
		   srv = "mysql:4860e49a@tcp(117.135.154.120:3306)/paladin?charset=utf8";
		   break;
		default: //sdxl
			 srv = "mysql:4860e49a@tcp(117.135.154.120:3306)/paladin?charset=utf8";
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

func main(){
	cr := get_customer("vt1103");
	fmt.Println (cr
	//cust_id , realm_id := get_customer ("vt1103");
	//fmt.Println("cust_id: ", cust_id);
	//fmt.Println("realm_id: ", realm_id);
	//player_id := get_player(realm_id, cust_id);
	//fmt.Println("player_id: ", player_id);
}
