package main
import (
"fmt"
"crypto/aes"
"strings"
)
func main(){
rb:=[]byte {1,2,3,4,5,6,7,8,9,0,1,2,3,4,5,6};
b:=make([]byte,16);
strings.NewReader("1234567890123456").Read(b);
// b=b[0:16];
fmt.Print("b:",b);
cip,err:= aes.NewCipher(b);
fmt.Print("cip:",cip,"\nerr:",err,"\n");
out:=make([]byte,len(rb));
cip.Encrypt (rb, out);

}

