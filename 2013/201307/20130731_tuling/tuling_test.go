package main 

import (
"testing"
"log"
)

func Test_tyxw(ts *testing.T){
    str := askTuling("体育新闻");
    log.Println(str);
    ts.Log(str);
}
