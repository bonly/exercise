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

// func Test_lock_init_marsh(ts *testing.T){
// 	lock := &Lock_Init{
// 		Lock:Lock{Appid:"123",Sign:"sssiiigggnnn"},
// 		Lock_id:111111,
// 	};
// 	str, err := json.MarshalIndent(lock, " ", " ");
// 	if err != nil{
// 		fmt.Println(err);
// 		return;
// 	}
// 	fmt.Println(string(str));

// 	// dat, err := ToMap(lock, "json");
// 	// fmt.Println(Gen_hmac(&dat));
// }

// func Test_lock_init_unmarsh(ts *testing.T){
// 	lock_str := `
// {
//   "appid": "123",
//   "sign": "sssiiigggnnn",
//   "lock_id": 111111
// }
// 	`;
// 	var lock Lock_Init;
// 	err := json.Unmarshal(([]byte)(lock_str), &lock);
// 	if err != nil{
// 		fmt.Println(err);
// 		return;
// 	}
// 	fmt.Printf("%#v", lock);
// }

// func Test_lock_init_encode(ts *testing.T){
// 	lock := &Lock_Init{
// 		Lock_id:111111,
// 	};

// 	by, err := lock.Marshal("app_xbed", "b64fe7cd7ea973c5077ce90aa675c12d");
// 	if err != nil{
// 		fmt.Println(err);
// 		return;
// 	}
// 	fmt.Printf("\nData: %s", string(by));
// }

func Test_lock_init_send(ts *testing.T){
	lock := &Lock_Init{
		Lock_id:1000001,
	};

	by, err := lock.Marshal("app_xbed", "b64fe7cd7ea973c5077ce90aa675c12d");
	if err != nil{
		fmt.Println(err);
		return;
	}

	Post(*Srv, "/third/lock/init/forth", string(by));
}