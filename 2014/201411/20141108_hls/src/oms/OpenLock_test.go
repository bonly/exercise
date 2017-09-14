package oms

import(
"testing"
"fmt"
)

func Test_srv_OpenLock_REQ(ts *testing.T){
	fmt.Printf("========== %s ==========\n", "OpenLock_REQ");
	defer fmt.Printf("========== End %s ==========\n", "OpenLock_REQ");
	var qry OpenLock_REQ;
	qry.New();
	qry.MidId = *MidId;
	qry.LockId = "1";
	pack, _ := qry.Encode();
	// fmt.Printf("%s\n", (string)(pack));

	// qry.Decode(pack);

	err := Post(*Srv, string(pack));
	if err != nil{
		ts.Errorf(err.Error());
	}	
}
