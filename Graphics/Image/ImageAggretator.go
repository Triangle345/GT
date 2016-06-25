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

var imageBuddy *partition

//TODO maybe not the best place for this? need to init somewhere else?
func Start() {
	AggrImg = &AggregateImage{sectionMap: map[string]*aggregateImageSection{}, fontsSectionMap: map[string]*FontImageSection{}}

	// TODO need to figure out a way to setup opengl and window context to use video card probe
	imageBuddy = NewBuddyAggregator(int(Opengl.Probe().MaxTextureSize))

	// use absolute pathing so that we can reference / map the images easily from anywhere
	imgsPath, err := filepath.Abs("../Assets/Images")
	if err != nil {
		panic(err)
	}

	// load all images from the beginning only. Never load images again after this
	LoadImages(imgsPath)

	Opengl.SetAggregateImage(AggrImg.aggregateImage)

	// for debug
	// AggrImg.Print("./aggr.png")
}

// FontImageSection keeps track of where our fonts are within the aggregate image
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

// AggregateImage is our class for grouping all images, sprites, etc. into one place.
// This will improve performance over loading an image every time it's needed.
// Instead we only need to select a space of our pre-rendered composite image
type AggregateImage struct {
	images     []*aggregateImageSection
	sectionMap map[string]*aggregateImageSection

	fonts           []*FontImageSection
	fontsSectionMap map[string]*FontImageSection

	aggregateImage image.Image
}

// loadTExtureImages populates image section maps
func loadTextureImages(location string) {
	AggrImg.fileWalker(location)

	maxHeight := 0
	maxWidth := 0

	for _, imgSec := range AggrImg.images {

		part := imageBuddy.Insert(imgSec.pathName, imgSec.Bounds().Dx(), imgSec.Bounds().Dy())

		fmt.Println("Partition: ", part)

		if part != nil {
			imgSec.section = image.Rectangle{part.Bounds().Min,
				part.Bounds().Min.Add(imgSec.Bounds().Size())}

			if imgSec.section.Bounds().Max.Y > maxHeight {
				maxHeight = imgSec.section.Bounds().Max.Y
			}
			if imgSec.Bounds().Max.X > maxWidth {
				maxWidth = imgSec.section.Bounds().Max.X
			}
		}

	}

	fmt.Printf("Height is %d\n", maxHeight)
	fmt.Printf("Width is %d\n", maxWidth)

}

// loadFontImages populates font sections
func loadFontImages() {
	fonts := Font.GetFonts()

	for _, f := range fonts {

		fontImg := f.GetImage()

		// find free space for fonts
		part := imageBuddy.Insert(f.Name(), fontImg.Bounds().Dx(), fontImg.Bounds().Dy())

		// create section
		sec := image.Rectangle{part.Bounds().Min, part.Bounds().Min.Add(fontImg.Bounds().Size())}

		// create font section with section created from insertion
		fontSec := &FontImageSection{&aggregateImageSection{fontImg, f.Name(), sec}, f.GetSectionMap()}

		// add font section to lists
		AggrImg.fonts = append(AggrImg.fonts, fontSec)
		AggrImg.fontsSectionMap[f.Name()] = fontSec

	}
}

/** NewAggregateImage
 *
 * Walk the directory and aggregate all images to one image and store in meemory
 * @param {[type]} location string [description]
 */
func LoadImages(location string) {
	fmt.Println("new aggr image")

	loadTextureImages(location)
	loadFontImages()

	// create empty image to hold all images
	texSize := int(Opengl.Probe().MaxTextureSize)
	finalImg := image.Rectangle{image.Point{0, 0}, image.Point{texSize, texSize}}

	rgbaFinal := image.NewRGBA(finalImg)

	// draw all images
	for _, imgSec := range AggrImg.images {

		draw.Draw(rgbaFinal, imgSec.section, imgSec, image.Point{0, 0}, draw.Src) // draw images

	}

	for _, fontSec := range AggrImg.fonts {

		draw.Draw(rgbaFinal, fontSec.section, fontSec, image.Point{0, 0}, draw.Src) // draw images

	}

	// store the final aggregate image
	AggrImg.aggregateImage = rgbaFinal
}

func (this *AggregateImage) loadImage(imgPath string) error {
	imgFile, err := os.Open(imgPath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	img, _, err := image.Decode(imgFile)

	if err != nil {
		fmt.Println("Cannot decode image: ", err)
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
