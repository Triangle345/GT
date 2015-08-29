// Scene
package Graphics

import (
	"GT/Graphics/Opengl"
	"GT/Window"
	"fmt"
	"github.com/go-gl/gl/v3.2-core/gl"
	"time"
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
		sprite := NewBasicSprite(area)
		s.entities[id] = &sprite
		s.spriteDraw = append(s.spriteDraw, &sprite)
	}

}

func (s *BaseScene) GetSprite(id string) *Sprite {
	//TODO add error handling if doesnt exist
	return s.entities[id]
}

func (s *BaseScene) Start() {

	// TODO load only needs to be done once!
	s.LoadHandler()

	for true {

		s.window.Clear()
		s.UpdateHandler()
		s.Draw()
		s.window.Refresh()
	}
}

func (s *BaseScene) Draw() {

	if len(s.entities) == 0 {
		fmt.Println("No Sprites to draw.")
		return
	}

	idx := uint32(0)

	vertexInfo := &Opengl.OpenGLVertexInfo{}

	// clone := &Opengl.OpenGLVertexInfo{}
	for _, v := range s.spriteDraw {

		//TODO: move this to sprite class and use interface instead.. what if i have "shape"
		// fmt.Printf("Img name: %s\n", k)
		// glInfo := v.getGLVertexInfo()
		w, h := v.GetImageSection().GetDimensions()

		for i := 0; i < 4; i++ {
			vertexInfo.Translations = append(vertexInfo.Translations, float32(v.x), float32(v.y), 0)
			vertexInfo.Rotations = append(vertexInfo.Rotations, 0, 0, 1, v.rot)
			vertexInfo.Scales = append(vertexInfo.Scales, v.xS, v.yS, 0)
			vertexInfo.Colors = append(vertexInfo.Colors, v.r, v.g, v.b, v.a)

		}

		vertexInfo.Vertices = append(vertexInfo.Vertices, -0.5*w, 0.5*h, 1.0, 0.5*w, 0.5*h, 1.0, 0.5*w, -0.5*h, 1.0, -0.5*w, -0.5*h, 1.0)

		// TODO: maybe put this inside struct to store since this will really never change
		x, y := s.spriteSheet.GetUVFromPosition(v.GetImageSection().BottomLeft())
		vertexInfo.Uvs = append(vertexInfo.Uvs, x, y)
		x, y = s.spriteSheet.GetUVFromPosition(v.GetImageSection().BottomRight())
		vertexInfo.Uvs = append(vertexInfo.Uvs, x, y)
		x, y = s.spriteSheet.GetUVFromPosition(v.GetImageSection().TopRight())
		vertexInfo.Uvs = append(vertexInfo.Uvs, x, y)
		x, y = s.spriteSheet.GetUVFromPosition(v.GetImageSection().TopLeft())
		vertexInfo.Uvs = append(vertexInfo.Uvs, x, y)

		vertexInfo.Elements = append(vertexInfo.Elements, uint32(idx*4), uint32(idx*4+1), uint32(idx*4+2), uint32(idx*4), uint32(idx*4+2), uint32(idx*4+3))

		// glInfo.AdjustElements(idx)

		// idx += 4
		// clone.Append(glInfo)
		idx = idx + 1
	}

	// vertexInfo.Print()
	Opengl.BindBuffers(vertexInfo)

	gl.BindTexture(gl.TEXTURE_2D, s.spriteSheet.textureId)

	Opengl.Draw(vertexInfo)

	Opengl.Cleanup()

	// calc fps
	if (int32(time.Now().Unix()) - s.timestart) >= 1 {
		fmt.Printf("FPS: %d \n", s.fps)
		s.fps = 0
		s.timestart = int32(time.Now().Unix())

	}

	s.fps = s.fps + 1

}
