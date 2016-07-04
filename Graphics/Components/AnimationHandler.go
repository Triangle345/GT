// Package Components ...
package Components

import (
	"time"
)

// if keeping track of nanoseconds for updates, use this to convert to seconds
const nanoToSeconds = 1000000000

// this helper is to be packaged with your specific Animation in order to keep track of image switching
type AnimationHandler struct {
	animationMap map[string]*FrameAnimation

	// animation properties and tracking
	IndexInAnimation int

	framesSinceLastToggle int
	timeOfLastToggle      float64
	Frequency             float64
	FrequencyIsInFrames   bool

	OneTimeOnly   bool
	ShouldAnimate bool
}

// Animation interface provides Model and Sprite the same method names
type Animation interface {

	// These will likely only differ by the list type
	Append(string)
	Reorder(int, int)
	Remove(int)
	SetCurrentAnimation(string)

	// These will likely be exactly the same for each struct (merely interfaced since this class has lower visibility)
	StartAnimation()
	StopAnimation()
	SetFrequency(float64, bool)
	SetOneTimeOnly(bool)
}

func NewAnimation() *AnimationHandler {

	// preset our defaults
	a := AnimationHandler{}
	a.IndexInAnimation = 0

	a.Frequency = 1
	a.FrequencyIsInFrames = true
	a.framesSinceLastToggle = 0
	a.timeOfLastToggle = float64(time.Now().UnixNano()) / nanoToSeconds

	a.OneTimeOnly = false
	a.ShouldAnimate = true

	return &a
}

// Update internally evaluates and increments toggle logic then returns true if we did swap images
func (a *AnimationHandler) Update(listSize int) bool {

	// verify we have stuff to animate, then check if we are ready to toggle
	if listSize > 0 && a.ShouldAnimate {

		// get the current time and check if our frequency has been met, if so then update the image
		timeNow := float64(time.Now().UnixNano()) / nanoToSeconds
		if a.FrequencyIsInFrames && a.framesSinceLastToggle/int(a.Frequency) == 1 ||
			!a.FrequencyIsInFrames && timeNow-a.timeOfLastToggle >= float64(a.Frequency) {

			// allow our animation to continue by resetting the index
			if a.IndexInAnimation == listSize-1 {
				a.IndexInAnimation = 0
				if a.OneTimeOnly {
					a.ShouldAnimate = false
				}
			} else {
				a.IndexInAnimation++
			}
			a.framesSinceLastToggle = 0
			a.timeOfLastToggle = float64(time.Now().UnixNano()) / nanoToSeconds

			// indicate we have updated our current image
			return true
		}
		a.framesSinceLastToggle++
	}
	return false
}
