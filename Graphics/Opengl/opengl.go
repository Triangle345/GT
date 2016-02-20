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

	gl.GenBuffers(1, &elementvbo)

}

type OpenGLVertexInfo struct {
	VertexData []float32
	Elements   []uint32
	Stride     int
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

func (i *OpenGLVertexInfo) Append(o *OpenGLVertexInfo) {

	i.VertexData = append(i.VertexData, o.VertexData...)
	i.Elements = append(i.Elements, o.Elements...)
	i.Stride += o.Stride

}

func (i OpenGLVertexInfo) Print() {
	fmt.Printf("--------------------------------------------------\n")
	for _, v := range i.VertexData {
		fmt.Printf("inside data:%f \n", v)
	}

	for _, v := range i.Elements {
		fmt.Printf("inside element:%f \n", v)
	}

}

func BindBuffers(data OpenGLVertexInfo) {

	// fmt.Println(program)
	gl.UseProgram(program)

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(data.VertexData)*4, gl.Ptr(data.VertexData), gl.DYNAMIC_DRAW)

	positionAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertexPosition_modelspace\x00")))
	gl.EnableVertexAttribArray(positionAttrib)
	gl.VertexAttribPointer(positionAttrib, 3, gl.FLOAT, false, 4*9, gl.PtrOffset(0))

	colorAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertexColor\x00")))
	gl.EnableVertexAttribArray(colorAttrib)
	gl.VertexAttribPointer(colorAttrib, 4, gl.FLOAT, false, 4*9, gl.PtrOffset(3*4))

	uvAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertexUV\x00")))
	gl.EnableVertexAttribArray(uvAttrib)
	gl.VertexAttribPointer(uvAttrib, 2, gl.FLOAT, false, 4*9, gl.PtrOffset(7*4))

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementvbo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(data.Elements)*4, gl.Ptr(data.Elements), gl.STATIC_DRAW)

}

func RepopulateVBO(data OpenGLVertexInfo) {

	//gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(data.VertexData)*4, gl.Ptr(data.VertexData))
	//data.Print()
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(data.VertexData)*4, gl.Ptr(data.VertexData), gl.STATIC_DRAW)

	// update for element array handled in above function for buffers
	// gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementvbo)
	// gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(data.Elements)*4, gl.Ptr(data.Elements), gl.STATIC_DRAW)

}

func Draw(data OpenGLVertexInfo) {
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(data.VertexData)*4, gl.Ptr(data.VertexData))
	// gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	// gl.BufferData(gl.ARRAY_BUFFER, len(data.VertexData)*4, gl.Ptr(data.VertexData), gl.STATIC_DRAW)
	MVP := projectionM.Mul4(viewM) //.Mul4(Model)

	gl.UniformMatrix4fv(MVPid, 1, false, &MVP[0])

	gl.DrawElements(gl.TRIANGLES, int32(len(data.Elements)), gl.UNSIGNED_INT, nil)
}

func Cleanup() {
	gl.DeleteBuffers(1, &vbo)
	gl.DeleteBuffers(1, &colorvbo)
	gl.DeleteBuffers(1, &uvvbo)
}
