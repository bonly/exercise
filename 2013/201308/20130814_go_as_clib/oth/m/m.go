package main  //import "m" 

import _ "p"

func main(){

}

/*建so
go install -buildmode=shared runtime sync/atomic #构建核心基本库
go install -buildmode=shared -linkshared p

go build -buildmode=c-shared -linkshared -ldflags "-r=."  
*/
