// Graphics project Graphics.go
package Graphics

import (
	// "errors"
	"fmt"
	"image"
	// gl "github.com/chsc/gogl/gl21"
	// "github.com/Jragonmiris/mathgl"

	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/go-gl/gl/v3.2-core/gl"
)

var imagecache map[string]Image = make(map[string]Image)

type Image struct {
	data          *image.Image
	height, width int
	textureId     uint32
	rect          RectangularArea
}

type SpriteSheetImage struct {
	*Image
}

type Drawable interface {
	Draw()
	GetId() uint32
}

func (this Image) GetId() uint32 {
	return this.textureId
}

func (this Image) GetRect() RectangularArea {
	return this.rect
}

func (img Image) Draw() {

	if img.data == nil {
		return
	}

	// gl.ActiveTexture(texId2)

	gl.BindTexture(gl.TEXTURE_2D, img.textureId)

}

func (img Image) GetUVFromPosition(x, y float32) (u, v float32) {

	u = x / float32(img.width)
	v = y / float32(img.height)

	return
}

func NewSpriteSheetImage(path string, rect RectangularArea) (retImg SpriteSheetImage, err error) {
	if newImg, err := NewImage(path); err != nil {
		//TODO maybe replace with default no-photo image
		ss := SpriteSheetImage{Image: &newImg}
		fmt.Println("Error in NewSpriteSheetImage(): cannot create")
		fmt.Println(err)
		return ss, err
	} else {
		newImg.rect = rect
		ss := SpriteSheetImage{Image: &newImg}
		var uvs []float32
		x, y := ss.GetUVFromPosition(rect.BottomLeft())
		uvs = append(uvs, x, y)
		x, y = ss.GetUVFromPosition(rect.BottomRight())
		uvs = append(uvs, x, y)
		x, y = ss.GetUVFromPosition(rect.TopRight())
		uvs = append(uvs, x, y)
		x, y = ss.GetUVFromPosition(rect.TopLeft())
		uvs = append(uvs, x, y)

		// set uv coords
		ss.rect.uvs = uvs

		return ss, nil
	}
}

func NewImage(path string) (retImg Image, err error) {

	// first look at cache for image
	if imgCache, ok := imagecache[path]; ok {
		return imgCache, nil
	}

	imgFile, err := os.Open(path)
	if err != nil {
		return Image{}, err
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return Image{}, err
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return Image{}, fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32

	gl.GenTextures(1, &texture)
	// gl.ActiveTexture(gl.TEXTURE1)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_R, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(rgba.Rect.Size().X), int32(rgba.Rect.Size().Y), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))

	if gl.GetError() != gl.NO_ERROR {

		return Image{}, fmt.Errorf("Failed to load texture: " + path)
	}

	imgRet := Image{data: &img, textureId: texture, width: rgba.Rect.Size().X, height: rgba.Rect.Size().Y}
	imagecache[path] = imgRet
	return imgRet, nil
}
