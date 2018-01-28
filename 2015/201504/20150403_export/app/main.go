//http://colobu.com/2017/05/12/call-private-functions-in-other-packages/
package main
import (
	"fmt"
	"b"
)
func main() {
	s := b.Greet("world")
	fmt.Println(s)
	s = b.Hi("world")
	fmt.Println(s)
}
