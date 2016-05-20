package Image

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"strings"

	"path/filepath"

	"github.com/go-gl/gl/v3.2-core/gl"
)

type AggregateImageSection struct {
	image.Image
	pathName string
	section  image.Rectangle
}

type AggregateImage struct {
	images     []*AggregateImageSection
	sectionMap map[string]*AggregateImageSection

	// font sections
	fontImages     []*AggregateImageSection
	fontSectionMap map[string]*AggregateImageSection

	aggregateImage image.Image
	textureId      uint32
	initialized    bool
}

//var images []image.Image = []image.Image{}

/** NewAggregateImage
 *
 * Walk the directory and aggregate all images to one image and store in meemory
 * @param {[type]} location string [description]
 */
func NewAggregateImage(location string) *AggregateImage {
	imgAgg := &AggregateImage{sectionMap: map[string]*AggregateImageSection{}, initialized: false}
	imgAgg.fileWalker(location)

	height := 0
	maxWidth := 0

	for _, imgSec := range imgAgg.images {
		height += imgSec.Bounds().Dy()
		if imgSec.Bounds().Dx() > maxWidth {
			maxWidth = imgSec.Bounds().Dx()
		}

		//fmt.Printf("Dy is: %d\n", img.Bounds().Dy())
	}

	fmt.Printf("Height is %d\n", height)
	fmt.Printf("Width is %d\n", maxWidth)

	//rectangle for the big image
	finalImg := image.Rectangle{image.Point{0, 0}, image.Point{maxWidth, height}}

	rgbaFinal := image.NewRGBA(finalImg)

	lastLocY := 0

	for _, imgSec := range imgAgg.images {

		fmt.Printf("lastLocY: %d\n", lastLocY)

		p1 := image.Point{0, lastLocY}
		p2 := p1.Add(imgSec.Bounds().Size())
		imgSec.section = image.Rectangle{p1, p2}
		fmt.Println(p1.Add(imgSec.Bounds().Size()))
		draw.Draw(rgbaFinal, imgSec.section, imgSec, image.Point{0, 0}, draw.Src) // draw first image
		lastLocY += imgSec.Bounds().Dy()

	}

	// store the final aggregate image
	imgAgg.aggregateImage = rgbaFinal

	return imgAgg
}

func (this * AggregateImage) AppendImage (image.Image img) {
	newDim := image.Rectangle{image.Point{0, 0}, this.Bounds().Max().Add(img.Bounds().Max())}

	rgbaFinal := image.NewRGBA(finalImg)

	draw.Draw(rgbaFinal, img.Bounds().Max(), imgSec, image.Point{0, 0}, draw.Src) // draw first image
}

func (this *AggregateImage) Bind2GL() {

	if this.initialized {
		return
	}

	if rgba, ok := this.aggregateImage.(*image.RGBA); ok {
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

			fmt.Println("Cannot load Image in location: ./")
			os.Exit(-1)
		}

		this.textureId = texture
	} else {
		fmt.Println("Image not RGBA at location: ./")
		os.Exit(-1)
	}

	this.initialized = true

}

func (this *AggregateImage) loadImage(imgPath string) error {
	imgFile, err := os.Open(imgPath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	img, _, err := image.Decode(imgFile)

	if err != nil {
		fmt.Println(err)
		return err
	}

	sec := &AggregateImageSection{img, imgPath, image.Rectangle{}}

	// populate both section map and image list with section
	this.images = append(this.images, sec)
	this.sectionMap[imgPath] = sec

	return nil
}

func (this *AggregateImage) fileVisitor(path string, f os.FileInfo, err error) error {
	if strings.Contains(path, ".png") {
		fmt.Printf("Processing: %s\n", path)
		this.loadImage(path)
	}
	return nil
}

func (this *AggregateImage) fileWalker(path string) {
	filepath.Walk(path, this.fileVisitor)
}

func (this *AggregateImage) GetImageSection(path string) *AggregateImageSection {
	return this.sectionMap[path]
}

func (this *AggregateImage) Print(path string) {
	outfirst, errfirst := os.Create(path)
	if errfirst != nil {
		fmt.Println(errfirst)
		//os.Exit(1)
	}

	if this.aggregateImage == nil {
		fmt.Println("err: no image created!")
	}

	png.Encode(outfirst, this.aggregateImage)
}
