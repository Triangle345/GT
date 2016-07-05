// Package Components ...
package Components

import "fmt"

// if keeping track of nanoseconds for updates, use this to convert to seconds
const nanoToSeconds = 1000000000

// this helper is to be packaged with your specific Animation in order to keep track of image switching
type AnimationHandler struct {
	animationsMap map[string]*FrameAnimation

	currentAnimation *FrameAnimation

	// // animation properties and tracking
	// IndexInAnimation int

	// framesSinceLastToggle int
	// timeOfLastToggle      float64
	// Frequency             float64
	// FrequencyIsInFrames   bool

	// OneTimeOnly   bool
	// ShouldAnimate bool
}

// // Animation interface provides Model and Sprite the same method names
// type Animation interface {

// 	// These will likely only differ by the list type
// 	Append(string)
// 	Reorder(int, int)
// 	Remove(int)
// 	SetCurrentAnimation(string)

// 	// These will likely be exactly the same for each struct (merely interfaced since this class has lower visibility)
// 	StartAnimation()
// 	StopAnimation()
// 	SetFrequency(float64, bool)
// 	SetOneTimeOnly(bool)
// }

// func NewAnimation() *AnimationHandler {

// 	// preset our defaults
// 	a := AnimationHandler{}
// 	a.IndexInAnimation = 0

// 	a.Frequency = 1
// 	a.FrequencyIsInFrames = true
// 	a.framesSinceLastToggle = 0
// 	a.timeOfLastToggle = float64(time.Now().UnixNano()) / nanoToSeconds

// 	a.OneTimeOnly = false
// 	a.ShouldAnimate = true

// 	return &a
// }

func (s *AnimationHandler) CurrentAnimation() *FrameAnimation {

	return s.currentAnimation
}

// SetCurrentAnimation takes a animation's name and tries to set that as the current animation
func (s *AnimationHandler) SetCurrentAnimation(mappedAnimationName string) {
	animationRetrieved, ok := s.animationsMap[mappedAnimationName]
	if !ok {
		fmt.Printf("couldn't find the animation" + mappedAnimationName)
	}

	s.currentAnimation = animationRetrieved
	fmt.Println(mappedAnimationName, ":Set current anim to", s.currentAnimation)
	fmt.Println("Animation Map: ", s.animationsMap)
	// s.img = s.currentAnimation.currentImage()

}

// AddAnimation maps a created animation inside the renderer
func (s *AnimationHandler) AddAnimation(animationToAdd *FrameAnimation, nameToMap string) {
	_, ok := s.animationsMap[nameToMap]
	if !ok {
		s.animationsMap[nameToMap] = animationToAdd
	} else {
		fmt.Printf("the animation " + nameToMap + " already exists, please try a different name")
	}
}

// StopAnimation tells our current animation that it should not animate (to be used when scripting)
func (s *AnimationHandler) StopAnimation() {
	s.currentAnimation.meta.ShouldAnimate = false
}

// StartAnimation tells our current animation that it should animate (to be used when scripting)
func (s *AnimationHandler) StartAnimation() {
	s.currentAnimation.meta.ShouldAnimate = true
}
