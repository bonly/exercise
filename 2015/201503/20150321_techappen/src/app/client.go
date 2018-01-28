package main 

/*
#cgo LDFLAGS: -pthread -fPIC
*/
import "C"

import (
	"Srv"
	"log"
	"User"
	"reflect"
	"Command"
)

func init(){
	log.Printf("核心加载\n");
}

func main(){
}

//export Run
func Run(){
	Srv.Start();
}

//export Exit
func Exit(){
	Srv.End();
}

//export Proc
func Proc(c_scope *C.char, c_data *C.char){
	scp := C.GoString(c_scope);
	data := C.GoString(c_data);
	var method reflect.Value;

	var cmd Command.ICmd; //解包分析包头
	cmd = &Command.TCmd{};
	cmd.(*Command.TCmd).Data = &User.Q_User{};
	err := cmd.Decode([]byte(data));
	if err != nil{
		log.Printf("数据错误: %v\n", err);
		return;
	}
	log.Printf("data %v \n", cmd.(*Command.TCmd).Data);

	var idata interface{};
	switch scp{
		case "User":
		   User.Start(); //初次调用时初始化连接设置
		   method = User.Fnc.MethodByName(cmd.(*Command.TCmd).Func);
		   idata = cmd.(*Command.TCmd).Data.(*User.Q_User);
		   break;
		 default:
		   log.Printf("未定义: %s\n", scp);
		   return;
	}

	if method.IsValid(){
		vdata := reflect.ValueOf(idata);
		method.Call([]reflect.Value{0:vdata});
	}else{
		log.Printf("无此方法: %v\n", cmd.(*Command.TCmd).Func);
		return;
	}
}