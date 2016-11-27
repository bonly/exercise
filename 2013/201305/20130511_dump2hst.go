package main

/*
#cgo CPPFLAGS: -g
#cgo LDFLAGS: -l dump2hst -L./
#include "20130511_dump2hst.h"
#include <stdlib.h>  //need for free
*/
import "C"
import "unsafe"
import "os"
import "path/filepath"
import "log"

func main(){
    // arg := os.Args;
    // Test1(arg);
    proc_dir(os.Args[1]);
}

func proc_dir(dir string){
    err := filepath.Walk(dir,
        func (path string, fl os.FileInfo, err error) error{
            if fl == nil{
                log.Panic(err);
                return err;
            }
            if (fl.IsDir()){
                return nil;
            }
            proc_file (path);
            return nil;
        });
    if err != nil{
        log.Println("path err: ", err);
    }
}

func proc_file( ifile string){
    arg := make([](*C.char), 0); //c的char**
    
    tmain := C.CString("main");  //c的string
    defer C.free(unsafe.Pointer(tmain));
    ptr := (*C.char)(unsafe.Pointer(tmain));  // c的char*
    arg = append(arg, (*C.char)(unsafe.Pointer(ptr))); //add 到 char**中
    
    tfile := C.CString(ifile);
    defer C.free(unsafe.Pointer(tfile));
    ptr = (*C.char)(unsafe.Pointer(tfile));
    arg = append(arg, (*C.char)(unsafe.Pointer(ptr)));
    
    C.cmain(C.int(len(arg)), (**C.char)(unsafe.Pointer(&arg[0])));
}

func Test(args[] string){
   arg := make([](*_Ctype_char), 0)  //C语言char*指针创建切片
   l := len(args)
   for i,_ := range args{
       char := C.CString(args[i])
       defer C.free(unsafe.Pointer(char)) //释放内存
       strptr := (*_Ctype_char)(unsafe.Pointer(char))
       arg = append(arg, strptr)  //将char*指针加入到arg切片
   }
                                                                                  
   C.cmain(C.int(l), (**_Ctype_char)(unsafe.Pointer(&arg[0])))  //即c语言的main(int argc,char**argv)
}

func Test1(args[] string){
   arg := make([](*C.char), 0)  //C语言char*指针创建切片
   l := len(args)
   for i,_ := range args{
       char := C.CString(args[i])
       defer C.free(unsafe.Pointer(char)) //释放内存
       strptr := (*C.char)(unsafe.Pointer(char))
       arg = append(arg, strptr)  //将char*指针加入到arg切片
   }
                                                                                  
   C.cmain(C.int(l), (**C.char)(unsafe.Pointer(&arg[0])))  //即c语言的main(int argc,char**argv)
}

/*
go build -v -gcflags="-N -l" 20130511_dump2hst.go 
*/
