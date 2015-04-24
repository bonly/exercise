package main

import (
        "fmt"
        "regexp"
)

func main() {
        re := regexp.MustCompile("fo.?")
        fmt.Printf("%q\n", re.FindString("seafood"))
        fmt.Printf("%q\n", re.FindString("meat"))
}
