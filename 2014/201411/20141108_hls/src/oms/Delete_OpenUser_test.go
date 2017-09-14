package oms

import(
"testing"
"fmt"
)

func Test_srv_Del_OpenUser_REQ(ts *testing.T){
	fmt.Printf("========== %s ==========\n", "Delete_OpenUser_REQ");
	defer fmt.Printf("========== End %s ==========\n", "Delete_OpenUser_REQ");
	var qry Delete_OpenUser_REQ;
	qry.New();
	qry.MidId = *MidId;
	qry.LockId = "1";
	qry.CardType = "3";
	qry.CardData = "898989";
	pack, _ := qry.Encode();
	err := Post("http://120.25.106.243:5010/cmd", string(pack));
	if err != nil{
		ts.Errorf(err.Error());
	}	
}
