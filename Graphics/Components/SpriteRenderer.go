// Package Components ...
package Components

import (
	"GT/Graphics/Image"
	"GT/Graphics/Opengl"
	"fmt"
	"image"

	mathgl "github.com/go-gl/mathgl/mgl32"
)

// Initialize is necessary for the renderer to be utilized as a Component
func (s *SpriteRenderer) Initialize() {

}

// NewSpriteRenderer creates a renderer and initializes its animation map
func NewSpriteRenderer() *SpriteRenderer {
	sprite := SpriteRenderer{a: 1}
	sprite.animationsMap = map[string]*spriteAnimation{}
	return &sprite
}

// SpriteRenderer is a component which allows a sprite to be drawn or animated
type SpriteRenderer struct {
	// the parent node
	//Parent *Components.Node
	ChildComponent

	// singular image
	img *Image.Image

	// map of animations representing possible visuals for a sprite
	animationsMap    map[string]*spriteAnimation
	currentAnimation *spriteAnimation

	// color
	r, g, b, a float32
}

// SetCurrentAnimation takes a animation's name and tries to set that as the current animation
func (s *SpriteRenderer) SetCurrentAnimation(mappedAnimationName string) {
	animationRetrieved, ok := s.animationsMap[mappedAnimationName]
	if !ok {
		fmt.Printf("couldn't find the animation" + mappedAnimationName)
	}
	s.currentAnimation = animationRetrieved
	s.img = s.currentAnimation.currentImage()

}

// AddAnimation maps a created animation inside the renderer
func (s *SpriteRenderer) AddAnimation(animationToAdd *spriteAnimation, nameToMap string) {
	_, ok := s.animationsMap[nameToMap]
	if !ok {
		s.animationsMap[nameToMap] = animationToAdd
	} else {
		fmt.Printf("the animation " + nameToMap + " already exists, please try a different name")
	}
}

// StopAnimation tells our current animation that it should not animate (to be used when scripting)
func (s *SpriteRenderer) StopAnimation() {
	s.currentAnimation.meta.shouldAnimate = false
}

// StartAnimation tells our current animation that it should animate (to be used when scripting)
func (s *SpriteRenderer) StartAnimation() {
	s.currentAnimation.meta.shouldAnimate = true
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

	if s.img == nil {
		return
	}

	r := s.r
	g := s.g
	b := s.b
	a := s.a

	vData := s.img.VertexData()

	Model := mathgl.Ident4()

	Model = s.GetParent().transform.GetUpdatedModel()

	// transform all vertex data and combine it with other data

	for j := 0; j < vData.NumVerts(); j++ {
		x, y, z := vData.GetVertex(j)
		transformation := mathgl.Vec4{x, y, z, 1}
		t := Model.Mul4x1(transformation)

		vData.SetVertex(j, t[0], t[1], t[2])
		vData.SetColor(j, r, g, b, a)
	}

	// send OpenGLVertex info to Opengl module
	Opengl.AddVertexData(1, vData)

	// run the animation update (if applicable) and set our renderer image if the animation toggled
	if s.currentAnimation != nil && s.currentAnimation.update() {
		s.img = s.currentAnimation.currentImage()

	}
}

// SetColor allows us to modify image coloring of whatever is set in the Renderer
func (s *SpriteRenderer) SetColor(r, g, b, a float32) {
	s.r = r
	s.g = g
	s.b = b
	s.a = a
}
