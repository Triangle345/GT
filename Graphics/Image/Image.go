package Image

import (
	"GT/Graphics/Opengl"
	"fmt"
	"image"

	_ "image/jpeg"
	_ "image/png"
)

// Color represents floating point unit color 0 to 1
type Color struct {
	R, G, B, A float32
}

type Image struct {
	// data          *image.Image
	*aggregateImageSection

	// uvs for image
	uvs []float32
}

type FontImage struct {
	Image
	R rune
}

func (this *FontImage) VertexData() *Opengl.OpenGLVertexInfo {
	v := Opengl.OpenGLVertexInfo{}

	w := float32(this.Bounds().Dx())
	h := float32(this.Bounds().Dy())

	// first tri
	idx := v.NewVertex(0*w, -0.5*h, 1.0)
	v.SetUV(idx, this.uvs[0], this.uvs[1])
	v.SetAggregateId(idx, this.aggrId)

	idx = v.NewVertex(1*w, -0.5*h, 1.0)
	v.SetUV(idx, this.uvs[2], this.uvs[3])
	v.SetAggregateId(idx, this.aggrId)

	idx = v.NewVertex(0*w, 0.5*h, 1.0)
	v.SetUV(idx, this.uvs[6], this.uvs[7])
	v.SetAggregateId(idx, this.aggrId)

	// second tri
	idx = v.NewVertex(0*w, 0.5*h, 1.0)
	v.SetUV(idx, this.uvs[6], this.uvs[7])
	v.SetAggregateId(idx, this.aggrId)

	idx = v.NewVertex(1*w, -0.5*h, 1.0)
	v.SetUV(idx, this.uvs[2], this.uvs[3])
	v.SetAggregateId(idx, this.aggrId)

	idx = v.NewVertex(1*w, 0.5*h, 1.0)
	v.SetUV(idx, this.uvs[4], this.uvs[5])
	v.SetAggregateId(idx, this.aggrId)

	return &v
}

// SpriteSheetImage TODO: may be useful for sprite grouping in aggregate, SpriteRenderer will handle animations though
type SpriteSheetImage struct {
	Image
	section image.Rectangle
}

// Bounds will "override" the normal image bounds
func (this Image) Bounds() image.Rectangle {
	return this.Section
}

func generateUVs(bounds image.Rectangle) []float32 {
	var uvs []float32

	topLeft := image.Point{bounds.Min.X, bounds.Min.Y}
	u, v := GetUVFromPosition(topLeft)
	uvs = append(uvs, u, v)

	topRight := image.Point{bounds.Max.X, bounds.Min.Y}
	u, v = GetUVFromPosition(topRight)
	uvs = append(uvs, u, v)

	bottomRight := image.Point{bounds.Max.X, bounds.Max.Y}
	u, v = GetUVFromPosition(bottomRight)
	uvs = append(uvs, u, v)

	bottomLeft := image.Point{bounds.Min.X, bounds.Max.Y}
	u, v = GetUVFromPosition(bottomLeft)
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
	v.SetAggregateId(idx, this.aggrId)

	idx = v.NewVertex(0.5*w, -0.5*h, 1.0)
	v.SetUV(idx, this.uvs[2], this.uvs[3])
	v.SetAggregateId(idx, this.aggrId)

	idx = v.NewVertex(-0.5*w, 0.5*h, 1.0)
	v.SetUV(idx, this.uvs[6], this.uvs[7])
	v.SetAggregateId(idx, this.aggrId)

	// second tri
	idx = v.NewVertex(-0.5*w, 0.5*h, 1.0)
	v.SetUV(idx, this.uvs[6], this.uvs[7])
	v.SetAggregateId(idx, this.aggrId)

	idx = v.NewVertex(0.5*w, -0.5*h, 1.0)
	v.SetUV(idx, this.uvs[2], this.uvs[3])
	v.SetAggregateId(idx, this.aggrId)

	idx = v.NewVertex(0.5*w, 0.5*h, 1.0)
	v.SetUV(idx, this.uvs[4], this.uvs[5])
	v.SetAggregateId(idx, this.aggrId)

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

	img.Section = image.Rectangle{origB.Min.Add(bounds.Min), origB.Min.Add(bounds.Max)}
	img.uvs = generateUVs(img.Section)

	return img, nil
}

// creates a new image
func NewImage(path string) (retImg Image, err error) {

	if newImg := GetImageSection(path); newImg != nil {

		imgRet := Image{newImg, generateUVs(newImg.Section)}
		return imgRet, nil
	}

	return Image{}, fmt.Errorf("Cannot get valid image section for path: %s", path)

}

// creates a new font rune image
//TODO create interface ImageSection that has a "recursive" getSection function
func NewFontImage(font string, r rune) (retImg FontImage, err error) {

	if fontSec := GetFontImageSection(font); fontSec != nil {

		fontImg := Image{fontSec.aggregateImageSection, generateUVs(fontSec.Section)}
		runeImg, _ := fontImg.SubImage(fontSec.FontSections[r])

		return FontImage{runeImg, r}, nil
	}

	return FontImage{}, fmt.Errorf("Cannot get valid font section for font: %s", font)

}
