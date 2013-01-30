package main

import "os"
import "io"
import "bufio"
import "bytes"
import "fmt"
import "regexp"
import "flag"
import "log"
import "path/filepath"

func ReadDir(path string){
   dirname := string(filepath.Separator) + path;
   dir, err := os.Open(dirname);
   if err != nil {
   	 log.Println(err);
   	 os.Exit(1);
   }
   flst, err := dir.Readdir(-1);
   if err != nil {
   	 log.Println(err);
   	 os.Exit(1);
   }
   for _, fi := range flst{
      //if fi.IsRegular(){
      	fmt.Println(fi.Name(), fi.Size(), "bytes");
      //}
   }
}

func readLines(path string, endflag string) (ibegin int, iend int, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)
	if file, err = os.Open(path); err != nil {
		panic(err)
		return
	}

	//正则表达式检查
	str_begin := "file_id: 1$"       //开始
	str_end := "file_id: " + endflag //结束

	rb, err := regexp.Compile(str_begin)
	re, err := regexp.Compile(str_end)

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 1024))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			//println(err);
			break
		}
		buffer.Write(part) // 把part的内容写到 buffer中
		if !prefix {
			if rb.MatchString(buffer.String()) == true {
				ibegin++
				//lines = append(lines, buffer.String());
			}
			if re.MatchString(buffer.String()) == true {
				iend++
				//lines = append(lines, buffer.String());
			}
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func main() {
  flag.Parse();
	ReadDir("tmp");
	/*
	if flag.NArg() == 2 {
		println("parse: ", flag.Arg(0))
		begin, end, err := readLines(flag.Arg(0), flag.Arg(1))
		if err != nil {
			fmt.Println("Error: %s\n", err)
			return
		}
		fmt.Println("begin download: ", begin)
		fmt.Println("end download: ", end)
		*/
		
		/*
			for _, line := range lines {
			    fmt.Println(line);
			}
		*/
	//}
}
