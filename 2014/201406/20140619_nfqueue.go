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

var current_ttl = 1;

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
			// Wg.Done();
			fmt.Println("end sig process");
		}();
		done := false;
		for ;!done; {
			fmt.Println("wait sig");
			select {
				case done = <-Done:
				    fmt.Println("sig get done: ", done);
				    if done == true{
					    Done <- true;
					}
					Wg.Done();
					return;
				case sig := <-sigs:
					process_sig(sig);
			}
		}
	}();

	Wg.Add(1);
	go func(){
		defer func(){
			Wg.Done();
			fmt.Println("end pack process");
		}();
		packets := nfq.GetPackets();
		done := false;
		for ;!done; {
			fmt.Println("wait pack");
			select {
			case done = <-Done:
				fmt.Println("pack get done: ", done);
				if done == true{
					Done<-true;
				}
				return;
			case p := <-packets:
				// fmt.Printf("数据包：%+v\n", p);
				ip := p.Packet.NetworkLayer().(*layers.IPv4);
				if ip.Protocol == layers.IPProtocolICMPv4{
					fmt.Println("ICMP From ", ip.SrcIP);
			    }else if ip.Protocol == layers.IPProtocolUDP{
			    	fmt.Println("UDP From ", ip.SrcIP);
			    	fmt.Println("UDP TO   ", ip.DstIP);
			    	if p.Packet.Layers()[2].LayerType() == layers.LayerTypeDNS{
			    		fmt.Printf("数据包：%+v\n", p);
			    		current_ttl += 1;
			    		dns := p.Packet.Layers()[2].(*layers.DNS);
			    		for idx :=0 ; idx < len(dns.Answers); idx++ {
    			    		// fmt.Printf("DNS包主要内容： %+v\n", dns.Answers[idx]);
    			    		fmt.Printf("DNS服务器解释后： %s\n", dns.Answers[idx].IP);
			    		}
			    	}else{
			    		fmt.Println("非DNS的UDP包,可能是出去的");
			    		// p.SetVerdict(netfilter.NF_DROP);
			    		p.SetVerdict(netfilter.NF_ACCEPT);
			    		return;
			    	}
			    }
				p.SetVerdict(netfilter.NF_ACCEPT);
			}
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


/*
dig +tries=30 +time=1 @8.8.8.8 twitter.com
数据包：{Packet:PACKET: 100 bytes
- Layer 1 (20 bytes) = IPv4	{Contents=[..20..] Payload=[..80..] Version=4 IHL=5 TOS=0 Length=100 Id=54605 Flags= FragOffset=0 TTL=44 Protocol=UDP Checksum=59071 SrcIP=8.8.8.8 DstIP=192.168.1.196 Options=[] Padding=[]}
- Layer 2 (08 bytes) = UDP	{Contents=[..8..] Payload=[..72..] SrcPort=53(domain) DstPort=60255 Length=80 Checksum=40736}
- Layer 3 (72 bytes) = DNS	{Contents=[..72..] Payload=[] ID=21555 QR=true OpCode=0 AA=false TC=false RD=true RA=true Z=0 ResponseCode=No Error QDCount=1 ANCount=2 NSCount=0 ARCount=1 Questions=[{Name=[..11..] Type=1 Class=1}] Answers=[{Name=[..11..] Type=1 Class=1 TTL=54 DataLength=4 Data=[104, 244, 42, 129] IP=104.244.42.129 NS=[] CNAME=[] PTR=[] TXTs=[] SOA={ MName=[] RName=[] Serial=0 Refresh=0 Retry=0 Expire=0 Minimum=0} SRV={ Priority=0 Weight=0 Port=0 Name=[]} MX={ Preference=0 Name=[]} TXT=[]}, {Name=[..11..] Type=1 Class=1 TTL=54 DataLength=4 Data=[104, 244, 42, 1] IP=104.244.42.1 NS=[] CNAME=[] PTR=[] TXTs=[] SOA={ MName=[] RName=[] Serial=0 Refresh=0 Retry=0 Expire=0 Minimum=0} SRV={ Priority=0 Weight=0 Port=0 Name=[]} MX={ Preference=0 Name=[]} TXT=[]}] Authorities=[] Additionals=[{Name=[] Type=41 Class=512 TTL=0 DataLength=0 Data=[] IP=<nil> NS=[] CNAME=[] PTR=[] TXTs=[] SOA={ MName=[] RName=[] Serial=0 Refresh=0 Retry=0 Expire=0 Minimum=0} SRV={ Priority=0 Weight=0 Port=0 Name=[]} MX={ Preference=0 Name=[]} TXT=[]}]}
 verdictChannel:0xc42115a120}

https://godoc.org/github.com/google/gopacket#Layer
http://fqrouter.tumblr.com/post/38463337823/%E4%BD%BF%E7%94%A8%E4%B8%AD%E5%9B%BD%E7%9A%84ip%E8%AE%BF%E9%97%AEtwitter1%E4%BD%BF%E7%94%A8nfqueue%E5%A4%84%E7%90%86packet
http://fqrouter.tumblr.com/post/38465016969/%E4%BD%BF%E7%94%A8%E4%B8%AD%E5%9B%BD%E7%9A%84ip%E8%AE%BF%E9%97%AEtwitter2nat%E7%9A%84%E5%9B%B0%E5%A2%83
http://fqrouter.tumblr.com/post/38468377418/%E4%BD%BF%E7%94%A8%E4%B8%AD%E5%9B%BD%E7%9A%84ip%E8%AE%BF%E9%97%AEtwitter3raw-socket%E7%9A%84%E4%BD%BF%E7%94%A8
http://fqrouter.tumblr.com/post/38469096940/%E4%BD%BF%E7%94%A8%E4%B8%AD%E5%9B%BD%E7%9A%84ip%E8%AE%BF%E9%97%AEtwitter4%E4%B8%8A%E8%A1%8C%E4%BB%A3%E7%90%86


http://fqrouter.tumblr.com/android-latest
*/

