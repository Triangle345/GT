// Scene
package Graphics

import (
	"GT/Graphics/Components"
	"GT/Graphics/Opengl"
	"GT/Window"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type Scene interface {
	Start()
	Load()  // unused?
	Clear() // unused?
	Draw()  // draw() to be private
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
	running := true
	s.update = true
	drawStart := int(time.Now().Unix())
	drawEnd := int(time.Now().Unix())

	// continue to render the scene unless signalled for termination
	for running {

		// poll any current events and handle accordingly
		// TODO: while implementing Issue #9 abstract this to an input handler
		var event sdl.Event
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyDownEvent:
				if t.Keysym.Sym == sdl.K_ESCAPE {
					running = false
				} else {
					// fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
					// t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
				}
			default:
				// fmt.Printf("Some event\n")
			}
		}

		s.window.Clear()
		drawEnd = int(time.Now().Unix())

		// TODO: implement or store this data. Delta is currently not being used...
		delta := float32(drawEnd - drawStart)
		s.RootNode.Update(delta)
		drawStart = int(time.Now().Unix())
		s.Draw()

		// sleep for 1ms to smooth out the image rendering
		// TODO: consider making this user defined (fps limit)
		time.Sleep(1000)
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
