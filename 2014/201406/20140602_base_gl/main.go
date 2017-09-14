package main 

import (
"golang.org/x/mobile/gl"
"golang.org/x/mobile/exp/gl/glutil"
"golang.org/x/mobile/app"
"golang.org/x/mobile/event/touch"
"golang.org/x/mobile/geom"
"golang.org/x/mobile/exp/f32"
"log"
"encoding/binary"
)

var (
	program 	gl.Program;
	position 	gl.Attrib;
	offset 		gl.Uniform;
	color		gl.Uniform;
	buf			gl.Buffer;

	green 		float32;
	touchLoc	geom.Point;

	glx			gl.Context;
);

func initGL(){
	var err error;
	//编译顶点着色器和颜色着色器
	program, err = glutil.CreateProgram(glx, vertexShader, fragmentShader);
	if err != nil{
		log.Printf("create GL program: %v", err);
		return;
	}

	buf = glx.CreateBuffer(); //构建OAB
	glx.BindBuffer(gl.ARRAY_BUFFER, buf);
	glx.BufferData(gl.ARRAY_BUFFER, triangleData, gl.STATIC_DRAW);

	//绑定编译器中的变量
	position = glx.GetAttribLocation(program, "position");
	color = glx.GetUniformLocation(program, "color");
	offset = glx.GetUniformLocation(program, "offset");

	//初始化触模点
	// touchLoc = geom.Point{geom.Width / 2, geom.Height / 2};
}

func get_touch(t touch.Event){
	// touchLoc = g{t.X, t.Y};
}

func draw(){
	if program.Value == 0{
		initGL(); //初始化
	}

	glx.ClearColor(1, 0, 0, 1);
	glx.Clear(gl.COLOR_BUFFER_BIT);
}

func main(){
	app.Main(app.Callbacks{
		Draw: draw,
		Touch: get_touch,
	});
}

var triangleData = f32.Bytes(binary.LittleEndian,
	0.0, 0.4, 0.0, // top left
	0.0, 0.0, 0.0, // bottom left
	0.4, 0.0, 0.0, // bottom right
);

const vertexShader = `#version 100
uniform vec2 offset;

attribute vec4 position;
void main() {
	// offset comes in with x/y values between 0 and 1.
	// position bounds are -1 to 1.
	vec4 offset4 = vec4(2.0*offset.x-1.0, 1.0-2.0*offset.y, 0, 0);
	gl_Position = position + offset4;
}`

const fragmentShader = `#version 100
precision mediump float;
uniform vec4 color;
void main() {
	gl_FragColor = color;
}`