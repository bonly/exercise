package oms

import(
"testing"
"fmt"
)

func Test_srv_Clean_Log_REQ(ts *testing.T){
	fmt.Printf("========== %s ==========\n", "Clean_Log_REQ");
	defer fmt.Printf("========== End %s ==========\n", "Clean_Log_REQ");
	var qry Clean_Log_REQ;
	qry.New();
	qry.MidId = *MidId;
	qry.LockId = "1";
	pack, _ := qry.Encode();
	// fmt.Printf("%s\n", (string)(pack));

	// qry.Decode(pack);

	err := Post("http://120.25.106.243:5010/cmd", string(pack));
	if err != nil{
		ts.Errorf(err.Error());
	}	
}
