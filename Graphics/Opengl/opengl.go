package Opengl

import (
	"errors"
	"fmt"
	"github.com/go-gl/gl/v3.2-core/gl"
	mathgl "github.com/go-gl/mathgl/mgl32"

	"strings"
)

// var elements = []uint32{
// 	0, 1, 2,
// 	2, 3, 0,
// }
// var quad_colours = []float32{
// 	0.583, 0.771, 0.014,
// 	0.609, 0.115, 0.436,
// 	0.327, 0.483, 0.844,
// }

// var translations = []float32{}
// var rotations = []float32{}
// var scales = []float32{}
// var vertices = []float32{}
// var uvs = []float32{}
// var colors = []float32{}

var program uint32
var viewM mathgl.Mat4
var projectionM mathgl.Mat4
var MVPid int32

var vao, vbo, colorvbo, uvvbo, tvbo, rvbo, svbo, elementvbo uint32

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

func MakeProgram(vert, frag string) (uint32, error) {
	vertexShader, err := compileShader(vert, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	fragmentShader, err := compileShader(frag, gl.FRAGMENT_SHADER)
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

func CreateBuffers(width, height int) {

	var err error
	program, err = MakeProgram(vertexShaderSource, fragmentShaderSource)
	// defer program.Delete()

	if err != nil {
		fmt.Println("Error loading shaders: " + err.Error())
		panic("error loading shaders")
	}

	gl.BindFragDataLocation(program, 0, gl.Str("color\x00"))
	gl.GetUniformLocation(program, gl.Str("myTextureSampler\x00"))

	matrixID := gl.GetUniformLocation(program, gl.Str("MVP\x00"))
	MVPid = matrixID

	// Projection := mathgl.Perspective(45.0, winWidth/winHeight, 0.1, 100.0)
	// Projection := mathgl.Ortho2D(0.0, winWidth, winHeight, 0.0)
	// aspect := float32(s.width / s.height)
	Projection := mathgl.Ortho(0.0, float32(width), float32(height), 0.0, -5.0, 5.0)

	projectionM = Projection

	View := mathgl.LookAt(0.0, 0.0, 5.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0)

	viewM = View

	gl.Disable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LEQUAL)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	// var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	// var vbo uint32
	gl.GenBuffers(1, &vbo)

	// var colorvbo uint32
	gl.GenBuffers(1, &colorvbo)

	// var uvvbo uint32
	gl.GenBuffers(1, &uvvbo)

	// var tvbo uint32
	gl.GenBuffers(1, &tvbo)

	// var rvbo uint32
	gl.GenBuffers(1, &rvbo)

	// var svbo uint32
	gl.GenBuffers(1, &svbo)

	// element buffer
	// var elementvbo uint32
	gl.GenBuffers(1, &elementvbo)
	// defer elementBuffer.Delete()
	// gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementvbo)
	// gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(elements)*4, gl.Ptr(elements), gl.STATIC_DRAW)

}

type OpenGLVertexInfo struct {
	Translations []float32
	Rotations    []float32
	Scales       []float32
	Vertices     []float32
	Uvs          []float32
	Colors       []float32
	Elements     []uint32
}

// func (i OpenGLVertexInfo) GetMaxElement() uint32 {
// 	max := uint32(0)

// 	for _, v := range i.Elements {
// 		if v > max {
// 			max = v
// 		}
// 	}

// 	return max
// }

func (i *OpenGLVertexInfo) AdjustElements(adj uint32) {

	for k, _ := range i.Elements {
		i.Elements[k] += adj
		// fmt.Println(i.Elements[k])
	}

}

func (i *OpenGLVertexInfo) Append(o OpenGLVertexInfo) {

	i.Translations = append(i.Translations, o.Translations...)
	i.Rotations = append(i.Rotations, o.Rotations...)
	i.Scales = append(i.Scales, o.Scales...)
	i.Vertices = append(i.Vertices, o.Vertices...)
	i.Uvs = append(i.Uvs, o.Uvs...)
	i.Colors = append(i.Colors, o.Colors...)
	i.Elements = append(i.Elements, o.Elements...)

}

func (i OpenGLVertexInfo) Print() {
	fmt.Printf("--------------------------------------------------\n")
	for _, v := range i.Elements {
		fmt.Printf("inside elements:%d \n", v)
	}

}

func BindBuffers(data *OpenGLVertexInfo) {

	// fmt.Println(program)
	gl.UseProgram(program)

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(data.Vertices)*4, gl.Ptr(data.Vertices), gl.DYNAMIC_DRAW)
	positionAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertexPosition_modelspace\x00")))
	gl.EnableVertexAttribArray(positionAttrib)
	gl.VertexAttribPointer(positionAttrib, 3, gl.FLOAT, false, 0, nil)
	// defer positionAttrib.DisableArray()

	gl.BindBuffer(gl.ARRAY_BUFFER, colorvbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(data.Colors)*4, gl.Ptr(data.Colors), gl.DYNAMIC_DRAW)
	colorAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertexColor\x00")))
	gl.EnableVertexAttribArray(colorAttrib)
	gl.VertexAttribPointer(colorAttrib, 4, gl.FLOAT, false, 0, nil)

	// defer colorAttrib.DisableArray()

	gl.BindBuffer(gl.ARRAY_BUFFER, uvvbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(data.Uvs)*4, gl.Ptr(data.Uvs), gl.DYNAMIC_DRAW)
	uvAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertexUV\x00")))
	gl.EnableVertexAttribArray(uvAttrib)
	gl.VertexAttribPointer(uvAttrib, 2, gl.FLOAT, false, 0, nil)

	//------------------------------------------------------------------
	gl.BindBuffer(gl.ARRAY_BUFFER, tvbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(data.Translations)*4, gl.Ptr(data.Translations), gl.DYNAMIC_DRAW)
	tAttrib := uint32(gl.GetAttribLocation(program, gl.Str("translation\x00")))
	gl.EnableVertexAttribArray(tAttrib)
	gl.VertexAttribPointer(tAttrib, 3, gl.FLOAT, false, 0, nil)

	gl.BindBuffer(gl.ARRAY_BUFFER, rvbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(data.Rotations)*4, gl.Ptr(data.Rotations), gl.DYNAMIC_DRAW)
	rAttrib := uint32(gl.GetAttribLocation(program, gl.Str("rotation\x00")))
	gl.EnableVertexAttribArray(rAttrib)
	gl.VertexAttribPointer(rAttrib, 4, gl.FLOAT, false, 0, nil)

	gl.BindBuffer(gl.ARRAY_BUFFER, svbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(data.Scales)*4, gl.Ptr(data.Scales), gl.DYNAMIC_DRAW)
	sAttrib := uint32(gl.GetAttribLocation(program, gl.Str("scale\x00")))
	gl.EnableVertexAttribArray(sAttrib)
	gl.VertexAttribPointer(sAttrib, 3, gl.FLOAT, false, 0, nil)
	//------------------------------------------------------------------

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementvbo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(data.Elements)*4, gl.Ptr(data.Elements), gl.STATIC_DRAW)
}

func Draw(data *OpenGLVertexInfo) {
	MVP := projectionM.Mul4(viewM) //.Mul4(Model)

	gl.UniformMatrix4fv(MVPid, 1, false, &MVP[0])

	gl.DrawElements(gl.TRIANGLES, int32(len(data.Elements)), gl.UNSIGNED_INT, nil)
}

func Cleanup() {
	gl.DeleteBuffers(1, &vbo)
	gl.DeleteBuffers(1, &colorvbo)
	gl.DeleteBuffers(1, &uvvbo)
}
