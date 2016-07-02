package Image

import (
	"GT/Graphics/Opengl"
	"fmt"
	"image"

	_ "image/jpeg"
	_ "image/png"
)

type GameImage interface {
	image.Image
	VertexData() *Opengl.OpenGLVertexInfo
}

type Image struct {
	// data          *image.Image
	*aggregateImageSection

	// uvs for image
	uvs []float32

	isFont bool
}

type FontImage struct {
	Image
	r rune
}

func (this *FontImage) VertexData() *Opengl.OpenGLVertexInfo {
	v := Opengl.OpenGLVertexInfo{}

	w := float32(this.Bounds().Dx())
	h := float32(this.Bounds().Dy())

	// first tri
	idx := v.NewVertex(0*w, -0.5*h, 1.0)
	v.SetUV(idx, this.uvs[0], this.uvs[1])

	idx = v.NewVertex(1*w, -0.5*h, 1.0)
	v.SetUV(idx, this.uvs[2], this.uvs[3])

	idx = v.NewVertex(0*w, 0.5*h, 1.0)
	v.SetUV(idx, this.uvs[6], this.uvs[7])

	// second tri
	idx = v.NewVertex(0*w, 0.5*h, 1.0)
	v.SetUV(idx, this.uvs[6], this.uvs[7])

	idx = v.NewVertex(1*w, -0.5*h, 1.0)
	v.SetUV(idx, this.uvs[2], this.uvs[3])

	idx = v.NewVertex(1*w, 0.5*h, 1.0)
	v.SetUV(idx, this.uvs[4], this.uvs[5])

	return &v
}

// SpriteSheetImage TODO: may be useful for sprite grouping in aggregate, SpriteRenderer will handle animations though
type SpriteSheetImage struct {
	Image
	section image.Rectangle
}

// Bounds will "override" the normal image bounds
func (this Image) Bounds() image.Rectangle {
	return this.section
}

func getUVFromPosition(point image.Point) (u, v float32) {

	u = float32(point.X) / float32(AggrImg.aggregateImage.Bounds().Dx())
	v = float32(point.Y) / float32(AggrImg.aggregateImage.Bounds().Dy())

	return
}

func getUVs(bounds image.Rectangle) []float32 {
	var uvs []float32

	topLeft := image.Point{bounds.Min.X, bounds.Min.Y}
	u, v := getUVFromPosition(topLeft)
	uvs = append(uvs, u, v)

	topRight := image.Point{bounds.Max.X, bounds.Min.Y}
	u, v = getUVFromPosition(topRight)
	uvs = append(uvs, u, v)

	bottomRight := image.Point{bounds.Max.X, bounds.Max.Y}
	u, v = getUVFromPosition(bottomRight)
	uvs = append(uvs, u, v)

	bottomLeft := image.Point{bounds.Min.X, bounds.Max.Y}
	u, v = getUVFromPosition(bottomLeft)
	uvs = append(uvs, u, v)

	return uvs
}

func (this *Image) VertexData() *Opengl.OpenGLVertexInfo {
	v := Opengl.OpenGLVertexInfo{}

	w := float32(this.Bounds().Dx())
	h := float32(this.Bounds().Dy())

	// first tri
	idx := v.NewVertex(-0.5*w, -0.5*h, 1.0)
	v.SetUV(idx, this.uvs[0], this.uvs[1])

	idx = v.NewVertex(0.5*w, -0.5*h, 1.0)
	v.SetUV(idx, this.uvs[2], this.uvs[3])

	idx = v.NewVertex(-0.5*w, 0.5*h, 1.0)
	v.SetUV(idx, this.uvs[6], this.uvs[7])

	// second tri
	idx = v.NewVertex(-0.5*w, 0.5*h, 1.0)
	v.SetUV(idx, this.uvs[6], this.uvs[7])

	idx = v.NewVertex(0.5*w, -0.5*h, 1.0)
	v.SetUV(idx, this.uvs[2], this.uvs[3])

	idx = v.NewVertex(0.5*w, 0.5*h, 1.0)
	v.SetUV(idx, this.uvs[4], this.uvs[5])

	return &v
}

// returns a new image that is just a sub image based on given bounds
func (this *Image) SubImage(bounds image.Rectangle) (Image, error) {

	// create totally new image with copy of section
	aggrImgCopy := *this.aggregateImageSection
	img := Image{aggregateImageSection: &aggrImgCopy}

	origB := img.Bounds()

	if origB.Bounds().Dy() < bounds.Dy() {
		return img, fmt.Errorf("Bounds of sub image exceeds image size, sub %d vs original %d", bounds.Dy(), origB.Dy())
	}

	// this is actually a sub image within a subimage

	img.section = image.Rectangle{origB.Min.Add(bounds.Min), origB.Min.Add(bounds.Max)}
	img.uvs = getUVs(img.section)
	img.isFont = this.isFont

	return img, nil
}

// creates a new image
func NewImage(path string) (retImg Image, err error) {

	if newImg := AggrImg.GetImageSection(path); newImg != nil {

		imgRet := Image{newImg, getUVs(newImg.section), false}

		return imgRet, nil
	}

	return Image{}, fmt.Errorf("Cannot get valid image section for path: %s", path)

}

// creates a new font rune image
func NewFontImage(font string, r rune) (retImg FontImage, err error) {

	if fontSec := AggrImg.GetFontImageSection(font); fontSec != nil {

		fontImg := Image{fontSec.aggregateImageSection, getUVs(fontSec.section), true}

		runeImg, _ := fontImg.SubImage(fontSec.FontSections[r])

		return FontImage{runeImg, r}, nil
	}

	return FontImage{}, fmt.Errorf("Cannot get valid font section for font: %s", font)

}
