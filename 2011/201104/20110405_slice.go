package main;
import "fmt";

//内置函数 len() cap()
//append追加元素, 越过容量,则重新分配并考贝原数组
//是一个引用(创建需要用make)..指向数组,是变长数组的替代方案
//如果new 就是指向数组的指针的指针
//copy时以短的数组个数为准

func main(){
	var s1 []int;  ///var是公声名,没...也没指定长度及初始化,也没有{},是空的slice
	fmt.Println(s1);
	
	//:=相当于已实例化,不再需要new内存
	a:=[10]int{1,2,3,4,5,6,7,8,9,10};
	fmt.Println(a);
	s2 :=a[9];
	fmt.Println(s2);
  s3 :=a[5:10]; //尾部可用len(a)取得或直接写 a[5:],如前5个则为a[:5]
  fmt.Println(s3);
  
  m1 := make ([]int,3,5); //3个成员的数组, 5为容量,如自动重新分配会长一倍,即为10容量
  fmt.Println(m1);
  
  p1 :=[]byte{'a','b','c','d','e','f','g','h','i','j','k'}; //没有var,通过:=声名并实例化了slice
  sa := p1[2:5];
  fmt.Println(sa);
  fmt.Println(string(sa));
  
  sb := sa[1:6]; //reslice时,旧slice的索引为准, 并且实际数据还是真实数组中的,所以不越界
  fmt.Println(string(sb));
  
  sc := make ([]int, 3, 6);
  fmt.Printf("%v %p\n", sc, sc); //打印地址
  sc = append(sc,1,2,3);
  fmt.Printf("%v %p\n",  sc, sc);//小于=于容量,地址不变
  sc = append(sc,4,5,6);
  fmt.Printf("%v %p\n",  sc, sc);//大于容量,地址变改,因为是一个新复制的数组
  //切片<=旧容量的,所指向的是旧数组,切片>=旧容量的,指向新复制的数组
}


