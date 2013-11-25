package main
import "fmt"
import "time"
import "runtime"

func main(){
  //wait_end();
  //use_chn();
  //use_noblock_chn();
  Go_for10();
}

func use_noblock_chn(){
	c := make(chan bool);
	go func(){
		fmt.Println("Go Go Go!!!");
		c <- true;
		close(c);  //必须关闭,否则for会一直取.无结束
	}();
	
	for v:=range c{  //非阻塞检查
		fmt.Println(v);
	}
}

func use_chn(){
	c := make(chan bool); //make(chan bool, 2) 2:缓存大小
	go func(){
		fmt.Println("Go Go Go!!!");
		c <- true;
	}();
	
	<-c; //无缓存是阻塞等待, 有缓存是非阻塞
}

func wait_end(){
  go Go();
	time.Sleep(2*time.Second);
}

func Go(){
	fmt.Println("Go Go Go!!!");
}

func Go_for10(){
	runtime.GOMAXPROCS(runtime.NumCPU());
	c:=make(chan bool, 10); //没10的缓存时..顺序不定.可能没有全部执行完
	for i:=0; i<10; i++{
		go Go_10(c,i);
	}
	for i:=0; i<10; i++{
		<-c;
	}
}

func Go_10(c chan bool, index int){
	a:=1;
	for i:=0; i<100000; i++{
		a+=i;
	}
	fmt.Println(index, a);
	//if index == 9{
		c<-true; //10个都发
	//}
}
