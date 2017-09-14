package main
import _ "net/http/pprof" 
import (
"net/http"
) 
  
func main() {  
    // go func() {  
        http.ListenAndServe("localhost:6060", nil)  
    // }()  
}  


/*
go tool pprof http://localhost:6060/debug/pprof/heap 
*/