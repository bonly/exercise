package main

import(
"fmt"
"golang.org/x/text/encoding/simplifiedchinese"
"golang.org/x/text/transform"
"bytes"
// "encoding/binary"
"io/ioutil"
)

func main(){
	fmt.Println(string(0x4f55));
	fmt.Println("无法显示的GBK：",string(0xbace));

	//不正确，应是大小端写入后就变了
	// var buf bytes.Buffer;
	// binary.Write(&buf, binary.LittleEndian, 0xbace);
	// enc := simplifiedchinese.GBK.NewDecoder();
	// rd := bytes.NewReader(buf.Bytes());
	// data, err := ioutil.ReadAll(transform.NewReader(rd, enc));
	// if err == nil{
	// 	fmt.Println(string(data));
	// }

	//不正确，大小端写入后不对
	// var buf bytes.Buffer;
	// binary.Write(&buf, binary.LittleEndian, 0xbace);
	// dt, _ := GBK2UTF8(buf.Bytes());
	// fmt.Printf("%x\n", dt);


	var au = []byte{0xba, 0xce};
	fmt.Printf("16进制的数据:%x\n", au);
	data,_ := GBK2UTF8(au);
	fmt.Printf("%X\n", data);
	fmt.Println("转换为utf8的:", string(data));

	str := "gbk的编码";
	gbk, _ := UTF82GBK([]byte(str));
	fmt.Println("gbk: ", string(gbk));
	utf8, _ := GBK2UTF8(gbk);
	fmt.Println("utf8: ", string(utf8));

	dat := 0xbace;
	fmt.Printf("原数据: %x\n", dat);
	dat_h := dat >> 8;
	fmt.Printf("高位: %x\n", dat_h);
	dat_l := dat & 0xFF ;
	fmt.Printf("低位: %x\n", dat_l);
	f8, _ := GBK2UTF8([]byte{byte(dat_h), byte(dat_l)});
	fmt.Printf("转换后: %s", string(f8));
}

func Encode(src string) (dst string){
	data, err := ioutil.ReadAll(
			transform.NewReader(bytes.NewReader([]byte(src)), 
				simplifiedchinese.GBK.NewEncoder()));
	if err == nil{
		dst = string(data);
	}
	return;
}

func Decode(src string) (dst string){
	data, err := ioutil.ReadAll(
			transform.NewReader(bytes.NewReader([]byte(src)), 
				simplifiedchinese.GBK.NewDecoder()));
	if err == nil{
		dst = string(data);
	}
	return;
}

func GBK2UTF8(sb []byte)([]byte, error){
	reader := transform.NewReader(bytes.NewReader(sb), simplifiedchinese.GBK.NewDecoder());
	dst, err := ioutil.ReadAll(reader);
	if err != nil{
		return nil, err;
	}
	return dst, nil;
}

func UTF82GBK(sb []byte)([]byte, error){
	reader := transform.NewReader(bytes.NewReader(sb), simplifiedchinese.GBK.NewEncoder());
	dst, err := ioutil.ReadAll(reader);
	if err != nil{
		return nil, err;
	}
	return dst, nil;
}