package main

import "io/ioutil"

func main(){
  b, err := ioutil.ReadFile("input.txt");
  if err != nil {
    panic(err);
  }

  err = ioutil.WriterFile("output.txt", b, 0644);
  if err != nil {
    panic(err);
  }
}

