package proto

import (
"testing"
"fmt"
)

func Test_srv_user_del_passwd_forever(ts *testing.T){
	var cmd User_del_passwd;
	cmd.New();

	cmd.Command.head.LockID = 0x1;
	cmd.Passwd_len = [1]byte{0x6};
	cmd.Passwd = [11]byte{0x03, 0x09, 0x01, 0x07, 0x03, 0x09};

	pack, _ := cmd.Encode();

	fmt.Println("发送数据包: ");
	Hex_Dump(pack, len(pack));

	ts.Logf("请求删除用户(密码)\n");
	var srv SRV;
	srv.pack = pack;
	srv.srv(ts);
}

func Test_srv_user_del_passwd_for_time(ts *testing.T){
	var cmd User_del_passwd;
	cmd.New();

	cmd.Command.head.LockID = 0x1;
	cmd.Passwd_len = [1]byte{0x6};
	cmd.Passwd = [11]byte{0x02, 0x05, 0x08, 0x01, 0x02, 0x03};

	pack, _ := cmd.Encode();

	fmt.Println("发送数据包: ");
	Hex_Dump(pack, len(pack));

	ts.Logf("请求增加用户(密码)\n");
	var srv SRV;
	srv.pack = pack;
	srv.srv(ts);
}