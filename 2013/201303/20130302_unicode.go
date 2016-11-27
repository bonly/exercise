package main
import(
    "fmt"
    "unicode/utf8"
)
func main() {
    str := "你好，世界！"
    bytes := 0
    for _, r := range str {
        bytes += utf8.RuneLen(r)
    }
    fmt.Printf("the byte size: %d", bytes)
}
