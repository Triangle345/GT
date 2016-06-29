// Package Graphics
package Components

import (
	"GT/Graphics/Opengl"

	mathgl "github.com/go-gl/mathgl/mgl32"
)

// Initialize is necessary for the renderer to be utilized as a Component
func (s *modelRenderer) Initialize() {

}

// NewSpriteRenderer creates a renderer and initializes its animation map
func NewModelRenderer() *modelRenderer {

	model := modelRenderer{}

	return &model
}

// SpriteRenderer is a component which allows a sprite to be drawn or animated
type modelRenderer struct {
	// the parent node
	//Parent *Components.Node
	ChildComponent

	// uvs - store uvs for speed
	uvs []float32
}

// Update gets called every frame and accounts for all settings in the renderer as well as shifts animations
func (s *modelRenderer) Update(delta float32) {

	r := float32(.3)
	g := float32(.3)
	b := float32(.3)
	a := float32(1.000)

	//vertexData := []float32{-0.5, 0.5, 1.0, 0.5, 0.5, 1.0, 0.5, -0.5, 1.0, -0.5, -0.5, 1.0}

	var elements = make([]uint32, 0, 3*12*4)

	elements = append(elements, uint32(1), uint32(3), uint32(0))
	elements = append(elements, uint32(7), uint32(5), uint32(4))
	elements = append(elements, uint32(4), uint32(1), uint32(0))
	elements = append(elements, uint32(5), uint32(2), uint32(1))
	elements = append(elements, uint32(2), uint32(7), uint32(3))
	elements = append(elements, uint32(0), uint32(7), uint32(4))
	elements = append(elements, uint32(1), uint32(2), uint32(3))
	elements = append(elements, uint32(7), uint32(6), uint32(5))
	elements = append(elements, uint32(4), uint32(5), uint32(1))
	elements = append(elements, uint32(5), uint32(6), uint32(2))
	elements = append(elements, uint32(2), uint32(6), uint32(7))
	elements = append(elements, uint32(0), uint32(3), uint32(7))

	// elements = append(elements, uint32(1), uint32(3), uint32(2), uint32(0))

	// elements = append(elements, uint32(3), uint32(7), uint32(6), uint32(2))
	// elements = append(elements, uint32(7), uint32(5), uint32(4), uint32(6))

	// elements = append(elements, uint32(5), uint32(1), uint32(0), uint32(4))
	// elements = append(elements, uint32(0), uint32(2), uint32(6), uint32(4))
	// elements = append(elements, uint32(5), uint32(7), uint32(3), uint32(1))

	Model := mathgl.Ident4()

	Model = s.GetParent().transform.GetUpdatedModel()

	// transform all vertex data and combine it with other data
	var data = make([]float32, 0, 9*4)
	// for j := 0; j < 4; j++ {
	// 	transformation := mathgl.Vec4{vertexData[j*3+0], vertexData[j*3+1], vertexData[j*3+2], 1}
	// 	t := Model.Mul4x1(transformation)

	// 	data = append(data, t[0], t[1], t[2], r, g, b, a, 0, 0)

	// }

	//TODO: try quads out!

	data = append(data, 3.000000, -3.000000, -3.000000, r+.1, g, b, a, -1, -1, 0.0)
	data = append(data, 3.000000, -3.000000, 3.000000, r+.1, g, b, a, -1, -1, 0.0)
	data = append(data, -3.000000, -3.000000, 3.000000, r, g+.1, b, a, -1, -1, 0.0)
	data = append(data, -3.000000, -3.000000, -3.000000, r, g, b, a+.1, -1, -1, 0.0)
	data = append(data, 3.000000, 3.000000, -3.0, r, g, b, a, 0, 0, 0.0)
	data = append(data, 3.0, 3.000000, 3.0, r, g, b, a, 0, 0, 0.0)
	data = append(data, -3.000000, 3.000000, 3.000000, r, g, b, a, 0, 0, 0.0)
	data = append(data, -3.000000, 3.000000, -3.000000, r, g, b-.1, a, 0, 0, 0.0)

	// data = append(data, -5.000000, -5.000000, 5.000000, r+.1, g, b, a, -1, -1, 0.0)
	// data = append(data, -5.000000, 5.000000, 5.000000, r+.1, g, b, a, -1, -1, 0.0)
	// data = append(data, -5.000000, -5.000000, -5.000000, r, g+.1, b, a, -1, -1, 0.0)
	// data = append(data, -5.000000, 5.000000, -5.000000, r, g, b, a+.1, -1, -1, 0.0)
	// data = append(data, 5.000000, -5.000000, 5.0, r, g, b, a, 0, 0, 0.0)
	// data = append(data, 5.0, 5.000000, 5.0, r, g, b, a, 0, 0, 0.0)
	// data = append(data, 5.000000, -5.000000, -5.000000, r, g, b, a, 0, 0, 0.0)
	// data = append(data, 5.000000, 5.000000, -5.000000, r, g, b-.1, a, 0, 0, 0.0)

	for i := 0; i < 8; i++ {
		transformation := mathgl.Vec4{data[i*10+0], data[i*10+1], data[i*10+2], 1}
		t := Model.Mul4x1(transformation)
		data[i*10+0] = t[0]
		data[i*10+1] = t[1]
		data[i*10+2] = t[2]
	}

	// package everything up in an OpenGLVertexInfo
	vertexInfo := Opengl.OpenGLVertexInfo{
		VertexData: data,
		Elements:   elements,
		Stride:     8,
	}

	// send OpenGLVertex info to Opengl module
	Opengl.AddVertexData(1, &vertexInfo)

}
