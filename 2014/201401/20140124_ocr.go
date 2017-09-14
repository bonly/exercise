package main

import (
	"fmt"
	"github.com/otiai10/gosseract"
)

func main() {
	// This is the simlest way :)
	out := gosseract.Must(gosseract.Params{Src: "your/img/file.png",Languages:"eng+heb"})
	fmt.Println(out)

	// Using client
	client, _ := gosseract.NewClient()
	out, _ = client.Src("your/img/file.png").Out()
	fmt.Println(out)
}
