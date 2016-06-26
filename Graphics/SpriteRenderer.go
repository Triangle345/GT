// Package Graphics
package Graphics

import (
	// "github.com/go-gl/gl/v3.2-core/gl"

	"GT/Graphics/Components"
	"GT/Graphics/Image"
	"GT/Graphics/Opengl"
	"fmt"
	"image"

	mathgl "github.com/go-gl/mathgl/mgl32"
)

func NewSpriteRenderer() *SpriteRenderer {

	sprite := SpriteRenderer{a: 1}
	return &sprite

}

type sheetImage struct {
	img image.Image
	uvs []float32
}

// SpriteRenderer is a component which allows a sprite to be drawn or animated
type SpriteRenderer struct {
	// the parent node
	//Parent *Components.Node
	Components.ChildComponent

	// the image
	img image.Image

	// list of images representing our spliced sprite sheet (animation)
	sheetImages      []*sheetImage
	indexInAnimation int

	// color
	r, g, b, a float32

	// uvs - store uvs for speed
	uvs []float32
}

// func (this *SpriteRenderer) SetImageSpriteSheet(imageLoc string) {
// 	//TODO: should probably cache this for speed
// 	img, err := NewSpriteSheetImage(imageLoc, NewRectangularArea(0, 0, 128, 128))
//
// 	if err != nil {
// 		fmt.Println("Cannot create image: " + err.Error())
// 	}
//
// 	this.img = img
// }
func (this *SpriteRenderer) SetImage(imageLoc string) {
	img, err := Image.NewImage(imageLoc)

	if err != nil {
		fmt.Println("Cannot create image: " + err.Error())
	}
	this.uvs = img.UVs()
	this.img = img
}

// SpliceAndSetFullSheet manually cuts up an entire sprite sheet based on user defined frame dimensions
func (s *SpriteRenderer) SpliceAndSetFullSheet(imageLoc string, frameHeight int, frameWidth int) {
	// to set an entire sheet as an animation, set the Frames and Row Numbers to 0
	s.SpliceAndSetAnimation(imageLoc, frameHeight, frameWidth, 0, 0)
}

// SpliceAndSetAnimation manually cuts up a row of a sprite sheet based on user defined dimensions and sets it as the current animation
func (s *SpriteRenderer) SpliceAndSetAnimation(imageLoc string, frameHeight int, frameWidth int, noOfFrames int, rowNum int) {

	img, err := Image.NewImage(imageLoc)
	if err != nil {
		fmt.Println("Cannot create image: " + err.Error())
	}

	// throw warnings for bad input
	numOfRows := float32(img.Bounds().Dy() / frameHeight)
	numOfColumns := float32(img.Bounds().Dx() / frameWidth)
	if float32(noOfFrames) > numOfColumns || numOfColumns < 1 {
		fmt.Println("WARNING: frames out of bounds")
	}
	if float32(rowNum) > numOfRows || numOfRows < 1 {
		fmt.Println("WARNING: row desired out of bounds")
	}

	for j := 0; j < img.Bounds().Dy(); j += frameHeight {

		// only use our desired row (if specified)
		if rowNum != 0 && j/frameHeight != rowNum-1 {
			continue
		}

		// splice the row by the amount of intended images
		for i := 0; i < img.Bounds().Dx(); i += frameWidth {

			// only grab our desired number of frames (if specified)
			if noOfFrames != 0 && i/frameWidth >= noOfFrames {
				continue
			}

			// splice image from row, and insert piece into array
			b := image.Rect(i, j, i+frameWidth, j+frameHeight)
			spriteSheetPart, err := img.SubImage(b)
			if err != nil {
				fmt.Println("Cannot create sub image: " + err.Error())
			}

			s.sheetImages = append(s.sheetImages, &sheetImage{spriteSheetPart, spriteSheetPart.UVs()})
		}
	}

	// set the current image to the first in the new animation
	s.indexInAnimation = 0
	s.uvs = s.sheetImages[0].uvs
	s.img = s.sheetImages[0].img
}

// SetSubImage
// Sets a designated part of an image for this sprite renderer
//  @param  {[string]} this *SpriteRenderer [the base image path]
//  @param  {[image.Rectangle]} this *SpriteRenderer [the rectangular bounds of designated part of image]
//  @return
func (this *SpriteRenderer) SetSubImage(imageLoc string, bounds image.Rectangle) {
	img, err := Image.NewImage(imageLoc)

	if err != nil {
		fmt.Println("Cannot create image: " + err.Error())
	}

	img, err = img.SubImage(bounds)
	if err != nil {
		fmt.Println("Cannot create sub image: " + err.Error())
	}

	// this.uvs = Image.GetUVs(img.Bounds())
	this.uvs = img.UVs()
	this.img = img
}

func (s *SpriteRenderer) Initialize() {

}

// Update gets called every frame and accounts for all settings in the renderer as well as shifts animations
func (s *SpriteRenderer) Update(delta float32) {

	if s.img == nil {
		return
	}

	// this gets teh bounds of the sub image
	w := float32(s.img.Bounds().Dx())
	h := float32(s.img.Bounds().Dy())

	vertex_data := []float32{-0.5 * w, 0.5 * h, 1.0, 0.5 * w, 0.5 * h, 1.0, 0.5 * w, -0.5 * h, 1.0, -0.5 * w, -0.5 * h, 1.0}

	elements := []uint32{uint32(0), uint32(1), uint32(2), uint32(0), uint32(2), uint32(3)}

	Model := mathgl.Ident4()

	// check to see if its a regular node to get the updated model
	// if n, ok := s.Parent.(*Components.Node); ok {

	Model = s.GetParent().Transform().GetUpdatedModel()
	// }

	// transform all vertex data and combine it with other data
	var data []float32 = make([]float32, 0, 9*4)
	for j := 0; j < 4; j++ {
		transformation := mathgl.Vec4{vertex_data[j*3+0], vertex_data[j*3+1], vertex_data[j*3+2], 1}
		t := Model.Mul4x1(transformation)

		data = append(data, t[0], t[1], t[2], s.r, s.g, s.b, s.a, s.uvs[j*2+0], s.uvs[j*2+1])

	}

	// package everything up in an OpenGLVertexInfo
	vertexInfo := Opengl.OpenGLVertexInfo{
		VertexData: data,
		Elements:   elements,
		Stride:     4,
	}

	// send OpenGLVertex info to Opengl module
	Opengl.AddVertexData(1, &vertexInfo)

	// iterate images if we have a list (animation)
	if len(s.sheetImages) > 0 {
		if s.indexInAnimation == len(s.sheetImages)-1 {
			s.indexInAnimation = 0
		} else {
			s.indexInAnimation++
		}
		s.img = s.sheetImages[s.indexInAnimation].img
		s.uvs = s.sheetImages[s.indexInAnimation].uvs
	}
}

// SetColor allows us to modify image coloring of whatever is set in the Renderer
func (s *SpriteRenderer) SetColor(r, g, b, a float32) {
	s.r = r
	s.g = g
	s.b = b
	s.a = a
}
