package proto

import (
"testing"
"fmt"
"time"
)

func Test_srv_user_add_passwd_forever(ts *testing.T){
	var cmd User_add_passwd;
	cmd.New();

	cmd.Command.head.LockID = 0x1;
	cmd.Passwd_len = [1]byte{0x6};
	cmd.Passwd = [11]byte{0x03, 0x09, 0x01, 0x07, 0x03, 0x09};

	pack, _ := cmd.Encode();

	fmt.Println("发送数据包: ");
	Hex_Dump(pack, len(pack));

	ts.Logf("请求增加用户(密码)\n");
	var srv SRV;
	srv.pack = pack;
	srv.srv(ts);
}

func Test_srv_user_add_passwd_for_time(ts *testing.T){
	var cmd User_add_passwd;
	cmd.New();

	cmd.Command.head.LockID = 0x1;
	cmd.Passwd_len = [1]byte{0x6};
	cmd.Passwd = [11]byte{0x02, 0x05, 0x08, 0x01, 0x02, 0x03};

	tn := time.Now();
	cmd.Period[0] = (uint8)(tn.Year() - 2000);
	cmd.Period[1] = (uint8)(tn.Month());
	cmd.Period[2] = (uint8)(tn.Day());
	cmd.Period[3] = (uint8)(tn.Hour());
	cmd.Period[4] = (uint8)(tn.Minute());
	cmd.Period[5] = (uint8)(tn.Second());

	tn_end := time.Now().Add(time.Hour * 24);
	cmd.Period[6] = (uint8)(tn_end.Year() - 2000);
	cmd.Period[7] = (uint8)(tn_end.Month());
	cmd.Period[8] = (uint8)(tn_end.Day());
	cmd.Period[9] = (uint8)(tn_end.Hour());
	cmd.Period[10] = (uint8)(tn_end.Minute());
	cmd.Period[11] = (uint8)(tn_end.Second());

	pack, _ := cmd.Encode();

	fmt.Println("发送数据包: ");
	Hex_Dump(pack, len(pack));

	ts.Logf("请求增加用户(密码)\n");
	var srv SRV;
	srv.pack = pack;
	srv.srv(ts);
}

