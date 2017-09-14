package oms

import(
"testing"
"fmt"
)

// func Test_Add_OpenUser_REQ(ts *testing.T){
// 	var qry Add_OpenUser_REQ;
// 	qry.New();
// 	pack, _ := qry.Encode();
// 	fmt.Printf("%s\n", (string)(pack));

// 	qry.Decode(pack);
// }

func Test_srv_Add_OpenUser_REQ(ts *testing.T){
	fmt.Printf("========== %s ==========\n", "Add_OpenUser_REQ");
	defer fmt.Printf("========== End %s ==========\n", "Add_OpenUser_REQ");
	var qry Add_OpenUser_REQ;
	qry.New();
	qry.MidId = *MidId;
	qry.LockId = "1";
	qry.CardType = "3";
	qry.CardData = "898989";
	qry.BeginTime = "2016-09-21 18:59:00";
	qry.EndTime = "2016-09-29 18:59:00";
	pack, _ := qry.Encode();
	err := Post("http://120.25.106.243:5010/cmd", string(pack));
	if err != nil{
		ts.Errorf(err.Error());
	}	
}
