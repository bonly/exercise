package main

import (
  "bytes"
  "fmt"
  "io"
  "log"
  "mime/multipart"
  "net/http"
  "os"
)

func main() {
  extraParams := map[string]string{
      "dirId":"17",
      "name":"bonly-test",
  }
  bodyBuf := &bytes.Buffer{};
  bodyWriter := multipart.NewWriter(bodyBuf);

  //关键数据
  fileWriter, err := bodyWriter.CreateFormFile("file", "/home/bonly/media/picture/wallpaper.jpg");
  if err != nil{
    fmt.Println("error writing to buffer ", err);
    return;
  }

  //打开文件
  fh, err := os.Open("/home/bonly/media/picture/wallpaper.jpg");
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
    "http://omstest.xbed.com.cn/Interface/index.php/Home/MaterialInterface/uploadPic",
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