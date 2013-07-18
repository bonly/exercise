package main;
import (
  "fmt"
  "os"
);

func Get_num(in int) (out int){
  fmt.Fprintf(os.Stderr,"in num: %d\n", in);
  out = 12;
  in = 11;
  return out;	
}

func Get_num_point(in *int)(out int){
	fmt.Fprintf(os.Stderr, "org in num: %d\n", *in);
	*in = 22;
	return *in;
}

func main(){
	ret := Get_num (13);
	fmt.Fprintf(os.Stdout, "out num: %d\n", ret);
	
	org := 34;
	ret = Get_num_point(&org);
	fmt.Fprintf(os.Stdout, "aft num: %d\n", ret);
	fmt.Fprintf(os.Stdout, "chg num: %d\n", org);
}
