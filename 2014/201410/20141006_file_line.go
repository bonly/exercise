package main 

import(
"fmt"
"runtime"
"strings"
"flag"
)

func fileInfo(callDepth int) string {
	// Inspect runtime call stack
	pc := make([]uintptr, callDepth)
	runtime.Callers(callDepth, pc)
	f := runtime.FuncForPC(pc[callDepth-1])
	file, line := f.FileLine(pc[callDepth-1])

	// Truncate abs file path
	if slash := strings.LastIndex(file, "/"); slash >= 0 {
		file = file[slash+1:]
	}

	// Truncate package name
	funcName := f.Name()
	if slash := strings.LastIndex(funcName, "."); slash >= 0 {
		funcName = funcName[slash+1:]
	}

	return fmt.Sprintf("%s:%d %s -", file, line, funcName)
}

var line = flag.Int("l", 1, "call depth");

func main(){
	flag.Parse();
	fmt.Println(fileInfo(*line));	
}