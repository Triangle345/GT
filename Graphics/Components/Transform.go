package Components

import (
	mathgl "github.com/go-gl/mathgl/mgl32"
)

type AXIS int

func NewTransform() Transform {
	return Transform{model: mathgl.Ident4(), XS: 1, YS: 1, ZS: 1}
}

type Transform struct {

	// location
	X, Y, Z float32

	// rotation radians
	Rot float32

	// rotation (Euler Angle)
	XR, YR, ZR float32

	//scale
	XS, YS, ZS float32

	// the stored transform model
	model mathgl.Mat4
}

// Rotate first rotates on x, then y, then z
func (this *Transform) Rotate(r, x, y, z float32) {
	this.Rot = r
	this.XR = x
	this.YR = y
	this.ZR = z

}

func (s *Transform) Scale(x, y, z float32) {
	s.XS = x
	s.YS = y
	s.ZS = z
}

func (s *Transform) Translate(x, y, z float32) {
	s.X = x
	s.Y = y
	s.Z = z
}

func (this *Transform) GetUpdatedModel() mathgl.Mat4 {
	Model := this.model
	Model = Model.Mul4(mathgl.Translate3D(float32(this.X), float32(this.Y), float32(this.Z)))
	// Model = Model.Mul4(mathgl.HomogRotate3DZ(float32(s.rot)))
	// Model = Model.Mul4(mathgl.HomogRotate3D(s.Rot, mathgl.Vec3{s.XR, s.YR, s.ZR}))

	// Euler rotation is easy.. no quaternions yet.. what a pain.
	Model = Model.Mul4(mathgl.Rotate3DX(this.XR * this.Rot).Mat4())
	Model = Model.Mul4(mathgl.Rotate3DY(this.YR * this.Rot).Mat4())
	Model = Model.Mul4(mathgl.Rotate3DZ(this.ZR * this.Rot).Mat4())

	Model = Model.Mul4(mathgl.Scale3D(float32(this.XS), float32(this.YS), float32(this.ZS)))

	return Model
}
