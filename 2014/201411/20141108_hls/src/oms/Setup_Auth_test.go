package oms

import(
"testing"
"fmt"
)
func Test_srv_Setup_Auth_REQ(ts *testing.T){
	fmt.Printf("========== %s ==========\n", "Setup_Auth_REQ");
	defer fmt.Printf("========== End %s ==========\n", "Setup_Auth_REQ");
	var qry Setup_Auth_REQ;
	qry.New();
	qry.MidId = *MidId;
	qry.LockId = "1";
	qry.Allow = "true";
	pack, _ := qry.Encode();

	err := Post("http://120.25.106.243:5010/cmd", string(pack));
	if err != nil{
		ts.Errorf(err.Error());
	}	
}
