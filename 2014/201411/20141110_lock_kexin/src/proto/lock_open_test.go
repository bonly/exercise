/*
auth: bonly
create: 2015.10.27
*/
package proto

import(
"testing"
// "encoding/json"
"fmt"
)

func Test_lock_open_send(ts *testing.T){
	lock := &Lock_Open{
		Lock_id:1000003,
	};

	by, err := lock.Marshal("app_xbed", "b64fe7cd7ea973c5077ce90aa675c12d");
	if err != nil{
		fmt.Println(err);
		return;
	}

	Post(*Srv, "/third/lock/open", string(by));
}