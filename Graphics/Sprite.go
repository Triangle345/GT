// Sprite
package Graphics

import (
	gl "github.com/chsc/gogl/gl21"
)

func NewBasicSprite(img image) Sprite {
	return Sprite{image: img, xS: 1, yS: 1}
}

type Sprite struct {
	image
	x, y, rot float32

	//scale
	xS, yS float32
}

func (s *Sprite) Move(x, y float32) {
	s.x += x
	s.y += y
}

func (s *Sprite) Rotate(r float32) {
	s.rot += r
}

func (s *Sprite) Scale(x, y float32) {
	s.xS *= x
	s.yS *= y
}

func (s *Sprite) SetLocation(x, y float32) {
	s.x = x
	s.y = y
}

func (s Sprite) Draw() {
	gl.MatrixMode(gl.MODELVIEW)

	gl.PushMatrix()

	gl.LoadIdentity()

	gl.Translatef(gl.Float(s.x), gl.Float(s.y), 0)

	gl.Rotatef(gl.Float(s.rot), 0.0, 0.0, 1.0)

	gl.Scalef(gl.Float(s.xS), gl.Float(s.yS), 0)

	s.image.Draw()

	gl.PopMatrix()
}
