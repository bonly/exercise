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
// "log"
"log"
"unicode/utf8"
// "errors"
// "flag"
"unsafe"
)

var Cli redis.Conn;
var logFile *os.File;
var debugLog *log.Logger;

//export sync_redis_init
func sync_redis_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.my_bool {
	// flag.Parse();
	// flag.Set("alsologtostderr", "true");
	// flag.Set("log_dir", "/tmp");

	var err error;
	fileName := "/tmp/xxx_debug.log";
  	// logFile, _ = os.Create(fileName);
  	logFile, _ = os.OpenFile(fileName,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666);
  	defer logFile.Sync();
    // defer logFile.Close();
   //  if err != nil {
   //      log.Fatalln("open file error !");
   //  }
    debugLog = log.New(logFile,"[Debug]", log.Llongfile);
    debugLog.Println("A debug message here");
    debugLog.SetPrefix("[Info]");
    debugLog.Println("A Info Message here ");
    debugLog.SetFlags(debugLog.Flags() | log.LstdFlags);
    debugLog.Println("A different prefix");

    // logFile.Write([]byte("test"));

	if args.arg_count != 2 {
		msg := "sync_redis(table, pkey) requires two string argument\n";
		C.strcpy(message, C.CString(msg));
		return 1;
	}

	// Cli, err = redis.Dial("tcp", "192.168.1.13:6379");
	// if Cli {
	// 	debugLog.Println("had connect");
	// }

	Cli, err = redis.Dial("tcp", "127.0.0.1:6379");
	if  err != nil{
		C.strcpy(message, C.CString(fmt.Sprintf("connect fail: ", err)));
		return 2;
	}	
	debugLog.Println("connect success");

	return 0;
}

//export sync_redis
func sync_redis(initid *C.UDF_INIT, arg *C.UDF_ARGS, result *C.char, 
	length *C.ulong, is_null *C.char, error *C.char) *C.char  {
	defer logFile.Sync();

	argc := int(arg.arg_count); //参数的个数  *arg.lengths是单个参数的长度
	argp := arg.args;

	if argc != 2{
		debugLog.Println("param err: ", argc);
		result = C.CString(string("param err"));
		*length = C.ulong(utf8.RuneCountInString(C.GoString(result)));		
		return result;
	}
	debugLog.Println("argv count: ", argc);

	argv := (*(*[1<<30]*C.char)(unsafe.Pointer(argp)))[:argc];//指向一个足够大的数组空间
	args := make([]string, argc);//建存放数组的地方
	for i := 0; i < argc; i++ {  //跌代访问数组个数
		args[i] = C.GoString(argv[i]); //转存到go中
	}
	for i := 0; i < len(args); i++ { //输出各个参数
		debugLog.Printf("argv[%d]: %s\n", i, args[i]);
	}
	debugLog.Println("in func");	

	if _, err := Cli.Do("SET", args[0], args[1]); err != nil{
		debugLog.Println("set fail: ", err);
		result = C.CString(err.Error());
		*length = C.ulong(utf8.RuneCountInString(C.GoString(result)));
		return result;		
	}

	result = C.CString(string("ok"));
	*length = C.ulong(utf8.RuneCountInString(C.GoString(result)));
	return result;
}

//export sync_redis_deinit
func sync_redis_deinit(initid *C.UDF_INIT) {
	log.Println("destory");
    Cli.Close();
}

func main() {}

/*
export CPATH=/home/opt/maria/include/mysql
go build -buildmode=c-shared -o libsync_redis.so mysql_udf.go
cp libsync_redis.so /home/opt/maria/lib/plugin/
create function sync_redis returns STRING soname 'libsync_redis.so';
select sync_redis("abc", "field");
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
