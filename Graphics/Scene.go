// Scene
package Graphics

import (
	"GT/Graphics/Components"
	"GT/Graphics/Opengl"
	"GT/Window"
	"fmt"
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
	RootNode Components.GameNode

	window *Window.Window
	//TODO can put fps into struct fps counter
	fps       int
	timestart int32
	update    bool
}

func NewBasicScene(window *Window.Window) (BaseScene, error) {

	s := BaseScene{RootNode: Components.NewNode("RootNode"), window: window}
	s.fps = 0
	s.timestart = int32(time.Now().Unix())
	s.update = true

	//TODO: find a better way to load all images in
	aggrImg.bind2GL()
	return s, nil

}

func (s *BaseScene) Start() {

	//s.LoadHandler()
	s.update = true
	for true {

		s.window.Clear()
		//s.UpdateHandler()
		//s.Draw()
		//TODO provide valid delta
		s.RootNode.Update(.34)
		s.Draw()
		s.window.Refresh()
	}
}

func (s *BaseScene) Draw() {

	// if len(s.entities) == 0 {
	// 	fmt.Println("No Sprites to draw.")
	// 	return
	// }
	//
	// for i := 0; i < len(s.spriteDraw); i++ {
	//
	// 	sprite := s.spriteDraw[i]
	//
	// 	glInfo := sprite.getGLVertexInfo()
	//
	// 	Opengl.AddVertexData(&glInfo)
	//
	// }

	// data.Print()

	if s.update {
		fmt.Println("In update")
		//TODO, remove hard coded resolution

		// gl.BindTexture(gl.TEXTURE_2D, s.spriteSheet.textureId)

		Opengl.BindBuffers()
		s.update = false
	}

	Opengl.Draw()

	// calc fps
	if (int32(time.Now().Unix()) - s.timestart) >= 1 {
		fmt.Printf("FPS: %d \n", s.fps)
		s.fps = 0
		s.timestart = int32(time.Now().Unix())
	}

	s.fps = s.fps + 1

}
