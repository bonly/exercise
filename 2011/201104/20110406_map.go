package main;
import "fmt";
import "sort"

func main(){
	var m1 map[int]string; //声名
	m1 = map[int]string{}; //实例化 注意 {}
	fmt.Println(m1);
	
	var m2 map[int]string;
	m2 = make(map[int]string);
	fmt.Println(m2);
	
	m1[1]="OK";
	m1[2]="failed";
	delete (m1, 2);
	fmt.Println(m1);
	
	var mp map[int]map[int]string;
	mp = make(map[int]map[int]string);
	mp[1] = make(map[int]string);
	mp[1][1]="pk";
	a, ok := mp[1][1];
	fmt.Println(a,ok);
	
	//for range的访问是复制值访问
	sm := make([]map[int]string, 5);
	for _,v := range sm{
		v=make(map[int]string, 1);
		v[1]="ok";
		fmt.Println(v);
	}
	fmt.Println(sm);

	sm1 := make([]map[int]string, 5);
	for i := range sm1{
		sm1[i]=make(map[int]string, 1);
		sm1[i][1]="ok";
		fmt.Println(sm1[i]);
	}
	fmt.Println(sm1);	
	
	ms := map[int]string{1:"a",2:"b",3:"c",4:"d",5:"e"};
	s := make([]int, len(ms));
	i := 0;
	for k, _ := range ms{
		s[i] = k;
		i++;
	}
	
	sort.Ints(s);
	fmt.Println(s);
	
	mss := make(map[string]int);
	for k, v := range ms{
		mss[v] = k
	}
	fmt.Println(mss);
	
}
