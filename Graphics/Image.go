// Graphics project Graphics.go
package Graphics

import (
	"errors"
	gl "github.com/chsc/gogl/gl21"
	Image "image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

type image struct {
	data      Image.Image
	textureId gl.Uint
}

type Drawable interface {
	Draw()
}

func (img image) Draw() {

	if img.data == nil {
		return
	}

	// transparency
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.BLEND)

	gl.Enable(gl.TEXTURE_2D)

	gl.MatrixMode(gl.MODELVIEW)

	gl.PushMatrix()

	//gl.LoadIdentity()

	gl.BindTexture(gl.TEXTURE_2D, img.textureId)

	//gl.Translatef(400, 400, 0)
	w := gl.Float(img.data.Bounds().Dx())
	h := gl.Float(img.data.Bounds().Dy())

	gl.Begin(gl.QUADS)

	gl.Color3f(1.0, 1.0, 1.0)

	// need to flip v coord because of opengl texture rendering
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(-w/2, -h/2, 0)
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(w/2, -h/2, 0)
	gl.TexCoord2f(1, 0)
	gl.Vertex3f(w/2, h/2, 0)
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(-w/2, h/2, 0)

	gl.End()
	gl.Disable(gl.TEXTURE_2D)

	gl.PopMatrix()

}

func NewImage(path string) (retImg image, err error) {
	f, e := os.Open(path)
	if e != nil {
		return retImg, e
	}

	img, _, e := Image.Decode(f)

	retImg.data = img
	retImg.textureId = 0
	if e != nil {
		return retImg, e
	}

	rgbaImg, ok := img.(*Image.NRGBA)
	if !ok {
		return retImg, errors.New("texture must be an NRGBA image")
	}

	gl.GenTextures(1, &retImg.textureId)

	gl.BindTexture(gl.TEXTURE_2D, retImg.textureId)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)

	// taken from github.com/chsc/gogl/examples
	// flip image: first pixel is lower left corner
	imgWidth, imgHeight := img.Bounds().Dx(), img.Bounds().Dy()
	data := make([]byte, imgWidth*imgHeight*4)
	lineLen := imgWidth * 4
	dest := len(data) - lineLen
	for src := 0; src < len(rgbaImg.Pix); src += rgbaImg.Stride {
		copy(data[dest:dest+lineLen], rgbaImg.Pix[src:src+rgbaImg.Stride])
		dest -= lineLen
	}

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, gl.Sizei(imgWidth), gl.Sizei(imgHeight), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Pointer(&data[0]))
	return retImg, nil
}
