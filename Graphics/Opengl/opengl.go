package Opengl

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/go-gl/gl/v3.3-core/gl"

	mathgl "github.com/go-gl/mathgl/mgl32"

	"image"
	"strings"
)

const (
	NUM_ATTRIBUTES int32 = 11
)

var program uint32
var viewM mathgl.Mat4
var projectionM mathgl.Mat4
var MVPid int32

var textures []uint32

// texture units probed from video card
var texUnits int
var vertexDataTest OpenGLVertexInfo = OpenGLVertexInfo{Stride: 4, Elements: make([]uint32, 0, 9999999), vertexData: make([]float32, 0, 9999999)}

var aggregateImages []image.Image

var vao, vbo uint32

func AddAggregateImage(img image.Image) {
	aggregateImages = append(aggregateImages, img)
}

func Start() {
	// call init right after creating context
	if err := gl.Init(); err != nil {
		fmt.Println("Cannot initialize OGL: " + err.Error())
	}

	texUnits = int(Probe().MaxTextureImageUnits)
	textures = make([]uint32, texUnits)
}

func SetViewPort(width, height int32) {
	// set viewport for window
	gl.Viewport(0, 0, int32(width), int32(height))
}

func bindAggregateImage(img image.Image, idx int) uint32 {

	newIdx := idx

	if rgba, ok := img.(*image.RGBA); ok {

		gl.GenTextures(1, &textures[newIdx])

		gl.ActiveTexture(gl.TEXTURE0 + uint32(newIdx))
		gl.BindTexture(gl.TEXTURE_2D, textures[newIdx])
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_R, gl.CLAMP_TO_EDGE)

		gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(rgba.Rect.Size().X), int32(rgba.Rect.Size().Y), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))
		gl.GenerateMipmap(gl.TEXTURE_2D)

		idxS := strconv.Itoa(newIdx)
		fmt.Println("idx: ", int32(newIdx), "myTextureSampler["+idxS+"]\x00")
		gl.Uniform1i(gl.GetUniformLocation(program, gl.Str("myTextureSampler["+idxS+"]\x00")), int32(newIdx))
		if ok := gl.GetError(); ok != gl.NO_ERROR {

			fmt.Println("1- Cannot load Image in location: ./: ", ok)
			os.Exit(-1)
		}

		return textures[newIdx]
	} else {
		fmt.Println("Image not RGBA at location: ./")
		os.Exit(-1)
	}

	return 0
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csource, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csource, nil)
	free()
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
	if ok := gl.GetError(); ok != gl.NO_ERROR {

		return 0, errors.New("Error in Make Program : " + string(ok))

	}
	return program, nil
}

func SetOrthographic(width, height int) {
	// Projection := mathgl.Perspective(45.0, winWidth/winHeight, 0.1, 100.0)
	// Projection := mathgl.Ortho2D(0.0, winWidth, winHeight, 0.0)
	// aspect := float32(s.width / s.height)
	Projection := mathgl.Ortho(0.0, float32(width), float32(height), 0.0, -5.0, 5.0)
	viewM = mathgl.LookAt(0.0, 0.0, 5.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0)
	projectionM = Projection
	gl.Disable(gl.DEPTH_TEST)
}

func SetPerspective(width, height int) {
	Projection := mathgl.Perspective(mathgl.DegToRad(45.0), float32(width/height), 0.1, 100.0)
	viewM = mathgl.LookAt(0.0, 0.0, 20.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0)
	projectionM = Projection

	gl.Disable(gl.CULL_FACE)
	gl.Enable(gl.DEPTH_TEST)

}

