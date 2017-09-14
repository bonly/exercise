package main

import (
  "bytes"
  "fmt"
  "io"
  "log"
  "mime/multipart"
  "net/http"
  "os"
  "path/filepath"
)

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, 
                          paramName, path string) (*http.Request, error) {
  file, err := os.Open(path)
  if err != nil {
      return nil, err
  }
  defer file.Close()

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body); //关键类型
  part, err := writer.CreateFormFile(paramName, filepath.Base(path)); //文件字段
  if err != nil {
      return nil, err
  }
  _, err = io.Copy(part, file); //传送文件内容

  for key, val := range params {
      _ = writer.WriteField(key, val); //加入其它字段
  }
  err = writer.Close()
  if err != nil {
      return nil, err
  }

  return http.NewRequest("POST", uri, body)
}

func main() {
  path, _ := os.Getwd()
  path += "/test.pdf"
  extraParams := map[string]string{
      "dirId":"17",
      "name":"bonly-test",
  }
  request, err := newfileUploadRequest(
    "http://omstest.xbed.com.cn/Interface/index.php/Home/MaterialInterface/uploadPic", 
    // "http://127.0.0.1:9090",
    extraParams, "file", "/home/bonly/media/picture/wallpaper.jpg")
  if err != nil {
      log.Fatal(err)
  }

  request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");
  request.Header.Set("Content-Type", "multipart/form-data");
  // fmt.Println("request:", request);
  client := &http.Client{}
  
  resp, err := client.Do(request)
  if err != nil {
      log.Fatal(err)
  } else {
      body := &bytes.Buffer{}
      _, err := body.ReadFrom(resp.Body)
    if err != nil {
          log.Fatal(err)
      }
    resp.Body.Close()
      fmt.Println(resp.StatusCode)
      fmt.Println(resp.Header)
      fmt.Println(body)
  }
}