package main

import (
	"errors"
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	WindowWidth  = 500
	WindowHeight = 500
)

type Box struct {
	x      float64
	y      float64
	width  float64
	height float64
	color  [4]float64
}

type Bar struct {
	modified bool
	boxes    []Box

	x      float64
	y      float64
	width  float64
	height float64

	vertices []float64
	colors   []float64

	program uint32

	vao uint32
	vbo uint32
	colorVbo uint32
}

func NewBar(program uint32, width, pos float64) *Bar {
	b := Bar{}
	b.modified = true
	b.width = width
	b.x = pos

	gl.GenVertexArrays(1, &b.vao)
	gl.BindVertexArray(b.vao)
	gl.GenBuffers(1, &b.vbo)
	gl.GenBuffers(1, &b.colorVbo)
	b.program = program

	return &b
}

func (bar *Bar) AddValue(value, r, g, b, a float64) {
	bar.modified = true

	bar.boxes = append(bar.boxes, Box{bar.x, bar.height, bar.width, value, [4]float64{r, g, b, a}})
	bar.height += value
}

func (bar *Bar) Update() {
	if !bar.modified {
		return
	}

	bar.vertices = make([]float64, len(bar.boxes)*18)
	bar.colors = make([]float64, len(bar.boxes)*24)

	for i, box := range bar.boxes {
		// XXX ugly shit
		bar.vertices[i*18] = box.x
		bar.vertices[i*18+1] = box.y
		bar.vertices[i*18+2] = 0
		bar.vertices[i*18+3] = box.x + box.width
		bar.vertices[i*18+4] = box.y + box.height
		bar.vertices[i*18+5] = 0
		bar.vertices[i*18+6] = box.x
		bar.vertices[i*18+7] = box.y + box.height
		bar.vertices[i*18+8] = 0
		bar.vertices[i*18+9] = box.x
		bar.vertices[i*18+10] = box.y
		bar.vertices[i*18+11] = 0
		bar.vertices[i*18+12] = box.x + box.width
		bar.vertices[i*18+13] = box.y
		bar.vertices[i*18+14] = 0
		bar.vertices[i*18+15] = box.x + box.width
		bar.vertices[i*18+16] = box.y + box.height
		bar.vertices[i*18+17] = 0

		bar.colors[i*24] = box.color[0]
		bar.colors[i*24+1] = box.color[1]
		bar.colors[i*24+2] = box.color[2]
		bar.colors[i*24+3] = box.color[3]
		bar.colors[i*24+4] = box.color[0]
		bar.colors[i*24+5] = box.color[1]
		bar.colors[i*24+6] = box.color[2]
		bar.colors[i*24+7] = box.color[3]
		bar.colors[i*24+8] = box.color[0]
		bar.colors[i*24+9] = box.color[1]
		bar.colors[i*24+10] = box.color[2]
		bar.colors[i*24+11] = box.color[3]
		bar.colors[i*24+12] = box.color[0]
		bar.colors[i*24+13] = box.color[1]
		bar.colors[i*24+14] = box.color[2]
		bar.colors[i*24+15] = box.color[3]
		bar.colors[i*24+16] = box.color[0]
		bar.colors[i*24+17] = box.color[1]
		bar.colors[i*24+18] = box.color[2]
		bar.colors[i*24+19] = box.color[3]
		bar.colors[i*24+20] = box.color[0]
		bar.colors[i*24+21] = box.color[1]
		bar.colors[i*24+22] = box.color[2]
		bar.colors[i*24+23] = box.color[3]
	}


	gl.BindVertexArray(bar.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, bar.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(bar.vertices)*8, gl.Ptr(bar.vertices), gl.STATIC_DRAW)
	vertAttrib := uint32(gl.GetAttribLocation(bar.program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.DOUBLE, false, 0, gl.PtrOffset(0))

	gl.BindBuffer(gl.ARRAY_BUFFER, bar.colorVbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(bar.colors)*8, gl.Ptr(bar.colors), gl.STATIC_DRAW)
	vertAttrib = uint32(gl.GetAttribLocation(bar.program, gl.Str("vertexColor\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 4, gl.DOUBLE, false, 0, gl.PtrOffset(0))

	bar.modified = false
}

func (bar *Bar) Render() {
	gl.BindVertexArray(bar.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(bar.vertices)))
}

type Chart struct {
	bars []*Bar
	lastPos float64
	program uint32
}

func NewChart(program uint32) *Chart {
	c := &Chart{}
	c.program = program
	c.bars = make([]*Bar, 0)

	return c
}

func (c *Chart) AddSeries(name string, values []float64, r, g, b, a float64) {
	for len(c.bars) < len(values) {
		c.bars = append(c.bars, NewBar(c.program, 1.0, c.lastPos))
		c.lastPos += 1.5
	}

	for i, v := range values {
		c.bars[i].AddValue(v, r, g, b, a)
	}
}

func (c *Chart) Update() {
	for _, b := range c.bars {
		b.Update()
	}
}

func (c *Chart) Render() {
	for _, b := range c.bars {
		b.Render()
	}
}

func main() {
	runtime.LockOSThread()
	if err := glfw.Init(); err != nil {
		log.Fatalln("Failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(WindowWidth, WindowHeight, "Cube", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		log.Fatal("Failed to initialize opengl")
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	program, err := newProgram(vertexShader, fragmentShader)
	if err != nil {
		log.Fatal(err)
	}
	gl.UseProgram(program)

	projection := mgl32.Ortho(-10, 10, -1, 40, -100, 100)
	projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	camera := mgl32.Ident4()
	cameraUniform := gl.GetUniformLocation(program, gl.Str("camera\x00"))
	gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])

	model := mgl32.Ident4()
	modelUniform := gl.GetUniformLocation(program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

	gl.ClearColor(0.1, 0.1, 0.1, 0.0)
	chart := NewChart(program)
	chart.AddSeries("derp", []float64{1.5, 10, 32, 4}, 0.2, 0.8, 0.2, 1.0)
	chart.AddSeries("herp", []float64{3.7, 1, 3.2, 14}, 0.2, 0.5, 0.4, 1.0)
	chart.AddSeries("bucket", []float64{3.7, 13, 13.2, 9}, 0.4, 0.3, 0.4, 1.0)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		camera = mgl32.Translate3D(0, 0, -5)
		gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])

		model = mgl32.Ident4()
		gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

		chart.Update()
		chart.Render()

		window.SwapBuffers()
		glfw.PollEvents()
	}

}
func newProgram(vertexShaderSource, fragmentShaderSource string) (uint32, error) {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, errors.New(fmt.Sprintf("failed to link program: %v", log))
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csource := gl.Str(source)
	gl.ShaderSource(shader, 1, &csource, nil)
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

var vertexShader string = `
#version 330
uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;
in vec3 vert;
layout(location = 1) in vec4 vertexColor;
out vec4 fragmentColor;
void main() {
    gl_Position = projection * camera * model * vec4(vert, 1);
    fragmentColor = vertexColor;
}
` + "\x00"

var fragmentShader = `
#version 330
in vec4 fragmentColor;
out vec4 outputColor;
void main() {
    outputColor = fragmentColor;
}
` + "\x00"