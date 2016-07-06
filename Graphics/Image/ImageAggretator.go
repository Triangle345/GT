package Image

import (
	"GT/Graphics/Font"
	"GT/Graphics/Opengl"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

//TODO: create some way of doing this automatically maybe based on some config file
var AggrImg *AggregateImage

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
	Section  image.Rectangle
}

// AggregateImage is our class for grouping all images, sprites, etc. into one place.
// This will improve performance over loading an image every time it's needed.
// Instead we only need to select a space of our pre-rendered composite image
type AggregateImage struct {
	images     []*aggregateImageSection
	sectionMap map[string]*aggregateImageSection

	fonts           []*FontImageSection
	fontsSectionMap map[string]*FontImageSection

	imageBuddy *partition

	aggregateImage image.Image
}

func (this *AggregateImage) GetUVFromPosition(point image.Point) (u, v float32) {

	u = float32(point.X) / float32(this.aggregateImage.Bounds().Dx())
	v = float32(point.Y) / float32(this.aggregateImage.Bounds().Dy())

	return
}

// loadTExtureImages populates image section maps
func loadTextureImages(location string) {
	AggrImg.fileWalker(location)

	maxHeight := 0
	maxWidth := 0

	for _, imgSec := range AggrImg.images {

		part := AggrImg.imageBuddy.Insert(imgSec.pathName, imgSec.Bounds().Dx(), imgSec.Bounds().Dy())

		fmt.Println("Partition: ", part)

		if part != nil {
			imgSec.Section = image.Rectangle{part.Bounds().Min,
				part.Bounds().Min.Add(imgSec.Bounds().Size())}

			if imgSec.Section.Bounds().Max.Y > maxHeight {
				maxHeight = imgSec.Section.Bounds().Max.Y
			}
			if imgSec.Bounds().Max.X > maxWidth {
				maxWidth = imgSec.Section.Bounds().Max.X
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
		part := AggrImg.imageBuddy.Insert(f.Name(), fontImg.Bounds().Dx(), fontImg.Bounds().Dy())

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
func LoadImages(path string) {

	AggrImg = &AggregateImage{sectionMap: map[string]*aggregateImageSection{}, fontsSectionMap: map[string]*FontImageSection{}}

	// TODO need to figure out a way to setup opengl and window context to use video card probe
	AggrImg.imageBuddy = NewBuddyAggregator(int(Opengl.Probe().MaxTextureSize))

	loadTextureImages(path)
	loadFontImages()

	// create empty image to hold all images
	texSize := int(Opengl.Probe().MaxTextureSize)
	finalImg := image.Rectangle{image.Point{0, 0}, image.Point{texSize, texSize}}

	rgbaFinal := image.NewRGBA(finalImg)

	// draw all images
	for _, imgSec := range AggrImg.images {

		draw.Draw(rgbaFinal, imgSec.Section, imgSec, image.Point{0, 0}, draw.Src) // draw images

	}

	for _, fontSec := range AggrImg.fonts {

		draw.Draw(rgbaFinal, fontSec.Section, fontSec, image.Point{0, 0}, draw.Src) // draw images

	}

	// store the final aggregate image
	AggrImg.aggregateImage = rgbaFinal

	Opengl.AddAggregateImage(AggrImg.aggregateImage)
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

// GetImageSEction returns the image section based on named image
func (this *AggregateImage) GetImageSection(path string) *aggregateImageSection {

	return this.sectionMap[path]
}

// GetFontImageSection returns FontImageSection based on named font
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
