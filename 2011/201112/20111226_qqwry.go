package main 
import (
    "github.com/yinheli/qqwry"
    "log"
    "os"
)
func main() {
    q := qqwry.NewQQwry("qqwry.dat")
    q.Find("180.89.94.90")
    log.Printf("ip:%v, Country:%v, City:%v", q.Ip, q.Country, q.City)
    // output: 
    // 2014/02/22 22:10:32 ip:180.89.94.90, Country:北京市, City:长城宽带

    q.Find(os.Args[1]);
    log.Printf("ip:%v, Country:%v, City:%v", q.Ip, q.Country, q.City)    
}