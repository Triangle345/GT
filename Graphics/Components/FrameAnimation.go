// Package Components ...
package Components

import (
	"GT/Graphics/Image"
	"GT/Graphics/Opengl"
	"GT/Logging"
	"time"
)

// if keeping track of nanoseconds for updates, use this to convert to seconds
const nanoToSeconds = 1000000000

type metaInfo struct {
	// animation properties and tracking
	indexInAnimation int

	framesSinceLastToggle int
	timeOfLastToggle      float64
	frequency             float64
	frequencyIsInFrames   bool

	oneTimeOnly   bool
	shouldAnimate bool
}

func newMeta() *metaInfo {
	// preset our defaults
	a := metaInfo{}
	a.indexInAnimation = 0

	a.frequency = 1
	a.frequencyIsInFrames = true
	a.framesSinceLastToggle = 0
	a.timeOfLastToggle = float64(time.Now().UnixNano()) / nanoToSeconds

	a.oneTimeOnly = false
	a.shouldAnimate = true

	return &a
}

// FrameAnimation is a sequence of images and settings used by the renderer to animate sprites
type FrameAnimation struct {
	// list of render objects representing our animation
	animationImages []Opengl.RenderObject
	meta            *metaInfo
}

func newFrameAnimation() *FrameAnimation {

	// preset our defaults
	animation := FrameAnimation{}
	animation.meta = newMeta()

	return &animation
}

// AppendAnimation adds a list of images to the end of our animation
func (f *FrameAnimation) AppendAnimation(animIn *FrameAnimation) {
	for i := 0; i < len(animIn.animationImages); i++ {
		imgToAdd := animIn.animationImages[i]
		f.animationImages = append(f.animationImages, imgToAdd)
	}
}

// Append simply adds an image from the aggregate to the end of our animation
func (f *FrameAnimation) Append(imageLoc string) {
	img, err := Image.NewImage(imageLoc)
	if err != nil {
		Logging.Debug("Cannot create image: " + err.Error())
	}
	f.animationImages = append(f.animationImages, &img)
}

// Remove an item from the animation via 1-based input
func (f *FrameAnimation) Remove(imageIdx int) {
	// convert to zero based
	imageIdx--

	// shift the image to the end, then pop it
	if imageIdx >= 0 && imageIdx < len(f.animationImages) {
		f.animationImages = append(f.animationImages[:imageIdx], f.animationImages[imageIdx+1:]...)
		f.animationImages = f.animationImages[:len(f.animationImages)-1]
	} else {
		Logging.Debug("cant remove from animation, out of bounds...")
	}
}

// Reorder simply swaps the order in which any two images are rendered ()
func (f *FrameAnimation) Reorder(imageOneIdx int, imageTwoIdx int) {

	// convert from one-based input to zero-based logic
	listLength := len(f.animationImages)
	lowerIdx := imageOneIdx - 1
	higherIdx := imageTwoIdx - 1

	// keep track of the higher and lower indexes
	if imageOneIdx > imageTwoIdx {
		lowerIdx = imageTwoIdx
		higherIdx = imageOneIdx
	}

	if imageOneIdx == imageTwoIdx || higherIdx > listLength || lowerIdx < 0 {
		Logging.Debug("invalid input, please ensure you are entering two unique dimensions within the list")
		return
	}

	imgTmp := f.animationImages[lowerIdx]

	f.animationImages[lowerIdx] = f.animationImages[higherIdx]
	f.animationImages[higherIdx] = imgTmp
}

// SetFrequency sets our animation's timing in either seconds or frames per toggle
// i.e. Frequency(4,true) sets to update every 4 frames
// i.e. Frequency(0.25,false) sets to update every 1/4 of a second
func (f *FrameAnimation) SetFrequency(freqIn float64, setFrequencyByTheFrame bool) {
	f.meta.frequency = freqIn
	f.meta.frequencyIsInFrames = setFrequencyByTheFrame
}

func (f *FrameAnimation) SetAsOneTimeOnly(setOneTime bool) {
	f.meta.oneTimeOnly = setOneTime
}

// currentImage returns the animation image associated with the current index in the animation
func (f *FrameAnimation) currentImage() Opengl.RenderObject {
	// TODO: possibly make this return a blank image when we shouldn't animate?
	return f.animationImages[f.meta.indexInAnimation]
}

func (f *FrameAnimation) update() bool {
	return f.meta.update(len(f.animationImages))
}

// meta update internally evaluates and increments toggle logic then returns true if we did swap images
func (meta *metaInfo) update(listSize int) bool {

	// verify we have stuff to animate, then check if we are ready to toggle
	if listSize > 0 && meta.shouldAnimate {

		// get the current time and check if our frequency has been met, if so then update the image
		timeNow := float64(time.Now().UnixNano()) / nanoToSeconds
		if meta.frequencyIsInFrames && meta.framesSinceLastToggle/int(meta.frequency) == 1 ||
			!meta.frequencyIsInFrames && timeNow-meta.timeOfLastToggle >= float64(meta.frequency) {

			// allow our animation to continue by resetting the index
			if meta.indexInAnimation == listSize-1 {
				meta.indexInAnimation = 0
				if meta.oneTimeOnly {
					meta.shouldAnimate = false
				}
			} else {
				meta.indexInAnimation++
			}
			meta.framesSinceLastToggle = 0
			meta.timeOfLastToggle = float64(time.Now().UnixNano()) / nanoToSeconds

			// indicate we have updated our current image
			return true
		}
		meta.framesSinceLastToggle++
	}
	return false
}
