// Sprite
package Graphics

import (
	gl "github.com/chsc/gogl/gl21"
)

func NewBasicSprite(img image) Sprite {
	return Sprite{image: img}
}

type Sprite struct {
	image
	x, y float64
}

func (s *Sprite) Move(x, y float64) {
	s.x += x
	s.y += y
}

func (s *Sprite) SetLocation(x, y float64) {
	s.x = x
	s.y = y
}

func (s Sprite) Draw() {
	gl.MatrixMode(gl.MODELVIEW)

	gl.PushMatrix()

	gl.LoadIdentity()

	gl.Translated(gl.Double(s.x), gl.Double(s.y), 0)

	s.image.Draw()

	gl.PopMatrix()
}
