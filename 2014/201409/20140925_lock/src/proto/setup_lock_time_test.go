package proto

import (
"testing"
"fmt"
"time"
"encoding/hex"
)

func Test_setup_lock_time(ts *testing.T){
	var cmd Setup_lock_time;
	cmd.New();
	cmd.head.LockID = 0x1; 

	cmd.DateTime[0] = 15;
	cmd.DateTime[1] = 5;
	cmd.DateTime[2] = 30;
	cmd.DateTime[3] = 22;
	cmd.DateTime[4] = 40;
	cmd.DateTime[5] = 10;

	cmd.verify = false;

	fmt.Printf("BCD: %s\n", hex.EncodeToString(cmd.DateTime[:]));//转成BCD显示
	bcd, _ := hex.DecodeString("23");
	fmt.Printf("BCD: %X\n", bcd);

	pack, _ := cmd.Encode();

	currect := "0102098090350F051E16280A0D";
	if fmt.Sprintf("%X", pack) == currect{
		ts.Logf("构造包len[%d]: %X\n", len(pack), pack);
	}else{
		ts.Errorf("构造包不正确,应为: %s", currect);
	}
	
	Hex_Dump(pack, len(pack));
}

func Test_decode_setup_lock_time(ts *testing.T){
	pack := []byte{0x01, 0x02, 0x09, 0x80,
			 0x90, 0x35, 0x0f, 0x05,
			 0x1e, 0x16, 0x28, 0x0a, 0x0d};

	var cmd Setup_lock_time;
	cmd.Decode(pack);

	ts.Logf("数据包为：%+v\n", cmd);
	
	Hex_Dump(pack, len(pack));
}

func Test_srv_setup_lock_time(ts *testing.T){
	var cmd Setup_lock_time;
	cmd.New();
	cmd.head.LockID = 0x1; 

	tn := time.Now();
	cmd.DateTime[0] = (uint8)(tn.Year() - 2000);
	cmd.DateTime[1] = (uint8)(tn.Month());
	cmd.DateTime[2] = (uint8)(tn.Day());
	cmd.DateTime[3] = (uint8)(tn.Hour());
	cmd.DateTime[4] = (uint8)(tn.Minute());
	cmd.DateTime[5] = (uint8)(tn.Second());

	cmd.verify = false;

	pack, _ := cmd.Encode();

	ts.Logf("构造包len[%d]: %X\n", len(pack), pack);

	Hex_Dump(pack, len(pack));

	ts.Logf("请求设置锁的时间\n");
	var srv SRV;
	srv.pack = pack;
	srv.srv(ts);	
}

func Test_setup_lock_for_time_bcd_bad(ts *testing.T){
	var cmd Setup_lock_time;
	cmd.New();
	cmd.head.LockID = 0x1; 

	cmd.DateTime[0] = 15;
	cmd.DateTime[1] = 5;
	cmd.DateTime[2] = 30;
	
	tn := time.Now();
	hour, _ := hex.DecodeString(fmt.Sprintf("%02d", tn.Hour()));
	cmd.DateTime[3] = (uint8)(hour[0]);
	minute, _ := hex.DecodeString(fmt.Sprintf("%02d", tn.Minute()));
	cmd.DateTime[4] = (uint8)(minute[0]);
	second, _ := hex.DecodeString(fmt.Sprintf("%02d", tn.Second()));
	cmd.DateTime[5] = (uint8)(second[0]);

	cmd.verify = false;

	pack, _ := cmd.Encode();

	ts.Logf("构造包len[%d]: %X\n", len(pack), pack);
	
	Hex_Dump(pack, len(pack));
}