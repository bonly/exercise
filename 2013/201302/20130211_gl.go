// -*- coding: utf-8 -*-
// 参考: 床井研究室 - (4) シェーダの追加
// http://marina.sys.wakayama-u.ac.jp/~tokoi/?date=20120909
package main

import (
	gl "github.com/go-gl/gl"
	"github.com/go-gl/glfw"
	"log"
	"unsafe"
)

const (
	Title  = "Hello Shader"
	Width  = 640
	Height = 480
)

const vertexShaderSource = `
#version 130
in vec4 pv;
void main(void)
{
	gl_Position = pv;
}
`

const fragmentShaderSource = `
#version 130
out vec4 fc;
void main(void)
{
	fc = vec4(1.0, 0.0, 0.0, 1.0);
}
`

var vertices = [][2]gl.Float{
	{-0.5, -0.5},
	{0.5, -0.5},
	{0.5, 0.5},
	{-0.5, 0.5},
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatal(err)
	}
	defer glfw.Terminate()

	glfw.OpenWindowHint(glfw.OpenGLVersionMajor, 3)
	glfw.OpenWindowHint(glfw.OpenGLVersionMinor, 0)
	glfw.OpenWindowHint(glfw.WindowNoResize, 1)

	if err := glfw.OpenWindow(Width, Height, 0, 0, 0, 0, 16, 0, glfw.Windowed); err != nil {
		log.Fatal(err)
	}
	defer glfw.CloseWindow()

	glfw.SetSwapInterval(1)
	glfw.SetWindowTitle(Title)

	if err := gl.Init(); err != nil {
		log.Fatal(err)
	}

	program := createProgram(vertexShaderSource, "pv", fragmentShaderSource, "fc")
	vao := createObject(vertices)

	for glfw.WindowParam(glfw.Opened) == 1 {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.UseProgram(program)
		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.LINE_LOOP, 0, gl.Sizei(len(vertices)))
		gl.BindVertexArray(0)
		gl.UseProgram(0)
		glfw.SwapBuffers()
	}
}

func createProgram(vsrc, pv, fsrc, fc string) gl.Uint {
	glvsrc := gl.GLString(vertexShaderSource)
	defer gl.GLStringFree(glvsrc)
	vobj := gl.CreateShader(gl.VERTEX_SHADER)
	gl.ShaderSource(vobj, 1, &glvsrc, nil)
	gl.CompileShader(vobj)
	defer gl.DeleteShader(vobj)

	glfsrc := gl.GLString(fragmentShaderSource)
	defer gl.GLStringFree(glfsrc)
	fobj := gl.CreateShader(gl.FRAGMENT_SHADER)
	gl.ShaderSource(fobj, 1, &glfsrc, nil)
	gl.CompileShader(fobj)
	defer gl.DeleteShader(fobj)

	program := gl.CreateProgram()
	gl.AttachShader(program, vobj)
	gl.AttachShader(program, fobj)

	glpv := gl.GLString(pv)
	defer gl.GLStringFree(glpv)
	gl.BindAttribLocation(program, 0, glpv)
	glfc := gl.GLString(fc)
	defer gl.GLStringFree(glfc)
	gl.BindFragDataLocation(program, 0, glfc)
	gl.LinkProgram(program)

	return program
}

func createObject(vertices [][2]gl.Float) gl.Uint {
	var vao gl.Uint
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	defer gl.BindVertexArray(0)

	var vbo gl.Uint
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	defer gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(int(unsafe.Sizeof([2]gl.Float{}))*len(vertices)), gl.Pointer(&vertices[0]), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 2, gl.FLOAT, gl.FALSE, 0, nil)
	gl.EnableVertexAttribArray(0)

	return vao
}