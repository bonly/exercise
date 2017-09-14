package proto

import (
"testing"
"fmt"
)

func Test_srv_user_del_all(ts *testing.T){
	var cmd User_del_all;
	cmd.New();

	cmd.Command.head.LockID = 0x1;

	pack, _ := cmd.Encode();

	fmt.Println("发送数据包: ");
	Hex_Dump(pack, len(pack));

	ts.Logf("请求删除所有用户\n");
	var srv SRV;
	srv.pack = pack;
	srv.srv(ts);
}
