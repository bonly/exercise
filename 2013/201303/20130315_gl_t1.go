package main

import (
	"encoding/binary"
	"fmt"

	"github.com/go-gl/glfw/v3.1/glfw"
	"golang.org/x/mobile/f32"
	"golang.org/x/mobile/gl"
	"golang.org/x/mobile/gl/glutil"
)

const (
	vshader = `
#version 120
attribute vec4 coord;
void main(void) {
  gl_Position = coord;
}
`
	fshader = `
#version 120
void main(void) {
  gl_FragColor[0] = 0.0;
  gl_FragColor[1] = 0.0;
  gl_FragColor[2] = 1.0;
}`
)

func onError(err glfw.ErrorCode, desc string) {
	fmt.Printf("%v: %v\n", err, desc)
}

func onKey(w *glfw.Window, key glfw.Key, scancode int,
	action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		w.SetShouldClose(true)
	}
}

func onResize(w *glfw.Window, width, height int) {
	gl.Viewport(0, 0, width, height)

	//heightUnif.Uniform1i(height)
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	w, err := glfw.CreateWindow(640, 480, "my first triangle", nil, nil)
	if err != nil {
		panic(err)
	}
	defer w.Destroy()
	w.SetKeyCallback(onKey)

	w.MakeContextCurrent()
	glfw.SwapInterval(1)

	//gl.Init()

	prog, err := glutil.CreateProgram(vshader, fshader)
	if err != nil {
		panic(err)
	}
	defer gl.DeleteProgram(prog)
	w.SetSizeCallback(onResize)
	w.GetSize()

	triangleData := []float32{
		+0.0, +0.8, 0, 1,
		-0.8, -0.8, 0, 1,
		+0.8, -0.8, 0, 1,
	}
	buf := gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, buf)
	gl.BufferData(gl.ARRAY_BUFFER,
		f32.Bytes(binary.LittleEndian, triangleData...),
		gl.STATIC_DRAW,
	)

	coord := gl.GetAttribLocation(prog, "coord")

	gl.UseProgram(prog)

	for !w.ShouldClose() {
		gl.ClearColor(0, 0, 0, 1)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		gl.BindBuffer(gl.ARRAY_BUFFER, buf)
		gl.EnableVertexAttribArray(coord)
		gl.VertexAttribPointer(coord, 4, gl.FLOAT, false, 0, 0)
		gl.DrawArrays(gl.TRIANGLES, 0, len(triangleData))

		w.SwapBuffers()
		glfw.PollEvents()
	}

	//gl.ProgramUnuse()
}