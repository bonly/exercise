package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image/png"
	"math"
	"runtime"
	"time"

	"github.com/go-gl/glfw/v3.1/glfw"
	"golang.org/x/mobile/f32"
	"golang.org/x/mobile/geom"
	"golang.org/x/mobile/gl"
	"golang.org/x/mobile/gl/glutil"
)

const (
	deg2rad = math.Pi / 180.0
	width   = 800
	height  = 600
)

var (
	start = time.Now()
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

	eltbuf  gl.Buffer
	elt     gl.Attrib
	eltdata []uint16

	img *glutil.Image

	texbuf   gl.Texture
	tex      gl.Uniform
	texcoord gl.Attrib

	fade  gl.Uniform
	trans gl.Uniform
}

func (ctx *context) Delete() {
	gl.DeleteProgram(ctx.prog)
	gl.DeleteBuffer(ctx.posbuf)
	gl.DeleteBuffer(ctx.colbuf)
	gl.DeleteBuffer(ctx.eltbuf)
	gl.DeleteTexture(ctx.img.Texture)
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

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	w, err := glfw.CreateWindow(width, height, "my first triangle", nil, nil)
	if err != nil {
		panic(err)
	}
	defer w.Destroy()

	w.MakeContextCurrent()
	w.SetSizeCallback(onResize)
	w.SetKeyCallback(onKey)

	glfw.SwapInterval(1)

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.Enable(gl.DEPTH_TEST)

	ctx := context{
		w: w,
		posdata: []float32{
			// front
			-1, -1, +1, 1,
			+1, -1, +1, 1,
			+1, +1, +1, 1,
			-1, +1, +1, 1,

			// back
			-1, -1, -1, 1,
			+1, -1, -1, 1,
			+1, +1, -1, 1,
			-1, +1, -1, 1,
		},
		coldata: []float32{
			// front colors
			1.0, 0.0, 0.0,
			0.0, 1.0, 0.0,
			0.0, 0.0, 1.0,
			1.0, 1.0, 1.0,

			// back colors
			1.0, 0.0, 0.0,
			0.0, 1.0, 0.0,
			0.0, 0.0, 1.0,
			1.0, 1.0, 1.0,
		},

		eltdata: []uint16{
			// front
			0, 1, 2,
			2, 3, 0,
			// top
			3, 2, 6,
			6, 7, 3,
			// back
			7, 6, 5,
			5, 4, 7,
			// bottom
			4, 5, 1,
			1, 0, 4,
			// left
			4, 0, 3,
			3, 7, 4,
			// right
			1, 5, 6,
			6, 2, 1,
		},
	}
	ctx.prog, err = glutil.CreateProgram(
		vtxShader,
		fragShader,
	)
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

	ctx.eltbuf = gl.CreateBuffer()
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ctx.eltbuf)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER,
		u16Bytes(binary.LittleEndian, ctx.eltdata...),
		gl.STATIC_DRAW,
	)

	img, err := png.Decode(bytes.NewReader(MustAsset("res_texture.png")))
	if err != nil {
		panic(err)
	}
	ctx.img = NewImage(img)
	/*
		ctx.texbuf = gl.CreateTexture()
		gl.BindTexture(gl.TEXTURE_2D, ctx.texbuf)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
		gl.TexImage2D(gl.TEXTURE_2D, 0, 256, 256, gl.RGB, gl.UNSIGNED_BYTE, MustAsset("res_texture.png"))
	*/
	ctx.tex = gl.GetUniformLocation(ctx.prog, "mytexture")
	ctx.texcoord = gl.GetAttribLocation(ctx.prog, "texcoord")

	ctx.col = gl.GetAttribLocation(ctx.prog, "v_color")
	//ctx.fade = gl.GetUniformLocation(ctx.prog, "fade")
	ctx.trans = gl.GetUniformLocation(ctx.prog, "m_transform")

	for !ctx.w.ShouldClose() {
		display(ctx)
		glfw.PollEvents()

	}

}

