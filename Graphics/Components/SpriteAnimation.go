// Package Sprite ...
package Components

import (
	"GT/Graphics/Image"
	"GT/Graphics/Opengl"
	"fmt"
	"image"
)

// SpriteAnimation is a sequence of images and settings used by the renderer to animate sprites
type FrameAnimation struct {
	Animation

	// list of images representing our spliced sprite sheet (animation)
	animationImages []Opengl.RenderObject
	meta            *AnimationHelper
}

// NewSpriteAnimation creates a renderer and initializes its animation map
func NewFrameAnimation() *FrameAnimation {

	// preset our defaults
	animation := FrameAnimation{}
	animation.meta = NewAnimation()

	return &animation
}

// AppendAnimation adds a list of images to the end of our animation
func (s *FrameAnimation) AppendAnimation(animIn *FrameAnimation) {
	for i := 0; i < len(animIn.animationImages); i++ {
		imgToAdd := animIn.animationImages[i]
		s.animationImages = append(s.animationImages, imgToAdd)
	}
}

// Append simply adds an image from the aggregate to the end of our animation
func (s *FrameAnimation) Append(imageLoc string) {
	img, err := Image.NewImage(imageLoc)
	if err != nil {
		fmt.Println("Cannot create image: " + err.Error())
	}
	s.animationImages = append(s.animationImages, &img)
}

func (s *FrameAnimation) Remove(imageIdx int) {
	// convert to zero based
	imageIdx--

	// shift the image to the end, then pop it
	if imageIdx >= 0 && imageIdx < len(s.animationImages) {
		s.animationImages = append(s.animationImages[:imageIdx], s.animationImages[imageIdx+1:]...)
		s.animationImages = s.animationImages[:len(s.animationImages)-1]
	} else {
		fmt.Printf("cant remove from animation, out of bounds...")
	}
}

// Reorder simply swaps the order in which any two images are rendered ()
func (s *FrameAnimation) Reorder(imageOneIdx int, imageTwoIdx int) {

	// convert from one-based input to zero-based logic
	listLength := len(s.animationImages)
	lowerIdx := imageOneIdx - 1
	higherIdx := imageTwoIdx - 1

	// keep track of the higher and lower indexes
	if imageOneIdx > imageTwoIdx {
		lowerIdx = imageTwoIdx
		higherIdx = imageOneIdx
	}

	if imageOneIdx == imageTwoIdx || higherIdx > listLength || lowerIdx < 0 {
		fmt.Printf("invalid input, please ensure you are entering two unique dimensions within the list")
		return
	}

	imgTmp := s.animationImages[lowerIdx]

	s.animationImages[lowerIdx] = s.animationImages[higherIdx]
	s.animationImages[higherIdx] = imgTmp
}

// SetFrequency sets our animation's timing in either seconds or frames per toggle
// i.e. Frequency(4,true) sets to update every 4 frames
// i.e. Frequency(0.25,false) sets to update every 1/4 of a second
func (s *FrameAnimation) SetFrequency(freqIn float64, setFrequencyByTheFrame bool) {
	s.meta.Frequency = freqIn
	s.meta.FrequencyIsInFrames = setFrequencyByTheFrame
}

func (s *FrameAnimation) SetAsOneTimeOnly(setOneTime bool) {
	s.meta.OneTimeOnly = setOneTime
}

// SpliceAndSetFullSheetAnimation manually cuts up an entire sprite sheet based on user defined frame dimensions
func (s *FrameAnimation) SpliceAndSetFullSheetAnimation(imageLoc string, frameWidth int, frameHeight int) {

	// to set an entire sheet as an animation, set the Frames and Row Numbers to 0
	s.SpliceAndSetAnimation(imageLoc, frameWidth, frameHeight, 0, 0)
}

// SpliceAndSetAnimation manually cuts up a row of a sprite sheet based on user defined dimensions and sets it as the current animation
func (s *FrameAnimation) SpliceAndSetAnimation(imageLoc string, frameWidth int, frameHeight int, noOfFrames int, rowNum int) {

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

			s.animationImages = append(s.animationImages, &spriteSheetPart)
		}
	}

	// set the current image to the first in the new animation
	s.meta.IndexInAnimation = 0
}

// currentImage returns the animation image associated with the current index in the animation
func (s *FrameAnimation) currentImage() Opengl.RenderObject {
	// TODO: possibly make this return a blank image when we shouldn't animate?
	return s.animationImages[s.meta.IndexInAnimation]
}

func (s *FrameAnimation) update() bool {
	return s.meta.Update(len(s.animationImages))
}
