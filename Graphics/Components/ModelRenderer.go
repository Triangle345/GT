// Package Graphics
package Components

import (
	"GT/Graphics/G3D"
	"GT/Graphics/Opengl"
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

	vertexData Opengl.OpenGLVertexInfo
}

// func (this *modelRenderer) SetModel(path string, mat string) {
// 	//TODO make aggregate models store them in memory
// 	mesh, err := G3D.ParseOBJ(path, mat)

// 	if err != nil {
// 		Logging.Info(err)
// 	}

// 	this.mesh = mesh
// 	this.mesh.RecalcElementStride()

// 	// r := float32(.3)
// 	// g := float32(.3)
// 	// b := float32(.3)
// 	// a := float32(1.000)

// 	this.vertexData.Stride = this.mesh.Stride()
// 	this.vertexData.Elements = make([]uint32, 0, this.mesh.Stride()+1)

// 	colors := make([]G3D.Color, this.mesh.Stride())
// 	fmt.Println("Materials Color: ", this.mesh.Materials)
// 	for _, face := range this.mesh.Faces {
// 		fmt.Println("Material: ", face)
// 		for _, vIdx := range face.V {
// 			this.vertexData.Elements = append(this.vertexData.Elements, uint32(vIdx))
// 			colors[vIdx] = this.mesh.Materials[face.Material].Diffuse
// 		}
// 	}

// 	// fmt.Println("Colors: ", colors)

// 	for idx, v := range this.mesh.Vs {
// 		r := colors[idx].R
// 		g := colors[idx].G
// 		b := colors[idx].B
// 		a := float32(1.0)

// 		this.vertexData.VertexData = append(this.vertexData.VertexData, v.X, v.Y, v.Z)     // vdata
// 		this.vertexData.VertexData = append(this.vertexData.VertexData, r, g, b, a)        // color
// 		this.vertexData.VertexData = append(this.vertexData.VertexData, -1, -1)            // texture
// 		this.vertexData.VertexData = append(this.vertexData.VertexData, Opengl.NO_TEXTURE) //mode

// 	}

// }

// Update gets called every frame and accounts for all settings in the renderer as well as shifts animations
func (this *modelRenderer) Update(delta float32) {

	// Model := mathgl.Ident4()

	// Model = this.GetParent().transform.GetUpdatedModel()

	// vdata := this.vertexData.VertexData
	// var data = make([]float32, len(this.vertexData.VertexData)*2)
	// copy(data, vdata)

	// for i := 0; i < this.vertexData.Stride; i++ {

	// 	transformation := mathgl.Vec4{vdata[i*10+0], vdata[i*10+1], vdata[i*10+2], 1}
	// 	t := Model.Mul4x1(transformation)
	// 	data[i*10+0] = t[0]
	// 	data[i*10+1] = t[1]
	// 	data[i*10+2] = t[2]

	// }

	// // package everything up in an OpenGLVertexInfo
	// vertexInfo := Opengl.OpenGLVertexInfo{
	// 	VertexData: data,
	// 	Elements:   this.vertexData.Elements,
	// 	Stride:     this.vertexData.Stride,
	// }

	// // send OpenGLVertex info to Opengl module
	// Opengl.AddVertexData(1, &vertexInfo)

}
