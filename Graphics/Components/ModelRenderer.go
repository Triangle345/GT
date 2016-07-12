// Package Graphics
package Components

import (
	"GT/Graphics/G3D"
	"GT/Graphics/Opengl"
	"GT/Logging"
)

// NewSpriteRenderer creates a renderer and initializes its animation map
func NewModelRenderer() *modelRenderer {

	model := modelRenderer{}
	model.AnimationManager = &AnimationManager{animationsMap: map[string]*FrameAnimation{}}

	return &model
}

// SpriteRenderer is a component which allows a sprite to be drawn or animated
type modelRenderer struct {
	Renderer

	mesh Opengl.RenderObject

	// holds our map of animations as well as actions necessary to control them
	AnimationManager *AnimationManager

	// uvs - store uvs for speed
	uvs []float32

	// vertexData Opengl.OpenGLVertexInfo
}

func (this *modelRenderer) SetModel(path string, mat string) {
	//TODO make aggregate models store them in memory
	mesh, err := G3D.ParseOBJ(path, mat)

	if err != nil {
		Logging.Info(err)
	}

	this.mesh = mesh

	// this.vertexData = vertexData

}

// Update gets called every frame and accounts for all settings in
// the renderer as well as shifts animations
func (this *modelRenderer) Update(delta float32) {

	// run the animation update (if applicable) and set our
	// renderer mesh if the animation toggled
	if this.AnimationManager.currentAnimation != nil {
		// just care about updating the animation handler, thats it. We
		// want current image always if available
		this.AnimationManager.currentAnimation.update()
		this.mesh = this.AnimationManager.CurrentAnimation().currentImage()
	}

	if this.mesh == nil {
		return
	}

	this.Render(this.mesh)

}