func display(ctx context) {
	// clear the background as black
	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	ctx.img.Upload()
	ctx.img.Draw(
		geom.Point{0, geom.Pt(height)},
		geom.Point{geom.Pt(width), geom.Pt(height)},
		geom.Point{0, 0},
		ctx.img.RGBA.Bounds(),
	)

	gl.UseProgram(ctx.prog)

	var id f32.Mat4
	id.Identity()

	//  glm::mat4 model = glm::translate(glm::mat4(1.0f), glm::vec3(0.0, 0.0, -4.0));
	var model f32.Mat4
	model.Translate(&id, 0, 0, -4)

	// glm::mat4 view = glm::lookAt(glm::vec3(0.0, 2.0, 0.0), glm::vec3(0.0, 0.0, -4.0), glm::vec3(0.0, 1.0, 0.0));
	var view f32.Mat4
	view.LookAt(&f32.Vec3{0, 0, 2}, &f32.Vec3{0, 1, -4}, &f32.Vec3{4, 1, 2})

	// glm::mat4 projection = glm::perspective(45.0f, 1.0f*screen_width/screen_height, 0.1f, 10.0f);
	var proj f32.Mat4
	proj.Identity()
	proj.Perspective(f32.Radian(60*deg2rad), float32(width)/float32(height), 0.1, 10)

	// float angle = glutGet(GLUT_ELAPSED_TIME) / 1000.0 * 45;  // 45° per second
	// glm::vec3 axis_y(0.0, 1.0, 0.0);
	// glm::mat4 anim = glm::rotate(glm::mat4(1.0f), angle, axis_y);
	angle := float32(time.Since(start).Seconds()) * 45 * deg2rad
	var anim f32.Mat4
	anim.Rotate(&id, f32.Radian(angle), &f32.Vec3{0, 1, 0})

	// glm::mat4 mvp = projection * view * model * anim;
	var trans f32.Mat4
	trans.Identity()
	//trans.Mul(&anim, &trans)
	trans.Mul(&model, &anim)
	trans.Mul(&view, &trans)
	trans.Mul(&proj, &trans)

	gl.UniformMatrix4fv(ctx.trans, flatten(&trans))

	/*
		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, ctx.texbuf)
		gl.Uniform1i(ctx.tex, gl.TEXTURE)
	*/

	gl.EnableVertexAttribArray(ctx.col)
	gl.EnableVertexAttribArray(ctx.pos)
	gl.EnableVertexAttribArray(ctx.elt)

	gl.BindBuffer(gl.ARRAY_BUFFER, ctx.colbuf)
	gl.VertexAttribPointer(ctx.col, 3, gl.FLOAT, false, 0, 0)

	gl.BindBuffer(gl.ARRAY_BUFFER, ctx.posbuf)
	gl.VertexAttribPointer(ctx.pos, 4, gl.FLOAT, false, 0, 0)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ctx.eltbuf)
	sz := gl.GetBufferParameteri(gl.ELEMENT_ARRAY_BUFFER, gl.BUFFER_SIZE)
	gl.DrawElements(gl.TRIANGLES, sz, gl.UNSIGNED_SHORT, 0)

	gl.DisableVertexAttribArray(ctx.col)
	gl.DisableVertexAttribArray(ctx.pos)
	gl.DisableVertexAttribArray(ctx.elt)
	ctx.img.Draw(
		geom.Point{0, geom.Pt(height)},
		geom.Point{geom.Pt(width), geom.Pt(height)},
		geom.Point{0, 0},
		ctx.img.RGBA.Bounds(),
	)

	// display result
	ctx.w.SwapBuffers()
}

func flatten(m *f32.Mat4) []float32 {
	o := make([]float32, 0, 16)
	o = append(o, m[0][0], m[1][0], m[2][0], m[3][0])
	o = append(o, m[0][1], m[1][1], m[2][1], m[3][1])
	o = append(o, m[0][2], m[1][2], m[2][2], m[3][2])
	o = append(o, m[0][3], m[1][3], m[2][3], m[3][3])
	return o
}

// u16Bytes returns the byte representation of uint16 values in the given
// byte order. byteOrder must be either binary.BigEndian or binary.LittleEndian.
func u16Bytes(byteOrder binary.ByteOrder, values ...uint16) []byte {
	le := false
	switch byteOrder {
	case binary.BigEndian:
	case binary.LittleEndian:
		le = true
	default:
		panic(fmt.Sprintf("invalid byte order %v", byteOrder))
	}

	b := make([]byte, 2*len(values))
	if le {
		for i, v := range values {
			b[2*i+0] = byte(v >> 0)
			b[2*i+1] = byte(v >> 8)
		}
	} else {
		for i, v := range values {
			b[2*i+0] = byte(v >> 8)
			b[2*i+1] = byte(v >> 0)
		}
	}
	return b
}

func init() {
	runtime.LockOSThread()
}

const (
	vtxShader = `// vertices shader
attribute vec4 coord;
attribute vec2 texcoord;
varying vec2 f_texcoord;
uniform mat4 m_transform;

void main(void) {
  gl_Position = m_transform * coord;
  f_texcoord = texcoord;
}
`

	fragShader = `// fragment shader
varying vec2 f_texcoord;
uniform sampler2D mytexture;

void main(void) {
  vec2 flipped_texcoord = vec2(f_texcoord.x, 1.0 - f_texcoord.y);
  gl_FragColor = texture2D(mytexture, flipped_texcoord);
}
`
)
