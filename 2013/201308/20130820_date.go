package main

import (
	"time"
	"fmt"
);

func main(){
      baseTime := time.Date(1980, 1, 6, 0, 0, 0, 0, time.UTC)
      date := baseTime.Add(1722*7*24*time.Hour + 24*time.Hour + 66355*time.Second)
      fmt.Println(date)

      const layout = "20060102";
      nowTime := time.Now();
      date = nowTime.AddDate(0, 0, -1);
      fmt.Println(date.Format(layout));
}

