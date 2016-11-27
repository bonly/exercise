package main

import (
	"fmt"
	"os/exec"
)

func command(tel string, text string) (string, error) {
	out, err := exec.Command("java", "-jar", "xwsms.jar", tel, text).Output();
	if err != nil {
		return "", err;
	}

	return string(out), nil;
}

func main() {
	ret, err := command("15360534225", "this is a test");
	if err != nil {
		fmt.Println("command: ", err);
		return;
	}
	fmt.Println(ret);
}