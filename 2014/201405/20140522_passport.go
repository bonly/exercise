package main 

/*
#cgo pkg-config: opencv
#include <opencv/cv.h>
#include "20140522_passport.c"
*/
import "C"

import(
"unsafe"
"fmt"
"os"
"github.com/otiai10/gosseract"
)

func main(){
	//删除旧文件
	// if err := os.Remove("roi*.jpg"); err != nil{
	// 	fmt.Println("rm failed: ", err);
	// }

	argc := len(os.Args);
	argv := make([](*C.char), argc);  //C语言char*指针创建切片
	
	for i, arg := range os.Args{
		argv[i] = C.CString(string(arg));
	}
    
	C.Main(C.int(argc), (**C.char)(unsafe.Pointer(&argv[0])));

	var txt string;
	for idx := 0; idx <= 30; idx++{
		fm := fmt.Sprintf("%s%d.jpg","roi", idx);
		_, err := os.Stat(fm);
		if err != nil{
			continue;
		}
		out := gosseract.Must(gosseract.Params{
			Src:fm, 
			Languages:"chi_sim+eng"});
		// fmt.Println(out);
		txt += out + "\n";
	}
	fmt.Println(txt);
}

//http://blog.giorgis.io/cgo-examples