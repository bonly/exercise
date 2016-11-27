package main

// #include <stdio.h>
// #include <sys/types.h>
// #include <sys/stat.h>
// #include <stdlib.h>
// #include <string.h>
// #include <mysql.h>
// #cgo CFLAGS: -I/opt/mysql/5.7.9-2/include -fabi-version=2 -fno-omit-frame-pointer
import "C"

import (
//"os"
"fmt"
"github.com/garyburd/redigo/redis"
// "errors"
)

var Cli redis.Conn;

//export sync_redis_init
func sync_redis_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.my_bool {
	if args.arg_count != 2 {
		msg := "sync_redis(table, pkey) requires two string argument\n";
		C.strcpy(message, C.CString(msg));
		return 1;
	}

	var err error;
	Cli, err = redis.Dial("tcp", "192.168.1.13:6379");
	if  err != nil{
		fmt.Println("connect fail: ", err);
		return 2;
	}	
	return 0;
}

//export sync_redis
func sync_redis(initid *C.UDF_INIT, args *C.UDF_ARGS, result *C.char, length *C.ulong, is_null *C.char, error *C.char) *C.char  {
	
	if args.arg_count != 2 {
		msg := "sync_redis(table, pkey) requires two string argument\n";
		C.strcpy(table, C.CString(msg));
		return 1;
	}

	filename := C.GoString(*args.args);
	//写入
	if _, err := Cli.Do("SET", table, field); err != nil{
		fmt.Println("set fail: ", err);
		return 1;
	}
	return 0;
	
  // out, err := exec.Command(C.GoString(*args.args)).Output()
  // if err != nil {
  //   fmt.Println(err)
  //   os.Exit(1)
  // }
  // result = C.CString(string(out))
  // *length = C.ulong(utf8.RuneCountInString(C.GoString(result)))
  // return result	
}

//export sync_redis_deinit
func sync_redis_deinit(initid *C.UDF_INIT) {
    Cli.Close();
}

func main() {}

/*
export CPATH=/home/opt/maria/include/mysql
go build -buildmode=c-shared -o libsync_redis.so 20131208_mysql_ufd.go
create function sync_redis returns STRING soname 'libsync_redis.so';
select sync_redis("abc", "field");
*/