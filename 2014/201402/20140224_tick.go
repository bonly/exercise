package main

import (
	"fmt"
	"time"
)

func main() {
	//第一种实现方式 
	ticker1 := time.NewTicker(1 * time.Second)
	i := 1
	for c := range ticker1.C {
		i++
		fmt.Println(c.Format("2006/01/02 15:04:05.999999999"))
		if i > 5 {
			ticker1.Stop()
			break
		}
	}
	fmt.Println(time.Now().Format("2006/01/02 15:04:05.999999999"), " 1 Finished.")

	//第二种实现方式 
	i = 1
	ticker2 := time.AfterFunc(1*time.Second, func() {
		i++
		fmt.Println(time.Now().Format("2006/01/02 15:04:05.999999999"))
	})

	for {
		select {
		case <-ticker2.C:
			fmt.Println("nsmei")
		case <-time.After(3 * time.Second):
			if i <= 5 {
				ticker2.Reset(1 * time.Second)
				continue
			}
			goto BRK
		}
		BRK:
		ticker2.Stop()
		break
	}
	fmt.Println(time.Now().Format("2006/01/02 15:04:05.999999999"), " 2 Finished.")
}
