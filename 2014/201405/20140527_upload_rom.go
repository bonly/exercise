/*
auth: bonly
create: 2016-6-14
*/
package main

import (
  "bytes"
  "fmt"
  "io"
  "log"
  "mime/multipart"
  "net/http"
  "os"
  "flag"
)

var file_name = flag.String("f", "1.jpg", "file name");
var srv = flag.String("s", "http://omstest.xbed.com.cn", "file server address");
var dirId = flag.String("i", "1", "dir id");
var name = flag.String("n", "picture.jpg", "file title name");

func main(){
  flag.Parse();
  upload(*dirId, *name, *file_name);
}

func upload(dirId string, name string, file_name string) {
  extraParams := map[string]string{
      "dirId":dirId,
      "name":name,
  }
  bodyBuf := &bytes.Buffer{};
  bodyWriter := multipart.NewWriter(bodyBuf);

  //关键数据
  fileWriter, err := bodyWriter.CreateFormFile("file", file_name);
  if err != nil{
    fmt.Println("error writing to buffer ", err);
    return;
  }

  //打开文件
  fh, err := os.Open(file_name);
  if err != nil{
    fmt.Println("err opening file ", err);
    return;
  }
  defer fh.Close();

  //iocopy
  _, err = io.Copy(fileWriter, fh);
  if err != nil{
    fmt.Println("io copy ", err);
    return;
  }

  for key, val := range extraParams {
      _ = bodyWriter.WriteField(key, val); //加入其它字段
  }

  contentType := bodyWriter.FormDataContentType();
  bodyWriter.Close();

  request, err := http.NewRequest("POST", 
    *srv + "/Interface/index.php/Home/MaterialInterface/uploadPic",
    bodyBuf);
  if err != nil {
      fmt.Println(err);
      return;
  }

  request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");
  request.Header.Set("Content-Type", contentType);

  // fmt.Println("request:", request);
  client := &http.Client{};
  
  resp, err := client.Do(request);
  if err != nil {
      log.Fatal(err);
  } else {
      body := &bytes.Buffer{};
      _, err := body.ReadFrom(resp.Body);
    if err != nil {
          log.Fatal(err);
    }
    resp.Body.Close();
    fmt.Println(resp.StatusCode);
    fmt.Println(resp.Header);
    fmt.Println(body);
  }
}

