package oms

import(
"testing"
"fmt"
)

func Test_srv_Lock_stat_REQ(ts *testing.T){
	fmt.Printf("========== %s ==========\n", "Lock_stat_REQ");
	defer fmt.Printf("========== End %s ==========\n", "Lock_stat_REQ");
	var qry Lock_stat_REQ;
	qry.New();
	qry.MidId = *MidId;
	qry.LockId = "1";
	pack, _ := qry.Encode();

	err := Post(*Srv, string(pack));
	if err != nil{
		ts.Errorf(err.Error());
	}
}
