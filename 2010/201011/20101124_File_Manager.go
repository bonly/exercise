package log_pkg
import (
  "os"
  "bufio"
  "log"
  "bytes"
  "io"
)

type FM interface{
	ReadLineAt(buf []byte, rd int64) int;
	Open(filename string) int;
};

type MyFile struct{
	Fh *os.File;
	Rder *bufio.Reader;
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
	
	mf.Rder = bufio.NewReader(mf.Fh);
	return 0;
}

func (mf *MyFile)Close(){
	mf.Fh.Close();
}

func (mf *MyFile)ReadLineAt(rd int64)(lines []string, err error) {
	_, err = mf.Fh.Seek(rd, 0);
	if err != nil {
		log.Println(err);
		return;
	}
	
	//log.Println("seek to: ", rd64);
	
	buf := bytes.NewBuffer(make([]byte, 0));
	var (
	  part []byte;
	  prefix bool;
	)
	
	for {
		if part, prefix, err = mf.Rder.ReadLine(); err != nil {
			log.Println(err);
      break;
    }
    buf.Write(part);

    ///获取当前游标位置
		mf.Pos, err = mf.Fh.Seek(1, 1);
		if err != nil{
			log.Println(err);
		}
		//log.Println("next begin: ", mf.Pos);
		
    if !prefix {
        lines = append(lines, buf.String());
        buf.Reset();
        break;
    }
  }
	if err == io.EOF {
	  //err = nil;
	  return;
	}
	return;
}