package main

import (
    "fmt"
    "html/template"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "regexp"
    "strconv"
    "time"
)

const  Upload_Dir   = "./upload/";

func upload(w http.ResponseWriter, r *http.Request) {
    filename := strconv.FormatInt(time.Now().Unix(), 10) + "data";

	f, _ := os.OpenFile(Upload_Dir+filename, os.O_CREATE|os.O_WRONLY, 0660);
    _, err = io.Copy(f, file);
    if err != nil {
            fmt.Fprintf(w, "%v", "上传失败");
            return;
    }
    filedir, _ := filepath.Abs(Upload_Dir + filename);
    fmt.Fprintf(w, "%v", filename+"上传完成,服务器地址:"+filedir);
}

func main(){
	http.HandleFunc("/upload", upload);
	http.Handle("/", http.FileServer(http.Dir(".")));
	
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err);
	}
}