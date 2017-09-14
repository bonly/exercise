package main

import (
	"fmt"
	"html"
	"net/url"
	//"testing"
)

func Test_Escape() {
	//url编码
	str := "中文-_."
	unstr := "%2f"
	fmt.Printf("url.QueryEscape:%s", url.QueryEscape(str))
	fmt.Println()
	s, _ := url.QueryUnescape(unstr)
	fmt.Printf("url.QueryUnescape:%s", s)
	fmt.Println()
	//字符转码
	hstr := "<"
	hunstr := "&lt"
	fmt.Printf("html.EscapeString:%s", html.EscapeString(hstr))
	fmt.Println()
	fmt.Printf("html.UnescapeString:%s", html.UnescapeString(hunstr))
	fmt.Println()
}

func main(){
	Test_Escape();
}
