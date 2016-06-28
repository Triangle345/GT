// Package Components ...
package Components

import (
	"GT/Graphics/Image"
	"fmt"
	"image"
)

type animationImage struct {
	// the image
	img image.Image

	// opengl uvs to be used in renderer
	uvs []float32

	// color
	//r, g, b, a float32
}

// SpriteAnimation is a sequence of images and settings used by the renderer to animate sprites
type SpriteAnimation struct {

	// list of images representing our spliced sprite sheet (animation)
	AnimationImages       []*animationImage
	IndexInAnimation      int
	Frequency             int
	FramesSinceLastToggle int
}

// NewSpriteAnimation creates a renderer and initializes its animation map
func NewSpriteAnimation() *SpriteAnimation {

	animation := SpriteAnimation{}

	return &animation
}

// SpliceAndSetFullSheetAnimation manually cuts up an entire sprite sheet based on user defined frame dimensions
func (s *SpriteAnimation) SpliceAndSetFullSheetAnimation(imageLoc string, frameHeight int, frameWidth int) {
	// to set an entire sheet as an animation, set the Frames and Row Numbers to 0
	s.SpliceAndSetAnimation(imageLoc, frameHeight, frameWidth, 0, 0)
}

// SpliceAndSetAnimation manually cuts up a row of a sprite sheet based on user defined dimensions and sets it as the current animation
func (s *SpriteAnimation) SpliceAndSetAnimation(imageLoc string, frameHeight int, frameWidth int, noOfFrames int, rowNum int) {

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

			s.AnimationImages = append(s.AnimationImages, &animationImage{spriteSheetPart, spriteSheetPart.UVs()})
		}
	}

	// set the current image to the first in the new animation
	s.IndexInAnimation = 0
}
