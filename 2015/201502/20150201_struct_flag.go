package main 

import (
"flag"
"fmt"
)

var Server struct{
	Log_level int;
	Log_path  string;
};


func init(){
	flag.IntVar(&Server.Log_level, "L", 0, "Log level");
	flag.StringVar(&Server.Log_path, "P", "/tmp/", "Log path");
	flag.Parse();
}

func main(){
	fmt.Printf("%#v\n", Server);
	fmt.Printf("Level: %d\n", Server.Log_level);
	fmt.Printf("Path: %s\n", Server.Log_path);
}

