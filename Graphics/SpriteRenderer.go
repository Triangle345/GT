// Sprite
package Graphics

import (
	// "github.com/go-gl/gl/v3.2-core/gl"

	"GT/Graphics/Components"
	"GT/Graphics/Opengl"
	"fmt"

	mathgl "github.com/go-gl/mathgl/mgl32"
)

func NewSpriteRenderer() *SpriteRenderer {

	sprite := SpriteRenderer{a: 1}
	return &sprite

}

type SpriteRenderer struct {
	// the parent node
	Parent Components.GameNode

	// the image
	img Drawable

	// color
	r, g, b, a float32
}

func (this *SpriteRenderer) SetImageSpriteSheet(imageLoc string) {
	//TODO: should probably cache this for speed
	img, err := NewSpriteSheetImage(imageLoc, NewRectangularArea(0, 0, 128, 128))

	if err != nil {
		fmt.Println("Cannot create image: " + err.Error())
	}

	this.img = img
}
func (this *SpriteRenderer) SetImage(imageLoc string) {

}

func (s *SpriteRenderer) Initialize() {

}

func (s *SpriteRenderer) SetParent(node Components.GameNode) {
	s.Parent = node
}

func (s *SpriteRenderer) GetParent() Components.GameNode {
	return s.Parent
}

func (s *SpriteRenderer) Update(delta float32) {

	var rect RectangularArea

	//TODO need more elegant way to handle this, what if its regular image?
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

	Model := mathgl.Ident4()

	// check to see if its a regular node to get the updated model
	if n, ok := s.Parent.(*Components.Node); ok {
		Model = n.GetUpdatedModel()
	}

	// transform all vertex data and combine it with other data
	var data []float32
	for j := 0; j < 4; j++ {
		transformation := mathgl.Vec4{vertex_data[j*3+0], vertex_data[j*3+1], vertex_data[j*3+2], 1}
		t := Model.Mul4x1(transformation)
		data = append(data, t[0], t[1], t[2], s.r, s.g, s.b, s.a, uvs[j*2+0], uvs[j*2+1])

	}

	// package everything up in an OpenGLVertexInfo
	vertexInfo := Opengl.OpenGLVertexInfo{
		VertexData: data,
		Elements:   elements,
		Stride:     4,
	}

	// send OpenGLVertex info to Opengl module
	Opengl.AddVertexData(s.img.GetId(), &vertexInfo)

	//TODO: need to fix this, does not draw two images at once, maybe put inside vertex info and have that sorted by texture hash
	// maybe add an AddVertexData("texture", vertexInfo)
	// s.img.Draw()
}

func (s *SpriteRenderer) SetColor(r, g, b, a float32) {
	s.r = r
	s.g = g
	s.b = b
	s.a = a
}
