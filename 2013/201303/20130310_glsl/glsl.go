package main
import (
	"log"
	"encoding/binary"

	"golang.org/x/mobile/f32"
	"golang.org/x/mobile/event"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/gl"
	"golang.org/x/mobile/gl/glutil"
)

var (
	program  gl.Program;
	buf      gl.Buffer;
	position gl.Attrib;
	color    gl.Uniform;
)

//四边形数据
var dr = f32.Bytes(binary.LittleEndian,
	0, 0,
	0, 0.1,
	0.1, 0,
	0.1, 0.1,
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

	buf = gl.CreateBuffer(); //创建gl缓冲区
	gl.BindBuffer(gl.ARRAY_BUFFER,buf); //绑定缓冲区
        //传递缓冲区内容
	gl.BufferData(gl.ARRAY_BUFFER, dr, gl.STATIC_DRAW); 

	position = gl.GetAttribLocation(program, "position");
	color = gl.GetUniformLocation(program, "color");
}

func stop(){
	gl.DeleteProgram(program);
	gl.DeleteBuffer(buf);
}

func draw(){
	gl.ClearColor(0, 0, 0, 1);
	gl.Clear(gl.COLOR_BUFFER_BIT);
	gl.UseProgram(program);
	drawButton(3, 4);
}

func drawButton(x, y float32){
	gl.Uniform4f(color, 0, 0, 1, 1); //上颜色 红 绿 蓝 透明度

	gl.BindBuffer(gl.ARRAY_BUFFER, buf); //绑定操作的缓冲对象
	gl.VertexAttribPointer(position, 2, gl.FLOAT, false, 0,0);
	/* (Index, Size, Type, Norm, Stride, Offset)
	Index: 当前绑定缓冲区所映射的属性索引
	Size: 存储于当前绑定缓冲区中的、各顶点的数据值数量
	Type: 确定存储于当前绑定缓冲区中的数据值类型，即FIXED/BYTE/UNSIGNED_BYTE/FLOAT/SHORT/UNSIGNED_SHORT
	Norm: 该参数可设置为true或false,并用于处理范围之外的数值转换问题。在实际操作过程中，该参数设置为false
	Stride: 若Stride设置为0，则数据元素在缓冲区中的以连续方式存储
	Offset: 针对相关属性，该参数表示缓冲区中数据值读取时的起始位置，且通常设置为0，即从缓冲区中的首个元素处开始读取数值
	*/
        //启用顶点着色器的各项属性
	gl.EnableVertexAttribArray(position); 
	/* 
	DrawArrays(Mode, First, Count) 使用顶点数据生成几何体
        Mode: 表示为渲染的图元类型POINTS/LINE_STRIP/LINE_LOOP/LINES/TRIANGLE_STRIP/TRIANGLE_FAN/TRIANGLES
	First: 定义活动数组中的起始元素
	Count: 定义了渲染时的元素数量
	*/
	gl.DrawArrays(gl.TRIANGLE_STRIP, 0, 4);
	/*
	DrawElements(Mode, Count, Type, Offset) 使用索引访问顶点数据缓冲区生成几何体
        Mode: 表示为渲染的图元类型POINTS/LINE_STRIP/LINE_LOOP/LINES/TRIANGLE_STRIP/TRIANGLE_FAN/TRIANGLES
	Count: 定义了渲染时的元素数量
	Type: 确定索引值类型，当处理（整数）索引时，该参数应为SIGNED_BYTE/UNSIGNED_SHORT
	Offset: 表示缓冲区元素的渲染起始点，且通常为首个元素（即0值）
	*/
	gl.DisableVertexAttribArray(position);
}

const vertexShader = `
#version 100
attribute vec4 position;
void main(){
	gl_Position = position;
}`;

const fragmentShader = `
#version 100
uniform vec4 color;
void main(){
	gl_FragColor = color;
}`;


