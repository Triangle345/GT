// Package Components ...
package Components

import (
	"GT/Graphics/Image"
	"GT/Graphics/Opengl"
	"fmt"
	"image"
)

// NewSpriteRenderer creates a renderer and initializes its animation map
func NewSpriteRenderer() *SpriteRenderer {
	sprite := SpriteRenderer{}
	// sprite.animationsMap = map[string]*FrameAnimation{}
	sprite.AnimationHandler = &AnimationHandler{animationsMap: map[string]*FrameAnimation{}}

	return &sprite
}

// SpriteRenderer is a component which allows a sprite to be drawn or animated
type SpriteRenderer struct {
	// ChildComponent
	Renderer

	AnimationHandler *AnimationHandler

	// map of animations representing possible visuals for a sprite
	// animationsMap    map[string]*FrameAnimation
	// currentAnimation *FrameAnimation

	// singular image
	img Opengl.RenderObject
}

// SpliceAndSetFullSheetAnimation manually cuts up an entire sprite sheet based on user defined frame dimensions
func (s *SpriteRenderer) SpliceAndSetFullSheetAnimation(imageLoc string, frameWidth int, frameHeight int) *FrameAnimation {

	// to set an entire sheet as an animation, set the Frames and Row Numbers to 0
	fa := s.SpliceAndSetAnimation(imageLoc, frameWidth, frameHeight, 0, 0)

	return fa
}

// SpliceAndSetAnimation manually cuts up a row of a sprite sheet based on user defined dimensions and sets it as the current animation
func (s *SpriteRenderer) SpliceAndSetAnimation(imageLoc string,
	frameWidth, frameHeight, noOfFrames, rowNum int) *FrameAnimation {

	fa := newFrameAnimation()

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
			fa.animationImages = append(fa.animationImages, &spriteSheetPart)
		}
	}

	// set the current image to the first in the new animation
	fa.meta.IndexInAnimation = 0

	return fa
}

// SetImage puts a designated image from the agregate into our image which will be rendered
func (s *SpriteRenderer) SetImage(imageLoc string) {
	img, err := Image.NewImage(imageLoc)

	if err != nil {
		fmt.Println("Cannot create image: " + err.Error())
	}

	s.img = &img

}

// SetSubImage sets a designated part of an image for this sprite renderer
//  @param  {[string]} this *SpriteRenderer [the base image path]
//  @param  {[image.Rectangle]} this *SpriteRenderer [the rectangular bounds of designated part of image]
//  @return
func (s *SpriteRenderer) SetSubImage(imageLoc string, bounds image.Rectangle) {
	img, err := Image.NewImage(imageLoc)

	if err != nil {
		fmt.Println("Cannot create image: " + err.Error())
	}

	img, err = img.SubImage(bounds)
	if err != nil {
		fmt.Println("Cannot create sub image: " + err.Error())
	}

	// this.uvs = Image.GetUVs(img.Bounds())
	s.img = &img

}

// Update gets called every frame and accounts for all settings in the renderer as well as shifts animations
func (s *SpriteRenderer) Update(delta float32) {

	// fmt.Println("Anim update: ", s.AnimationHandler.currentAnimation.update())
	// run the animation update (if applicable) and set our renderer image if the animation toggled
	if s.AnimationHandler.currentAnimation != nil {
		// just care about updating the animation handler, thats it. We
		// want current image always if available
		s.AnimationHandler.currentAnimation.update()
		s.img = s.AnimationHandler.CurrentAnimation().currentImage()
		// fmt.Println("New image: ", s.img.VertexData())

	}
	// s.SetImage(GT.AssetsImages + "smiley.png")
	if s.img == nil {
		return
	}

	// fmt.Println("CURRENT animation: ", s.img.VertexData())
	s.Render(s.img)

}

// SetColor allows us to modify image coloring of whatever is set in the Renderer
func (s *SpriteRenderer) SetColor(r, g, b, a float32) {
	s.Color = &Image.Color{r, g, b, a}
}
