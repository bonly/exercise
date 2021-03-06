
package main

import (
	"encoding/binary"
	"log"
	"fmt"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/app/debug"
	"golang.org/x/mobile/event"
	"golang.org/x/mobile/f32"
_	"golang.org/x/mobile/geom"
	"golang.org/x/mobile/gl"
	"golang.org/x/mobile/gl/glutil"
)

var (
	program  gl.Program
	position gl.Attrib
	color    gl.Uniform
	buf      gl.Buffer
	idb      gl.Buffer

)

func main() {
	app.Run(app.Callbacks{
		Draw:  draw,
		Touch: touch,
	})
}

func initGL() {
	var err error;
	program, err = glutil.CreateProgram(vertexShader, fragmentShader);
	if err != nil {
		log.Printf("error creating GL program: %v", err);
		return;
	}


//  var idx = []uint16{ 3, 2, 1, 3, 1, 0 };
//	adx := make([]byte, 6);
//	binary.LittleEndian.PutUint16(adx, idx);

//  idx := []uint16{3, 2, 1, 3, 1, 0};
//  var adx []byte = idx[:];
	
//	idx := []uint16{3,2,1,3,1,0};
	
	buf = gl.CreateBuffer();//创建buffer
	gl.BindBuffer(gl.ARRAY_BUFFER, buf); //绑定缓冲区
	//传递数据
	gl.BufferData(gl.ARRAY_BUFFER, triangleData, gl.STATIC_DRAW)

    idb = gl.CreateBuffer(); //创建索引缓冲  
    gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, idb); //绑定索引缓冲区
	//传递索引数据
    gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, idx, gl.STATIC_DRAW);
	
	position = gl.GetAttribLocation(program, "position")
	color = gl.GetUniformLocation(program, "color")
}

func touch(t event.Touch) {
}

func u16Bytes(byteOrder binary.ByteOrder, values ...uint16) []byte {
	le := false;
	switch byteOrder {
		case binary.BigEndian:
		case binary.LittleEndian:
			le = true;
		default:
			panic(fmt.Sprintf("invalid byte order %v", byteOrder));
	}

	b := make([]byte, 2*len(values));
	if le {
		for i, v := range values {
			b[2*i+0] = byte(v >> 0);
			b[2*i+1] = byte(v >> 8);
		}
	} else {
		for i, v := range values {
			b[2*i+0] = byte(v >> 8);
			b[2*i+1] = byte(v >> 0);
		}
	}
	return b;
}

func draw() {
	if program.Value == 0 {
		initGL();
	}

	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)
//	gl.Viewport(0, 0, 120, 120);

	gl.UseProgram(program); //使用glsl

	gl.BindBuffer(gl.ARRAY_BUFFER, buf); //绑定缓冲区

    //定义每次从数组中取的数量数量（以每个点所包括的数据为单位)
	gl.VertexAttribPointer(position, 2, gl.FLOAT, false, 0, 0);
	gl.EnableVertexAttribArray(position); //gl.position ==> buf

//	sz := gl.GetBufferParameteri(gl.ARRAY_BUFFER, gl.BUFFER_SIZE); 
//	gl.DrawArrays(gl.TRIANGLE_STRIP, 0, sz);
   
    gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, idb); //绑定索引缓冲区
    sz := gl.GetBufferParameteri(gl.ELEMENT_ARRAY_BUFFER, gl.BUFFER_SIZE);
    gl.DrawElements(gl.TRIANGLES, sz, gl.UNSIGNED_SHORT,0);
	
	log.Printf("===================size: %v===========", sz);
	
	gl.DisableVertexAttribArray(position);

	debug.DrawFPS();
}

var triangleData = f32.Bytes(binary.LittleEndian,
-0.5, 0.5,  // 0.0,  //vertex 0
-0.5, -0.5, // 0.0,  //vertex 1
0.5,  -0.5, // 0.0,  //vertex 2
0.5,  0.5,  // 0.0,  //vertex 3
);

var idx = u16Bytes(binary.LittleEndian, 
3,2,1,3,1,0,
);

const vertexShader = `
#version 100

attribute vec4 position;
void main() {
	gl_Position = position;
}`

const fragmentShader = `
#version 100
precision mediump float;
uniform vec4 color;
void main() {
	gl_FragColor = vec4(1,0,0,1);
}`
