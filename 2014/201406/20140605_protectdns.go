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
// var Run = true;

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
		packets := nfq.GetPackets();
		for{
			fmt.Println("wait pack");
			select {
			case done := <-Done:
				fmt.Println("pack get done: ", done);
				if done == true{
					Done<-true;
				}
				return;
			case p := <-packets:
				fmt.Printf("数据包：%+v\n", p);
				// fmt.Printf("Payload: %+v\n", p.Packet.ApplicationLayer().Payload());
				ip := p.Packet.NetworkLayer().(*layers.IPv4);
				fmt.Println("Got SIFF packet for ", ip.DstIP);
				fmt.Println("From ", ip.SrcIP);
			}
		}
	}();

	Wg.Add(1);
	go func(){
		defer func(){
			Wg.Done();
			fmt.Println("end loop func");
		}();
		for{
			select {
			case done := <-Done:
				fmt.Println("loop get done: ", done);
				if done == true{
					Done<-true;
					return;
				}
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
           // Run = false; 
           Done <-true;
           break;
        case syscall.SIGTERM:
           fmt.Println("Get SigTerm");
            Done <-true;
           // Run = false;
           break;
        case syscall.SIGUSR1:
            fmt.Println("Get SIGUSR1. Flash log");
            // Run = true;
            break;
        default:
           fmt.Println("Get Sig: ", sig);
           // Run = true;
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
修改 /etc/config/firewall 加入
config include
        option path /var/g.firewall.user
以便在防火墙重启后重新添加规则

创建启动脚本 /etc/init.d/protectdns 以便随系统自动启动
#!/bin/sh /etc/rc.common

START=99
RULEFILENAME=/var/g.firewall.user

start() {
        /etc/protectdns.sh &
        }

stop() {
        local pid
        pid=`ps w | grep protectdns.sh | grep -v grep | awk '{print $1}'`
        for i in $pid;do kill -9 $i;done
        [[ -s $RULEFILENAME ]] && {
                sed -ie 's/iptables \+-I \+/iptables -D /' $RULEFILENAME
                . $RULEFILENAME
                rm $RULEFILENAME
                }
        }
执行 /etc/init.d/protectdns enable 以将该脚本设置成自动启动，执行 /etc/init.d/protectdns stop 会移除所有已添加的iptables规则并将 /var/g.firewall.user 删除，
执行 /etc/init.d/protectdns start 会查找新的被污染IP并添加iptables规则

为了定时更新，可以在 /etc/crontabs/root 加入
0 6 * * * /etc/init.d/protectdns start


0>>22&0x3C@8&0x810F=0x8000
“0>>22&0x3C”这部分是很常用的一个规则，你在别的地方也应该经常看到，表示取出IP报文头的长度，后面的@表示跳过这么多长度，也就是跳过IP头到UDP头部分。
8表示从第8个字节开始抓4个字节，其中第10个字节的第一位如果是1就表示这是一个响应包，0对应的是查询报，这个我们不用管；
第十个字节的最低位是RD位，劫持包的RD位是0；
第11个字节的低4位就是Reply Code的位置。
将这4个字节和0x810F做与操作就提出了响应标志位，RD位和Reply Code的值，判断出是一个RD=0，Reply Code为无错误的响应包，再继续和下面的规则匹配。
这三个条件缺一不可，尤其是对RD位的检查，这几乎是GFW劫持包最重要的特征。

-m string表示启用string模块进行匹配，—algo bm表示启用贝叶(Boyer-Moore)字符串搜索算法，另一种算法是kmp(Knuth-Pratt-Morris)，具体应该只是效率和不同应用的区别，算法原理我就不细究了。
—hex-string就是要匹配的IP地址的HEX格式了，注意这里前面不要有0x表示16进制了，直接写16进制数据就行，后面跟着的就是IP地址的16进制格式。
—from 60 —to 180都是表示搜索的偏置范围(offset)，从第60 byte开始到180 byte结束，一般来说DNS返回包的长度很少有超过180字节，而IP值的位置因为协议格式的关系，基本上不可能出现在60字节之前，
所以指定这个范围是足够的，同时也可以大大减轻搜索算法的运算压力。如果你担心有什么奇葩DNS返回的结果在这范围之外，删除这两个偏置选项就行了，这样会默认搜索整个字符串，但相应的搜索消耗的CPU资源就更多了。
*/