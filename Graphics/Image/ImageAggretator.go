package Image

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"strings"

	"path/filepath"

	"GT/Graphics/Font"
	"GT/Graphics/Opengl"
)

//TODO: create some way of doing this automatically maybe based on some config file
var AggrImg *AggregateImage

//TODO maybe not the best place for this? need to init somewhere else?
func init() {

	LoadImages("./")
	Font.LoadFonts("../Assets/Fonts")

	fonts := Font.GetFonts()

	for _, f := range fonts {
		fontImg := f.GetImage()
		mainSec := image.Rect(0,
			AggrImg.aggregateImage.Bounds().Dy(),
			fontImg.Bounds().Dx(),
			AggrImg.aggregateImage.Bounds().Dy()+fontImg.Bounds().Dy())

		fontSec := &FontImageSection{&aggregateImageSection{fontImg, f.Name(), mainSec}, f.GetSectionMap()}

		AggrImg.fonts = append(AggrImg.fonts, fontSec)
		AggrImg.fontsSectionMap[f.Name()] = fontSec

		AggrImg.AppendImage(fontSec.Image, f.Name())
	}

	Opengl.SetAggregateImage(AggrImg.aggregateImage)

	// for debug
	// AggrImg.Print("./aggr.png")
}

//Where to load all images from.
func LoadImages(path string) {
	AggrImg = NewAggregateImage("./")
}

type FontImageSection struct {
	// image.Image
	// name         string
	*aggregateImageSection
	FontSections map[rune]image.Rectangle
}

type aggregateImageSection struct {
	image.Image
	pathName string
	section  image.Rectangle
}

type AggregateImage struct {
	images     []*aggregateImageSection
	sectionMap map[string]*aggregateImageSection

	fonts           []*FontImageSection
	fontsSectionMap map[string]*FontImageSection

	aggregateImage image.Image
}

//var images []image.Image = []image.Image{}

/** NewAggregateImage
 *
 * Walk the directory and aggregate all images to one image and store in meemory
 * @param {[type]} location string [description]
 */
func NewAggregateImage(location string) *AggregateImage {
	fmt.Println("new aggr image")
	imgAgg := &AggregateImage{sectionMap: map[string]*aggregateImageSection{}, fontsSectionMap: map[string]*FontImageSection{}}
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

		draw.Draw(rgbaFinal, imgSec.section, imgSec, image.Point{0, 0}, draw.Src) // draw first image
		lastLocY += imgSec.Bounds().Dy()

	}

	// store the final aggregate image
	imgAgg.aggregateImage = rgbaFinal

	return imgAgg
}

func (this *AggregateImage) AppendImage(img image.Image, imgTag string) {
	fmt.Println("append img")
	//TODO : clean this up .. too much repetition of variables
	newDim := image.Rectangle{image.Point{0, 0},
		this.aggregateImage.Bounds().Size().Add(img.Bounds().Size())}

	rgbaFinal := image.NewRGBA(newDim)

	//TODO need to check this...
	draw.Draw(rgbaFinal,
		image.Rectangle{image.Point{0, 0}, this.aggregateImage.Bounds().Size()}, this.aggregateImage, image.Point{0, 0}, draw.Src) // draw first image
	draw.Draw(rgbaFinal,
		image.Rectangle{image.Point{0, this.aggregateImage.Bounds().Dy()}, this.aggregateImage.Bounds().Size().Add(img.Bounds().Size())}, img, image.Point{0, 0}, draw.Src) // draw first image

	sec := image.Rectangle{image.Point{0, this.aggregateImage.Bounds().Dy()}, image.Point{img.Bounds().Dx(), this.aggregateImage.Bounds().Dy() + img.Bounds().Dy()}}
	this.sectionMap[imgTag] = &aggregateImageSection{img, "", sec}
	this.images = append(this.images, this.sectionMap[imgTag])
	this.aggregateImage = rgbaFinal
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

	sec := &aggregateImageSection{img, imgPath, image.Rectangle{}}

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

func (this *AggregateImage) GetImageSection(path string) *aggregateImageSection {
	return this.sectionMap[path]
}

func (this *AggregateImage) GetFontImageSection(font string) *FontImageSection {
	return this.fontsSectionMap[font]
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
