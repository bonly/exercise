package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image/png"
	"runtime"

	"github.com/go-gl/glfw/v3.1/glfw"
	"golang.org/x/mobile/f32"
	"golang.org/x/mobile/geom"
	"golang.org/x/mobile/gl"
	"golang.org/x/mobile/gl/glutil"
)

type context struct {
	w    *glfw.Window
	prog gl.Program

	coord   gl.Attrib
	xoff    gl.Uniform
	xscale  gl.Uniform
	sprite  gl.Uniform
	texture gl.Uniform

	tex gl.Texture
	vbo gl.Buffer
	img *glutil.Image
}

func (ctx *context) Delete() {
	gl.DeleteProgram(ctx.prog)
	gl.DeleteBuffer(ctx.vbo)
	gl.DeleteTexture(ctx.img.Texture)
	ctx.w.Destroy()
}

func init() {
	runtime.LockOSThread()
}

func main() {
	w, err := New(width, height, "graph")
	if err != nil {
		panic(err)
	}
	defer Delete()

	w.SetKeyCallback(onKey)

	prog, err := glutil.CreateProgram(vtxShader, fragShader)
	if err != nil {
		panic(err)
	}

	img, err := png.Decode(bytes.NewReader(MustAsset("res_texture.png")))
	if err != nil {
		panic(err)
	}

	ctx := context{
		w:       w,
		prog:    prog,
		coord:   gl.GetAttribLocation(prog, "coord2d"),
		xoff:    gl.GetUniformLocation(prog, "offset_x"),
		xscale:  gl.GetUniformLocation(prog, "scale_x"),
		sprite:  gl.GetUniformLocation(prog, "sprite"),
		texture: gl.GetUniformLocation(prog, "mytexture"),

		tex: gl.CreateTexture(),
		vbo: gl.CreateBuffer(),
		img: NewImage(img),
	}
	defer ctx.Delete()

	//	gl.ActiveTexture(gl.TEXTURE0) gl.BindTexture(gl.TEXTURE_2D, ctx.tex)
	//	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	//	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	//	gl.TexImage2D( gl.TEXTURE_2D, 0, 15, 15, gl.RGBA, gl.UNSIGNED_BYTE,
	//	_res_texture_png,)

	gl.BindBuffer(gl.ARRAY_BUFFER, ctx.vbo)
	points := make([]float32, npoints*2)
	for i := 0; i < npoints; i += 2 {
		//		x := float32((i - npoints/2.0) / (npoints / 1))
		x := float32((i - npoints/2) / (npoints / 200))
		y := f32.Sin(x*10) / (1 + x*x)

		points[i] = x
		points[i+1] = y
	}
	gl.BufferData(
		gl.ARRAY_BUFFER,
		f32.Bytes(binary.LittleEndian, points...),
		gl.STATIC_DRAW,
	)

	for !ctx.w.ShouldClose() {
		display(ctx)
		glfw.PollEvents()
	}
}

func display(ctx context) {
	// clear the background as black
	gl.ClearColor(1, 1, 1, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.UseProgram(ctx.prog)

	gl.Uniform1i(ctx.texture, 0)
	gl.Uniform1f(ctx.xoff, xoff)
	gl.Uniform1f(ctx.xscale, xscale)

	gl.BindBuffer(gl.ARRAY_BUFFER, ctx.vbo)

	gl.EnableVertexAttribArray(ctx.coord)
	gl.VertexAttribPointer(ctx.coord, 2, gl.FLOAT, false, 0, 0)

	switch mode {
	case 0:
		gl.Uniform1f(ctx.sprite, 0)
		gl.DrawArrays(gl.LINE_STRIP, 0, npoints)

	case 1:
		gl.Uniform1f(ctx.sprite, 1)
		gl.DrawArrays(gl.POINTS, 0, npoints)

	case 2:
		ctx.img.Upload()
		ctx.img.Draw(
			geom.Point{0, geom.Height},
			geom.Point{geom.Width, geom.Height},
			geom.Point{0, 0},
			ctx.img.RGBA.Bounds(),
		)
		// gl.Uniform1f(ctx.sprite, float32(ctx.img.Bounds().Dy()))
		gl.DrawArrays(gl.POINTS, 0, npoints)
	}

	gl.DisableVertexAttribArray(ctx.coord)

	// display result
	ctx.w.SwapBuffers()
}

func onKey(w *glfw.Window, key glfw.Key, scancode int,
	action glfw.Action, mods glfw.ModifierKey) {
	switch {
	case key == glfw.KeyEscape && action == glfw.Press,
		key == glfw.KeyQ && action == glfw.Press:
		w.SetShouldClose(true)
		return
	}

	if action == glfw.Press {
		switch key {
		case glfw.KeyF1:
			fmt.Printf("now drawing using lines...\n")
			mode = 0
			return

		case glfw.KeyF2:
			fmt.Printf("now drawing using points...\n")
			mode = 1
			return

		case glfw.KeyF3:
			fmt.Printf("now drawing using point sprites...\n")
			mode = 2
			return

		case glfw.KeyLeft:
			xoff -= 0.1

		case glfw.KeyRight:
			xoff += 0.1

		case glfw.KeyUp:
			xscale *= 1.5
		case glfw.KeyDown:
			xscale /= 1.5

		case glfw.KeyHome:
			xoff = 0
			xscale = 1
		}
	}
}

var (
	xoff   = float32(0)
	xscale = float32(1)
	mode   = 0
)

const (
	width   = 800
	height  = 600
	npoints = 20000

	vtxShader = `
#version 120

attribute vec2 coord2d;
varying vec4 f_color;
uniform float offset_x;
uniform float scale_x;
uniform float sprite;

void main(void) {
	gl_Position = vec4((coord2d.x + offset_x) * scale_x, coord2d.y, 0, 1);
	f_color = vec4(coord2d.xy / 2.0 + 0.5, 1, 1);
	gl_PointSize = max(1.0, sprite);
}
`

	fragShader = `
#version 120

uniform sampler2D mytexture;
varying vec4 f_color;
uniform float sprite;

void main(void) {
	if (sprite > 1.0)
		gl_FragColor = texture2D(mytexture, gl_PointCoord) * f_color;
	else
		gl_FragColor = f_color;
}
`
)
