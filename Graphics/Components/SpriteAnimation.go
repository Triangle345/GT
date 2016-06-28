// Package Components ...
package Components

import (
	"GT/Graphics/Image"
	"fmt"
	"image"
)

// AnimationImage contains relevant gl info for use with renderer via CurrentImage()
type AnimationImage struct {
	img image.Image
	uvs []float32
}

// SpriteAnimation is a sequence of images and settings used by the renderer to animate sprites
type SpriteAnimation struct {

	// list of images representing our spliced sprite sheet (animation)
	animationImages []*AnimationImage

	indexInAnimation      int
	frequency             int
	framesSinceLastToggle int
}

// NewSpriteAnimation creates a renderer and initializes its animation map
func NewSpriteAnimation() *SpriteAnimation {

	animation := SpriteAnimation{}
	animation.indexInAnimation = 0
	return &animation
}

// CurrentImage returns the animation image associated with the current index in the animation
func (s *SpriteAnimation) CurrentImage() *AnimationImage {
	return s.animationImages[s.indexInAnimation]
}

// TODO: possibly add another option for freqency in seconds
// Frequency sets our animation's toggle frequency (in frames per toggle)
func (s *SpriteAnimation) Frequency(freqIn int) {
	s.frequency = freqIn
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

			s.animationImages = append(s.animationImages, &AnimationImage{spriteSheetPart, spriteSheetPart.UVs()})
		}
	}

	// set the current image to the first in the new animation
	s.indexInAnimation = 0
}

// Update internally evaluates and increments toggle logic then returns true if we did swap images
func (s *SpriteAnimation) Update() bool {

	// verify we have stuff to animate, then check if we are ready to toggle
	if len(s.animationImages) > 0 {
		if s.framesSinceLastToggle/s.frequency == 1 {

			// allow our animation to continue by resetting it
			// TODO: possibly allow for a one-time only vs. recurring animation
			if s.indexInAnimation == len(s.animationImages)-1 {
				s.indexInAnimation = 0
			} else {
				s.indexInAnimation++
			}
			s.framesSinceLastToggle = 0

			// indicate we have updated our current image
			return true
		}
		s.framesSinceLastToggle++
	}
	return false
}
