package main

/*
#include <stdint.h>

#pragma pack(push, 1)
typedef struct {
	uint16_t size;
	uint16_t msgtype;
	uint32_t sequnce;
	uint8_t data1;
	uint32_t data2;
	uint16_t data3;
} mydata;
#pragma pack(pop)

mydata foo = {
	1, 2, 3, 4, 5, 6,
};

int size() {
	return sizeof(mydata);
}
*/
import "C"
import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"unsafe"
)

func main() {
	bs := C.GoBytes(unsafe.Pointer(&C.foo), C.size())
	fmt.Printf("len %d data %v\n", len(bs), bs)
	var data struct {
		Size, Msytype uint16
		Sequence      uint32
		Data1         uint8
		Data2         uint32
		Data3         uint16
	}
	err := binary.Read(bytes.NewReader(bs), binary.LittleEndian, &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", data) // {1 2 3 4 5 6}

	buf := new(bytes.Buffer)
	// binary.Write(buf, binary.BigEndian, data)
	binary.Write(buf, binary.LittleEndian, data)
	fmt.Printf("%d %v\n", buf.Len(), buf.Bytes()) // 15 [0 1 0 2 0 0 0 3 4 0 0 0 5 0 6]
}