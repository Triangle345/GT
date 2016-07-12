// Package Components ...
package Components

import (
	"GT/Graphics/Image"
	"GT/Graphics/Opengl"
	"GT/Logging"
	"image"
)

// NewSpriteRenderer creates a renderer and initializes its animation map
func NewSpriteRenderer() *SpriteRenderer {
	sprite := SpriteRenderer{}
	sprite.AnimationManager = &AnimationManager{animationsMap: map[string]*FrameAnimation{}}

	return &sprite
}

// SpriteRenderer is a component which allows a sprite to be drawn or animated
type SpriteRenderer struct {
	// ChildComponent
	Renderer

	// holds our map of animations as well as actions necessary to control them
	AnimationManager *AnimationManager

	// singular image
	img Opengl.RenderObject
}

// SetImage puts a designated image from the agregate into our image which will be rendered
func (s *SpriteRenderer) SetImage(imageLoc string) {
	img, err := Image.NewImage(imageLoc)

	if err != nil {
		Logging.Debug("Cannot create image: " + err.Error())
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
		Logging.Debug("Cannot create image: " + err.Error())
	}

	img, err = img.SubImage(bounds)
	if err != nil {
		Logging.Debug("Cannot create sub image: " + err.Error())
	}

	s.img = &img
}

// Update gets called every frame and accounts for all settings in the renderer as well as shifts animations
func (s *SpriteRenderer) Update(delta float32) {

	// run the animation update (if applicable) and set our renderer image if the animation toggled
	if s.AnimationManager.currentAnimation != nil {
		// just care about updating the animation handler, thats it. We
		// want current image always if available
		s.AnimationManager.currentAnimation.update()
		s.img = s.AnimationManager.CurrentAnimation().currentImage()
	}

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
