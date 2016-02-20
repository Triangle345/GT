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

type RectangularArea struct {
	startPos      mathgl.Vec2
	height, width float32
	uvs           []float32
}

func (i RectangularArea) GetUVs() []float32 {
	return i.uvs
}

func (i RectangularArea) TopLeft() (x, y float32) {
	x, y = i.startPos[0], i.startPos[1]

	return
}

func (i RectangularArea) TopRight() (x, y float32) {
	x, y = i.startPos[0]+i.width, i.startPos[1]

	return
}

func (i RectangularArea) BottomLeft() (x, y float32) {
	x, y = i.startPos[0], i.startPos[1]+i.height

	return
}

func (i RectangularArea) BottomRight() (x, y float32) {
	x, y = i.startPos[0]+i.width, i.startPos[1]+i.height

	return
}

func (i RectangularArea) GetDimensions() (width, height float32) {
	width, height = i.width, i.height
	return
}

func NewRectangularArea(startX, startY, width, height float32) RectangularArea {

	return RectangularArea{startPos: mathgl.Vec2{startX, startY}, width: width, height: height}
}
