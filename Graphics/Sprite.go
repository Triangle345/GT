// Sprite
package Graphics

import (
	// "github.com/go-gl/gl/v3.2-core/gl"
	"GT/Graphics/Opengl"
)

func NewBasicSprite(img RectangularArea) Sprite {

	sprite := Sprite{img: img, xS: 1, yS: 1, a: 1}
	return sprite
}

type Sprite struct {
	// img
	img RectangularArea
	//TODO: maybe add real image

	// location
	x, y float32

	// rotation
	rot float32

	// color
	r, g, b, a float32

	//scale
	xS, yS float32
}

func (s Sprite) getGLVertexInfo() Opengl.OpenGLVertexInfo {
	// vertexInfo := Opengl.OpenGLVertexInfo{}
	w, h := s.GetImageSection().GetDimensions()

	vertexInfo := Opengl.OpenGLVertexInfo{
		Translations: []float32{float32(s.x), float32(s.y), 0, float32(s.x), float32(s.y), 0, float32(s.x), float32(s.y), 0, float32(s.x), float32(s.y), 0},
		Rotations:    []float32{0, 0, 1, s.rot, 0, 0, 1, s.rot, 0, 0, 1, s.rot, 0, 0, 1, s.rot},
		Scales:       []float32{s.xS, s.yS, 0, s.xS, s.yS, 0, s.xS, s.yS, 0, s.xS, s.yS, 0},
		Colors:       []float32{s.r, s.g, s.b, s.a, s.r, s.g, s.b, s.a, s.r, s.g, s.b, s.a, s.r, s.g, s.b, s.a},
	}

	// for i := 0; i < 4; i = i + 1 {
	// 	vertexInfo.Translations = append(vertexInfo.Translations, float32(s.x), float32(s.y), 0)
	// 	vertexInfo.Rotations = append(vertexInfo.Rotations, 0, 0, 1, s.rot)
	// 	vertexInfo.Scales = append(vertexInfo.Scales, s.xS, s.yS, 0)
	// 	vertexInfo.Colors = append(vertexInfo.Colors, s.r, s.g, s.b, s.a)

	// }

	vertexInfo.Vertices = []float32{-0.5 * w, 0.5 * h, 1.0, 0.5 * w, 0.5 * h, 1.0, 0.5 * w, -0.5 * h, 1.0, -0.5 * w, -0.5 * h, 1.0}

	vertexInfo.Elements = []uint32{uint32(0), uint32(1), uint32(2), uint32(0), uint32(2), uint32(3)}

	return vertexInfo
}

func (s *Sprite) SetColor(r, g, b, a float32) {
	s.r = r
	s.g = g
	s.b = b
	s.a = a
}

func (s *Sprite) GetImageSection() RectangularArea {
	return s.img
}

func (s *Sprite) Move(x, y float32) {
	s.x += x
	s.y += y
}

func (s *Sprite) GetLocation() (x, y float32) {
	x = s.x
	y = s.y

	return
}

func (s *Sprite) Rotate(r float32) {
	s.rot += r
}

func (s *Sprite) Scale(x, y float32) {
	s.xS = x
	s.yS = y
}

func (s *Sprite) SetLocation(x, y float32) {
	s.x = x
	s.y = y
}
