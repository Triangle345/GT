package Opengl

import "fmt"

type OpenGLVertexInfo struct {
	vertexData []float32
}

type RenderObject interface {
	VertexData() *OpenGLVertexInfo
}

func (this *OpenGLVertexInfo) Clone() *OpenGLVertexInfo {
	vinfo := OpenGLVertexInfo{}
	vData := make([]float32, len(this.vertexData))
	copy(vData, this.vertexData)

	vinfo.vertexData = vData
	return &vinfo
}

func (this *OpenGLVertexInfo) NewVertex(x, y, z float32) int {
	this.vertexData = append(this.vertexData,
		x, y, z, // vertex
		0, 0, 0, 1, // diffuse color
		0, 0, // uv
		0, 0, 0, //model view normal
		0, 0, 0, //w worldNormal
		TEXTURED, 0) // mode and sampler idx
	return this.NumVerts() - 1
}

func (this *OpenGLVertexInfo) SetMNormal(vIdx int, x, y, z float32) {
	this.vertexData[vIdx*int(NUM_ATTRIBUTES)+9] = x
	this.vertexData[vIdx*int(NUM_ATTRIBUTES)+10] = y
	this.vertexData[vIdx*int(NUM_ATTRIBUTES)+11] = z

}

func (this *OpenGLVertexInfo) GetMNormal(vIdx int) (x, y, z float32) {
	x = this.vertexData[vIdx*int(NUM_ATTRIBUTES)+9]
	y = this.vertexData[vIdx*int(NUM_ATTRIBUTES)+10]
	z = this.vertexData[vIdx*int(NUM_ATTRIBUTES)+11]

	return x, y, z

}

func (this *OpenGLVertexInfo) GetWNormal(vIdx int) (x, y, z float32) {
	x = this.vertexData[vIdx*int(NUM_ATTRIBUTES)+12]
	y = this.vertexData[vIdx*int(NUM_ATTRIBUTES)+13]
	z = this.vertexData[vIdx*int(NUM_ATTRIBUTES)+14]

	return x, y, z

}

func (this *OpenGLVertexInfo) SetWNormal(vIdx int, x, y, z float32) {
	this.vertexData[vIdx*int(NUM_ATTRIBUTES)+12] = x
	this.vertexData[vIdx*int(NUM_ATTRIBUTES)+13] = y
	this.vertexData[vIdx*int(NUM_ATTRIBUTES)+14] = z

}

func (this *OpenGLVertexInfo) SetUV(vIdx int, u, v float32) {
	this.vertexData[vIdx*int(NUM_ATTRIBUTES)+7] = u
	this.vertexData[vIdx*int(NUM_ATTRIBUTES)+8] = v
}

func (this *OpenGLVertexInfo) SetMode(vIdx int, m float32) {
	this.vertexData[vIdx*int(NUM_ATTRIBUTES)+15] = m
}

func (this *OpenGLVertexInfo) SetColor(vIdx int, r, g, b, a float32) {
	this.vertexData[vIdx*int(NUM_ATTRIBUTES)+3] = r
	this.vertexData[vIdx*int(NUM_ATTRIBUTES)+4] = g
	this.vertexData[vIdx*int(NUM_ATTRIBUTES)+5] = b
	this.vertexData[vIdx*int(NUM_ATTRIBUTES)+6] = a
}

func (this *OpenGLVertexInfo) SetVertex(vIdx int, x, y, z float32) {
	this.vertexData[vIdx*int(NUM_ATTRIBUTES)+0] = x
	this.vertexData[vIdx*int(NUM_ATTRIBUTES)+1] = y
	this.vertexData[vIdx*int(NUM_ATTRIBUTES)+2] = z
}

func (this *OpenGLVertexInfo) SetAggregateId(vIdx, aId int) {
	// add .1 to int and convert to float so we can round down in shader
	this.vertexData[vIdx*int(NUM_ATTRIBUTES)+16] = float32(aId) + float32(0.1)
}

func (this *OpenGLVertexInfo) GetVertex(vIdx int) (x, y, z float32) {
	x = this.vertexData[vIdx*int(NUM_ATTRIBUTES)+0]
	y = this.vertexData[vIdx*int(NUM_ATTRIBUTES)+1]
	z = this.vertexData[vIdx*int(NUM_ATTRIBUTES)+2]

	return x, y, z
}

func (this *OpenGLVertexInfo) NumVerts() int {
	return len(this.vertexData) / int(NUM_ATTRIBUTES)
}

func (i *OpenGLVertexInfo) append(o *OpenGLVertexInfo) {

	i.vertexData = append(i.vertexData, o.vertexData...)

}

func (i *OpenGLVertexInfo) Clear() {

	i.vertexData = []float32{}

}

func (i OpenGLVertexInfo) Print() {
	fmt.Printf("--------------------------------------------------\n")
	for _, v := range i.vertexData {
		fmt.Printf("inside data:%f \n", v)
	}
	fmt.Printf("+++++++++++++++++++++++++++++++++++++++++++++++++++++\n")

}
