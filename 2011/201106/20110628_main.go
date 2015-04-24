//+build ignore

package main 

import (
    "."
    "flag"
    "fmt"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func Init(){

}

func main() {
    flag.Parse() // Scan the arguments list 

    if *versionFlag {
        fmt.Println("Version:", APP_VERSION)
    }

    test();
    fmt.Println("exit main");
}

func test(){
    pomelo.Client = pomelo.Connect("192.168.1.111", 4010);
    if (pomelo.Client == nil){
        fmt.Println("connect failed");
        return;
    }
    pomelo.Notify(pomelo.Client);
    pomelo.WaitJoin(pomelo.Client);
}



