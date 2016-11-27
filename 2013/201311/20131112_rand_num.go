package main
 
import (
        "fmt"
        "math/rand"
)
 
func randSeq(n int) string {
        //letters := []rune("abcdefghijklmnopqrstuvwxyz")
        letters := []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
        b := make([]rune, n)
        for i := range b {
                b[i] = letters[rand.Intn(len(letters))]
        }
        return string(b)
}
 
func main() {
        for i := 1; i < 1000000; i++ {
                fmt.Println(randSeq(6))
                randSeq(5)
        }
}