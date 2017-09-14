package oms

import(
"testing"
"fmt"
)

func Test_Add_OpenUser_REQ(ts *testing.T){
	var qry Add_OpenUser_REQ;
	qry.New();
	pack, _ := qry.Encode();
	fmt.Printf("%s\n", (string)(pack));

	qry.Decode(pack);
	var srv PMS_Web;
	srv.srv(ts);
}
