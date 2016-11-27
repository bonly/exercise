package main
import "fmt"

func main(){
	fetchDemo();
	fmt.Println("the main function is executed.");
}

func fetchDemo(){
	defer func(){
		if v:= recover(); v!=nil{
			fmt.Printf("recovered a panic.[index=%d]\n",v);
		}
	}();
	ss := []string{"A", "B", "C"};
	fmt.Printf("getch the elements in %v one by one...\n", ss);
	fetchElement(ss, 0);
	fmt.Println("the elements fetching is done.");
}


func fetchElement(ss []string, index int)(element string){
	if index >= len(ss){
		fmt.Printf("occur a panic! [index=%d]\n", index);
		panic(index);
	}
	fmt.Printf("fetching the element... [index=%d]\n", index);
	element = ss[index];
	defer fmt.Printf("the element is \"%s\". [index=%d]\n", element, index);
	fetchElement(ss, index+1);
	return;
}
