// Package Graphics
package Model

import (
	"GT/Graphics/Components"
	"GT/Graphics/G3D"
	"GT/Graphics/Opengl"
	"GT/Logging"
)

// NewSpriteRenderer creates a renderer and initializes its animation map
func NewModelRenderer() *modelRenderer {

	model := modelRenderer{}

	return &model
}

// SpriteRenderer is a component which allows a sprite to be drawn or animated
type modelRenderer struct {
	Components.Renderer

	mesh Opengl.RenderObject

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

// Update gets called every frame and accounts for all settings in the renderer as well as shifts animations
func (this *modelRenderer) Update(delta float32) {

	// vertexData := this.mesh.VertexData()

	// Model := mathgl.Ident4()

	// Model = this.GetParent().Transform().GetUpdatedModel()

	// for i := 0; i < vertexData.NumVerts(); i++ {
	// 	x, y, z := vertexData.GetVertex(i)
	// 	transformation := mathgl.Vec4{x, y, z, 1}
	// 	t := Model.Mul4x1(transformation)
	// 	vertexData.SetVertex(i, t[0], t[1], t[2])

	// }

	// // send OpenGLVertex info to Opengl module
	// Opengl.AddVertexData(1, vertexData)

	this.Render(this.mesh)

}
