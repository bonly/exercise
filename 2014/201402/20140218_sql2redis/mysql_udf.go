package main

// #include <stdio.h>
// #include <sys/types.h>
// #include <sys/stat.h>
// #include <stdlib.h>
// #include <string.h>
// #include <mysql.h>
// #cgo CFLAGS: -I/home/opt/maria/include -fabi-version=2 -fno-omit-frame-pointer
import "C"

import (
"os"
// "os/exec"
"fmt"
"github.com/garyburd/redigo/redis"
"encoding/json"
"log"
"unicode/utf8"
// "errors"
// "flag"
"unsafe"
)

type Config struct{
	Redis_srv string;
};

var cfg Config;
var Cli redis.Conn;

//export sync_redis_init
func sync_redis_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.my_bool {
	log.Println("Begin Sync to redis");
	
	var err error;
  	fl, err := os.OpenFile("sql2redis.json",os.O_CREATE|os.O_RDONLY,0666);
  	if err != nil{
  		log.Println("open config file: ", err.Error());
  		return 1;
  	}
  	defer func(){
  		fl.Close();
  	}();

  	jsp := json.NewDecoder(fl);
  	if err = jsp.Decode(&cfg); err != nil{
  		log.Println("parse config file: ", err.Error());
  		return 1;
  	}

  	log.Println("redis_srv: ", cfg.Redis_srv);

	if (args.arg_count < 2)||(args.arg_count % 2 != 0) {
		msg := "sync_redis(table, pkey) requires two string argument\n";
		log.Println(msg);
		C.strcpy(message, C.CString(msg));
		return 1;
	}

	Cli, err = redis.Dial("tcp", cfg.Redis_srv);
	if  err != nil{
		C.strcpy(message, C.CString(fmt.Sprintf("connect fail: ", err)));
		log.Println(err);
		return 2;
	}	
	log.Println("connect success");

	return 0;
}

//export sync_redis
func sync_redis(initid *C.UDF_INIT, arg *C.UDF_ARGS, result *C.char, 
	length *C.ulong, is_null *C.char, error *C.char) *C.char  {

	argc := int(arg.arg_count); //参数的个数  *arg.lengths是单个参数的长度
	// argp := arg.args;

	// if argc != 2{
	// 	debugLog.Println("param err: ", argc);
	// 	result = C.CString(string("param err"));
	// 	*length = C.ulong(utf8.RuneCountInString(C.GoString(result)));		
	// 	return result;
	// }
	// debugLog.Println("argv count: ", argc);

	argv := (*(*[1<<30]*C.char)(unsafe.Pointer(arg.args)))[:argc];//指向一个足够大的数组空间
	args := make([]string, argc);//建存放数组的地方
	for i := 0; i < argc; i++ {  //跌代访问数组个数
		args[i] = C.GoString(argv[i]); //转存到go中
	}
	for i := 0; i < len(args); i++ { //输出各个参数
		log.Printf("args[%d]: %s\n", i, args[i]);
		if (i != 0) && ((i+1) % 2 == 0){
			if _, err := Cli.Do("SET", args[i-1], args[i]); err != nil{
				log.Println("set fail: ", err);
				result = C.CString(err.Error());
				*length = C.ulong(utf8.RuneCountInString(C.GoString(result)));
				return result;		
			}
			log.Printf("set [%s]: %s\n", args[i-1], args[i]);		
		}
	}

	result = C.CString(string("ok"));
	*length = C.ulong(utf8.RuneCountInString(C.GoString(result)));
	return result;
}

//export sync_redis_deinit
func sync_redis_deinit(initid *C.UDF_INIT) {
	log.Println("End Sync to redis");
    Cli.Close();
}

func main() {}

/*
export CPATH=/home/opt/maria/include/mysql
go build -buildmode=c-shared -o libsync_redis.so mysql_udf.go
cp libsync_redis.so /home/opt/maria/lib/plugin/
create function sync_redis returns STRING soname 'libsync_redis.so';
select sync_redis("abc", "field");
mysqlcheck -r -q -B mysql
*/

func cpp2go(argc C.int, argv_ **C.char) {
	//指向一个足够大的数组空间
    argv := (*(*[1 << 30]*C.char)(unsafe.Pointer(argv_)))[:int(argc)];
    args := make([]string, int(argc));//建存放数组的地方
    for i := 0; i < int(argc); i++ {//跌代访问数组个数
        args[i] = C.GoString(argv[i]);//转存到go中
    }
    for i := 0; i < len(args); i++ { //输出各个参数
        fmt.Printf("argv[%d]: %s\n", i, args[i])
    }
}
