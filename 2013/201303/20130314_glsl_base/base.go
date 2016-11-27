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
	pic_buf  gl.Buffer;
	pic      gl.Attrib;
//	color    gl.Uniform;
)

//四边形数据
var pic_data = f32.Bytes(binary.LittleEndian,
	0,   0,
	0,   0.5,
	0.5, 0,
	0.5, 0.5,
) //都大于0时， 可见是个长方形
//var dr = f32.Bytes(binary.LittleEndian,
//	0,   0,
//	0,   0.1,
//	0.1, 0,
//	0.1, 0.1,
//)
//var dr = f32.Bytes(binary.LittleEndian,
//-0.5, 0.5,  // 0.0,  //vertex 0
//-0.5, -0.5, // 0.0,  //vertex 1
//0.5,  -0.5, // 0.0,  //vertex 2
//0.5,  0.5,  // 0.0,  //vertex 3
//)  //自动连线到0点，成了一个开口的缺口四边形

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

	pic_buf = gl.CreateBuffer(); //创建gl缓冲区
	gl.BindBuffer(gl.ARRAY_BUFFER,pic_buf); //绑定缓冲区
    //传递缓冲区内容
	gl.BufferData(gl.ARRAY_BUFFER, pic_data, gl.STATIC_DRAW); 

//	position = gl.GetAttribLocation(program, "position");//从程序GLSL的变量关联
//	color = gl.GetUniformLocation(program, "color");//从程序GLSL的变量关联
   
    gl.Viewport(0, 0, int(app.GetConfig().Width), 
	                  int(app.GetConfig().Height));
					  
//  gl.MatrixMode(gl.PROJECTION); GLSL2.0弃用了
//	gl.LoadIdentity(); GLSL2.0弃用了

}

func stop(){
	gl.DeleteProgram(program); //清理GLSL
	gl.DeleteBuffer(pic_buf); //清理缓冲
}

func draw(){
	gl.ClearColor(0, 0, 0, 1);
	gl.Clear(gl.COLOR_BUFFER_BIT);
	
	gl.Disable(gl.DEPTH_TEST);//GLSL2.0开始可以关闭这个
	
	gl.UseProgram(program);
	drawButton(3, 4);
}

func drawButton(x, y float32){
//	gl.Uniform4f(color, 0, 0, 1, 1); //上颜色 红 绿 蓝 透明度

	gl.BindBuffer(gl.ARRAY_BUFFER, pic_buf); //绑定操作的缓冲对象
	gl.VertexAttribPointer(pic, 2, gl.FLOAT, false, 0,0);
	/* (Index, Size, Type, Norm, Stride, Offset)
	Index: 当前绑定缓冲区所映射的属性索引
	Size: 存储于当前绑定缓冲区中的、各顶点的数据值数量
	Type: 确定存储于当前绑定缓冲区中的数据值类型，即FIXED/BYTE/UNSIGNED_BYTE/FLOAT/SHORT/UNSIGNED_SHORT
	Norm: 该参数可设置为true或false,并用于处理范围之外的数值转换问题。在实际操作过程中，该参数设置为false
	Stride: 若Stride设置为0，则数据元素在缓冲区中的以连续方式存储
	Offset: 针对相关属性，该参数表示缓冲区中数据值读取时的起始位置，且通常设置为0，即从缓冲区中的首个元素处开始读取数值
	*/
    //启用顶点着色器的各项属性
	gl.EnableVertexAttribArray(pic); 
	/* 
	DrawArrays(Mode, First, Count) 使用顶点数据生成几何体
    Mode: 表示为渲染的图元类型POINTS/LINE_STRIP/LINE_LOOP/LINES/TRIANGLE_STRIP/TRIANGLE_FAN/TRIANGLES
	First: 定义活动数组中的起始元素
	Count: 定义了渲染时的元素数量
	*/
	sz := gl.GetBufferParameteri(gl.ARRAY_BUFFER, gl.BUFFER_SIZE); 
	gl.DrawArrays(gl.TRIANGLE_STRIP, 0, sz);
	/*
	DrawElements(Mode, Count, Type, Offset) 使用索引访问顶点数据缓冲区生成几何体
    Mode: 表示为渲染的图元类型POINTS/LINE_STRIP/LINE_LOOP/LINES/TRIANGLE_STRIP/TRIANGLE_FAN/TRIANGLES
	Count: 定义了渲染时的元素数量
	Type: 确定索引值类型，当处理（整数）索引时，该参数应为SIGNED_BYTE/UNSIGNED_SHORT
	Offset: 表示缓冲区元素的渲染起始点，且通常为首个元素（即0值）
	*/
	gl.DisableVertexAttribArray(pic);
}

