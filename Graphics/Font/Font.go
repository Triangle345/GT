package Font

import (
	"GT/Logging"
	"fmt"
	"image"
	"image/draw"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var fonts []*FontInfo = []*FontInfo{}
var fontsMap map[string]*FontInfo = map[string]*FontInfo{}

type FontInfo truetype.Font

func GetFonts() []*FontInfo {
	return fonts
}

func GetFont(font string) *FontInfo {
	return fontsMap[font]
}

var (
	dpi float64 = 72
	//dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	// fontfile string = "../Graphics/Font/Times_New_Roman_Normal.ttf"
	//fontfile = flag.String("fontfile", "./Times_New_Roman_Normal.ttf", "filename of the ttf font")
	hinting string = "full"
	//hinting  = flag.String("hinting", "none", "none | full")
	size      float64       = 100
	sizeFixed fixed.Int26_6 = 100
	//size     = flag.Float64("size", 100, "font size in points")
	spacing float64 = 1.5
	//spacing  = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
	wonb bool = false
	//wonb     = flag.Bool("whiteonblack", false, "white text on a black background")

	// items per row for font spread
	itemsPerRow int = 20

	insetPadding int = 30

	text string = " ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`0123456789-=~!@#$%^&*()_+[]\\{}|;':\",./<>?"
)

// returns official font name: eg. "Times New Roman Normal""
func (this FontInfo) Name() string {
	f := truetype.Font(this)
	return f.Name(truetype.NameIDFontFullName)
}

//GetSectionMap
//Returns section mapping for each letter in ttf in form of image.Rect with y=0
func (this FontInfo) GetSectionMap() map[rune]image.Rectangle {
	f := truetype.Font(this)

	secMap := map[rune]image.Rectangle{}

	curW := 0
	curH := 0

	// fmt.Println("Font: ", this.Name())

	for idx, v := range text {
		if idx > 0 && int(math.Mod(float64(idx), float64(itemsPerRow))) == 0 {
			curH += int(f.Bounds(sizeFixed).Max.Sub(f.Bounds(sizeFixed).Min).Y) + insetPadding
			curW = 0
		}
		i := f.Index(v)
		hmet := f.HMetric(sizeFixed, i)
		// vmet := f.VMetric(sizeFixed, i)

		// advance width is how much font advances on x which next font starts AFTER that.
		minX := curW //+ 1
		minY := curH //+ 1
		maxX := curW + int(hmet.AdvanceWidth)
		maxY := curH + int(f.Bounds(sizeFixed).Max.Sub(f.Bounds(sizeFixed).Min).Y)

		curW += int(hmet.AdvanceWidth) + insetPadding

		secMap[v] = image.Rect(minX, minY, maxX, maxY)

		// fmt.Println("Letter: ", string(v))
		// fmt.Println("Bounds: ", secMap[v])
	}

	return secMap

}

// returns image of rendered font
func (this FontInfo) GetImage() image.Image {

	f := truetype.Font(this)

	maxWidth := 0
	curW := 0
	maxHeight := int(f.Bounds(sizeFixed).Max.Sub(f.Bounds(sizeFixed).Min).Y) + insetPadding
	for idx, val := range text {
		if idx > 0 && int(math.Mod(float64(idx), float64(itemsPerRow))) == 0 {
			maxHeight += int(f.Bounds(sizeFixed).Max.Sub(f.Bounds(sizeFixed).Min).Y) + insetPadding
			curW = 0

		}

		i := f.Index(val)
		hmet := f.HMetric(sizeFixed, i)

		curW += int(hmet.AdvanceWidth) + insetPadding

		if curW > maxWidth {
			maxWidth = curW
		}

	}

	fmt.Print("Font ", this.Name())
	fmt.Println("Max width: ", maxWidth)
	fmt.Println("Max height: ", maxHeight)

	// Draw the background and the guidelines.
	fg, bg := image.Black, image.Transparent

	// diffY := int(f.Bounds(sizeFixed).Max.Sub(f.Bounds(sizeFixed).Min).Y)

	// add a constant to the y term because of subtle bleeding between vertical space, probably because of advance height
	rgba := image.NewRGBA(image.Rect(0, 0, maxWidth, maxHeight)) //diffY+5))

	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)

	// Draw the text.
	h := font.HintingNone
	switch hinting {
	case "full":
		h = font.HintingFull
	}
	d := &font.Drawer{
		Dst: rgba,
		Src: fg,
		Face: truetype.NewFace(&f, &truetype.Options{
			Size:    size,
			DPI:     dpi,
			Hinting: h,
		}),
	}

	// y := int(math.Abs(float64(f.Bounds(sizeFixed).Max.Y))) //10 + int(math.Ceil(*size**dpi/72))
	// dy := int(math.Ceil(size * spacing * dpi / 72))

	tWidth := 0
	tHeight := int(math.Abs(float64(f.Bounds(sizeFixed).Max.Y)))
	for idx, r := range text {
		if idx > 0 && int(math.Mod(float64(idx), float64(itemsPerRow))) == 0 {
			tHeight += int(f.Bounds(sizeFixed).Max.Sub(f.Bounds(sizeFixed).Min).Y) + insetPadding
			tWidth = 0
		}
		i := f.Index(r)
		hmet := f.HMetric(sizeFixed, i)

		d.Dot = fixed.P(tWidth, tHeight)
		tWidth += int(hmet.AdvanceWidth) + insetPadding

		d.DrawString(string(r))

	}

	return rgba
}

func loadFont(fontFile string) FontInfo {

	// Read the font data.
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		log.Println(err)
		return FontInfo{}
	}
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Println(err)
		return FontInfo{}
	}

	return FontInfo(*f)
}

func fileVisitor(path string, f os.FileInfo, err error) error {
	if strings.Contains(path, ".ttf") {
		fmt.Printf("Processing font: %s\n", path)
		fInfo := loadFont(path)
		fonts = append(fonts, &fInfo)

		Logging.Info("Found Font: " + fInfo.Name())
		fontsMap[fInfo.Name()] = &fInfo
	}

	return nil
}

/**
 * ReadFonts
 * Read fonts from directory
 * @param {[string]} path string [the directory where fonts are located]
 */
func LoadFonts(path string) {

	filepath.Walk(path, fileVisitor)
}