// Clear instructs opengl to clear the background to a certain color
func Clear(r, g, b, a float32) {
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func CreateBuffers() {

	var err error
	program, err = MakeProgram(VertexShader(), FragmentShader())
	// defer program.Delete()

	if err != nil {
		fmt.Println("Error loading shaders: " + err.Error())
		panic("error loading shaders")
	}

	gl.BindFragDataLocation(program, 0, gl.Str("color\x00"))

	MVPid = gl.GetUniformLocation(program, gl.Str("MVP\x00"))

	// View := mathgl.LookAt(0.0, 0.0, 5.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0)

	// viewM = View

	gl.DepthFunc(gl.LEQUAL)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	// var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	// var vbo uint32
	gl.GenBuffers(1, &vbo)

	// gl.GenBuffers(1, &elementvbo)
	// fmt.Println(program)
	gl.UseProgram(program)

	for idx, img := range aggregateImages {
		bindAggregateImage(img, idx)

	}

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

func BindBuffers() { //vertexData *OpenGLVertexInfo) {

	vertexData := &vertexDataTest

	// check to see if there are any vertices at all to bind
	if len(vertexData.vertexData) == 0 {
		return
	}

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertexData.vertexData)*4, gl.Ptr(vertexData.vertexData), gl.DYNAMIC_DRAW)

	positionAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertexPosition_modelspace\x00")))
	gl.EnableVertexAttribArray(positionAttrib)
	gl.VertexAttribPointer(positionAttrib, 3, gl.FLOAT, false, 4*NUM_ATTRIBUTES, gl.PtrOffset(0))

	colorAttrib := uint32(gl.GetAttribLocation(program, gl.Str("diffuseColor\x00")))
	gl.EnableVertexAttribArray(colorAttrib)
	gl.VertexAttribPointer(colorAttrib, 4, gl.FLOAT, false, 4*NUM_ATTRIBUTES, gl.PtrOffset(3*4))

	uvAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertexUV\x00")))
	gl.EnableVertexAttribArray(uvAttrib)
	gl.VertexAttribPointer(uvAttrib, 2, gl.FLOAT, false, 4*NUM_ATTRIBUTES, gl.PtrOffset(7*4))

	shaderMode := uint32(gl.GetAttribLocation(program, gl.Str("mode\x00")))
	gl.EnableVertexAttribArray(shaderMode)
	gl.VertexAttribPointer(shaderMode, 1, gl.FLOAT, false, 4*NUM_ATTRIBUTES, gl.PtrOffset(9*4))

	samplerIdx := uint32(gl.GetAttribLocation(program, gl.Str("samplerIdx\x00")))
	gl.EnableVertexAttribArray(samplerIdx)
	gl.VertexAttribPointer(samplerIdx, 1, gl.FLOAT, false, 4*NUM_ATTRIBUTES, gl.PtrOffset(10*4))

	// gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementvbo)
	// gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(vertexData.Elements)*4, gl.Ptr(vertexData.Elements), gl.STATIC_DRAW)

	// gl.ActiveTexture(gl.TEXTURE0)
	// gl.BindTexture(gl.TEXTURE_2D, 1)
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
	if len(vertexData.vertexData) == 0 {
		return
	}

	//	vertexData.Print()

	// vertexData := vertexDataTest

	//vertexData.Print()
	// BindBuffers(vertexData)
	// vertexData.Print()
	//gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(vertexData.VertexData)*4, gl.Ptr(vertexData.VertexData))

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertexData.vertexData)*4, gl.Ptr(vertexData.vertexData), gl.DYNAMIC_DRAW)
	// gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, elementvbo)
	// gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(vertexData.Elements)*4, gl.Ptr(vertexData.Elements), gl.DYNAMIC_DRAW)

	MVP := projectionM.Mul4(viewM) //.Mul4(Model)

	gl.UniformMatrix4fv(MVPid, 1, false, &MVP[0])

	// vertexData.Print()
	// gl.DrawElements(gl.TRIANGLES, int32(len(vertexData.Elements)), gl.UNSIGNED_INT, nil)
	// }

	numTriVerts := int32((len(vertexData.vertexData) / (int(NUM_ATTRIBUTES) * 2)) * 3)

	gl.DrawArrays(gl.TRIANGLES, 0, numTriVerts)

	ClearVertexData()

}

func Cleanup() {
	gl.DeleteBuffers(1, &vbo)

}
