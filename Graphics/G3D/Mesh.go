package G3D

import (
	"GT"
	"GT/Graphics/Image"
	"GT/Graphics/Opengl"
	"GT/Logging"
	"image"
)

type vertex struct {
	X, Y, Z float32
}

type vertexNormal struct {
	X, Y, Z float32
}

type vertexTexture struct {
	U, V float32
}

type face struct {
	V, UV, VN []int
	Material  string
}

func (this *Mesh) VertexData() *Opengl.OpenGLVertexInfo {
	vertexData := Opengl.OpenGLVertexInfo{}

	for _, face := range this.Faces {
		for idx, vIdx := range face.V {
			c := this.Materials[face.Material].Diffuse
			v := this.Vs[vIdx]
			vdID := vertexData.NewVertex(v.X, v.Y, v.Z)
			vertexData.SetColor(vdID, c.R, c.G, c.B, 1)

			tex := this.Materials[face.Material].DiffuseTex

			if len(face.UV) == 0 || tex == "" {
				vertexData.SetMode(vdID, Opengl.NO_TEXTURE)
			} else {

				imgSec := Image.AggrImg.GetImageSection(GT.AssetsImages + tex)

				if imgSec == nil {
					Logging.Info("Cannot open: ", GT.AssetsImages+tex, " For Mat: ", face.Material)

				}

				u := this.VTs[face.UV[idx]].U
				v := this.VTs[face.UV[idx]].V

				// x starts from left to right
				locX := int(float32(imgSec.Bounds().Dx()) * u)
				// y starts bottom to top so we need to convert.
				locY := int(float32(imgSec.Bounds().Max.Y) - float32(imgSec.Bounds().Dy())*v)

				locX += imgSec.Section.Min.X
				locY += imgSec.Section.Min.Y

				// TODO move GetUVFromPosition into maybe image aggregator
				newU, newV := Image.AggrImg.GetUVFromPosition(image.Point{locX, locY})

				vertexData.SetUV(vdID, newU, newV)
				vertexData.SetMode(vdID, Opengl.TEXTURED)
			}
		}
	}
	return &vertexData
}

type Mesh struct {
	Name      string
	File      string
	Vs        []vertex
	VNs       []vertexNormal
	VTs       []vertexTexture
	Faces     []face
	Materials map[string]*Material
}

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
	Name       string
	File       string
	Ambient    Color
	Diffuse    Color
	Specular   Color
	Emission   Color
	DiffuseTex string
	AmbientTex string
}
