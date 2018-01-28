package main 

/*
#cgo LDFLAGS: -pthread -fPIC
*/
import "C"

import (
	"srv"
	"log"
	"user"
	// "reflect"
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
func Proc(c_name *C.char, c_data *C.char){
	scp := C.GoString(c_name);
	data := C.GoString(c_data);

	var cmd proto.ICmd; //解包处理
	cmd = &proto.Cmd{};
	err := cmd.Decode(scp, []byte(data), &(cmd.(*proto.Cmd).Req));
	if err != nil{
		log.Printf("数据错误: %v\n", err);
		return;
	}

	log.Printf("cmd: %#v \n", cmd.(*proto.Cmd).Req.(*user.Q_User));

	// var idata interface{};
	// switch scp{
	// 	case "User":
	// 	   User.Start(); //初次调用时初始化连接设置
	// 	   method = User.Fnc.MethodByName(cmd.(*Command.TCmd).Func);
	// 	   idata = cmd.(*Command.TCmd).Data.(*User.Q_User);
	// 	   break;
	// 	 default:
	// 	   log.Printf("未定义: %s\n", scp);
	// 	   return;
	// }

	// if method.IsValid(){
	// 	vdata := reflect.ValueOf(idata);
	// 	method.Call([]reflect.Value{0:vdata});
	// }else{
	// 	log.Printf("无此方法: %v\n", cmd.(*Command.TCmd).Func);
	// 	return;
	// }
}