/*顶点shader负责完成顶点变换。这里将按照固定功能的方程完成顶点变换
固定功能流水线中一个顶点通过模型视图矩阵以及投影矩阵进行变换，使用如下公式：
vTrans = projection * modelview *incomingVertex 

首先GLSL需要访问OpenGL状态，获得公式中的前两个矩阵。前面讲过，GLSL可以获取某些OpenGL状态信息的，这两个矩阵当然包括在内。
可以通过预先定义的一致变量来获取它们：
uniform mat4 gl_ModelViewMatrix;  
uniform mat4 gl_ProjectionMatrix;

接下来需要得到输入的顶点。通过预先定义的属性变量，所有的顶点将可以一个个传入顶点shader中。
attribute vec4 gl_Vertex;  

为了输出变换后的顶点，shader必须写入预先定义的vec4型变量gl_Position中，注意这个变量没有修饰符。
现在我们可以写一个仅仅进行顶点变换的顶点shader了。注意所有其他功能都将丧失，比如没有光照计算。
顶点shader必须有一个main函数，如下面的代码所示：
void main(){
    gl_Position =gl_ProjectionMatrix * gl_ModelViewMatrix * gl_Vertex;
}

上面代码中变换每个顶点时，投影矩阵都将乘上模型视图矩阵，这显然非常浪费时间，因为这些矩阵不是随每个顶点变化的。注意这些矩阵是一致变量。
GLSL提供一些派生的矩阵，也就是说gl_ModelViewProjectionMatrix是上面两个矩阵的乘积，所以顶点shader也可以写成下面这样：
void main(){
    gl_Position =gl_ModelViewProjectionMatrix * gl_Vertex;
}

上面的操作能够获得和固定功能流水线相同的结果吗？理论上是如此，但实际上对顶点变换操作的顺序可能会不同。
顶点变换通常在显卡中是高度优化的任务，所以有一个利用了这种优化的特定函数用来处理这个任务。这个神奇的函数如下：
vec4 ftransform(void);

使用这个函数的另一个原因是float数据类型的精度限制。
由于数据精度的限制，当使用不同的顺序计算时，可能得到不同的结果，因此GLSL提供这个函数保证获得最佳性能的同时，还能得到与固定功能流水线相同的结果。
这个函数按照与固定功能相同的步骤对输入顶点进行变换，然后返回变换后的顶点。所以shader可以重新写成如下形式：
void main()  {  
    gl_Position =ftransform();  
}  
*/

/*实际上GLSL2.0开始需要自己通过GLSL设置投影和视图矩阵变量,即:
uniform mat4 uMVPMatrix;
attribute vec4 vPosition;
void main(){
    gl_Position = uMVPMatrix * vPosition;
}
*/
const vertexShader = `
#version 100
attribute vec4 pic;
void main(){
	gl_Position = vec4(pic.x, pic.y, 0.0, 1.0);
}`;


/*
片断shader也有预先定义的变量gl_FragColor，可以向其中写入片断的颜色值
*/
const fragmentShader = `
#version 100
uniform vec4 color;
void main(){
	gl_FragColor = vec4(1,0,0,1);
}`;


