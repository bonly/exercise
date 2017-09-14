package main

import(
"fmt"
"os"
"os/exec"
"github.com/AkihiroSuda/go-netfilter-queue"
)

func main(){
	build();

	var err error;
	nfq, err := netfilter.NewNFQueue(0, 100, 
		netfilter.NF_DEFAULT_PACKET_SIZE);
	if err != nil{
		fmt.Println("new queue:", err);
		os.Exit(1);
	}
	defer nfq.Close();

	packets := nfq.GetPackets();

	for {
		select {
		case p := <-packets:
			fmt.Println("有包经过：", p.Packet);
			p.SetVerdict(netfilter.NF_ACCEPT);
		}
	}
	clean();
}


func clean(){
	if out, err := exec.Command("iptables",
		"-D", "OUTPUT",
		"-p", "udp",
		"--dst", "8.8.8.8",
		"-j", "QUEUE").Output(); err != nil{
		fmt.Println("del output: ", err);
	}else{
		fmt.Println("del out success", out);
	}

	if out, err := exec.Command("iptables",
		"-D", "INPUT",
		"-p", "udp",
		"--src", "8.8.8.8",
		"-j", "QUEUE").Output(); err != nil{
		fmt.Println("del input: ", err);
	}else{
		fmt.Println("delete input success", out);
	}
}

func build(){
	if out, err := exec.Command("iptables",
		"-I", "OUTPUT",
		"-p", "udp",
		"--dst", "8.8.8.8",
		"-j", "QUEUE").Output();err != nil{
		fmt.Println("create output: ", err);
	}else{
		fmt.Println("create output success ", out);
	}

	if out, err := exec.Command("iptables",
		"-I", "INPUT",
		"-p", "udp",
		"--src", "8.8.8.8",
		"-j", "QUEUE").Output(); err != nil{
		fmt.Println("create input: ", err);
	}else{
		fmt.Println("create input success ", out);
	}	
}