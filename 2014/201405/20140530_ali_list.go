package main 
import (
"github.com/aliyun/aliyun-oss-go-sdk/oss"
"fmt"
)

func main(){
    client, err := oss.New("http://oss-cn-shenzhen.aliyuncs.com", 
        "GKDPHaw0FzDkdWSy", "oKqB4iENxztqsX0jLG1WDg7bVJbin8")
    if err != nil {
        // HandleError(err)
        fmt.Println("client: ", err);
    }

    // err = client.CreateBucket("xbed.test");
    // if err != nil{
    //     fmt.Println("create: ", err);
    // }

    lsRes, err := client.ListBuckets()
    if err != nil {
        // HandleError(err)
        fmt.Println("list: ", err);
    }

    for _, bucket := range lsRes.Buckets {
        fmt.Println("bucket:", bucket.Name)
    }

    bucket, err := client.Bucket("xbed")
    if err != nil {
        // HandleError(err)
        fmt.Println("bind: ", err);
    }

    err = bucket.PutObjectFromFile("bonly-test.jpg", "/home/bonly/media/picture/wallpaper.jpg")
    if err != nil {
        // HandleError(err)
        fmt.Println("put: ", err);
    }   

    err = bucket.DeleteObject("bonly-test")
    if err != nil {
        // HandleError(err)
        fmt.Println("put: ", err);
    }     
}