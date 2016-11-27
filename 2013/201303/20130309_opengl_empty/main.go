package main
import (
	"log"

	"golang.org/x/mobile/event"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/gl"
	"golang.org/x/mobile/gl/glutil"
)

var (
	program gl.Program
)

func main(){
	app.Run(app.Callbacks{
		Start: start,
		Stop: stop,
		Draw: draw,
		Touch: touch,
	});
}


func touch(t event.Touch){
}

func start(){
	var err error;
	program, err = glutil.CreateProgram(vertexShader, 
	fragmentShader);

	if err != nil{
		log.Printf("error create GL program: %v", err);
		return;
	}

}

func stop(){
	gl.DeleteProgram(program);
}

func draw(){
	gl.ClearColor(0, 0, 0, 1);
	gl.Clear(gl.COLOR_BUFFER_BIT);
	gl.UseProgram(program);
}

const vertexShader = `
#version 100
void main(){
	gl_Position = 0; //opengl内置变量，把经过变换的点写入该变量，传入渲染管线中进行后续处理
}`;

const fragmentShader = `
void main(){
	gl_FragColor = 0;
}`;


