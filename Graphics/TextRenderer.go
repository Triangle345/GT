// Package Graphics
package Graphics

import (
	// "github.com/go-gl/gl/v3.2-core/gl"

	"GT/Graphics/Components"
	"GT/Graphics/Font"
	"GT/Graphics/Image"
	"GT/Graphics/Opengl"
	"fmt"
	"image"

	"github.com/golang/freetype/truetype"

	mathgl "github.com/go-gl/mathgl/mgl32"
)

func NewTextRenderer() *TextRenderer {

	text := TextRenderer{a: 1, size: 100, font: "Times New Roman Normal"}
	return &text

}

type fontImage struct {
	runeImg image.Image
	uvs     []float32
}

type TextRenderer struct {
	// the parent node
	//Parent *Components.Node
	Components.ChildComponent

	size int

	font string

	text string

	runeImgs []*fontImage

	// color
	r, g, b, a float32
}

func (this *TextRenderer) SetFont(font string) {

	this.font = font

}

func (this *TextRenderer) SetSize(size int) {

	this.size = size

}

func (this *TextRenderer) SetText(text string) {

	this.text = text

	for _, t := range text {
		img, err := Image.NewFontImage(this.font, t)

		if err != nil {
			fmt.Println("Cannot create image: " + err.Error())
		}

		this.runeImgs = append(this.runeImgs, &fontImage{img, img.UVs()})

	}

}

func (s *TextRenderer) Initialize() {

}

func (this *TextRenderer) Update(delta float32) {

	if len(this.runeImgs) == 0 {
		return
	}

	fInfo := truetype.Font(*Font.GetFont(this.font))
	totalWidth := float32(0.0)

	for i, img := range this.runeImgs {

		// get advance width
		fIdx := fInfo.Index(rune(this.text[i]))

		// this gets teh bounds of the sub image
		w := float32(img.runeImg.Bounds().Dx())
		h := float32(img.runeImg.Bounds().Dy())

		vertex_data := []float32{-0.5*w + totalWidth, 0.5 * h, 1.0, 0.5*w + totalWidth, 0.5 * h, 1.0, 0.5*w + totalWidth, -0.5 * h, 1.0, -0.5*w + totalWidth, -0.5 * h, 1.0}

		advance := float32(fInfo.HMetric(100, fIdx).AdvanceWidth)
		//fmt.Println("Advance: ", advance)
		totalWidth += advance

		elements := []uint32{uint32(0), uint32(1), uint32(2), uint32(0), uint32(2), uint32(3)}

		Model := mathgl.Ident4()

		// check to see if its a regular node to get the updated model
		// if n, ok := s.Parent.(*Components.Node); ok {

		Model = this.GetParent().Transform().GetUpdatedModel()
		// }

		// transform all vertex data and combine it with other data
		var data []float32 = make([]float32, 0, 9*4)
		for j := 0; j < 4; j++ {
			transformation := mathgl.Vec4{vertex_data[j*3+0], vertex_data[j*3+1], vertex_data[j*3+2], 1}
			t := Model.Mul4x1(transformation)

			data = append(data, t[0], t[1], t[2], this.r, this.g, this.b, this.a, img.uvs[j*2+0], img.uvs[j*2+1])

		}

		// package everything up in an OpenGLVertexInfo
		vertexInfo := Opengl.OpenGLVertexInfo{
			VertexData: data,
			Elements:   elements,
			Stride:     4,
		}

		// send OpenGLVertex info to Opengl module
		Opengl.AddVertexData(1, &vertexInfo)

	}

}

func (s *TextRenderer) SetColor(r, g, b, a float32) {
	s.r = r
	s.g = g
	s.b = b
	s.a = a
}
