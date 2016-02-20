// Scene
package Graphics

import (
	"GT/Graphics/Opengl"
	"GT/Window"
	"fmt"
	"time"

	"github.com/go-gl/gl/v3.2-core/gl"
)

var elements = []uint32{
	0, 1, 2,
	2, 3, 0,
}
var quad_colours = []float32{
	0.583, 0.771, 0.014, 1,
	0.609, 0.115, 0.436, 1,
	0.327, 0.483, 0.844, 1,
	0.327, 0.483, 0.844, 1,
}
var quad_uvs = []float32{
	0.0, 0.0,
	1.0, 0.0,
	1.0, 1.0,
	0.0, 1.0,
}

type Scene interface {
	Load()
	Clear()
	Draw()
}

type BaseScene struct {
	entities      map[string]*Sprite
	spriteDraw    []*Sprite
	window        *Window.Window
	spriteSheet   *image
	LoadHandler   func()
	UpdateHandler func()
	//TODO can put fps into struct fps counter
	fps       int
	timestart int32
	update    bool
}

func NewBasicScene(spriteSheet string, window *Window.Window) (BaseScene, error) {
	img, err := NewImage(spriteSheet)

	if err != nil {
		fmt.Println("Cannot create image: " + err.Error())
	}

	s := BaseScene{window: window, spriteSheet: &img}
	s.init()
	s.fps = 0
	s.timestart = int32(time.Now().Unix())

	return s, err

}

func (s *BaseScene) init() {

	s.entities = make(map[string]*Sprite)
	s.LoadHandler = func() {}
	s.UpdateHandler = func() {}

}

func (s *BaseScene) AddSprite(id string, area RectangularArea) {

	if s.entities[id] == nil {
		fmt.Println("adding sprites: " + id)
		sprite := NewBasicSprite(area)
		s.entities[id] = &sprite

		var uvs []float32
		x, y := s.spriteSheet.GetUVFromPosition(area.BottomLeft())
		uvs = append(uvs, x, y)
		x, y = s.spriteSheet.GetUVFromPosition(area.BottomRight())
		uvs = append(uvs, x, y)
		x, y = s.spriteSheet.GetUVFromPosition(area.TopRight())
		uvs = append(uvs, x, y)
		x, y = s.spriteSheet.GetUVFromPosition(area.TopLeft())
		uvs = append(uvs, x, y)

		// set uv coords
		sprite.img.uvs = uvs

		s.spriteDraw = append(s.spriteDraw, &sprite)

		s.update = true
	}

}

func (s *BaseScene) GetSprite(id string) *Sprite {
	//TODO add error handling if doesnt exist
	return s.entities[id]
}

func (s *BaseScene) Start() {

	s.LoadHandler()

	for true {

		s.window.Clear()
		//s.UpdateHandler()
		s.Draw()
		s.window.Refresh()
	}
}

func (s *BaseScene) Draw() {

	if len(s.entities) == 0 {
		fmt.Println("No Sprites to draw.")
		return
	}

	// idx := uint32(0)

	data := Opengl.OpenGLVertexInfo{}
	stride := 0
	// // clone := &Opengl.OpenGLVertexInfo{}
	for i := 0; i < len(s.spriteDraw); i++ {

		sprite := s.spriteDraw[i]

		glInfo := sprite.getGLVertexInfo()

		glInfo.AdjustElements(uint32(stride))
		data.Append(&glInfo)

		stride += glInfo.Stride

	}

	// data.Print()

	if s.update {
		fmt.Println("In update")
		//TODO, remove hard coded resolution

		gl.BindTexture(gl.TEXTURE_2D, s.spriteSheet.textureId)

		Opengl.BindBuffers(data)
		s.update = false
	} else {

		//	Opengl.RepopulateVBO(data)

	}

	Opengl.Draw(data)

	// Opengl.Cleanup()

	// calc fps
	if (int32(time.Now().Unix()) - s.timestart) >= 1 {
		fmt.Printf("FPS: %d \n", s.fps)
		s.fps = 0
		s.timestart = int32(time.Now().Unix())

	}

	s.fps = s.fps + 1

}
