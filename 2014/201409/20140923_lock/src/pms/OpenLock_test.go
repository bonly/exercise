package pms

import(
"testing"
"fmt"
"manage"
)

func Test_OpenLock_REQ(ts *testing.T){
	var qry OpenLock_REQ;
	qry.New();
	qry.MidId = "00000001";
	qry.LockId = "3";
	pack, _ := qry.Encode();
	fmt.Printf("%s\n", (string)(pack));

	qry.Decode(pack);

	qry.Process();
}
