package Image

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"math"
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

	Font.LoadFonts("../Assets/Fonts")

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

//var images []image.Image = []image.Image{}

/** NewAggregateImage
 *
 * Walk the directory and aggregate all images to one image and store in meemory
 * @param {[type]} location string [description]
 */
func LoadImages(location string) {
	fmt.Println("new aggr image")
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

	//rectangle for the big image
	texSize := int(Opengl.Probe().MaxTextureSize)
	finalImg := image.Rectangle{image.Point{0, 0}, image.Point{texSize, texSize}}

	rgbaFinal := image.NewRGBA(finalImg)

	lastLocY := 0

	for _, imgSec := range AggrImg.images {

		draw.Draw(rgbaFinal, imgSec.section, imgSec, image.Point{0, 0}, draw.Src) // draw images
		lastLocY += imgSec.Bounds().Dy()

	}

	fonts := Font.GetFonts()

	for _, f := range fonts {

		fontImg := f.GetImage()

		fontSec := &FontImageSection{&aggregateImageSection{fontImg, f.Name(), fontImg.Bounds()}, f.GetSectionMap()}
		AggrImg.fonts = append(AggrImg.fonts, fontSec)
		AggrImg.fontsSectionMap[f.Name()] = fontSec

		fontSec = AggrImg.fontsSectionMap[f.Name()]

		fontSec.section = image.Rect(0, lastLocY, fontSec.Bounds().Dx(), lastLocY+fontSec.Bounds().Dy())
		draw.Draw(rgbaFinal, fontSec.section, fontSec, image.Point{0, 0}, draw.Src) // draw fonts image
		lastLocY += fontSec.Bounds().Dy()
	}

	// store the final aggregate image
	AggrImg.aggregateImage = rgbaFinal
}

// AppendImage adds, and keeps track of, a new image within the aggregate
func (this *AggregateImage) AppendImage(img image.Image, imgTag string) {
	fmt.Println("append img", imgTag)

	width := math.Max(float64(this.aggregateImage.Bounds().Dx()),
		float64(img.Bounds().Dx()))

	//TODO : add binary square image
	// newDim := image.Rect(0, 0,
	// 	int(width),
	// 	this.aggregateImage.Bounds().Max.Y+img.Bounds().Dy())
	texSize := int(Opengl.Probe().MaxTextureSize)
	newDim := image.Rect(0, 0, texSize, texSize)

	rgbaFinal := image.NewRGBA(newDim)

	//TODO need to check this...
	draw.Draw(rgbaFinal,
		image.Rectangle{image.Point{0, 0}, this.aggregateImage.Bounds().Size()}, this.aggregateImage, image.Point{0, 0}, draw.Src) // draw first image
	draw.Draw(rgbaFinal,
		image.Rect(0, this.aggregateImage.Bounds().Dy(),
			int(width), this.aggregateImage.Bounds().Max.Y+img.Bounds().Dy()),
		img, image.Point{0, 0}, draw.Src) // draw second image

	sec := image.Rectangle{
		image.Point{0, this.aggregateImage.Bounds().Dy()},
		image.Point{img.Bounds().Dx(), this.aggregateImage.Bounds().Dy() + img.Bounds().Dy()}}

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
