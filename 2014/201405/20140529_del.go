/*
auth: bonly
create: 2016-6-14
*/
package main

import (
  "bytes"
  "fmt"
  "log"
  "mime/multipart"
  "net/http"
  "flag"
)

var srv = flag.String("s", "http://omstest.xbed.com.cn", "file server address");
var Id = flag.String("i", "1", "dir id");

func main(){
  flag.Parse();
  del(*Id);
}

func del(Id string) {
  extraParams := map[string]string{
      "id":Id,
  }
  bodyBuf := &bytes.Buffer{};
  bodyWriter := multipart.NewWriter(bodyBuf);

  for key, val := range extraParams {
      _ = bodyWriter.WriteField(key, val); //加入其它字段
  }

  contentType := bodyWriter.FormDataContentType();
  bodyWriter.Close();

  request, err := http.NewRequest("POST", 
    *srv + "/Interface/index.php/Home/MaterialInterface/delBaseMaterial",
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

