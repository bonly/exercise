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

var REMOTE_OPEN_DOOR_CMD_PARAM = []byte{0x90, 0x36};

type Remote_open_door struct{
	Command;
	Passwd []uint8;
};

func (this *Remote_open_door)New()(cmd interface{}){
	this.Command.New();
	this.head.Type = SRV_2_LOCK;
	this.head.Cmd = REMOTE_OPEN_DOOR_CMD_PARAM[0];
	this.head.Param = REMOTE_OPEN_DOOR_CMD_PARAM[1];

	return this;
}

func (this *Remote_open_door)Encode()(pack []byte, err error){
	this.Command.Add_data(this.Passwd);

	return this.Command.Encode();
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