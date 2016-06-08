package Opengl

import (
	"errors"
	"fmt"

	"github.com/go-gl/gl/v3.2-core/gl"
	mathgl "github.com/go-gl/mathgl/mgl32"

	"strings"
)

var program uint32
var viewM mathgl.Mat4
var projectionM mathgl.Mat4
var MVPid int32

var textureHash map[uint32]*OpenGLVertexInfo = make(map[uint32]*OpenGLVertexInfo)
var layers []uint32 = []uint32{}

var vertexDataTest OpenGLVertexInfo = OpenGLVertexInfo{Stride: 4, Elements: make([]uint32, 0, 9999999), VertexData: make([]float32, 0, 9999999)}

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

func SetOrthographic(width, height int) {
	// Projection := mathgl.Perspective(45.0, winWidth/winHeight, 0.1, 100.0)
	// Projection := mathgl.Ortho2D(0.0, winWidth, winHeight, 0.0)
	// aspect := float32(s.width / s.height)
	Projection := mathgl.Ortho(0.0, float32(width), float32(height), 0.0, -5.0, 5.0)

	projectionM = Projection
}

func CreateBuffers() {

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

func AddVertexData(id uint32, o *OpenGLVertexInfo) {
	// if val, ok := textureHash[id]; ok {
	// 	// fmt.Println("ading to  layer")
	// 	layers = append(layers, id)
	// 	// vertexData.append(o)
	// 	val.append(o)
	// } else {
	// 	// fmt.Print("creating layer ")
	// 	// fmt.Println(id)
	// 	// o.Print()
	// 	layers = append(layers, id)
	// 	textureHash[id] = o
	//
	// }

	vertexDataTest.append(o)

}

func ClearVertexData() {
	// vertexData.Clear()
	// layers = nil
	// layers = []uint32{}
	// textureHash = nil
	// textureHash = make(map[uint32]*OpenGLVertexInfo)
	vertexDataTest.Clear()
}

type OpenGLVertexInfo struct {
	VertexData []float32
	Elements   []uint32

	// TODO: make stride private and set it statically using new OpenGLVertexInfo2D or something
	Stride      int
	totalStride int
}

// TODO: need to put this in as part of vertex info and into hashhash
func (this *OpenGLVertexInfo) adjustElements(o *OpenGLVertexInfo) {

	for k, _ := range o.Elements {
		o.Elements[k] += uint32(this.totalStride)
		// fmt.Println(i.Elements[k])
	}

	//TODO: add argument of stride to add stride of other vertexinfo
	this.totalStride += o.Stride
}

func (i *OpenGLVertexInfo) append(o *OpenGLVertexInfo) {
	i.adjustElements(o)

	i.VertexData = append(i.VertexData, o.VertexData...)
	i.Elements = append(i.Elements, o.Elements...)

	//i.Stride += o.Stride

}

func (i *OpenGLVertexInfo) Clear() {

	i.VertexData = make([]float32, 0, 1000000)
	i.Elements = make([]uint32, 0, 1000000)
	i.totalStride = 0

}

func (i OpenGLVertexInfo) Print() {
	fmt.Printf("--------------------------------------------------\n")
	for _, v := range i.VertexData {
		fmt.Printf("inside data:%f \n", v)
	}
	fmt.Printf("+++++++++++++++++++++++++++++++++++++++++++++++++++++\n")
	for _, v := range i.Elements {
		fmt.Printf("inside element:%f \n", v)
	}
	fmt.Printf("stride %d \n", i.Stride)
	fmt.Printf("Total stride %d \n", i.totalStride)

}

func BindBuffers() { //vertexData *OpenGLVertexInfo) {

	// fmt.Println(program)
	gl.UseProgram(program)

	vertexData := &vertexDataTest

	// check to see if there are any vertices at all to bind
	if len(vertexData.VertexData) == 0 {
		return
	}

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertexData.VertexData)*4, gl.Ptr(vertexData.VertexData), gl.DYNAMIC_DRAW)

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
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(vertexData.Elements)*4, gl.Ptr(vertexData.Elements), gl.STATIC_DRAW)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, 1)
}

func Draw() {

	// if len(layers) == 0 {
	// 	fmt.Println("layers empty!!!")
	// 	return
	// }

	// for _, layer := range layers {

	// vertexData := textureHash[layer]
	vertexData := &vertexDataTest

	// check to see if there are any vertices at all to draw
	if len(vertexData.VertexData) == 0 {
		return
	}

	//	vertexData.Print()

	// vertexData := vertexDataTest

	//vertexData.Print()
	// BindBuffers(vertexData)
	// vertexData.Print()
	//gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(vertexData.VertexData)*4, gl.Ptr(vertexData.VertexData))

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertexData.VertexData)*4, gl.Ptr(vertexData.VertexData), gl.DYNAMIC_DRAW)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementvbo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(vertexData.Elements)*4, gl.Ptr(vertexData.Elements), gl.DYNAMIC_DRAW)

	MVP := projectionM.Mul4(viewM) //.Mul4(Model)

	gl.UniformMatrix4fv(MVPid, 1, false, &MVP[0])
	// vertexData.Print()
	gl.DrawElements(gl.TRIANGLES, int32(len(vertexData.Elements)), gl.UNSIGNED_INT, nil)
	// }
	ClearVertexData()

}

func Cleanup() {
	gl.DeleteBuffers(1, &vbo)
	gl.DeleteBuffers(1, &colorvbo)
	gl.DeleteBuffers(1, &uvvbo)
}
