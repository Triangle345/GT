package G3D

type vertex struct {
	X, Y, Z float32
}

type vertexNormal struct {
	X, Y, Z float32
}

type face struct {
	V, UV, VN []int
	Material  string
}

type Mesh struct {
	Name  string
	File  string
	Vs    []vertex
	VNs   []vertexNormal
	Faces []face
	// stride    int
	Materials map[string]*Material
}

// func (this *Mesh) RecalcElementStride() {
// 	maxIdx := 0
// 	for _, val := range this.Faces {
// 		for _, val2 := range val.V {
// 			if val2 > maxIdx {

// 				maxIdx = val2
// 			}
// 		}
// 	}

// 	// add one since its the idx, we need count
// 	this.stride = maxIdx + 1
// }

// func (this *Mesh) Stride() int {
// 	if this.stride == 0 {
// 		this.RecalcElementStride()
// 	}
// 	return this.stride
// }

// newmtl Material
// Ns 96.078431
// Ka 1.000000 1.000000 1.000000
// Kd 0.640000 0.640000 0.640000
// Ks 0.500000 0.500000 0.500000
// Ke 0.000000 0.000000 0.000000
// Ni 1.000000
// d 1.000000

type Color struct {
	R, G, B float32
}

type Material struct {
	Name     string
	File     string
	Ambient  Color
	Diffuse  Color
	Specular Color
	Emission Color
}
