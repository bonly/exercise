package file_manager
import (
  "os"
  "bufio"
  "log"
)

type FM interface{
	ReadLineAt(buf []byte, rd int64) int;
	Open(filename string) int;
};

type MyFile struct{
	Fh *os.File;
	Pos int64;
};

func (mf *MyFile)Open(filename string)(int){
	fh, err := os.Open(filename);
	if err != nil{
	  log.Fatal("open fail: ", err);
	  return -1;
	}
	mf.Fh = fh;
	mf.Pos = 0;
	return 0;
}

func (mf *MyFile)Close(){
	mf.Fh.Close();
}

func (mf *MyFile)ReadLineAt(buf []byte, rd int64)(int64){
	rd64, err := mf.Fh.Seek(rd, 0);
	if err != nil {
		log.Println(err);
		return -1;
	}
	
	log.Println("seek to: ", rd64);
	
	rder := bufio.NewReader(mf.Fh);
	tmpbuf, _, err := rder.ReadLine();
	if err != nil{
		log.Println(err);
		return -1;
	}
	
	copy (buf, tmpbuf);
	mf.Pos, err = mf.Fh.Seek(0, 1);
	if err != nil{
		log.Println(err);
	}
	log.Println("next begin: ", rd64);
	return mf.Pos;
}