func main(){
  if len(os.Args) > 1{
     Tree(os.Args[1], 1, map[int]bool{1:true});
  }
}

func Tree(dirname string, curHier int, hierMap map[int]bool) error{
  dirAbs, err := filepath.Abs(dirname);
  if err != nil{
    return err;
  }
  fileInfos, err := ioutil.ReadDir(dirAbs);
  if err != nil{
    return err;
  }
  
  fileNum := len(fileInfos);
  for i, fileInfo := range fileInfos{
     for j:=1;