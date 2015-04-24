package main

import (
    "fmt"
    "sort"
)

// func main() {
//     m := map[string]string{"b": "15", "z": "123123", "x": "sdf", "a": "12"}
//     mk := make([]string, len(m))
//     i := 0
//     for k, _ := range m {
//         mk[i] = k
//         i++
//     }
//     sort.Strings(mk)
//     fmt.Println(mk)
// }

func main() {
    m := map[string]string{"b": "15", "z": "123123", "x": "sdf", "a": "12"}
    keys := make([]string, 0, len(m))
    for key := range m {
        keys = append(keys, key)
    }
    sort.Strings(keys)
    fmt.Println(keys)
}