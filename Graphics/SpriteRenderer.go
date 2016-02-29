// Sprite
package Graphics

import (
	// "github.com/go-gl/gl/v3.2-core/gl"

	"GT/Graphics/Components"
	"GT/Graphics/Opengl"
	"fmt"

	mathgl "github.com/go-gl/mathgl/mgl32"
)

type Renderer interface {
	getGLVertexInfo() Opengl.OpenGLVertexInfo
}

func NewSpriteRenderer() SpriteRenderer {

	sprite := SpriteRenderer{a: 1, Transform: Components.NewTransform()}
	return sprite
}

type SpriteRenderer struct {
	Transform Components.Transform
	// img
	//img Graphics.RectangularArea
	//TODO: maybe add real image
	img Drawable
	// color
	r, g, b, a float32
}

func (this *SpriteRenderer) SetImageSpriteSheet(imageLoc string) {
	img, err := NewSpriteSheetImage(imageLoc, NewRectangularArea(0, 0, 128, 128))

	if err != nil {
		fmt.Println("Cannot create image: " + err.Error())
	}

	this.img = img
}
func (this *SpriteRenderer) SetImage(imageLoc string) {

}

func (s SpriteRenderer) Initialize() {

}

func (s SpriteRenderer) Update(delta float32) {

	var rect RectangularArea

	//TODO need more elegant way to handle this
	if img, ok := s.img.(SpriteSheetImage); ok {
		rect = img.GetRect()
	} else {
		fmt.Println("Cannot Get rect: SpriteRenderer:Update()")
	}

	// vertexInfo := Opengl.OpenGLVertexInfo{}
	w, h := rect.GetDimensions()

	vertex_data := []float32{-0.5 * w, 0.5 * h, 1.0, 0.5 * w, 0.5 * h, 1.0, 0.5 * w, -0.5 * h, 1.0, -0.5 * w, -0.5 * h, 1.0}

	elements := []uint32{uint32(0), uint32(1), uint32(2), uint32(0), uint32(2), uint32(3)}
	uvs := rect.GetUVs()

	Model := s.Transform.GetUpdatedModel()

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
	Opengl.AddVertexData(&vertexInfo)
}

func (s *SpriteRenderer) SetColor(r, g, b, a float32) {
	s.r = r
	s.g = g
	s.b = b
	s.a = a
}
