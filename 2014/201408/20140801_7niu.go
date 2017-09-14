package main

import (
  "qiniupkg.com/api.v7/kodo"
  "qiniupkg.com/api.v7/conf"
  "qiniupkg.com/api.v7/kodocli"
  "fmt"
)

var (
    //设置上传到的空间
    bucket = "xbed-private"
    //设置上传文件的key
    key = "TEST/abc.jpg"
)

//构造返回值字段
type PutRet struct {
    Hash    string `json:"hash"`
    Key     string `json:"key"`
}

func main() {
    //初始化AK，SK
    conf.ACCESS_KEY = "50mt9hMHzqvlXi8ma4D7uYoVSW4c8lwaOOhU7e9-"
    conf.SECRET_KEY = "Qxxv09qWyex5ewwyf8Y_5Tkvl2YQGBx-RbUPNrBG"

    //创建一个Client
    c := kodo.New(0, nil)

    //设置上传的策略
    policy := &kodo.PutPolicy{
        Scope:   bucket+ ":" + key,
        //设置Token过期时间
        Expires: 3600,
    }
    //生成一个上传token
    token := c.MakeUptoken(policy)

    //构建一个uploader
    zone := 0
    uploader := kodocli.NewUploader(zone, nil)

    var ret PutRet
    //设置上传文件的路径
    filepath := "/home/bonly/media/picture/qiuqian.jpg"
    //调用PutFile方式上传，这里的key需要和上传指定的key一致
    res := uploader.PutFile(nil, &ret, token, key, filepath, nil)
    //打印返回的信息
    fmt.Println(ret)
    //打印出错信息
    if res != nil {
        fmt.Println("io.Put failed:", res)
        return
    }   

}
