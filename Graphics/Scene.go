// Scene
package Graphics

import (
	"GT/Graphics/Components"
	"GT/Graphics/Opengl"
	"GT/Window"
	"fmt"
	"time"
)

type Scene interface {
	Load()
	Clear()
	Draw()
}

type BaseScene struct {
	RootNode *Components.Node

	window *Window.Window
	//TODO can put fps into struct fps counter
	fps       int
	timestart int32
	update    bool
}

// orthographic scene
func NewBasicScene(window *Window.Window) (BaseScene, error) {

	s := BaseScene{RootNode: Components.NewNode("RootNode"), window: window}
	s.fps = 0
	s.timestart = int32(time.Now().Unix())
	s.update = true

	Opengl.SetOrthographic(window.Width, window.Height)

	return s, nil

}

func (s *BaseScene) Start() {

	//s.LoadHandler()
	s.update = true
	for true {

		s.window.Clear()

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
