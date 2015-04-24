package main
 
import (
    "fmt"
    "sort"
)
 

func main() {
    m := map[string]string{"b": "15", "z": "123123", "x": "sdf", "a": "12"}
    mk := make([]string, len(m))
    i := 0
    for k, _ := range m {
        mk[i] = k
        i++
    }
    sort.SortStrings(mk)
    fmt.Println(mk)
}