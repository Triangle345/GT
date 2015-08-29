// Graphics project Graphics.go
package Graphics

import (
	// "errors"
	// "fmt"
	// gl "github.com/chsc/gogl/gl21"
	mathgl "github.com/go-gl/mathgl/mgl32"
	// "github.com/go-gl/gl/v3.2-core/gl"

	// "os"
)

type RectangularArea interface {
	TopLeft() (x, y float32)
	TopRight() (x, y float32)
	BottomLeft() (x, y float32)
	BottomRight() (x, y float32)
	GetDimensions() (width, height float32)
}

type imageSection struct {
	startPos      mathgl.Vec2
	height, width float32
}

func (i imageSection) TopLeft() (x, y float32) {
	x, y = i.startPos[0], i.startPos[1]

	return
}

func (i imageSection) TopRight() (x, y float32) {
	x, y = i.startPos[0]+i.width, i.startPos[1]

	return
}

func (i imageSection) BottomLeft() (x, y float32) {
	x, y = i.startPos[0], i.startPos[1]+i.height

	return
}

func (i imageSection) BottomRight() (x, y float32) {
	x, y = i.startPos[0]+i.width, i.startPos[1]+i.height

	return
}

func (i imageSection) GetDimensions() (width, height float32) {
	width, height = i.width, i.height
	return
}

func NewImageSection(startX, startY, width, height float32) imageSection {

	return imageSection{startPos: mathgl.Vec2{startX, startY}, width: width, height: height}
}
