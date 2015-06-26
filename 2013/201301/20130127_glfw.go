package main 

import (
glfw "github.com/go-gl/glfw3"
     // "github.com/tedsta/gosfml"
)

var window *glfw.Window;

func main(){
	glfw.Init();
	window, err := glfw.CreateWindow(800, 600, "my window", nil, nil);
	if err != nil {
		panic(err);
	}
    
	for !window.ShouldClose() {
		// window.SwapBuffers();
	}
}
//  most simple window
