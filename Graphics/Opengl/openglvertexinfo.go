package Opengl

import "fmt"

type OpenGLVertexInfo struct {
	vertexData []float32
	Elements   []uint32

	// TODO: make stride private and set it statically using new OpenGLVertexInfo2D or something
	Stride      int
	totalStride int
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
	this.vertexData = append(this.vertexData, x, y, z, 0, 0, 0, 1, 0, 0, TEXTURED, 0)
	return this.NumVerts() - 1
}

func (this *OpenGLVertexInfo) SetUV(vIdx int, u, v float32) {
	this.vertexData[vIdx*int(NUM_ATTRIBUTES)+7] = u
	this.vertexData[vIdx*int(NUM_ATTRIBUTES)+8] = v
}

func (this *OpenGLVertexInfo) SetMode(vIdx int, m float32) {
	this.vertexData[vIdx*int(NUM_ATTRIBUTES)+9] = m
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
	this.vertexData[vIdx*int(NUM_ATTRIBUTES)+10] = float32(aId) + float32(0.1)
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

// TODO: need to put this in as part of vertex info and into hashhash
func (this *OpenGLVertexInfo) adjustElements(o *OpenGLVertexInfo) {

	for k, _ := range o.Elements {
		o.Elements[k] += uint32(this.totalStride)
		// fmt.Println(i.Elements[k])
	}

	//TODO: add argument of stride to add stride of other vertexinfo
	this.totalStride += o.Stride
}

func (i *OpenGLVertexInfo) append(o *OpenGLVertexInfo) {
	i.adjustElements(o)

	i.vertexData = append(i.vertexData, o.vertexData...)
	i.Elements = append(i.Elements, o.Elements...)

	//i.Stride += o.Stride

}

func (i *OpenGLVertexInfo) Clear() {

	i.vertexData = make([]float32, 0, 1000000)
	i.Elements = make([]uint32, 0, 1000000)
	i.totalStride = 0

}

func (i OpenGLVertexInfo) Print() {
	fmt.Printf("--------------------------------------------------\n")
	for _, v := range i.vertexData {
		fmt.Printf("inside data:%f \n", v)
	}
	fmt.Printf("+++++++++++++++++++++++++++++++++++++++++++++++++++++\n")
	for _, v := range i.Elements {
		fmt.Printf("inside element:%f \n", v)
	}
	fmt.Printf("stride %d \n", i.Stride)
	fmt.Printf("Total stride %d \n", i.totalStride)

}
