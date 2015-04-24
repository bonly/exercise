package main
import (
    "path/filepath"
    "os"
    "fmt"
    "flag"
)

func getFilelist(path string) {
        err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
                if ( f == nil ) {return err}
                if f.IsDir() {return nil}
                println(path)
                return nil
        })
        if err != nil {
                fmt.Printf("filepath.Walk() returned %v\n", err)
        }
}

func main(){
        flag.Parse()
        root := flag.Arg(0)
        getFilelist(root)
}
