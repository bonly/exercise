package main 

import (
"fmt"
"github.com/AkihiroSuda/go-netfilter-queue"
"github.com/google/gopacket/layers"
"os"
"os/exec"
"os/signal"
"syscall"
"sync"
)
var Wg sync.WaitGroup;
var Done = make(chan bool);

func main(){
	build();
	defer clean();

	sigs := make(chan os.Signal, 1);
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM);

	var err error;

	nfq, err := netfilter.NewNFQueue(0, 100, 
		netfilter.NF_DEFAULT_PACKET_SIZE);
	if err != nil{
		fmt.Println("new queue:", err);
		os.Exit(1);
	}
	defer nfq.Close();

	Wg.Add(1);
	go func(){
		defer func(){
			Wg.Done();
			fmt.Println("end sig process");
		}();
		for{
			fmt.Println("wait sig");
			select {
				case done := <-Done:
				    fmt.Println("sig get done: ", done);
					return;
				case sig := <-sigs:
					process_sig(sig);
					// break; //需要继续执行后面的case?其实无执行,自动break
			}
		}
	}();

	Wg.Add(1);
	go func(){
		defer func(){
			Wg.Done();
			fmt.Println("end pack process");
		}();
		
		for packet := range nfq.GetPackets(){ //不知如何中断
			fmt.Printf("数据包: %+v\n", packet);
			ip := packet.Packet.NetworkLayer().(*layers.IPv4);
			fmt.Println("Got SIFF packet for ", ip.DstIP);
			fmt.Println("From ", ip.SrcIP);
			packet.SetVerdict(netfilter.NF_ACCEPT);
		}
	}();

	Wg.Wait();
}

/// 处理系统信号
func process_sig(sig os.Signal){
    switch(sig){
        case syscall.SIGINT:
           fmt.Println("Get SigInt");
           Done <-true;
           break;
        case syscall.SIGTERM:
           fmt.Println("Get SigTerm");
            Done <-true;
           break;
        case syscall.SIGUSR1:
            fmt.Println("Get SIGUSR1. Flash log");
            break;
        default:
           fmt.Println("Get Sig: ", sig);
           break;
    }
}

func clean(){
	//iptables -D OUTPUT -p udp --dst 8.8.8.8 -j QUEUE
	if out, err := exec.Command("iptables",
		"-D", "OUTPUT",
		"-p", "udp",
		"--dst", "8.8.8.8",
		"-j", "QUEUE").Output(); err != nil{
		fmt.Println("del output: ", err);
	}else{
		fmt.Println("delete output success", out);
	}

	//iptables -D INPUT -p udp --src 8.8.8.8 -j QUEUE
	if out, err := exec.Command("iptables",
		"-D", "INPUT",
		"-p", "udp",
		"--src", "8.8.8.8",
		"-j", "QUEUE").Output(); err != nil{
		fmt.Println("del input: ", err);
	}else{
		fmt.Println("delete input success", out);
	}

	//iptables -D INPUT -p icmp -m icmp --icmp-type 11 -j QUEUE
	if out, err := exec.Command("iptables",
		"-D", "INPUT",
		"-p", "icmp",
		"-m", "icmp",
		"--icmp-type", "11",
		"-j", "QUEUE").Output(); err != nil{
		fmt.Println("del input: ", err);
	}else{
		fmt.Println("delete input success", out);
	}
}

func build(){
	//iptables -I INPUT -p icmp -m icmp --icmp-type 11 -j QUEUE
	if out, err := exec.Command("iptables",
		"-I", "INPUT",
		"-p", "icmp",
		"--icmp-type", "11",
		"-j", "QUEUE").Output();err != nil{
		fmt.Println("create input: ", err);
	}else{
		fmt.Println("create input success ", out);
	}

	//iptables -I INPUT -p udp --src 8.8.8.8 -j QUEUE
	if out, err := exec.Command("iptables",
		"-I", "INPUT",
		"-p", "udp",
		"--src", "8.8.8.8",
		"-j", "QUEUE").Output(); err != nil{
		fmt.Println("create input: ", err);
	}else{
		fmt.Println("create input success ", out);
	}	

	//iptables -I OUTPUT -p udp --dst 8.8.8.8 -j QUEUE
	if out, err := exec.Command("iptables",
		"-I", "OUTPUT",
		"-p", "udp",
		"--dst", "8.8.8.8",
		"-j", "QUEUE").Output(); err != nil{
		fmt.Println("create input: ", err);
	}else{
		fmt.Println("create input success ", out);
	}	
}
