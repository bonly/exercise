package main 

import (
    "os"
    "flag"
    "fmt"
    "io/ioutil"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "encoding/base64"
    "path/filepath"
    "regexp"
)

const APP_VERSION = "0.1";

// The flag package provides a default help printer via -h switch
var file_prefix *string = flag.String("p", "", "File prefix in database.");
var versionFlag *bool = flag.Bool("v", false, "Print the version number.");
var ver *string = flag.String("n", "1", "Version in database.");
var dir_name *string = flag.String("d", "", "Dir name.");
var str_db *string = flag.String("d", "doctor.db", "Database name.");

func scan(dir string){
    err := filepath.Walk(dir, func(dir string, f os.FileInfo, err error) error{
       if f == nil {return err}
       if f.IsDir() {return nil}
       //println(dir);
       matched, err := regexp.MatchString("\\.jpg", f.Name());
       if matched == true {
          proc(dir, "jpg");
       }
       matched, err = regexp.MatchString("\\.png", f.Name());
       if matched == true {
          proc(dir, "png");
       }
       matched, err = regexp.MatchString("\\.gif", f.Name());
       if matched == true {
          proc(dir, "gif");
       }
       matched, err = regexp.MatchString("\\.bmp", f.Name());
       if matched == true {
          proc(dir, "bmp");
       }
       return nil;
    });

    if err != nil {
      fmt.Printf("dir return: %v\n", err);
    }
}

func main() {
    flag.Parse(); // Scan the arguments list 

    if *versionFlag {
        fmt.Println("Version:", APP_VERSION);
    }
    
}

func proc(file_name string, pic_type string){
    arr, err := ioutil.ReadFile(file_name);
    if err != nil {
      panic(err);
    }
    
    str := base64.StdEncoding.EncodeToString(arr);
    
    db, err := sql.Open("sqlite3", *str_db);
    if err != nil {
      panic(err);
    }
    
    _, err = db.Exec("CREATE TABLE IF NOT EXISTS AppPic(pic_name TEXT UNIQUE primary key, pic_type TEXT, version int, pic TEXT)");
    if err != nil {
      panic(err);
    }
    
    _, err = db.Exec("replace into AppPic values(?, ?, ?, ?)", *file_prefix + file_name, pic_type, *ver, string(str));
    if err != nil {
      panic(err);
    }
}



