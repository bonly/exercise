package main

import (
"fmt"
"syscall"
"unsafe"
)

func Call_dll(){
    var resultCode int;
    user32 := syscall.NewLazyDLL("user32.dll");
    endDialog := user32.NewProc("EndDialog");
    endDialog.Call(uintptr(resultCode));
    fmt.Println("end");
}


func Init(){
   dll := syscall.NewLazyDLL("RayNetSdk.dll");
//    defer syscall.FreeLibrary(dll);
   
   fmt.Println("dll name:", dll.Name);

   prc := dll.NewProc("_COM_AVD_DEV_Init@0");
   prc.Call();    
}

func deInit(){
   dll := syscall.NewLazyDLL("RayNetSdk.dll");
//    defer syscall.FreeLibrary(dll);
   
   fmt.Println("dll name:", dll.Name);

   prc := dll.NewProc("_COM_AVD_DEV_DeInit@0");
   prc.Call();     
}

func main(){
//    win();
//    Call_dll();
//    Drv();
    Init();
    defer deInit();
}


func Drv(){
    dll32 := syscall.NewLazyDLL("kernel32.dll");
    gr := dll32.NewProc("GetDriveTypeA");
    ret, _, _ := gr.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("C:"))));
    fmt.Println(uint16(ret));
}

func win(){
    var mod = syscall.NewLazyDLL("user32.dll");
    var proc = mod.NewProc("MessageBoxW");
    var MB_YESNOCANCEL = 0x00000003;
    
    ret, _, _ := proc.Call(0,
    uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Done Title"))),
    uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("This test is done"))),
    uintptr(MB_YESNOCANCEL));
    fmt.Printf("Return: %d\n", ret);
}

/*
GOARCH=386
*/