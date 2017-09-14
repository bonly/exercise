package proto

import (
"testing"
"fmt"
)

func Test_srv_setup_authority_all(ts *testing.T){
	var cmd Setup_lock_authority;
	cmd.New();
	
	cmd.head.LockID = 0x1; 
	cmd.Authority = ALLOW_OPEN;

	pack, _ := cmd.Encode();

	currect := "010204809038300D";
	if fmt.Sprintf("%X", pack) == currect{
		ts.Logf("构造包len[%d]: %X\n", len(pack), pack);
	}else{
		ts.Errorf("构造包不正确%s", currect);
	}
	
	Hex_Dump(pack, len(pack));

	ts.Logf("设置门锁权限\n");
	var srv SRV;
	srv.pack = pack;
	srv.srv(ts);	
}

func Test_srv_setup_authority_none(ts *testing.T){
	var cmd Setup_lock_authority;
	cmd.New();
	
	cmd.head.LockID = 0x1; 
	cmd.Authority = FORBIT_OPEN;

	pack, _ := cmd.Encode();

	currect := "010204809038310D";
	if fmt.Sprintf("%X", pack) == currect{
		ts.Logf("构造包len[%d]: %X\n", len(pack), pack);
	}else{
		ts.Errorf("构造包不正确%s", currect);
	}
	
	Hex_Dump(pack, len(pack));

	ts.Logf("设置门锁权限\n");
	var srv SRV;
	srv.pack = pack;
	srv.srv(ts);	
}