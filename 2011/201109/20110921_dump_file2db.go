package main 

import (
    "flag"
    "fmt"
    "io/ioutil"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

const APP_VERSION = "0.1";

// The flag package provides a default help printer via -h switch
var file_prefix *string = flag.String("p", "", "File prefix in database.");
var versionFlag *bool = flag.Bool("v", false, "Print the version number.");
var ver *string = flag.String("n", "1", "Version in database.");
var file_name *string = flag.String("f", "", "File name.");
var str_db *string = flag.String("d", "doctor.db", "Database name.");

func main() {
    flag.Parse(); // Scan the arguments list 

    if *versionFlag {
        fmt.Println("Version:", APP_VERSION);
    }
    
    arr, err := ioutil.ReadFile(*file_name);
    if err != nil {
      panic(err);
    }
    
    str := string(arr);
    
    db, err := sql.Open("sqlite3", *str_db);
    if err != nil {
      panic(err);
    }
    
    _, err = db.Exec("CREATE TABLE IF NOT EXISTS AppScript(script_name TEXT UNIQUE primary key, version int, script TEXT)");
    if err != nil {
      panic(err);
    }
    
    _, err = db.Exec("replace into AppScript values(?, ?, ?)", *file_prefix + *file_name, *ver, str);
    if err != nil {
      panic(err);
    }
}



