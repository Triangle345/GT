// Graphics project Graphics.go
package Graphics

import (
	// "errors"
	"fmt"
	// gl "github.com/chsc/gogl/gl21"
	// "github.com/Jragonmiris/mathgl"
	"github.com/go-gl/gl/v3.2-core/gl"
	Image "image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

type image struct {
	data          *Image.Image
	height, width int
	textureId     uint32
}

type Drawable interface {
	Draw()
}

func (img image) Draw() {

	if img.data == nil {
		return
	}

	// gl.ActiveTexture(texId2)

	gl.BindTexture(gl.TEXTURE_2D, img.textureId)

}

func (img image) GetUVFromPosition(x, y float32) (u, v float32) {

	u = x / float32(img.width)
	v = y / float32(img.height)

	return
}

func NewImage(path string) (retImg image, err error) {

	imgFile, err := os.Open(path)
	if err != nil {
		return image{}, err
	}
	img, _, err := Image.Decode(imgFile)
	if err != nil {
		return image{}, err
	}

	rgba := Image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return image{}, fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, Image.Point{0, 0}, draw.Src)

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

		return image{}, fmt.Errorf("Failed to load texture: " + path)
	}

	return image{data: &img, textureId: texture, width: rgba.Rect.Size().X, height: rgba.Rect.Size().Y}, nil
}
