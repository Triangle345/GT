// Package Components ...
package Components

import (
	"GT/Graphics/G3D"
	"GT/Graphics/Image"
	"GT/Logging"
	"image"
	"io/ioutil"
	"os"
	"strings"
)

// AnimationManager contains a map of animations, and allows the user to interface with them
type AnimationManager struct {
	animationsMap    map[string]*FrameAnimation
	currentAnimation *FrameAnimation
}

func (a *AnimationManager) CurrentAnimation() *FrameAnimation {
	return a.currentAnimation
}

// SetCurrentAnimation takes a animation's name and tries to set that as the current animation
func (a *AnimationManager) SetCurrentAnimation(mappedAnimationName string) {
	animationRetrieved, ok := a.animationsMap[mappedAnimationName]
	if !ok {
		Logging.Debug("couldn't find the animation: " + mappedAnimationName)
	}

	a.currentAnimation = animationRetrieved
	Logging.Info(mappedAnimationName, ":Set current anim to ", a.currentAnimation)
	Logging.Info("Animation Map: ", a.animationsMap)
}

// AddAnimation maps a created animation inside the renderer
func (a *AnimationManager) AddAnimation(animationToAdd *FrameAnimation, nameToMap string) {
	_, ok := a.animationsMap[nameToMap]
	if !ok {
		a.animationsMap[nameToMap] = animationToAdd
	} else {
		Logging.Info("the animation " + nameToMap + " already exists, please try a different name")
	}
}

// GetAnimation grabs a mapped animation and returns it for the user to manipulate if desired
func (a *AnimationManager) GetAnimation(mappedName string) *FrameAnimation {
	animation, exists := a.animationsMap[mappedName]
	if !exists {
		Logging.Info("the animation " + mappedName + " wasn't found, please try a different name")
	}
	return animation
}

// StopAnimation tells our current animation that it should not animate (to be used when scripting)
func (a *AnimationManager) StopAnimation() {
	a.currentAnimation.meta.shouldAnimate = false
}

// StartAnimation tells our current animation that it should animate (to be used when scripting)
func (a *AnimationManager) StartAnimation() {
	a.currentAnimation.meta.shouldAnimate = true
}

// SpliceFullSpriteSheetAnimation manually cuts up an entire sprite sheet based on user defined frame dimensions
func SpliceFullSpriteSheetAnimation(imageLoc string, frameWidth int, frameHeight int) *FrameAnimation {

	// to set an entire sheet as an animation, set the Frames and Row Numbers to 0
	fa := SpliceSpriteSheetAnimation(imageLoc, frameWidth, frameHeight, 0, 0)

	return fa
}

// SpliceSpriteSheetAnimation manually cuts up a row of a sprite sheet based on user defined dimensions and sets it as the current animation
func SpliceSpriteSheetAnimation(imageLoc string,
	frameWidth, frameHeight, noOfFrames, rowNum int) *FrameAnimation {

	fa := newFrameAnimation()

	img, err := Image.NewImage(imageLoc)
	if err != nil {
		Logging.Debug("Cannot create image: " + err.Error())
	}

	// throw warnings for bad input
	numOfRows := float32(img.Bounds().Dy() / frameHeight)
	numOfColumns := float32(img.Bounds().Dx() / frameWidth)
	if float32(noOfFrames) > numOfColumns || numOfColumns < 1 {
		Logging.Debug("WARNING: frames out of bounds")
	}
	if float32(rowNum) > numOfRows || numOfRows < 1 {
		Logging.Debug("WARNING: row desired out of bounds")
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
				Logging.Debug("Cannot create sub image: " + err.Error())
			}
			fa.animationImages = append(fa.animationImages, &spriteSheetPart)
		}
	}

	return fa
}

// OBJAnimation takes the folder to OBJ files and creates animation from
// them
func OBJAnimation(loc string) *FrameAnimation {
	//TODO need to cache all models
	files, err := ioutil.ReadDir(loc)
	if err != nil {
		Logging.Info(err)
		return nil
	}

	fa := newFrameAnimation()

	for _, file := range files {
		// Logging.Info("Parsing anim OBJ: ", file.Name())
		splitExt := strings.Split(file.Name(), ".")
		if splitExt[1] == "obj" {
			obj := loc + string(os.PathSeparator) + splitExt[0] + ".obj"
			mtl := loc + string(os.PathSeparator) + splitExt[0] + ".mtl"
			mesh, err := G3D.ParseOBJ(obj, mtl)

			if err != nil {
				Logging.Info(err)
				continue
			}

			fa.animationImages = append(fa.animationImages, mesh)

		}

	}

	return fa
}
