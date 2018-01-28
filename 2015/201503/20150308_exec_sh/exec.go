package main 

import (
	"fmt"
	"os/exec"
	"os"
	"bufio"
	// "io/ioutil"  //一般io会连一个bytes.Buffer来读取
	// "bytes"
)

func main(){
	// cmd := exec.Command("./child.sh")
	cmd_argv := []string{"-l", "-ha"};
	// cmd := exec.Command("ls", cmd_argv...)
	cmd := exec.Command("./child.sh", cmd_argv...)
	cmdStdin, err := cmd.StdinPipe()
	exitOnError("cmd.StdinPipe()", err)
	cmdStdout, err := cmd.StdoutPipe()
	// cmdStdout, err := cmd.StderrPipe()
	exitOnError("cmd.StdoutPipe()", err)

	err = cmd.Start()
	exitOnError("cmd.Start()", err)

	// scanner := bufio.NewScanner(cmdStdout)
	// go func() {
	// 	for scanner.Scan() {
	// 		fmt.Printf("docker build out | %s\n", scanner.Text())
	// 	}
	// }()



	iBuf := bufio.NewWriter(cmdStdin)
	oBuf := bufio.NewReader(cmdStdout)

	// out, _ := ioutil.ReadAll(cmdStdout); //会err或eof了才结束，所以不适用
	// fmt.Printf("get: %s\n", string(out)); //wait

	// out := new (bytes.Buffer);
	// out.ReadFrom(cmdStdout); //会err或eof了才结束，所以不适用
	// out.ReadString('\n'); //无法绑定io.Reader，也不适用
	// fmt.Printf("get: %s\n", out.String());

	line, err := oBuf.ReadString('\n')
	exitOnError("oBuf.ReadString()", err)
	fmt.Printf("parent read: %s\n", line) // wait

	n, err := iBuf.WriteString("Knock, Knock\n")
	exitOnError("iBuf.WriteString(Knock, Knock)", err)
	fmt.Fprintf(os.Stderr, "parent says: Knock, Knock. Wrote %d bytes\n", n)
	iBuf.Flush();


	line, err = oBuf.ReadString('\n')
	exitOnError("oBuf.ReadString()", err)
	fmt.Printf("parent read: %s\n", line) // Who's there?

	n, err = iBuf.WriteString("Canoe\n")
	exitOnError("iBuf.WriteString(Canoe)", err)
	fmt.Fprintf(os.Stderr, "parent says: Canoe. Wrote %d bytes\n", n)
	iBuf.Flush();

	line, err = oBuf.ReadString('\n')
	exitOnError("oBuf.ReadString()", err)
	fmt.Printf("parent read: %s\n", line) // Canoe who?

	n, err = iBuf.WriteString("Canoe help me figure this out?\n")
	exitOnError("iBuf.WriteString(Canoe)", err)
	fmt.Fprintf(os.Stderr, "parent says: Canoe. Wrote %d bytes\n", n)
	iBuf.Flush();

	line, err = oBuf.ReadString('\n')
	exitOnError("oBuf.ReadString()", err)
	fmt.Printf("parent read: %s\n", line) // Groan...

	// err = cmd.Start()  // 必须在前面执行start
	// exitOnError("cmd.Start()", err)

	cmd.Wait()
	exitOnError("cmd.Wait()", err)	
}

func exitOnError(head string, err error){
	if err != nil{
	   fmt.Printf("%s: %s\n", head, err.Error());
	   os.Exit(1);
	}
}

/*
io.Reader 可用 bytes.Buffer关联去读； ioutil.ReadAll只能一次读完;
		  可用 bufio.NewReader 关联去读

逆转：
strings.NewReader 和 bytes.NewBufferString 及 bufio.NewReader(os.Stdin)
可转得io.Reader

*/