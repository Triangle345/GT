// Sprite
package Components

import (
	// "github.com/go-gl/gl/v3.2-core/gl"
	"GT/Graphics"
	"GT/Graphics/Opengl"

	mathgl "github.com/go-gl/mathgl/mgl32"
)

type Renderer interface {
	getGLVertexInfo() Opengl.OpenGLVertexInfo
}

func NewSpriteRenderer(img Graphics.RectangularArea) SpriteRenderer {

	sprite := SpriteRenderer{img: img, a: 1}
	return sprite
}

type SpriteRenderer struct {
	Node
	// img
	img Graphics.RectangularArea
	//TODO: maybe add real image

	// color
	r, g, b, a float32
}

func (s SpriteRenderer) getGLVertexInfo() Opengl.OpenGLVertexInfo {
	// vertexInfo := Opengl.OpenGLVertexInfo{}
	w, h := s.GetImageSection().GetDimensions()

	vertex_data := []float32{-0.5 * w, 0.5 * h, 1.0, 0.5 * w, 0.5 * h, 1.0, 0.5 * w, -0.5 * h, 1.0, -0.5 * w, -0.5 * h, 1.0}

	elements := []uint32{uint32(0), uint32(1), uint32(2), uint32(0), uint32(2), uint32(3)}
	uvs := s.img.GetUVs()

	Model := s.Transform.GetModel()

	var data []float32
	for j := 0; j < 4; j++ {
		transformation := mathgl.Vec4{vertex_data[j*3+0], vertex_data[j*3+1], vertex_data[j*3+2], 1}
		t := Model.Mul4x1(transformation)
		data = append(data, t[0], t[1], t[2], s.r, s.g, s.b, s.a, uvs[j*2+0], uvs[j*2+1])

	}

	vertexInfo := Opengl.OpenGLVertexInfo{
		VertexData: data,
		Elements:   elements,
		Stride:     4,
	}

	return vertexInfo
}

func (s *SpriteRenderer) SetColor(r, g, b, a float32) {
	s.r = r
	s.g = g
	s.b = b
	s.a = a
}

func (s *SpriteRenderer) GetImageSection() Graphics.RectangularArea {
	return s.img
}
