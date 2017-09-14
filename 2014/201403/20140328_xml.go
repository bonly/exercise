package main

import "fmt"
import "encoding/xml"

var text = `<?xml version="1.0" encoding="utf-8"?><string xmlns="http://tempuri.org/">4884983927739945376</string>`;
//var text = `<data>abc</data>`;

func main(){
	data := struct{
		Sn string `xml:",chardata"`;
	}{};
	xml.Unmarshal([]byte(text), &data);
	fmt.Println(data);
}
