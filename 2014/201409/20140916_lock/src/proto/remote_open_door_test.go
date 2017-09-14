package proto

import (
"testing"
"fmt"
)

func Test_e_remote_open_door(ts *testing.T){
	var rop Remote_open_door;
	rop.New();

	rop.Passwd = []byte{0x01, 0x02, 0x03, 0x04, 0x05};

	pack, _ := rop.Encode();

	if fmt.Sprintf("%X", pack) == "01020880903601020304050D00" {
		ts.Logf("构造包len[%d]: %X", len(pack), pack);
	}else{
		ts.Errorf("构造包不正确%X", pack);
	}
}

func Test_e_remote_open_door_no_passwd(ts *testing.T){
	var rop Remote_open_door;
	rop.New();

	pack, _ := rop.Encode();

	if fmt.Sprintf("%X", pack) == "0102038090360D00" {
		ts.Logf("构造包len[%d]: %X", len(pack), pack);
	}else{
		ts.Errorf("构造包不正确%X", pack);
	}
}

func Test_d_r_remote_open_door(ts *testing.T){
	smp := []byte{0x01, 0x02, 0x03, 0x86, 0x90, 0x36, 0x0d, 0x00};

	var rop R_Remote_open_door;

	// cmd := rop.Decode(smp).(R_Remote_open_door);
	cd := rop.Decode(smp);
	if cd == nil{
		ts.Fatalf("数据包不正确%x", smp);
	}
	cmd := cd.(R_Remote_open_door)
	if cmd.head.Len == 3{
		ts.Logf("解释包len[%d] data: %+v\n", len(smp), cmd);
	}else{
		ts.Errorf("解释包不正确%d", cmd.head.Len);
	}
}