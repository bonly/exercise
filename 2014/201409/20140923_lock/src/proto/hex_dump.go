/*
auth: bonly
create: 2016.9.16
desc: 十六进制输出
*/
package proto

import (
// "encoding/binary"
// "bytes"
// "unsafe"
"fmt"
"strconv"
)

const LINE_SIZE = 16;

//转换一行数据
func Hex_Line(data []byte, begin *int, end *int)(ret []byte){
	if *end <= 0{
		return;
	}

	var line [LINE_SIZE * 4 + 4]byte; //定义输出的一行buf
	
	data_idx := 0;
	line_idx := 0;
	for ; data_idx < LINE_SIZE; data_idx++{
		if *end > 0{
			line[line_idx] = (([]byte)(fmt.Sprintf("%02X", data[*begin + data_idx])))[0]; //输出Hex在左边
			line[line_idx+1] = (([]byte)(fmt.Sprintf("%02X", data[*begin + data_idx])))[1];
			line[line_idx+2] = ' ';
			line_idx = line_idx + 3;
			char := strconv.IsPrint(rune(data[*begin + data_idx]));
			if char == true{//转换后输出到右边，不可显字符显示为'.'
				line[LINE_SIZE * 3 + 3 + data_idx] = data[*begin + data_idx];
			}else{
				line[LINE_SIZE * 3 + 3 + data_idx] = '.';
			}
		}else{
			line[line_idx] = ' '; //Hex左边补空格
			line[line_idx+1] = ' ';
			line[line_idx+2] = ' ';
			line_idx = line_idx + 3;

			//右边补空格
			line[LINE_SIZE * 3 + 3 + data_idx] = ' ';
		}
		*end = *end - 1;
	}
	ret = line[:];
	*begin += data_idx;
	return ret;
}

func Hex_Dump(data []byte, size int){
	begin := 0;
	end := size;
	for ; end > 0;{
		line := Hex_Line(data, &begin, &end);
		fmt.Printf("%s\n", line);
	}
}