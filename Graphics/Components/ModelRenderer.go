// Package Graphics
package Components

import (
	"GT/Graphics/G3D"
	"GT/Graphics/Opengl"
	"GT/Logging"

	mathgl "github.com/go-gl/mathgl/mgl32"
)

// Initialize is necessary for the renderer to be utilized as a Component
func (s *modelRenderer) Initialize() {

}

// NewSpriteRenderer creates a renderer and initializes its animation map
func NewModelRenderer() *modelRenderer {

	model := modelRenderer{}

	return &model
}

// SpriteRenderer is a component which allows a sprite to be drawn or animated
type modelRenderer struct {
	ChildComponent

	mesh *G3D.Mesh

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

	vertexData := Opengl.OpenGLVertexInfo{}

	for _, face := range this.mesh.Faces {
		for _, vIdx := range face.V {
			c := this.mesh.Materials[face.Material].Diffuse
			v := this.mesh.Vs[vIdx]
			vdID := vertexData.NewVertex(v.X, v.Y, v.Z)
			vertexData.SetColor(vdID, c.R, c.G, c.B, 1)
			vertexData.SetMode(vdID, Opengl.NO_TEXTURE)
		}

	}

	Model := mathgl.Ident4()

	Model = this.GetParent().transform.GetUpdatedModel()

	for i := 0; i < vertexData.NumVerts(); i++ {
		x, y, z := vertexData.GetVertex(i)
		transformation := mathgl.Vec4{x, y, z, 1}
		t := Model.Mul4x1(transformation)
		vertexData.SetVertex(i, t[0], t[1], t[2])

	}

	// send OpenGLVertex info to Opengl module
	Opengl.AddVertexData(1, &vertexData)

}
