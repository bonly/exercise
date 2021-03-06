package oms

import(
"testing"
"fmt"
)

func Test_srv_Time_REQ(ts *testing.T){
	fmt.Printf("========== %s ==========\n", "Time_REQ");
	defer fmt.Printf("========== End %s ==========\n", "Time_REQ");
	var qry Time_REQ;
	qry.New();
	qry.MidId = *MidId;
	qry.LockId = "1";
	pack, _ := qry.Encode();
	err := Post("http://120.25.106.243:5010/cmd", string(pack));
	if err != nil{
		ts.Errorf(err.Error());
	}	
}
