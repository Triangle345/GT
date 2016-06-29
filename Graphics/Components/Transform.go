package Components

import (
	mathgl "github.com/go-gl/mathgl/mgl32"
)

func NewTransform() Transform {
	return Transform{model: mathgl.Ident4(), XS: 1, YS: 1, ZS: 1}
}

type Transform struct {

	// location
	X, Y, Z float32

	// rotation radians
	Rot float32

	// rotation axis
	XR, YR, ZR float32

	//scale
	XS, YS, ZS float32

	// the stored transform model
	model mathgl.Mat4
}

func (s *Transform) Rotate(r, x, y, z float32) {
	s.Rot += r
	s.XR = x
	s.YR = y
	s.ZR = z
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

func (s *Transform) GetUpdatedModel() mathgl.Mat4 {
	Model := s.model
	Model = Model.Mul4(mathgl.Translate3D(float32(s.X), float32(s.Y), float32(s.Z)))
	// Model = Model.Mul4(mathgl.HomogRotate3DZ(float32(s.rot)))
	Model = Model.Mul4(mathgl.HomogRotate3D(s.Rot, mathgl.Vec3{s.XR, s.YR, s.ZR}))

	Model = Model.Mul4(mathgl.Scale3D(float32(s.XS), float32(s.YS), float32(s.ZS)))

	return Model
}
