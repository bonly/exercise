package main

import (
"runtime/pprof"
"os"
"log"
)

func main(){
	//代码加载需要些heap memory profiler 日志的地方
    f, err := os.Create("HeapProfile.txt")
    if err != nil {
        log.Fatal(err)
    }
    pprof.WriteHeapProfile(f)
    f.Close()
}
/*
启动程序后， 运行到这个地方会生成 HeapProfile.txt 文件
go tool pprof  -text  -inuse_objects  20140927_pprof HeapProfile.txt

进入交互模式，可以输入top， peek 命令那个查看更详细一些
go tool pprof  20140927_pprof  HeapProfile.txt

GODEBUG='gctrace=1' 20140927_pprof 可将gc信息列印出来


*/
