package main 

/*
#cgo LDFLAGS: -pthread -fPIC
*/
import "C"

import (
	srv "cli"
	"log"
	_ "cli/user"
	"reflect"
	"proto"
)

func init(){
	log.Printf("核心加载\n");
}

func main(){
}

//export Run
func Run(){
	srv.Start();
}

//export Exit
func Exit(){
	srv.End();
}

//export Proc
func Proc(c_json *C.char)(ret *C.char){
	data := C.GoString(c_json);

	var cmd proto.ICmd; //解包处理
	cmd = &proto.Cmd{};
	err := cmd.Decode([]byte(data));
	if err != nil{
		log.Printf("数据错误: %v\n", err);
		return C.CString("{}");
	}

	// log.Printf("cmd: %#v \n", cmd.(*proto.Cmd).Data.(*user.Q_User));

	method := cmd.(*proto.Cmd).Key.Fnc.MethodByName(cmd.(*proto.Cmd).Key.Func);
	if method.IsValid(){
		vdata := reflect.ValueOf(cmd.(*proto.Cmd).Data); //Data结构转换成参数
		method.Call([]reflect.Value{0:vdata});
	}else{
		log.Printf("无此方法: %v\n", cmd.(*proto.Cmd).Key.Func);
		return C.CString("{}");
	}
	return C.CString("{}");
}
