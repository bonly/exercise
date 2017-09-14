package main 

import (
"log"
"fmt"
"regexp"
"os"
"bufio"
"flag"
)

var logFile = flag.String("f", "/tmp/2016-08-15/HLS_info.log", "豪力仕日志文件");

func main(){
	flag.Parse();

	infile, err := os.OpenFile(*logFile, os.O_RDONLY, 0666);
	if err != nil{
		log.Printf("打开日志[%s]失败\n", *logFile, err);
		return;
	}
	defer infile.Close();

	scanner := bufio.NewScanner(infile);

	door_log := regexp.MustCompile(`锁\d.*在.*被密码\d.*开门成功`);

	for scanner.Scan(){
		text := scanner.Text();
		match_idx := door_log.FindString(text);
		fmt.Println(len(match_idx));
		if len(match_idx) > 1 {
			fmt.Println(string(text[match_idx[0]]));
		}
	}

}

