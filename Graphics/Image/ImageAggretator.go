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

var images []*aggregateImageSection
var sectionMap = map[string]*aggregateImageSection{}

var fonts []*FontImageSection
var fontsSectionMap = map[string]*FontImageSection{}

//TODO: create some way of doing this automatically maybe based on some config file
// var AggrImg *AggregateImage
var aggregateImages []*AggregateImage

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
	aggrId   int
}

// AggregateImage is our class for grouping all images, sprites, etc. into one place.
// This will improve performance over loading an image every time it's needed.
// Instead we only need to select a space of our pre-rendered composite image
type AggregateImage struct {
	imageBuddy *partition

	aggregateImage image.Image

	id int
}

func GetUVFromPosition(point image.Point) (u, v float32) {
	//TODO make clean and safe
	bounds := aggregateImages[0].aggregateImage.Bounds()
	u = float32(point.X) / float32(bounds.Dx())
	v = float32(point.Y) / float32(bounds.Dy())

	return
}
func newAggregateImage() *AggregateImage {
	newIdx := len(aggregateImages)

	aggrImg := &AggregateImage{}
	aggrImg.id = newIdx

	aggrImg.imageBuddy = NewBuddyAggregator(int(Opengl.Probe().MaxTextureSize))

	aggregateImages = append(aggregateImages, aggrImg)

	return aggrImg

}

// loadTExtureImages populates image section maps
func loadTextureImages(location string) {
	walker := pngWalker{}
	walker.fileWalker(location)

	maxHeight := 0
	maxWidth := 0

	aggrImg := newAggregateImage()

	for _, imgSec := range walker.images {

		part := aggrImg.imageBuddy.Insert(imgSec.pathName, imgSec.Bounds().Dx(), imgSec.Bounds().Dy())

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
			imgSec.aggrId = aggrImg.id
			images = append(images, imgSec)
			sectionMap[imgSec.pathName] = imgSec
		} else {
			// no more room in this partition, need to get new one
			aggrImg = newAggregateImage()
		}

	}

	fmt.Printf("Height is %d\n", maxHeight)
	fmt.Printf("Width is %d\n", maxWidth)

}

// loadFontImages populates font sections
func loadFontImages() {
	fontsInfo := Font.GetFonts()

	aggrImg := newAggregateImage()

	for _, f := range fontsInfo {

		fontImg := f.GetImage()

		// find free space for fonts
		part := aggrImg.imageBuddy.Insert(f.Name(), fontImg.Bounds().Dx(), fontImg.Bounds().Dy())

		// create section
		sec := image.Rectangle{part.Bounds().Min, part.Bounds().Min.Add(fontImg.Bounds().Size())}

		// create font section with section created from insertion
		fontSec := &FontImageSection{&aggregateImageSection{fontImg, f.Name(), sec, aggrImg.id}, f.GetSectionMap()}

		// add font section to lists
		fonts = append(fonts, fontSec)
		fontsSectionMap[f.Name()] = fontSec

	}
}

/** NewAggregateImage
 *
 * Walk the directory and aggregate all images to one image and store in meemory
 * @param {[type]} location string [description]
 */
func LoadImages(path string) {

	loadTextureImages(path)
	loadFontImages()

	// create empty image to hold all images
	texSize := int(Opengl.Probe().MaxTextureSize)
	for idx, aggrImg := range aggregateImages {

		finalImg := image.Rectangle{image.Point{0, 0}, image.Point{texSize, texSize}}

		rgbaFinal := image.NewRGBA(finalImg)

		if idx < len(aggregateImages)-1 {

			// draw all images
			for _, imgSec := range images {

				draw.Draw(rgbaFinal, imgSec.Section, imgSec, image.Point{0, 0}, draw.Src) // draw images

			}
		} else {
			// draw fonts on only the last sheet as a brand new sheet is made for just them
			for _, fontSec := range fonts {

				draw.Draw(rgbaFinal, fontSec.Section, fontSec, image.Point{0, 0}, draw.Src) // draw images

			}
		}
		// store the final aggregate image
		aggrImg.aggregateImage = rgbaFinal
		// fmt.Println("Idx: " + strconv.Itoa(idx))
		// aggrImg.Print("aggrimg" + strconv.Itoa(idx) + ".png")
		Opengl.AddAggregateImage(aggrImg.aggregateImage)
	}
}

func (this *pngWalker) loadImage(imgPath string) error {
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

	sec := &aggregateImageSection{img, imgPath, image.Rectangle{}, -1}

	// populate both section map and image list with section
	this.images = append(this.images, sec)

	//sectionMap[imgPath] = sec

	return nil
}

type pngWalker struct {
	images []*aggregateImageSection
}

func (this *pngWalker) fileVisitor(path string, f os.FileInfo, err error) error {

	if strings.Contains(path, ".png") {
		fmt.Printf("Processing: %s\n", path)
		this.loadImage(path)
	}
	return nil
}

func (this *pngWalker) fileWalker(path string) {
	filepath.Walk(path, this.fileVisitor)
}

// GetImageSEction returns the image section based on named image
func GetImageSection(path string) *aggregateImageSection {

	return sectionMap[path]
}

// GetFontImageSection returns FontImageSection based on named font
func GetFontImageSection(font string) *FontImageSection {

	return fontsSectionMap[font]
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
