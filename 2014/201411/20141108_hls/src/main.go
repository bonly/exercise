/*
auth: bonly
create: 2016.9.20
desc: 主程序
todo: 等所有线程停止了再结束
*/
package main 

import(
"flag"
"manage"
"oms"
log "glog"
"config"
"runtime/pprof"
"runtime"
"os"
"fmt"
"syscall"
"os/signal"
"sync"
"callback"
)

var Version string;
var Code string;
var App_time string;

var Ver *bool = flag.Bool("V", false, "版本");

func init(){
  flag.Set("alsologtostderr", "true");
  flag.Set("v", "99");
  flag.Set("log_dir", "./");  
}

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

  log.Info("门锁服务起动"); 
  defer log.Info("门锁服务结束"); 
	runtime.GOMAXPROCS(runtime.NumCPU());

	/// 建立要收集的信号
	sigs := make(chan os.Signal, 1);
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM);

  var Wg sync.WaitGroup;
	go func(){
		var pms oms.SRV;
		pms.Srv();
	}();

  cbc := callback.Run();
	
  go func(){
		var box manage.SRV;
		box.Srv(cbc);
	}();

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
           log.Info("Get SigInt");
           config.Run = false; 
           break;
        case syscall.SIGTERM:
           log.Info("Get SigTerm");
           config.Run = false;
           break;
        // case syscall.SIGUSR1:
        //     log.Info("Get SIGUSR1.");
        //     config.Run = true;
        //     break;
        default:
           log.Info("Get Sig: ", sig);
           config.Run = true;
           break;
    }
}