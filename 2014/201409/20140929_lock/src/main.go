/*
auth: bonly
create: 2016.9.20
desc: 主程序
*/
package main 

import(
"flag"
"manage"
"oms"
"log"
"config"
"runtime/pprof"
"runtime"
"os"
"fmt"
"syscall"
"os/signal"
"sync"
)

var Version string;
var Code string;
var App_time string;

var Ver *bool = flag.Bool("V", false, "版本");

func main(){
	flag.Parse();

	if *Ver{ //版本打印
		fmt.Println("Version: ", Version);
		fmt.Println("Code: ", Code);
		fmt.Println("Time: ", App_time);
		return;
	}

    if *config.Cpuprf != ""{
    	var err error;
    	config.CpuFile, err = os.Create(*config.Cpuprf);
    	if err != nil{
    		log.Fatal("创建CPU文件", err);
    	}
    	if err = pprof.StartCPUProfile(config.CpuFile); err != nil{
    		log.Fatal("写入CPU文件", err);
    	}
    	defer pprof.StopCPUProfile();
    }

	if *config.Memprf != "" {
        var err error;
        config.MemFile, err = os.Create(*config.Memprf);
        if err != nil {
            log.Fatal("创建内存文件", err);
        } 
        runtime.GC();
        if err = pprof.WriteHeapProfile(config.MemFile); err != nil{
        	log.Fatal("写入内存文件", err);
        }
        defer config.MemFile.Close();
    }

	runtime.GOMAXPROCS(runtime.NumCPU());

	/// 建立要收集的信号
	sigs := make(chan os.Signal, 1);
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM);

	go func(){
		var pms oms.SRV;
		pms.Srv();
	}();

	go func(){
		var box manage.SRV;
		box.Srv();
	}();

	var Wg sync.WaitGroup;
	Wg.Add(1);
	go func(){
		defer Wg.Done();
		for{
		  sig := <-sigs;
		  process_sig(sig);
		  if config.Run == false {
		    break;
		  }
		}
	}();	
	Wg.Wait();	
}

/// 处理系统信号
func process_sig(sig os.Signal){
    switch(sig){
        case syscall.SIGINT:
           log.Println("Get SigInt");
           config.Run = false; 
           break;
        case syscall.SIGTERM:
           log.Println("Get SigTerm");
           config.Run = false;
           break;
        case syscall.SIGUSR1:
            log.Println("Get SIGUSR1.");
            config.Run = true;
            break;
        default:
           log.Println("Get Sig: ", sig);
           config.Run = true;
           break;
    }
}