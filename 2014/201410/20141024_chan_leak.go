package main

import(
"runtime/pprof"
"log"
"time"
"flag"
"os"
)

var t = flag.Int("t", 5, "sleep time");
var Cpuprf = flag.String("c", "cpu.prof", "cpu profile");
var Memprf = flag.String("m", "mem.prof", "mem profile");

func worker(c <-chan int) {
    var i int

    for {
        i += <-c
    }
}

func wrapper() {
    c := make(chan int)

    go worker(c)

    for i := 0; i < 0xff; i++ {
        c <- i
    }
}

var CpuFile *os.File;
var MemFile *os.File;

func main() {
    flag.Parse();

    if *Cpuprf != ""{
        var err error;
        CpuFile, err = os.Create(*Cpuprf);
        if err != nil{
            log.Fatal("创建CPU文件", err);
        }
        if err = pprof.StartCPUProfile(CpuFile); err != nil{
            log.Fatal("写入CPU文件", err);
        }
        defer pprof.StopCPUProfile();
    }    

    if *Memprf != "" {
        var err error;
        MemFile, err = os.Create(*Memprf);
        if err != nil {
            log.Fatal("创建内存文件", err);
        } 
        // runtime.GC();
        if err = pprof.WriteHeapProfile(MemFile); err != nil{
            log.Fatal("写入内存文件", err);
        }
        defer MemFile.Close();
    }
    for ix :=0; ix < 10; ix++ {
        time.Sleep((time.Duration)(*t)*time.Second);
        wrapper()
    }
}