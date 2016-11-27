package main

import "fmt"

func f(args ...string) {
	fmt.Println(len(args))
	for i, s := range args {
		fmt.Println(i, s)
	}
}

func main() {

	args := []string{"a", "b"}

	f(args...);

	fk(3, "kid", "farther", "mother", "ok", "failed");
	Update_payment_status("pk", "133", "t1", "13334", "t2", "field");
}


func fk(at int, para ...string){
	fmt.Println(at);
	fmt.Println("len: ", len(para));
	for i, s := range para{
		fmt.Println(i, ": ", s);
	}
}

func Update_payment_status(key string, val string, field_val ...string){
	var con string;
	for idx, str := range field_val{
		if idx%2==0{
			if idx == 0{
				con = str;
			}else{
				con = con + "," + str;
			}
		}else{
			con = con + "='" + str + "'";
		}
	}
	all_str := fmt.Sprintf("update xb_payment set %s where %s='%s'", con, key, val);
	fmt.Println(all_str);
}

