/*
auth: bonly
create: 2016.9.14
desc: 远程开门
*/
package proto

import(
// "log"
// "encoding/binary"
// "bytes"
);

type Remote_open_door struct{
	Command;
	Passwd []uint8;
};

func (this *Remote_open_door)New()(cmd interface{}){
	this.Command.New();
	this.head.LockID = 0x1; //todo 外部设置
	this.head.Type = SRV_2_LOCK;
	this.head.Cmd = 0x90;
	this.head.Param = 0x36;

	return this;
}

func (this *Remote_open_door)Encode()(pack []byte, err error){
	this.Command.Add_data(this.Passwd);

	return this.Command.Encode();;
}

type R_Remote_open_door struct{
	Command;
};

func (this *R_Remote_open_door)Decode(pack []byte)(cmd interface{}){
	cmd = this.Command.Decode(pack);
	if cmd != nil{
		return *this;
	}else{
		return nil;
	}
	
}

/*
    defer func() {
        if err := recover(); err != nil {
            fmt.Println(err);
            return;
        }
    }();	
*/