// Sprite
package Graphics

import (
	// "github.com/go-gl/gl/v3.2-core/gl"
	"GT/Graphics/Opengl"
	"fmt"

	mathgl "github.com/go-gl/mathgl/mgl32"
)

func NewBasicSprite(img Drawable) Sprite {

	sprite := Sprite{img: img, xS: 1, yS: 1, a: 1}
	return sprite
}

type Sprite struct {
	//img
	img Drawable

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

	var rect RectangularArea

	if img, ok := s.img.(*SpriteSheetImage); ok {
		rect = img.rect
	} else {
		fmt.Println("Cannot get image rect: getGLVertexInfo()")
	}

	w, h := rect.GetDimensions()

	vertex_data := []float32{-0.5 * w, 0.5 * h, 1.0, 0.5 * w, 0.5 * h, 1.0, 0.5 * w, -0.5 * h, 1.0, -0.5 * w, -0.5 * h, 1.0}

	elements := []uint32{uint32(0), uint32(1), uint32(2), uint32(0), uint32(2), uint32(3)}

	uvs := rect.uvs

	//uvs := s.img.uvs

	Model := mathgl.Ident4()
	Model = Model.Mul4(mathgl.Translate3D(float32(s.x), float32(s.y), float32(0.0)))
	Model = Model.Mul4(mathgl.HomogRotate3DZ(float32(s.rot)))
	Model = Model.Mul4(mathgl.Scale3D(float32(s.xS), float32(s.yS), float32(1)))

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

	//vertexInfo.Print()

	return vertexInfo
}

func (s *Sprite) SetColor(r, g, b, a float32) {
	s.r = r
	s.g = g
	s.b = b
	s.a = a
}

func (s *Sprite) GetImage() Drawable {
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
