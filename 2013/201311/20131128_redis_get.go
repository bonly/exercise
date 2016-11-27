package main  

import (
"fmt"
"github.com/garyburd/redigo/redis"
"os"
)

func main(){
    if len(os.Args) != 2{
    	fmt.Println(os.Args[0] + " key\n");
    	return;
    }
 	cli, err := redis.Dial("tcp", "192.168.1.13:6379");
	if  err != nil{
		fmt.Println("connect fail: ", err);
		return;
	}
	defer cli.Close();

	//读出
	username, err := redis.String(cli.Do("GET", os.Args[1]));
	if  err != nil{
		fmt.Println("set fail: ", err);
	}else{
		fmt.Printf("get: %v\n", username);
	}
}