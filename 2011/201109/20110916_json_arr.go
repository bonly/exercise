package main

import (
    "encoding/json"
    "fmt"
)

type Root struct {
    Cmd string
    Operator string
    Change []Op
}

type Ar struct {
    Bag_id string
    slot_num string
}
type Op struct {
    Flag string
    Bag_id string
    Slot_num string
}

func main() {
    var s Root
    str := `{"Cmd":"27","Operator":"0","Change":[{"Flag":"from","Bag_id":"1001","Slot_num":"3"},{"Flag":"to","Bag_id":"1002","Slot_num":"1"}]}`
    json.Unmarshal([]byte(str), &s)
    fmt.Println(s)
}
