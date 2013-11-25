package main
import "fmt"
import "sync"
import "runtime"

func main(){
  Go_for10();
}

func Go_for10(){
	runtime.GOMAXPROCS(runtime.NumCPU());
	wg := sync.WaitGroup{};
	wg.Add(10);

	for i:=0; i<10; i++{
		go Go_10(&wg,i);
	}
  wg.Wait();
}

func Go_10(wg *sync.WaitGroup, index int){
	a:=1;
	for i:=0; i<100000; i++{
		a+=i;
	}
	fmt.Println(index, a);
  wg.Done();
}
