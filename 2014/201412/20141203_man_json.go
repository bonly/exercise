/*
手工订制解包
*/
package main
import (
    "encoding/json"
    "fmt"
)

type Output struct {
    A int `json:"a"`
    B int `json:"b"`
    C int `json:"c"`
}

type JsonOutput Output
type Out struct {
    Output JsonOutput `json:"output"`
}

func (o *JsonOutput) UnmarshalJSON(data []byte) error {
    if len(data) == 2 && data[0] == '"' && data[1] == '"' {
        return nil
    }
    return json.Unmarshal(data, (*Output)(o))
}
func main() {
    txt := []byte(`
    {
        "output": ""
    }
    `)
    var out Out
    if err := json.Unmarshal(txt, &out); err != nil {
        panic(err)
    }
    fmt.Println(out)
}