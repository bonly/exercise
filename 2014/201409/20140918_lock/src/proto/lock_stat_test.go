package proto

import (
"testing"
"fmt"
)

func Test_srv_lock_stat(ts *testing.T){
	var stat Lock_stat;
	stat.New();
	stat.head.LockID = 0x1; 
	pack, _ := stat.Encode();

	ts.Logf("构造包len[%d]: %X\n", len(pack), pack);
	
	Hex_Dump(pack, len(pack));

	ts.Logf("请求门锁状态查询\n");
	var srv SRV;
	srv.pack = pack;
	srv.srv(ts);	
}

func Test_lock_stat(ts *testing.T){
	var stat Lock_stat;
	stat.New();
	stat.head.LockID = 0x2; 
	pack, _ := stat.Encode();

	currect := "0202038091300323";
	if fmt.Sprintf("%X", pack) == currect{
		ts.Logf("构造包len[%d]: %X\n", len(pack), pack);
	}else{
		ts.Errorf("构造包不正确%s", currect);
	}
	
	Hex_Dump(pack, len(pack));
}
