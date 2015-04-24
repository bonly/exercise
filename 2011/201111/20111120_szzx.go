package main 

import (
	"net/http"
	"log"
	"io/ioutil"
	"io"
	"compress/gzip"
	"bytes"
	"strings"
)

func main(){
   post ("http://pay3.shenzhoufu.com/interface/version3/serverconnszx/entry-noxml.aspx","	", "abc");
}

func post(url string,cookie string,postStr string) []byte{
 
	log.Println("let's post :"+url)
 
	client := &http.Client{
		CheckRedirect: nil,
	}
 
	postBytesReader := bytes.NewReader([]byte(postStr))
	reqest, _ := http.NewRequest("POST", url, postBytesReader)
 
	reqest.Header.Set("User-Agent"," Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36")
	reqest.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Set("Accept-Charset","GBK,utf-8;q=0.7,*;q=0.3")
	reqest.Header.Set("Accept-Encoding","gzip,deflate,sdch")
//	reqest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	reqest.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	reqest.Header.Set("Accept-Language","zh-CN,zh;q=0.8")
	reqest.Header.Set("Cache-Control","max-age=0")
	reqest.Header.Set("Connection","keep-alive")
	reqest.Header.Set("Referer", url)
 
 
 
	if(len(cookie) > 0){
		log.Println("dealing with cookie:"+cookie)
		array:=strings.Split(cookie,";")
		for item:= range array{
			array2:=strings.Split(array[item],"=")
			if(len(array2)==2){
				cookieObj:= http.Cookie{}
				cookieObj.Name=array2[0]
				cookieObj.Value=array2[1]
				reqest.AddCookie(&cookieObj)
			}else{
				log.Println("error,index out of range:"+array[item])
			}
		}
	}
 
	resp, err := client.Do(reqest)
 
	if err != nil {
		log.Println(url,err)
		return nil
	}
 
	defer resp.Body.Close()
 
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			log.Println(url,err)
			return nil
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}
 
 
	if(reader!=nil){
		body, err := ioutil.ReadAll(reader)
		if err != nil {
			log.Println(url,err)
			return nil
		}
		return body
	}
	return nil
}

func get(url string,cookie string) []byte{
 
	log.Println("let's get :"+url)
 
	client := &http.Client{
		CheckRedirect: nil,
	}
	reqest, _ := http.NewRequest("GET", url, nil)
 
	reqest.Header.Set("User-Agent"," Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36")
	reqest.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Set("Accept-Charset","GBK,utf-8;q=0.7,*;q=0.3")
	reqest.Header.Set("Accept-Encoding","gzip,deflate,sdch")
	reqest.Header.Set("Accept-Language","zh-CN,zh;q=0.8")
	reqest.Header.Set("Cache-Control","max-age=0")
	reqest.Header.Set("Connection","keep-alive")
	reqest.Header.Set("Referer", url)
 
 
 
	if(len(cookie) > 0){
		log.Println("dealing with cookie:"+cookie)
		array:=strings.Split(cookie,";")
		for item:= range array{
			array2:=strings.Split(array[item],"=")
			if(len(array2)==2){
				cookieObj:= http.Cookie{}
				cookieObj.Name=array2[0]
				cookieObj.Value=array2[1]
				reqest.AddCookie(&cookieObj)
			}else{
				log.Println("error,index out of range:"+array[item])
			}
		}
	}
 
	resp, err := client.Do(reqest)
 
	if err != nil {
		log.Println(url,err)
		return nil
	}
 
	defer resp.Body.Close()
 
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			log.Println(url,err)
			return nil
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}
 
 
	if(reader!=nil){
		body, err := ioutil.ReadAll(reader)
		if err != nil {
			log.Println(url,err)
			return nil
		}
		return body
 
	}
	return nil
}
 