package main
import (
	// "bytes"
	"fmt"
	"github.com/gmallard/stompngo"
	"log"
	"net"
	"os"
	// "os/exec"
)

// 存储日志信息的变量
var mylog *log.Logger;

// 启动初始化
func init() {
	file := "./log.txt";
	//t := time.Now()
	//file := "./log_" + strings.Replace(t.String()[:19], ":", "_", 3) + ".txt"
	hander, err := os.OpenFile(file, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666);
	if err != nil {
		log.Println(err);
	}
	mylog = log.New(hander, "\r\n", log.Ldate|log.Ltime|log.Llongfile);
}

// 主程序
func main() {
	host := "127.0.0.1"; 
	port := "61613";
	n, e := net.Dial("tcp", net.JoinHostPort(host, port));
	if e != nil {
		fmt.Println(e);
	}
	// STOMP 1.0 的标准头
	//h := stompngo.Headers{};
	// STOMP 1.1 的标准头
	h := stompngo.Headers{"accept-version", "1.1"};
	// @todo 强化网络断开之后重试
	c, e := stompngo.Connect(n, h);
	if e != nil {
		fmt.Println(e);
	}
	// 必须客户端响应才可以删除MQ队列数据
	f := stompngo.Headers{"destination", "/queue/bbg_ordercache", "ack", "client"};
	// 自动删除MQ队列的数据
	//f := stompngo.Headers{"destination", "/queue/bbg_ordercache"};
	// s, e := c.Subscribe(f);
	_, e = c.Subscribe(f);
	if e != nil {
		fmt.Println(e);
	}
	// 设置通道的容量
	//fmt.Println(c.SubChanCap());
	//c.SetSubChanCap(1);
	/*
	for {
		//r := <-s
		//fmt.Println(r)
		run(c, s);
	}
	*/
}

/*
// 外部shell脚本调用，成功处理删除相应队列
func run(c *stompngo.Connection, s <-chan stompngo.MessageData) {
	phproot, _ := config.GetValue("php", "phpbin")
	filepath, _ := config.GetValue("php", "filepath")
	params, _ := config.GetValue("php", "params")
	r := <-s
	// 记录结果
	mylog.Println(r)
	order_id := r.Message.Headers.Value("order_id")
	//fmt.Println(order_id)
	cmd := exec.Command(phproot, filepath, params, "order_id", order_id)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	mylog.Println("Waiting for command to finish...")
	err = cmd.Wait()
	if err != nil {
		mylog.Printf("Command finished with error: %v", err)
	}
	str := out.String()
	//fmt.Println(str)
	if str == "success" {
		e := c.Ack(r.Message.Headers)
		if e != nil {
			mylog.Println(e)
		}
	}
}
*/
