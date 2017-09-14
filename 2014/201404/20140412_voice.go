package main
import (
    "io/ioutil"
    "net/http"
    "net/url"
    "fmt"
    "encoding/json"
)
  
//----------------------------------
// 在线接口文档：http://www.juhe.cn/docs/61
//----------------------------------
  
const APPKEY = "*******************" //您申请的APPKEY
  
func main(){
  
    //1.发送语音验证码
    Request1()
  
}
  
//1.发送语音验证码
func Request1(){
    //请求地址
    juheURL :="http://op.juhe.cn/yuntongxun/voice"
  
    //初始化参数
    param:=url.Values{}
  
    //配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
    param.Set("valicode","") //验证码内容，字母、数字 4-8位
    param.Set("to","") //接收手机号码
    param.Set("playtimes","") //验证码播放次数，默认3
    param.Set("key",APPKEY) //应用APPKEY(应用详细页查询)
    param.Set("dtype","") //返回数据的格式,xml或json，默认json
  
  
    //发送请求
    data,err:=Get(juheURL,param)
    if err!=nil{
        fmt.Errorf("请求失败,错误信息:\r\n%v",err)
    }else{
        var netReturn map[string]interface{}
        json.Unmarshal(data,&netReturn)
        if netReturn["error_code"].(float64)==0{
            fmt.Printf("接口返回result字段是:\r\n%v",netReturn["result"])
        }
    }
}
  
  
  
// get 网络请求
func Get(apiURL string,params url.Values)(rs[]byte ,err error){
    var Url *url.URL
    Url,err=url.Parse(apiURL)
    if err!=nil{
        fmt.Printf("解析url错误:\r\n%v",err)
        return nil,err
    }
    //如果参数中有中文参数,这个方法会进行URLEncode
    Url.RawQuery=params.Encode()
    resp,err:=http.Get(Url.String())
    if err!=nil{
        fmt.Println("err:",err)
        return nil,err
    }
    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}
  
// post 网络请求 ,params 是url.Values类型
func Post(apiURL string, params url.Values)(rs[]byte,err error){
    resp,err:=http.PostForm(apiURL, params)
    if err!=nil{
        return nil ,err
    }
    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}