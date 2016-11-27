package main

import (
"fmt"
"net/url"
)

func main(){
  str := "json_data=%7B%22roomId%22%3A%221%22%2C%22lodger_name%22%3A%22%22%2C%22hotelId%22%3A%221%22%2C%22lodger_id%22%3A%22%22%7D"

  	u, err := url.Parse(str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u.Path)
	fmt.Println(u.RawPath)
	fmt.Println(u.String())
}
