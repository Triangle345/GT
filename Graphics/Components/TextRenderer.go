// Package Graphics
package Components

import (
	// "github.com/go-gl/gl/v3.2-core/gl"

	"GT/Graphics/Image"
	"GT/Graphics/Opengl"
	"fmt"

	mathgl "github.com/go-gl/mathgl/mgl32"
)

func NewTextRenderer() *TextRenderer {

	text := TextRenderer{a: 1, size: 100, font: "Raleway"}
	return &text

}

type fontImage struct {
	r       rune
	runeImg Image.Image
	uvs     []float32
}

type TextRenderer struct {
	ChildComponent

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

// size of text in pixels (current max 100)
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

		this.runeImgs = append(this.runeImgs, &fontImage{t, img, img.UVs()})

	}

}

func (s *TextRenderer) Initialize() {

}

func (this *TextRenderer) Update(delta float32) {

	if len(this.runeImgs) == 0 {
		return
	}

	Model := mathgl.Ident4()
	Model = this.GetParent().transform.GetUpdatedModel()

	// fInfo := truetype.Font(*Font.GetFont(this.font))
	totalWidth := float32(0.0)
	scale := float32(this.size) / 100.0

	for _, img := range this.runeImgs {
		vData := img.runeImg.VertexData()
		// get advance width
		// fIdx := fInfo.Index(rune(this.text[i]))
		// this gets teh bounds of the sub image
		w := float32(img.runeImg.Bounds().Dx())
		// h := float32(img.runeImg.Bounds().Dy())

		//vertex_data := []float32{-0.5 * w, 0.5 * h, 1.0, 0.5 * w, 0.5 * h, 1.0, 0.5 * w, -0.5 * h, 1.0, -0.5 * w, -0.5 * h, 1.0}
		// vertex_data := []float32{0, 0.5 * h, 1.0, w, 0.5 * h, 1.0, w, -0.5 * h, 1.0, 0, -0.5 * h, 1.0}

		// elements := []uint32{uint32(0), uint32(1), uint32(2), uint32(0), uint32(2), uint32(3)}

		// transform all vertex data and combine it with other data
		// var data []float32 = make([]float32, 0, 9*4)
		for j := 0; j < vData.NumVerts(); j++ {
			x, y, z := vData.GetVertex(j)
			transformation := mathgl.Vec4{x, y, z, 1}

			// apply font transforms
			t := mathgl.Translate3D(totalWidth, 0, 0).Mul4x1(transformation)

			t = mathgl.Scale3D(scale, scale, 1).Mul4x1(t)

			t = Model.Mul4x1(t)

			// data = append(data, t[0], t[1], t[2], this.r, this.g, this.b, this.a, img.uvs[j*2+0], img.uvs[j*2+1], 1.0)
			vData.SetVertex(j, t[0], t[1], t[3])
			vData.SetColor(j, this.r, this.g, this.b, this.a)

		}

		totalWidth += w
		// package everything up in an OpenGLVertexInfo
		// vertexInfo := Opengl.OpenGLVertexInfo{
		// // VertexData: data,
		// // Elements:   elements,
		// // Stride:     4,
		// }

		// send OpenGLVertex info to Opengl module
		Opengl.AddVertexData(1, vData)

	}

}

func (s *TextRenderer) SetColor(r, g, b, a float32) {
	s.r = r
	s.g = g
	s.b = b
	s.a = a
}
