package Components

import (
	"GT/Graphics/Image"
	"GT/Graphics/Opengl"

	mathgl "github.com/go-gl/mathgl/mgl32"
)

// ScriptComponent becomes a user defined script
// Contains helper methods
type Renderer struct {
	ChildComponent
	Color           *Image.Color
	customTransform func(mathgl.Vec4) mathgl.Vec4
}

func (this *Renderer) Initialize() {
	// this.Color.A = 1

}

// CustomTransforms allows the caller to apply a transformation to the current
// vertex data
func (this *Renderer) CustomTransform(f func(mathgl.Vec4) mathgl.Vec4) {
	this.customTransform = f
}

// Render writes vertex data of RenderObjects to opengl
func (this *Renderer) Render(obj Opengl.RenderObject) {
	Model := mathgl.Ident4()
	Model = this.GetParent().transform.GetUpdatedModel()

	vData := obj.VertexData()
	for j := 0; j < vData.NumVerts(); j++ {
		x, y, z := vData.GetVertex(j)
		transformation := mathgl.Vec4{x, y, z, 1}

		if this.customTransform != nil {
			transformation = this.customTransform(transformation)
		}

		t := Model.Mul4x1(transformation)

		vData.SetVertex(j, t[0], t[1], t[2])

		// set color if we have one
		if this.Color != nil {
			c := this.Color
			vData.SetColor(j, c.R, c.G, c.B, c.A)
		}

		// set normals for both model view and world (bump mapping and lighting)

		// world normal
		nX, nY, nZ := vData.GetWNormal(j)
		normalMat := Model.Mat3().Inv().Transpose()
		normal := normalMat.Mul3x1(mathgl.Vec3{nX, nY, nZ})
		vData.SetWNormal(j, normal.X(), normal.Y(), normal.Z())

	}

	// fmt.Println("Vertex Data: ", vData)

	// send OpenGLVertex info to Opengl module
	Opengl.AddVertexData(1, vData)

}

// Update is called just in case it is not overloaded
func (this *Renderer) Update(delta float32) {
	panic("Update in Renderer Needs to be overridden!")

}
