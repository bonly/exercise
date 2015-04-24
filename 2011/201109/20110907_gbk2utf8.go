import (
“bufio”
“code.google.com/p/mahonia”
“log”
“os”
)
func checkError(err interface{}) {
if err != nil {
log.Fatal(err)
}
}
func main() {
f, err := os.Open(“test.txt”)
checkError(err)
defer f.Close()
decoder := mahonia.NewDecoder(“gb18030″)
r := bufio.NewReader(decoder.NewReader(f))
line, _, err := r.ReadLine()
checkError(err)
println(string(line))
}