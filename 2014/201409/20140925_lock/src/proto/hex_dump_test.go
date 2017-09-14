package proto

import(
"testing"
)

func Test_hex_print(ts *testing.T){
	arr := []byte{0x34, 0x1}; 

	begin := 0;
	end := len(arr);
	ret := Hex_Line(arr, &begin, &end);
	ts.Logf("%s", ret);

	Hex_Dump(arr, len(arr));
}