// Package Components ...
package Components

import (
	"time"
)

// if keeping track of nanoseconds for updates, use this to convert to seconds
const nanoToSeconds = 1000000000

// this helper is to be packaged with your specific Animation in order to keep track of image switching
type animationHelper struct {

	// animation properties and tracking
	indexInAnimation int

	framesSinceLastToggle int
	timeOfLastToggle      float64
	frequency             float64
	frequencyIsInFrames   bool

	oneTimeOnly   bool
	shouldAnimate bool
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

func newAnimation() *animationHelper {

	// preset our defaults
	a := animationHelper{}
	a.indexInAnimation = 0

	a.frequency = 1
	a.frequencyIsInFrames = true
	a.framesSinceLastToggle = 0
	a.timeOfLastToggle = float64(time.Now().UnixNano()) / nanoToSeconds

	a.oneTimeOnly = false
	a.shouldAnimate = true

	return &a
}

// Update internally evaluates and increments toggle logic then returns true if we did swap images
func (a *animationHelper) update(listSize int) bool {

	// verify we have stuff to animate, then check if we are ready to toggle
	if listSize > 0 && a.shouldAnimate {

		// get the current time and check if our frequency has been met, if so then update the image
		timeNow := float64(time.Now().UnixNano()) / nanoToSeconds
		if a.frequencyIsInFrames && a.framesSinceLastToggle/int(a.frequency) == 1 ||
			!a.frequencyIsInFrames && timeNow-a.timeOfLastToggle >= float64(a.frequency) {

			// allow our animation to continue by resetting the index
			if a.indexInAnimation == listSize-1 {
				a.indexInAnimation = 0
				if a.oneTimeOnly {
					a.shouldAnimate = false
				}
			} else {
				a.indexInAnimation++
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
