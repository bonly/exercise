package main

import (
    "fmt"
    "os/exec"
)

func main() {
    argv := []string{"20111206_DES.php", "50", "1111", "2222", "123456"}
    c := exec.Command("php", argv...)
    d, _ := c.Output()
    fmt.Println(string(d)) 

}