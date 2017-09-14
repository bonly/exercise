package proto

import (
"testing"
"fmt"
)

func Test_srv_setup_clean_log(ts *testing.T){
	var cmd Setup_lock_clean_log;
	cmd.New();
	
	cmd.head.LockID = 0x1; 

	pack, _ := cmd.Encode();

	currect := "010203809039032B";
	if fmt.Sprintf("%X", pack) == currect{
		ts.Logf("构造包len[%d]: %X\n", len(pack), pack);
	}else{
		ts.Errorf("构造包不正确%s", currect);
	}
	
	Hex_Dump(pack, len(pack));

	ts.Logf("设置门锁日志清空\n");
	var srv SRV;
	srv.pack = pack;
	srv.srv(ts);	
}