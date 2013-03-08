package file_manager
import (
  "os"
  "bufio"
  "log"
)

type FM interface{
	ReadLineAt() int
};

type MyFile struct{
	
};

func (mf MyFile)ReadLineAt(filename string, rd int64, buf []byte)(int){
	fh, err := os.Open(filename);
	if err != nil{
		log.Fatal("open fail: ", err);
		return -1;
	}
	
	rd64, err := fh.Seek(rd, 0);
	if err != nil {
		log.Println(err);
		return -1;
	}
	
	log.Println("seek to: ", rd64);
	
	rder := bufio.NewReader(fh);
	tmpbuf, _, err := rder.ReadLine();
	if err != nil{
		log.Println(err);
		return -1;
	}
	return copy (buf, tmpbuf);
}