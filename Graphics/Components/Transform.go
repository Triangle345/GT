package Components

import (
	mathgl "github.com/go-gl/mathgl/mgl32"
)

func NewTransform() Transform {
	return Transform{model: mathgl.Ident4(), xS: 1, yS: 1}
}

type Transform struct {

	// location
	x, y float32

	// rotation
	rot float32

	//scale
	xS, yS float32

	// the stored transform model
	model mathgl.Mat4
}

func (s *Transform) Rotate(r float32) {
	s.rot += r
}

func (s *Transform) Scale(x, y float32) {
	s.xS = x
	s.yS = y
}

func (s *Transform) SetLocation(x, y float32) {
	s.x = x
	s.y = y
}

func (s Transform) GetUpdatedModel() mathgl.Mat4 {
	Model := s.model
	Model = Model.Mul4(mathgl.Translate3D(float32(s.x), float32(s.y), float32(0.0)))
	Model = Model.Mul4(mathgl.HomogRotate3DZ(float32(s.rot)))
	Model = Model.Mul4(mathgl.Scale3D(float32(s.xS), float32(s.yS), float32(1)))

	return Model
}
