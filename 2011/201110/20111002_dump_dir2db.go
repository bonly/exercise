package main 

import (
    "flag"
    "os"
    "fmt"
    "io/ioutil"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "path/filepath"
    "regexp"
)

const APP_VERSION = "0.1";

// The flag package provides a default help printer via -h switch
var file_prefix *string = flag.String("p", "/Page/", "File prefix in database.");
var versionFlag *bool = flag.Bool("v", false, "Print the version number.");
var ver *string = flag.String("n", "1", "Version in database.");
var dir_name *string = flag.String("r", ".", "Dir name.");
var str_db *string = flag.String("d", "doctor.db", "Database name.");

func main() {
    flag.Parse(); // Scan the arguments list 

    if *versionFlag {
        fmt.Println("Version:", APP_VERSION);
    }
    println ("dir name: ", *dir_name);
    scan (*dir_name);
}

func scan(dir string){
    err := filepath.Walk(dir, func(dir string, f os.FileInfo, err error) error{
       if f == nil {return err}
       if f.IsDir() {return nil}
       //println(dir);
       matched, err := regexp.MatchString("\\.qml|\\.js", f.Name());
       if matched == true {
          proc(dir);
       }
       return nil;
    });

    if err != nil {
      fmt.Printf("dir return: %v\n", err);
    }
}

func proc(file_name string){    
    arr, err := ioutil.ReadFile(file_name);
    if err != nil {
      panic(err);
    }
    
    fmt.Println("processing file " + file_name);
    str := string(arr);
    
    db, err := sql.Open("sqlite3", *str_db);
    if err != nil {
      panic(err);
    }
    
    _, err = db.Exec("CREATE TABLE IF NOT EXISTS AppScript(script_name TEXT UNIQUE primary key, version int, script TEXT)");
    if err != nil {
      panic(err);
    }
    
    _, err = db.Exec("replace into AppScript values(?, ?, ?)", *file_prefix + file_name, *ver, str);
    if err != nil {
      panic(err);
    }
}



