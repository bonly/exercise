package main

import (
	"encoding/binary"
	"fmt"
	"runtime"

	"github.com/go-gl/glfw/v3.1/glfw"
	"golang.org/x/mobile/f32"
	"golang.org/x/mobile/gl"
	"golang.org/x/mobile/gl/glutil"
)

type context struct {
	w    *glfw.Window
	prog gl.Program

	posbuf  gl.Buffer
	pos     gl.Attrib
	posdata []float32

	colbuf  gl.Buffer
	col     gl.Attrib
	coldata []float32

	fade gl.Uniform
}

func (ctx *context) Delete() {
	gl.DeleteProgram(ctx.prog)
	gl.DeleteBuffer(ctx.posbuf)
	gl.DeleteBuffer(ctx.colbuf)
}

func onError(err glfw.ErrorCode, desc string) {
	fmt.Printf("%v: %v\n", err, desc)
}

func onKey(w *glfw.Window, key glfw.Key, scancode int,
	action glfw.Action, mods glfw.ModifierKey) {
	switch {
	case key == glfw.KeyEscape && action == glfw.Press,
		key == glfw.KeyQ && action == glfw.Press:
		w.SetShouldClose(true)
	}
}

func onResize(window *glfw.Window, w, h int) {
	gl.Viewport(0, 0, w, h)
}

func display(ctx context) {
	// clear the background as black
	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.UseProgram(ctx.prog)

	gl.Uniform1f(ctx.fade, 0.5)

	gl.EnableVertexAttribArray(ctx.col)
	gl.EnableVertexAttribArray(ctx.pos)

	gl.BindBuffer(gl.ARRAY_BUFFER, ctx.colbuf)
	gl.VertexAttribPointer(ctx.col, 3, gl.FLOAT, false, 0, 0)

	gl.BindBuffer(gl.ARRAY_BUFFER, ctx.posbuf)
	gl.VertexAttribPointer(ctx.pos, 4, gl.FLOAT, false, 0, 0)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)

	gl.DisableVertexAttribArray(ctx.col)
	gl.DisableVertexAttribArray(ctx.pos)

	// display result
	ctx.w.SwapBuffers()
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

	w.MakeContextCurrent()
	glfw.SwapInterval(1)

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	ctx := context{
		w: w,
		posdata: []float32{
			+0.0, +0.8, 0, 1,
			-0.8, -0.8, 0, 1,
			+0.8, -0.8, 0, 1,
		},
		coldata: []float32{
			1.0, 1.0, 0.0,
			0.0, 0.0, 1.0,
			1.0, 0.0, 0.0,
		},
	}
	ctx.prog, err = glutil.CreateProgram(
		`
#version 120

attribute vec4 coord;
attribute vec3 v_color;
varying   vec3 f_color;

void main(void) {
  gl_Position = coord;
  f_color = v_color;
}
`,
		`
#version 120
varying vec3 f_color;
uniform float fade;
void main(void) {
  gl_FragColor = vec4(f_color.x, f_color.y, f_color.z, fade);
}
`)
	if err != nil {
		panic(err)
	}
	defer ctx.Delete()

	ctx.posbuf = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, ctx.posbuf)
	gl.BufferData(gl.ARRAY_BUFFER,
		f32.Bytes(binary.LittleEndian, ctx.posdata...),
		gl.STATIC_DRAW,
	)
	ctx.pos = gl.GetAttribLocation(ctx.prog, "coord")

	ctx.colbuf = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, ctx.colbuf)
	gl.BufferData(gl.ARRAY_BUFFER,
		f32.Bytes(binary.LittleEndian, ctx.coldata...),
		gl.STATIC_DRAW,
	)
	ctx.col = gl.GetAttribLocation(ctx.prog, "v_color")
	ctx.fade = gl.GetUniformLocation(ctx.prog, "fade")

	ctx.w.SetSizeCallback(onResize)
	ctx.w.SetKeyCallback(onKey)

	for !ctx.w.ShouldClose() {
		display(ctx)
		glfw.PollEvents()

	}

}

func init() {
	runtime.LockOSThread()
}