// Package Text
package Components

import (
	// "github.com/go-gl/gl/v3.2-core/gl"

	"GT/Graphics/Image"
	"GT/Graphics/Opengl"
	"fmt"

	"GT/Graphics/Font"

	mathgl "github.com/go-gl/mathgl/mgl32"
	"github.com/golang/freetype/truetype"
)

var curTotWidth float32 = 0

func NewTextRenderer() *TextRenderer {

	text := TextRenderer{size: 100, font: "Raleway"}
	text.CustomTransform(text.TextTransform)
	return &text

}

type TextRenderer struct {
	Renderer

	size int

	font string

	text string

	runeImgs []Opengl.RenderObject

	// color
	// r, g, b, a float32
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

		this.runeImgs = append(this.runeImgs, &img)

	}

}

func (this *TextRenderer) TextTransform(vertex mathgl.Vec4) mathgl.Vec4 {
	// apply font transforms
	scale := float32(this.size) / 100.0
	t := mathgl.Translate3D(curTotWidth, 0, 0).Mul4x1(vertex)

	t = mathgl.Scale3D(scale, scale, 1).Mul4x1(t)

	return t
}

func (this *TextRenderer) Update(delta float32) {

	if len(this.runeImgs) == 0 {
		return
	}

	// Model := mathgl.Ident4()
	// Model = this.GetParent().transform.GetUpdatedModel()

	fInfo := truetype.Font(*Font.GetFont(this.font))

	curTotWidth = float32(0.0)
	// scale := float32(this.size) / 100.0

	for idx, img := range this.runeImgs {
		// vData := img.VertexData()
		i := fInfo.Index(rune(this.text[idx]))

		hmetric := fInfo.HMetric(100, i)
		w := float32(hmetric.AdvanceWidth) //float32(img.Bounds().Dx())
		// h := float32(img.runeImg.Bounds().Dy())

		// for j := 0; j < vData.NumVerts(); j++ {
		// 	x, y, z := vData.GetVertex(j)
		// 	transformation := mathgl.Vec4{x, y, z, 1}

		// 	// apply font transforms
		// 	t := mathgl.Translate3D(totalWidth, 0, 0).Mul4x1(transformation)

		// 	t = mathgl.Scale3D(scale, scale, 1).Mul4x1(t)

		// 	t = Model.Mul4x1(t)

		// 	// data = append(data, t[0], t[1], t[2], this.r, this.g, this.b, this.a, img.uvs[j*2+0], img.uvs[j*2+1], 1.0)
		// 	// vData.SetVertex(j, t[0], t[1], t[3])
		// 	// vData.SetColor(j, this.r, this.g, this.b, this.a)

		// }

		this.Render(img)

		curTotWidth += w

		// send OpenGLVertex info to Opengl module
		// Opengl.AddVertexData(1, vData)

	}

}

func (s *TextRenderer) SetColor(r, g, b, a float32) {
	s.Color = &Image.Color{r, g, b, a}
}
