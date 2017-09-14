package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if !Exists("redis.tar.gz") {
		Download("http://download.redis.io/redis-stable.tar.gz", "redis.tar.gz")
	}
	if Exists("redis") {
		V(os.RemoveAll("redis"))
	}
	V(os.MkdirAll("redis", 0755))
	Run("tar zxvf redis.tar.gz -C redis --strip-components=1")
	V(os.Chdir("redis"))
	Run("make")
}

func V(err error) {
	if err != nil {
		panic(err)
	}
}

func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		V(err)
	}
	return true
}

func Download(fileURL string, filename string) {
	resp, err := http.Get(fileURL)
	V(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	V(err)
	V(ioutil.WriteFile(filename, body, 0644))
}

//拆分参数
func Run(command string) {
	fmt.Printf("run: %s\n", command)
	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("command failed: %s\n", command)
		V(err)
	}
}

//带结果值输出
func RunExit(command string) int {
	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			return exiterr.Sys().(syscall.WaitStatus).ExitStatus()
		} else {
			fmt.Printf("command failed: %s\n", command)
			panic(err)
		}
	}
	return 0
}

//带执行结果输出
func Capture(command string) string {
	args := strings.Split(command, " ")
	output, err := exec.Command(args[0], args[1:]...).CombinedOutput()
	if err != nil {
		fmt.Printf("command failed: %s\n output=%s\n", command, output)
		panic(err)
	}
	return string(output)
}

//可处理命令中带有空格git commit -am "a test commit"
func Run(command string, args ...string) {
	var cmd *exec.Cmd
	if len(args) == 0 {
		fmt.Printf("run: %s\n", command)
		args := strings.Split(command, " ")
		cmd = exec.Command(args[0], args[1:]...)
	} else {
		fmt.Printf("run: %s %s\n", command, strings.Join(args, " "))
		cmd = exec.Command(command, args...)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("command failed: %s\n", command)
		panic(err)
	}
}

//http://goroutines.com/shell