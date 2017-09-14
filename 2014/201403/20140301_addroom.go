package main  

import(
"fmt"
"log"
"github.com/jmoiron/sqlx"
_ "github.com/go-sql-driver/mysql"
)

const constr = "db_admin:db_admin2015@tcp(120.25.106.243:3306)/xbed_service?charset=utf8";

var Db *sqlx.DB;

func main(){
	var err error;
	if Db, err = sqlx.Open("mysql", constr); err != nil{
		log.Printf(fmt.Sprintf("connect db fail: %s", err.Error()));
		return;
	}
	Db.SetMaxIdleConns(5);
	defer Db.Close();

	var temp_room Xb_temp_room;
	if temp_room, err = get_temp_room(); err != nil{
		return;
	}

	if err := ins_xb_room(&temp_room); err != nil{
		return;
	}


}

type Xb_temp_room struct{
Xb_room_id,
Update_room,
Update_at,
Up_date,
Stress_address,
Status,
Room_type,
Room_id,
Province,
Parent_id,
Owner_name,
Owner_mobile,
Operation_date,
Name,
Img_urls,
Id,
Down_date,
District,
Create_at,
City string;
};

func get_temp_room() (info Xb_temp_room, err error){
	sql := `select ifnull(xb_room_id,'')xb_room_id,
ifnull(update_room,'')update_room,
ifnull(update_at,'')update_at,
ifnull(up_date,'')up_date,
ifnull(stress_address,'')stress_address,
ifnull(status,'')status,
ifnull(room_type,'')room_type,
ifnull(room_id,'')room_id,
ifnull(province,'')province,
ifnull(parent_id,'')parent_id,
ifnull(owner_name,'')owner_name,
ifnull(owner_mobile,'')owner_mobile,
ifnull(operation_date,'')operation_date,
ifnull(name,'')name,
ifnull(img_urls,'')img_urls,
ifnull(id,'')id,
ifnull(down_date,'')down_date,
ifnull(district,'')district,
ifnull(create_at,'')create_at,
ifnull(city,'')city
from xb_temp_room 
where status='wait_check' limit 1`;
	rows, err := Db.Queryx(sql);
	if err != nil{
		log.Println("sql: ", err);
		return;
	}

	for rows.Next(){
		// var info Xb_temp_room;
		if err := rows.StructScan(&info); err != nil{
			log.Println("row: ", err);
			return info, err;
		}
		log.Println(fmt.Sprintf("%+v",info));
	}
	return info, err;
}


type Xb_room struct{
	Room_id,Room_name,Chain_id,
	Title,Addr,Flag,Room_type_name,
	Building_id,Room_type_id,Room_floor,
	Area,House_type,Stat,Price,Locate,
	Province,City,District,Descri,Pic_id,
	Pic_count,Currency,Tag,Checkin,Checkout,
	Checkout_plot,Lodger_count string;
};

/*
insert into xb_room(room_id,room_name,chain_id,title,addr,flag,room_type_name,building_id,room_type_id,room_floor,area,house_type,stat,price,locate,province,city,district,descri,pic_id,pic_count,currency,tag,checkin,checkout,checkout_plot,lodger_count) 
values(
)
*/
func ins_xb_room(tmp *Xb_temp_room) (err error){
	sql := `insert into xb_room(owner_room_id,room_name,
		chain_id,title,addr,flag,room_type_name,
		building_id,room_type_id,room_floor,area,
		house_type,stat,price,locate,province,city,
		district,descri,pic_id,pic_count,currency,
		tag,checkin,checkout,checkout_plot,lodger_count) 
	values(:Xb_room_id, :Name, 
		'1',:Name, :Stress_address, '1', '测试房',
		'0','0', '11', '80方',
		'一房一厅', '1', '36800', '0.00,0.00',:Province, :City,
		:District, '新房开售', 0, 0, '人民币',
		'热销', '14:00', '12:00', '退房政策', '2人房'
	)`;
	_, err = Db.NamedExec(sql, *tmp);
	if err != nil{
		log.Println("insert: ", err);
		return err;
	}
	return nil;
}