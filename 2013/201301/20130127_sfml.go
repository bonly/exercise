package main 

import (
glfw "github.com/go-gl/glfw3"
     "fmt"
)

func errCallback(err glfw.ErrorCode, desc string){
	fmt.Println("%v: %v\n", err, desc);
}

func main(){
	glfw.SetErrorCallback(errCallback);

	if !glfw.Init(){
		panic("can't init glfw");
	}
	defer glfw.Terminate();

	window, err := glfw.CreateWindow(800, 600, "my window", nil, nil);
	if err != nil {
		panic(err);
	}
    
    window.MakeContextCurrent();

	for !window.ShouldClose() {
		window.SwapBuffers();
        glfw.PollEvents();
	}
}
//  most simple window
