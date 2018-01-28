package main

import (
"fmt"
"os"
"path/filepath"
// "io/ioutil"
"strings"
)

// func List_Dir(dirPth string, suffix string) (files []string, err error){
// 	files = make([]string, 0, 10);
// 	dir, err := ioutil.ReadDir(dirPth);
// 	if err != nil{
// 		return nil, err;
// 	}
// }
type Pair struct{
	First string;
	Second string;
};

func Walk_Dir(dirPth string, suffix string)(files []string, err error){
	files = make([]string, 0, 30);
	suffix = strings.ToUpper(suffix);

	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error{ //遍历目录
			if fi.IsDir(){
				return nil;
			}
			
			if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix){
				files = append(files, filename);
			}

			return nil;
	});
	return files, err;
}

func Check(dirPth string)(files []string, err error){
	files = make([]string, 0, 30);

	lst := make(map[string]*Pair); 
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error{
		if fi.IsDir(){
			// return nil;
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), ".META"){ //meta文件
			val, ok := lst[filename]; 
			if ok{ //有文件
				val.First = fi.Name();
			}else{ //无文件
				lst[filename] = &Pair{fi.Name(), ""};
			}
			// fmt.Printf("have: %#v\n", val);
		}else{ // 实体文件
			val, ok := lst[filename + ".meta"];
			if ok{
				val.Second = fi.Name();
			}else{
				lst[filename + ".meta"] = &Pair{"", fi.Name()};  
			}
			// fmt.Printf("has: %#v\n", val);
		}

		return nil;
	});

	for im, it := range lst{
		// fmt.Printf("%#v = %#v \n", im, it);
		if it.Second == ""{
			files = append(files, im);
		}
	}
	return files, err;
}

func main(){
	test2();
}

func test1(){
	files, err := Walk_Dir(".", ".meta");
	if err != nil{
		return;
	}
	for idx, it := range files{
		fmt.Printf("%d: %s \n", idx, it);
    }	
}

func test2(){
	files, _ := Check(".");
	for _, it := range files{
		fmt.Printf("rm -rf '%s'\n", it);
	}
}