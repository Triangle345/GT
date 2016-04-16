package Image

import (
	// "errors"
	"fmt"
	"image"
	// gl "github.com/chsc/gogl/gl21"
	// "github.com/Jragonmiris/mathgl"

	_ "image/jpeg"
	_ "image/png"
)

//TODO: create some way of doing this automatically maybe based on some config file
var AggrImg *AggregateImage = NewAggregateImage("./")

// func Init() {
// 	if rgba, ok := aggrImg.aggregateImage.(*image.RGBA); ok {
// 		var texture uint32
//
// 		gl.GenTextures(1, &texture)
// 		// gl.ActiveTexture(gl.TEXTURE1)
// 		gl.BindTexture(gl.TEXTURE_2D, texture)
// 		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
// 		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
// 		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
// 		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
// 		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_R, gl.CLAMP_TO_EDGE)
//
// 		gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(rgba.Rect.Size().X), int32(rgba.Rect.Size().Y), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))
// 		if gl.GetError() != gl.NO_ERROR {
//
// 			fmt.Println("Cannot load Image in location: ./")
// 			os.Exit(-1)
// 		}
//
// 		aggrImg.textureId = texture
// 	} else {
// 		fmt.Println("Image not RGBA at location: ./")
// 		os.Exit(-1)
// 	}
//
// }

// var imagecache map[string]Image = make(map[string]Image)

type Image struct {
	// data          *image.Image
	*AggregateImageSection

	// uvs for image
	uvs []float32
}

type SpriteSheetImage struct {
	Image
	section image.Rectangle
}

func (this Image) Bounds() image.Rectangle {
	return this.section
}

func (this *Image) SubImage(bounds image.Rectangle) (Image, error) {
	img, err := NewImage(this.pathName)

	if err != nil {
		return img, err
	}

	origB := img.Bounds()

	if origB.Bounds().Dy() < bounds.Dy() {
		return img, fmt.Errorf("Bounds of sub image exceeds image size, sub %d vs original %d", bounds.Dy(), origB.Dy())
	}

	// this is actually a sub image within a subimage

	img.section = image.Rectangle{origB.Min.Add(bounds.Min), origB.Min.Add(bounds.Max)}
	fmt.Println(img.section)
	return img, nil
}

func GetUVFromPosition(point image.Point) (u, v float32) {

	u = float32(point.X) / float32(AggrImg.aggregateImage.Bounds().Dx())
	v = float32(point.Y) / float32(AggrImg.aggregateImage.Bounds().Dy())

	return
}

func GetUVs(bounds image.Rectangle) []float32 {
	var uvs []float32

	bottomLeft := image.Point{bounds.Min.X, bounds.Max.Y}
	u, v := GetUVFromPosition(bottomLeft)
	uvs = append(uvs, u, v)

	bottomRight := image.Point{bounds.Max.X, bounds.Max.Y}
	u, v = GetUVFromPosition(bottomRight)
	uvs = append(uvs, u, v)

	topRight := image.Point{bounds.Max.X, bounds.Min.Y}
	u, v = GetUVFromPosition(topRight)
	uvs = append(uvs, u, v)

	topLeft := image.Point{bounds.Min.X, bounds.Min.Y}
	u, v = GetUVFromPosition(topLeft)
	uvs = append(uvs, u, v)

	return uvs
}

// func NewSpriteSheetImage(path string, rect image.Rectangle) (retImg SpriteSheetImage, err error) {
//
// 	if newImg, err := NewImage(path); err == nil {
//
// 		// set uv coords
//
// 		entireBound := newImg.data.section
// 		newMin := newImg.data.section.Min.Add(entireBound.Min)
// 		newMax := newImg.data.section.Max.Sub(entireBound.Max)
// 		subBound := image.Rectangle{newMin, newMax}
// 		newImg.uvs = GetUVs(subBound)
// 		ss := SpriteSheetImage{image: newImg, section: rect}
// 		// fmt.Println("uvs")
// 		// fmt.Println(uvs)
// 		return ss, nil
// 	}
//
// 	return SpriteSheetImage{}, nil
// }

func NewImage(path string) (retImg Image, err error) {

	if newImg := AggrImg.GetImageSection(path); newImg != nil {

		imgRet := Image{newImg, GetUVs(newImg.section)}

		return imgRet, nil
	}

	return Image{}, fmt.Errorf("Cannot get valid image section for path: %s", path)

}
