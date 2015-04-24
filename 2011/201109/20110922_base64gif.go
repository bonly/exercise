package main

import (
       "net/http"
       "io"
       "encoding/base64"
)

const base64GifPixel = "R0lGODlhAQABAIAAAP///wAAACwAAAAAAQABAAACAkQBADs="

func respHandler(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Content-Type","image/gif")
    output,_ := base64.StdEncoding.DecodeString(base64GifPixel)
    io.WriteString(res,string(output))
}


func main() {
    http.HandleFunc("/", respHandler)
    http.ListenAndServe(":8089", nil)
}
/*
wget -q -O file.gif http://localhost:8086
file file.gif
file.gif: GIF image data, version 89a, 1 x 1
*/
