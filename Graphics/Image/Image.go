package Image

import (
	"fmt"
	"image"

	_ "image/jpeg"
	_ "image/png"
)

type Image struct {
	// data          *image.Image
	*aggregateImageSection

	// uvs for image
	uvs []float32
}

// UVs will return the uv coords for this image (based on aggregate image)
func (this *Image) UVs() []float32 {
	return this.uvs
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

// func (this *Image) GetUVs() (BLU, BLV, BRU, BRV, TRU, TRV, TLU, TLV float32) {
// 	BLU = this.uvs[0]
// 	BLV = this.uvs[1]

// 	BRU = this.uvs[2]
// 	BRV = this.uvs[3]

// 	TRU = this.uvs[4]
// 	TRV = this.uvs[5]

// 	TLU = this.uvs[6]
// 	TLV = this.uvs[7]

// 	return BLU, BLV, BRU, BRV, TRU, TRV, TLU, TLV
// }

// func (this *Image) GetVertexData() []float32 {
// 	vData := []float32{}

// 	w := this.Bounds().Dx()
// 	h := this.Bounds().Dy()

// 	// TODO maybe add rgba to Image?
// 	//first tri
// 	vData = append(vData,
// 		-0.5*w, -0.5*h, 1.0, 0, 0, 0, 0, this.uvs[0], this.uvs[1], Opengl.TEXTURED)

// 	vData = append(vData,
// 		0.5*w, -0.5*h, 1.0, 0, 0, 0, 0, this.uvs[2], this.uvs[3], Opengl.TEXTURED)

// 	vData = append(vData,
// 		-0.5*w, 0.5*h, 1.0, 0, 0, 0, 0, this.uvs[6], this.uvs[7], Opengl.TEXTURED)

// 	// second tri
// 	vData = append(vData,
// 		-0.5*w, 0.5*h, 1.0, 0, 0, 0, 0, this.uvs[6], this.uvs[7], Opengl.TEXTURED)

// 	vData = append(vData,
// 		0.5*w, -0.5*h, 1.0, 0, 0, 0, 0, this.uvs[2], this.uvs[3], Opengl.TEXTURED)

// 	vData = append(vData,
// 		0.5*w, 0.5*h, 1.0, 0, 0, 0, 0, this.uvs[4], this.uvs[5], Opengl.TEXTURED)

// 	// first tri
// 	// -0.5 * w, -0.5 * h, 1.0, this.uvs[0], this.uvs[1]
// 	// 0.5 * w, -0.5 * h, 1.0,
// 	// -0.5 * w, 0.5 * h, 1.0,
// 	// // second tri
// 	// -0.5 * w, 0.5 * h, 1.0,
// 	// 0.5 * w, -0.5 * h, 1.0,
// 	// 0.5 * w, 0.5 * h, 1.0

// 	return vData
// }

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

	return img, nil
}

// creates a new image
func NewImage(path string) (retImg Image, err error) {

	if newImg := AggrImg.GetImageSection(path); newImg != nil {

		imgRet := Image{newImg, getUVs(newImg.section)}

		return imgRet, nil
	}

	return Image{}, fmt.Errorf("Cannot get valid image section for path: %s", path)

}

// creates a new font rune image
func NewFontImage(font string, r rune) (retImg Image, err error) {

	if fontSec := AggrImg.GetFontImageSection(font); fontSec != nil {

		fontImg := Image{fontSec.aggregateImageSection, getUVs(fontSec.section)}
		runeImg, _ := fontImg.SubImage(fontSec.FontSections[r])

		return runeImg, nil
	}

	return Image{}, fmt.Errorf("Cannot get valid font section for font: %s", font)

}
